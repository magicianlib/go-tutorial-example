package zip

import (
	"archive/zip"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// CreateArchive create a zip file
//
// compress the file specified by src into a zip(dst) file
//
func CreateArchive(dst string, src ...string) error {

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

	return CreateArchiveW(archive, src...)
}

func CreateArchiveW(w io.Writer, src ...string) error {

	// zip archive writer
	zw := zip.NewWriter(w)
	defer zw.Close()

	for _, f := range src {

		stat, err := os.Stat(f)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			if err := addArchiveDir(zw, f); err != nil {
				return err
			}
		} else {
			if err := addArchiveFile(zw, "", f, &stat); err != nil {
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

func CreateArchiveF(w io.Writer, files []File) error {

	zw := zip.NewWriter(w)
	defer zw.Close()

	for _, file := range files {

		f, err := zw.Create(file.Name)
		if err != nil {
			return err
		}

		_, err = f.Write([]byte(file.Body))
		if err != nil {
			return err
		}
	}

	return nil
}
