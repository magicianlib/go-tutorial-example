package csv

import (
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

// Writer write data into dst(file)
//
func Writer(dst string, data interface{}) error {

	f, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	err = gocsv.MarshalFile(data, f)
	if err != nil {
		return err
	}

	return nil
}

// WriterToW write data with io.Writer
//
func WriterToW(w io.Writer, data interface{}) error {

	err := gocsv.Marshal(data, w)
	if err != nil {
		return err
	}

	return nil
}

func WriterString(data interface{}) (string, error) {
	return gocsv.MarshalString(data)
}
