package models

import (
	"crypto/sha256"
	"database/sql/driver"
	"encoding/hex"
	"errors"
)

type UserHash string

func (hash UserHash) Valid() error {
	if len(hash) != 64 {
		return errors.New("hash length has to be 64")
	}
	return nil
}

func (UserHash) GormDataType() string { return "varchar(64)" }
func (vo UserHash) GetValue() (driver.Value, error) {
	if err := vo.Valid(); err != nil {
		return nil, err
	}
	return string(vo), nil
}

// this can be placed in userPassword
func (password UserPassword) Hash() UserHash {
	hasher := sha256.New()
	hasher.Write([]byte(string(password)))
	return UserHash(hex.EncodeToString(hasher.Sum(nil)))
}

func (password UserPassword) Matches(hash UserHash) bool {
	comparedHash := password.Hash()
	return comparedHash == hash
}
