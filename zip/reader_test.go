package zip

import (
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestUnpackZip(t *testing.T) {

	dir, _ := filepath.Abs(".")

	src := path.Join(dir, "testdata", "testdata.zip")
	t.Log("archive src:", src)

	dst, _ := os.MkdirTemp("", "testdata*")
	t.Log("archive unpack dst:", dst)

	err := UnpackZip(dst, src)
	if err != nil {
		t.Error(err)
	}
}

func TestShowArchiveFiles(t *testing.T) {

	dir, _ := filepath.Abs(".")
	src := path.Join(dir, "testdata", "testdata.zip")

	t.Log("archive src:", src)

	err := ShowArchiveFiles(src)
	if err != nil {
		t.Error(err)
	}
}
