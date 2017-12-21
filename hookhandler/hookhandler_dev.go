package hookhandler

import (
	"io/ioutil"
	"bitbucket.org/level11consulting/ocelot/admin/models"
	"bitbucket.org/level11consulting/ocelot/admin/handler"
	"os"
)

//has necessary functions for running hookhandler in dev mode

type MockHookHandlerContext struct {
	HookHandlerContext //embedding hookhandler context cause didn't want to stub out getters/setters
}

func (hhc *MockHookHandlerContext) GetBitbucketClient(cfg *models.Credentials) (handler.VCSHandler, string, error) {
	mockVcsHandler := &MockVCSHandler{}
	return mockVcsHandler , "", nil
}

type MockVCSHandler struct {}

func (mvh *MockVCSHandler) Walk() error {
	return nil
}

//TODO: ****WARNING**** this assumes you're running this inside of the hookhandler folder
func (mvh *MockVCSHandler) GetFile(filePath string, fullRepoName string, commitHash string) (bytez []byte, err error) {
	pwd, _ := os.Getwd()
	return ioutil.ReadFile(pwd + "/test-fixtures/dev-ocelot.yml")
}

func (mvh *MockVCSHandler) CreateWebhook(webhookURL string) error {
	return nil
}

func (mvh *MockVCSHandler) GetCallbackURL() string {
	return ""
}

func (mvh *MockVCSHandler) SetCallbackURL(callbackURL string) {}

func (mvh *MockVCSHandler) SetBaseURL(baseURL string) {}

func (mvh *MockVCSHandler) GetBaseURL() string {
	return ""
}

func (mvh *MockVCSHandler) FindWebhooks(getWebhookURL string) bool {
	return true
}