package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func UnpackZip(dst, src string) error {

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {

		filename := filepath.Join(dst, f.Name)

		// It's a dir?
		i := strings.LastIndex(f.Name, "/")
		if i == (len(f.Name) - 1) {
			_ = os.Mkdir(filename, 0777)
			continue
		}

		fileEntry, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return err
		}

		r, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(fileEntry, r)
		if err != nil {
			return err
		}
	}

	return nil
}

func ShowArchiveFiles(src string) error {

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	_, archive := filepath.Split(src)

	fmt.Printf("Archive: %s\n", archive)
	fmt.Printf("%10s | %-19s | %-20s\n", "Length", "Modified", "Name")
	fmt.Println("---------- + ------------------- + --------------------")

	var num int8
	var size uint64
	for _, f := range r.File {

		num++
		size += f.CompressedSize64

		fmt.Printf("%10d | %-19s | %-20s\n",
			f.CompressedSize64, f.Modified.Format("2006-01-02 15:04:05"), f.Name)
	}

	fmt.Println("------------                     ----------------------")
	fmt.Printf("%10d                         %d files\n", size, num)

	return nil
}
