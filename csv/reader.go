package csv

import (
	"errors"
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

func Reader(src string, dst interface{}) error {

	if src == "" {
		return errors.New("csv: file(src) not exist")
	}

	f, err := os.Open(src)
	if err != nil {
		return err
	}

	err = gocsv.Unmarshal(f, dst)
	if err != nil {
		return err
	}

	return nil
}

func ReaderFromR(r io.Reader, dst interface{}) error {

	err := gocsv.Unmarshal(r, dst)
	if err != nil {
		return err
	}

	return nil
}

func ReaderString(src *string, dst interface{}) error {

	err := gocsv.UnmarshalString(*src, dst)
	if err != nil {
		return err
	}

	return nil
}
