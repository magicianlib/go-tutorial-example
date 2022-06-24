package zip

import (
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestCreateZip(t *testing.T) {

	tDir, err := os.MkdirTemp("", "testdata*")
	if err != nil {
		t.Error(err)
	}

	zipPath := path.Join(tDir, "testdata.zip")

	t.Logf("writer zip: %s\n", zipPath)

	dir, _ := filepath.Abs(".")
	testdata := path.Join(dir, "testdata")

	var files = []string{
		path.Join(testdata, "order-1.csv"),
		path.Join(testdata, "dir"),
	}

	err = CreateZip(zipPath, files...)
	if err != nil {
		t.Error(err)
	}
}
