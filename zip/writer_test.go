package zip

import (
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestCreateZipArchive(t *testing.T) {

	tDir, err := os.MkdirTemp("", "testdata*")
	if err != nil {
		t.Error(err)
	}

	dst := path.Join(tDir, "testdata.zip")

	t.Log("archive dst:", dst)

	dir, _ := filepath.Abs(".")
	testdata := path.Join(dir, "testdata")

	var src = []string{
		path.Join(testdata, "order-1.csv"),
		path.Join(testdata, "dir"),
	}

	for _, f := range src {
		t.Log("archive src:", f)
	}

	err = CreateZipArchive(dst, src...)
	if err != nil {
		t.Error(err)
	}
}
