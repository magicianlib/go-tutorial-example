package zip

import (
	"archive/zip"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// CreateZipArchive create a zip file
//
// compress the file specified by src into a zip(dst) file
//
func CreateZipArchive(dst string, src ...string) error {

	if len(src) == 0 {
		return errors.New("zip: not found need archive file")
	}

	if dst == "" {
		return errors.New("zip: archive pack path is empty")
	}

	// remove dst if exist
	_ = os.Remove(dst)

	// create zip archive
	archive, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer archive.Close()

	// zip archive writer
	w := zip.NewWriter(archive)
	defer w.Close()

	for _, f := range src {

		stat, err := os.Stat(f)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			if err := addArchiveDir(w, f); err != nil {
				return err
			}
		} else {
			if err := addArchiveFile(w, "", f, &stat); err != nil {
				return err
			}
		}
	}

	return nil
}

func addArchiveDir(w *zip.Writer, dir string) error {

	err := filepath.Walk(dir, func(path string, fi fs.FileInfo, err error) error {

		//if strings.Compare(dir, path) == 0 {
		//	return nil
		//}

		return addArchiveFile(w, dir, path, &fi)
	})

	return err
}

func addArchiveFile(w *zip.Writer, dir string, path string, fi *fs.FileInfo) error {

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
