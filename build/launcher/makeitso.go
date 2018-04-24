package launcher

import (
	"context"
	"fmt"
	"time"

	ocelog "bitbucket.org/level11consulting/go-til/log"
	"bitbucket.org/level11consulting/ocelot/build"
	"bitbucket.org/level11consulting/ocelot/build/valet"
	"bitbucket.org/level11consulting/ocelot/models"
	"bitbucket.org/level11consulting/ocelot/models/pb"
	"bitbucket.org/level11consulting/ocelot/storage"
)

// watchForResults sends the *Transport object over the transport channel for stream functions to process
func (w *launcher) WatchForResults(hash string, dbId int64) {
	ocelog.Log().Debugf("adding hash ( %s ) & infochan to transport channel", hash)
	transport := &models.Transport{Hash: hash, InfoChan: w.infochan, DbId: dbId}
	w.StreamChan <- transport
}


// MakeItSo will call appropriate builder functions
func (w *launcher) MakeItSo(werk *pb.WerkerTask, builder build.Builder, finish, done chan int) {
	// right off the bat we know all the environment variables that will exist for the
	// lifetime of the container, so just add them now
	w.addGlobalEnvVars(werk, builder)

	ocelog.Log().Debug("hash build ", werk.CheckoutHash)
	w.BuildValet.RegisterDoneChan(werk.CheckoutHash, done)
	defer w.BuildValet.MakeItSoDed(finish)
	defer w.BuildValet.UnregisterDoneChan(werk.CheckoutHash)
	defer func(){
		ocelog.Log().Info("calling done for nsqpb")
		done <- 1
	}()

	// create context for entire build
	ctx, cancel := context.WithCancel(context.Background())

	//send build context off, build kills are performed by calling cancel on the cancellable context
	w.BuildCtxChan <- &models.BuildContext{
		Hash: werk.CheckoutHash,
		Context: ctx,
		CancelFunc: cancel,
	}

	defer cancel()
	defer func(){
		ocelog.Log().Info("closing infochan for ", werk.Id)
		close(w.infochan)
	}()

	w.WatchForResults(werk.CheckoutHash, werk.Id)

	//update consul with active build data
	consul := w.RemoteConf.GetConsul()
	// if we can't register with consul, bail, just exit out. the maintainer will soon be pausing message flow anyway
	if err := w.BuildValet.StartBuild(consul, werk.CheckoutHash, werk.Id); err != nil {
		return
	}

	dockerIdChan := make (chan string)
	go w.listenForDockerUuid(dockerIdChan, werk.CheckoutHash)

	// do setup stage
	setupStart := time.Now()
	w.BuildValet.Reset("setup", werk.CheckoutHash)
	setupResult, dockerUUid := builder.Setup(ctx, w.infochan, dockerIdChan, werk, w.RemoteConf, w.ServicePort)
	defer w.BuildValet.Cleanup(ctx, dockerUUid, w.infochan)
	setupDura := time.Now().Sub(setupStart)
	if err := storeStageToDb(w.Store, werk.Id, setupResult, setupStart, setupDura.Seconds()); err != nil {
		ocelog.IncludeErrField(err).Error("couldn't store build output")
		return
	}
	if setupResult.Status == pb.StageResultVal_FAIL {
		handleFailure(setupResult, w.Store, "setup", setupDura, werk.Id)
		return
	}

	// run integrations, executable download, codebase download
	if err := w.preFlight(ctx, werk, builder); err != nil { return }

	// run the actual stages outlined in the ocelot.yml
	fail, dura, err := w.runStages(ctx, werk, builder)
	if err != nil {
		return
	}

	//update build_summary table
	if err := w.Store.UpdateSum(fail, dura.Seconds(), werk.Id); err != nil {
		ocelog.IncludeErrField(err).Error("couldn't update summary in database")
		return
	}
	ocelog.Log().Infof("finished building id %s", werk.CheckoutHash)
}


// addGlobalEnvVars sets the global env vars on builders, these are the variables that will live for the duration of the build.
//	- `GIT_HASH`
//	- `BUILD_ID`
//	- `GIT_HASH_SHORT`
//	- `GIT_BRANCH`
func (w *launcher) addGlobalEnvVars(werk *pb.WerkerTask, builder build.Builder) {
	paddedEnvs := []string{
		fmt.Sprintf("GIT_HASH=%s", werk.CheckoutHash),
		fmt.Sprintf("BUILD_ID=%d", werk.Id),
		fmt.Sprintf("GIT_HASH_SHORT=%s", werk.CheckoutHash[:7]),
		fmt.Sprintf("GIT_BRANCH=%s", werk.Branch),
	}
	paddedEnvs = append(paddedEnvs, werk.BuildConf.Env...)
	builder.SetGlobalEnv(paddedEnvs)
}


func (w *launcher) listenForDockerUuid(dockerChan chan string, checkoutHash string) error {
	dockerUuid := <- dockerChan

	if err := valet.RegisterBuild(w.RemoteConf.GetConsul(), w.Uuid.String(), checkoutHash, dockerUuid); err != nil {
		ocelog.IncludeErrField(err).Error("couldn't register build")
		return err
	}

	return nil
}

//storeStageToDb is a helper function for storing stages to db - this runs on completion of every stage
func storeStageToDb(storage storage.BuildStage, buildId int64, stageResult *pb.Result, start time.Time, dur float64) error {
	err := storage.AddStageDetail(&models.StageResult{
		BuildId:       buildId,
		Stage:         stageResult.Stage,
		Status:        int(stageResult.Status),
		Error:         stageResult.Error,
		Messages:      stageResult.Messages,
		StartTime:     start,
		StageDuration: dur,
	})

	if err != nil {
		return err
	}

	return nil
}

func handleFailure(result *pb.Result, store storage.OcelotStorage, stageName string, duration time.Duration, id int64) {
	errStr := fmt.Sprintf("%s stage failed", stageName)
	if len(result.Error) > 0 {
		errStr = errStr + result.Error
	}
	ocelog.Log().Error(errStr)
	if err := store.UpdateSum(true, duration.Seconds(), id); err != nil {
		ocelog.IncludeErrField(err).Error("couldn't update summary in database")
	}
}

