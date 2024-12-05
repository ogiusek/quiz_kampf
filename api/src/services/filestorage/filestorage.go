package filestorage

import (
	"fmt"
	"io"
	"lib/common/errs"
	"lib/common/id"
	"os"
	"path/filepath"
)

type fileStorage struct{}

func NewFileStorage() FileStorage {
	return &fileStorage{}
}

func save(fileName File, stream io.Reader) error {
	filePath := Config.StoragePath + string(fileName)
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("cannot create directories: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return errs.Exists("file already exists")
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return fmt.Errorf("cannot copy file: %w", err)
	}
	return nil
}

func (*fileStorage) Save(stream io.Reader) (File, error) {
	fileName := File("generated/" + string(id.New()))
	if err := save(fileName, stream); err != nil {
		return File(""), err
	}
	return fileName, nil
}

func (f *fileStorage) SaveAs(fileName File, stream io.Reader) error {
	if err := save(fileName, stream); err != nil {
		return err
	}
	return nil
}

func (*fileStorage) Remove(fileName File) error {
	filePath := Config.StoragePath + string(fileName)
	err := os.Remove(filePath) // Attempt to remove the file
	if err != nil {
		return errs.NotFound("File do not exists")
	}
	return nil
}
