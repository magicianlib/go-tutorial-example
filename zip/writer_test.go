package zip

import (
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestCreateArchive(t *testing.T) {

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

	err = CreateArchive(dst, src...)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateArchiveF(t *testing.T) {

	var files = []File{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	f, err := os.CreateTemp("", "*test.zip")
	if err != nil {
		t.Error(err)
	}

	t.Log("archive dst:", f.Name())

	err = CreateArchiveF(f, files)
	if err != nil {
		t.Error(err)
	}
}
