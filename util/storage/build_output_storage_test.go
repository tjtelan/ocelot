package storage

import (
	"bytes"
	"github.com/mitchellh/go-homedir"
	"github.com/shankj3/ocelot/util"
	"path/filepath"
	"testing"
)


func testGetHomeDirec() string{
	direc, err := homedir.Dir()
	if err != nil {
		panic("need home direc capability")
	}
	return direc
}
var hash = "12345678"
var HomeDirec = testGetHomeDirec()

var filebuildstorages = []struct{
	initSaveDirec   string
	actualSaveLoc string
}{
	{"", filepath.Join(HomeDirec, ".ocelot", "build-output", hash)},
	{"/tmp/test/one", filepath.Join("/tmp", "test", "one", hash)},
}

func TestFileBuildStorage(t *testing.T) {
	testBytes := []byte("woooooooeooooeooeeeoeo!!!!!")

	for _, fb := range filebuildstorages {
		fbs := &FileBuildStorage{
			saveDirec: fb.initSaveDirec,
		}
		fbs.setup()
		defer fbs.Clean()
		err := fbs.Store(hash, testBytes)
		if err != nil {
			t.Fatal(err)
		}
		if filep := fbs.GetTempFile(hash); filep != fb.actualSaveLoc {
			t.Error(util.StrFormatErrors("file path", fb.actualSaveLoc, filep))
		}

		actualData, err := fbs.Retrieve(hash)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(actualData, testBytes){
			t.Error(util.GenericStrFormatErrors("file data", string(testBytes), string(actualData)))
		}
	}


}