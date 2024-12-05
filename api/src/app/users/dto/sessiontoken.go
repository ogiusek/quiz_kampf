package dto

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

var sessionDuration time.Duration = time.Minute * 5

type SessionToken string

func (sessionToken SessionToken) Hash() string {
	hasher := sha256.New()
	hasher.Write([]byte(string(sessionToken)))
	return string(hex.EncodeToString(hasher.Sum(nil)))
}

func (sessionToken SessionToken) HashMatchesSessionToken(hash string) bool {
	return sessionToken.Hash() == hash
}
