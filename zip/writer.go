package zip

import (
	"archive/zip"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func CreateZip(zipPath string, filepath ...string) error {

	if len(filepath) == 0 {
		return errors.New("zip: filepath is empty")
	}

	// remove zipPath if exist
	_ = os.Remove(zipPath)

	// create zip
	f, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	// zip writer
	w := zip.NewWriter(f)
	defer func(w *zip.Writer) {
		_ = w.Close()
	}(w)

	for _, filename := range filepath {

		stat, err := os.Stat(filename)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			err := writerDir(w, filename)
			if err != nil {
				return err
			}
		} else {
			err := writerFile(w, filename, &stat)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func writerFile(w *zip.Writer, filename string, fi *fs.FileInfo) error {
	return writer(w, "", filename, fi)
}

func writerDir(w *zip.Writer, dir string) error {

	err := filepath.Walk(dir, func(path string, fi fs.FileInfo, err error) error {

		//if strings.Compare(dir, path) == 0 {
		//	return nil
		//}

		return writer(w, dir, path, &fi)
	})

	return err
}

func writer(w *zip.Writer, dir string, path string, fi *fs.FileInfo) error {

	header, err := zip.FileInfoHeader(*fi)
	if err != nil {
		return err
	}

	// determine filename
	var filename string
	if dir == "" {
		_, filename = filepath.Split(path)
	} else {
		dir = filepath.Dir(dir)
		filename, err = filepath.Rel(dir, path)
		if err != nil {
			return err
		}
	}

	// Store OR Deflate
	header.Method = zip.Store
	header.Name = filename

	if (*fi).IsDir() {
		header.Name += "/"
	}

	hw, err := w.CreateHeader(header)
	if err != nil {
		return err
	}

	if (*fi).IsDir() {
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	_, err = io.Copy(hw, f)
	if err != nil {
		return err
	}

	return nil
}
