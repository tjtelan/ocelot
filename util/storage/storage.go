package storage

import (
	"bitbucket.org/level11consulting/ocelot/util/storage/models"
	"fmt"
	"time"
)

type Dest int

const (
	FileSystem Dest = iota
	Postgres
)

type Stringy interface {
	StorageType() string
}

type BuildOut interface {
	AddOut(output *models.BuildOutput) error
	RetrieveOut(buildId int64) (models.BuildOutput, error)
	RetrieveLastOutByHash(gitHash string) (models.BuildOutput, error)
}


type BuildSum interface {
	// AddSumStart will
	AddSumStart(hash string, starttime time.Time, account string, repo string, branch string) (int64, error)
	UpdateSum(failed bool, duration float64, id int64) error
	RetrieveSum(gitHash string) ([]models.BuildSummary, error)
	RetrieveLatestSum(gitHash string) (models.BuildSummary, error)
	RetrieveHashStartsWith(partialGitHash string) ([]models.BuildSummary, error)
	RetrieveLastFewSums(repo string, account string, limit int32) ([]models.BuildSummary, error)
}

type BuildStage interface {
	AddStageDetail(stageResult *models.StageResult) error
	RetrieveStageDetail(buildId int64) ([]models.StageResult, error)
}


type OcelotStorage interface {
	BuildOut
	BuildSum
	BuildStage
	Stringy
}

var (
	BUILD_SUM_404 = "no build summary found for %s"
	STAGE_REASON_404 = "no stages found for %s"
	BUILD_OUT_404 = "no build output found for %s"
)


func BuildSumNotFound(id string) *ErrNotFound {
	return &ErrNotFound{fmt.Sprintf(BUILD_SUM_404, id)}
}

func BuildOutNotFound(id string) *ErrNotFound {
	return &ErrNotFound{fmt.Sprintf(BUILD_OUT_404, id)}
}

func StagesNotFound(id string) *ErrNotFound {
	return &ErrNotFound{fmt.Sprintf(STAGE_REASON_404, id)}
}

type ErrNotFound struct {
	msg string
}

func (e *ErrNotFound) Error() string {
	return e.msg
}