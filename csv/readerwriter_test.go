package csv

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"
)

type member struct {
	Id       string    `csv:"账号"`
	Name     string    `csv:"昵称"`
	Age      uint8     `csv:"年龄"`
	Birthday time.Time `csv:"生日"`
	Gender   string    `csv:"gender"`
	Ignore   bool      `csv:"-"` // Ignore by csv

}

var members = []member{
	{
		Id:       "4BAB0D03-07D4-418B-B0B2-B38EAE4F4743",
		Name:     "Han Meimei",
		Age:      20,
		Birthday: time.Now(),
		Gender:   "女",
		Ignore:   false,
	},
	{
		Id:       "2D0D34BE-5457-4211-96D6-F0A6505645F6",
		Name:     "翠花",
		Age:      18,
		Birthday: time.Now(),
		Gender:   "girl",
		Ignore:   false,
	}, {
		Id:       "402A61F4-297C-465E-9F4F-CB7EFC6F1DA0",
		Name:     "Li Lei",
		Age:      20,
		Birthday: time.Now(),
		Gender:   "男",
		Ignore:   false,
	}, {
		Id:       "C0331985-90F8-40BF-A5E3-D5C38DBC2733",
		Name:     "大锤",
		Age:      20,
		Birthday: time.Now(),
		Gender:   "boy",
		Ignore:   false,
	},
}

type DateTime struct {
	time.Time
}

func (t *DateTime) MarshalCSV() (string, error) {
	return t.Format("2006-01-02 15:04:05"), nil
}

type memberC struct {
	Id       string    `csv:"账号"`
	Name     string    `csv:"昵称"`
	Age      uint8     `csv:"年龄"`
	Birthday time.Time `csv:"生日"`
	Gender   string    `csv:"gender"`
	Ignore   bool      `csv:"-"` // Ignore by csv
}

var membersC = []memberC{
	{
		Id:       "4BAB0D03-07D4-418B-B0B2-B38EAE4F4743",
		Name:     "Han Meimei",
		Age:      20,
		Birthday: time.Now(),
		Gender:   "女",
		Ignore:   false,
	},
	{
		Id:       "2D0D34BE-5457-4211-96D6-F0A6505645F6",
		Name:     "翠花",
		Age:      18,
		Birthday: time.Now(),
		Gender:   "girl",
		Ignore:   false,
	}, {
		Id:       "402A61F4-297C-465E-9F4F-CB7EFC6F1DA0",
		Name:     "Li Lei",
		Age:      20,
		Birthday: time.Now(),
		Gender:   "男",
		Ignore:   false,
	}, {
		Id:       "C0331985-90F8-40BF-A5E3-D5C38DBC2733",
		Name:     "大锤",
		Age:      20,
		Birthday: time.Now(),
		Gender:   "boy",
		Ignore:   false,
	},
}

func TestWriter(t *testing.T) {

	f, err := os.CreateTemp("", "*members.csv")
	if err != nil {
		t.Error(err)
	}

	t.Log("csv dst:", f.Name())

	err = Writer(f.Name(), &members)
	if err != nil {
		t.Error(err)
	}
}

func TestWriterToW(t *testing.T) {

	f, err := os.CreateTemp("", "*members.csv")
	if err != nil {
		t.Error(err)
	}

	t.Log("csv dst:", f.Name())

	err = WriterToW(f, &members)
	if err != nil {
		t.Error(err)
	}
}

func TestWriterString(t *testing.T) {

	// w := new(bytes.Buffer)
	w := os.Stdout

	err := WriterToW(w, &members)
	if err != nil {
		t.Error(err)
	}
}

func TestWriterCustom(t *testing.T) {

	gocsv.SetCSVWriter(func(w io.Writer) *gocsv.SafeCSVWriter {

		nw := csv.NewWriter(w)
		nw.Comma = '|' // Custom Separator

		return gocsv.NewSafeCSVWriter(nw)
	})

	err := WriterToW(os.Stdout, &membersC)
	if err != nil {
		t.Error(err)
	}
}

func TestReader(t *testing.T) {

	dir, _ := filepath.Abs(".")

	src := path.Join(dir, "testdata", "members.csv")

	var dst []*member

	err := Reader(src, &dst)
	if err != nil {
		t.Error(err)
	}
}

func TestReaderFromR(t *testing.T) {

	dir, _ := filepath.Abs(".")

	src := path.Join(dir, "testdata", "members.csv")
	f, err := os.Open(src)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	var dst []*member

	err = ReaderFromR(f, &dst)
	if err != nil {
		t.Error(err)
	}
}

func TestReaderCustom(t *testing.T) {

	gocsv.SetCSVReader(func(r io.Reader) gocsv.CSVReader {
		nw := csv.NewReader(r)
		nw.Comma = '|'
		return nw
	})

	dir, _ := filepath.Abs(".")

	src := path.Join(dir, "testdata", "members_custom.csv")

	var dst []*memberC

	err := Reader(src, &dst)
	if err != nil {
		t.Error(err)
	}
}
