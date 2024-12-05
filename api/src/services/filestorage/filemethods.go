package filestorage

import (
	"database/sql/driver"
	"errors"
	"io"
	"os"

	"github.com/h2non/filetype"
	"github.com/ogiusek/hw/src/hw"
)

func (file File) Valid() error {
	val := string(file)
	if len(val) == 0 {
		return errors.New("File cannot be empty")
	}
	if val[0] == '/' {
		return errors.New("File can be only relative path")
	}
	return nil
}

func (fileRelativePath File) Response() hw.Resp {
	if Config.StoragePath == "" {
		return nil
	}
	resp := hw.NewResponse()

	filePath := Config.StoragePath + string(fileRelativePath)
	file, err := os.Open(filePath)
	if err != nil {
		resp.WriteHeader(404)
		return resp
	}
	defer file.Close()

	fileType, err := filetype.MatchFile(filePath)
	if err == nil {
		resp.WriteHeader(500)
		resp.Write([]byte("cannot decode file type"))
		return resp
	}
	io.Copy(resp, file)
	resp.Header().Set("Content-Type", fileType.MIME.Value)
	return resp
}

func (file File) Url() FileUrl { return FileUrl(Config.ApiUrl + string(file)) }

func (file File) GormDataType() string { return "varchar(128)" }
func (file File) GetValue() (driver.Value, error) {
	if err := file.Valid(); err != nil {
		return nil, err
	}
	return string(file), nil
}
