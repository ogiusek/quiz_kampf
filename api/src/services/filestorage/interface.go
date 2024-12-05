package filestorage

import (
	"io"
)

type Configuration struct {
	StoragePath string // start and end with /
	ApiUrl      string // end with /
}

var Config Configuration

type FileUrl string
type File string

type FileStorage interface {
	Save(io.Reader) (File, error)
	SaveAs(File, io.Reader) error
	Remove(File) error
}
