package models

import "errors"

var ErrPlayerLimitOutOfRange error = errors.New("players limit has to be between 2 and 50")

type PlayersLimit int

func (limit PlayersLimit) Valid() error {
	if limit < 2 && limit > 50 {
		return ErrPlayerLimitOutOfRange
	}
	return nil
}
