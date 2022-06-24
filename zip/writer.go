package zip

import (
	"archive/zip"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// CreateZip create a zip file
//
// compress the file specified by src into a zip(dest) file
//
func CreateZip(dest string, src ...string) error {

	if len(src) == 0 {
		return errors.New("zip: filepath(src) is empty")
	}

	// remove dest zip file if exist
	_ = os.Remove(dest)

	// create a new zip file
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	// create a new zip writer
	w := zip.NewWriter(f)
	defer f.Close()

	for _, filename := range src {

		stat, err := os.Stat(filename)
		if err != nil {
			return err
		}

		if stat.IsDir() {
			if err := walkDir(w, filename); err != nil {
				return err
			}
		} else {
			if err := packFile(w, "", filename, &stat); err != nil {
				return err
			}
		}
	}

	return nil
}

func walkDir(w *zip.Writer, dir string) error {

	err := filepath.Walk(dir, func(path string, fi fs.FileInfo, err error) error {

		//if strings.Compare(dir, path) == 0 {
		//	return nil
		//}

		return packFile(w, dir, path, &fi)
	})

	return err
}

func packFile(w *zip.Writer, absDir string, absPath string, fi *fs.FileInfo) error {

	header, err := zip.FileInfoHeader(*fi)
	if err != nil {
		return err
	}

	// determine filename
	var filename string
	if absDir == "" {
		_, filename = filepath.Split(absPath)
	} else {
		absDir = filepath.Dir(absDir)
		filename, err = filepath.Rel(absDir, absPath)
		if err != nil {
			return err
		}
	}

	// Store OR Deflate
	header.Method = zip.Deflate
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

	f, err := os.Open(absPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(hw, f)
	if err != nil {
		return err
	}

	return nil
}
