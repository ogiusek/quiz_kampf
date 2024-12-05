package dto

import (
	"lib/common/id"
	"time"
)

var refreshDuration time.Duration = time.Hour * 24 * 7

type RefreshToken string

func (sessionToken SessionToken) RefreshToken(id id.ID) (RefreshToken, error) {
	return RefreshTokenPayload{
		UserId: id,
		Hash:   sessionToken.Hash(),
	}.Encode()
}
