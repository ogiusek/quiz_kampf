package models

import (
	"errors"
	"lib/app/users/models"
	"lib/common/id"
)

var ErrUserNotFound error = errors.New("this account do not exists")

type PlayerData struct {
	UserId  id.ID       `json:"user_id"`
	User    models.User `json:"user"`
	IsHost  bool        `json:"is_host"`
	IsReady bool        `json:"is_ready"`
	Score   int         `json:"score"`
}

func NewPlayer(user models.User) PlayerData {
	return PlayerData{
		UserId:  user.UserId,
		User:    user,
		IsHost:  false,
		IsReady: false,
		Score:   0,
	}
}

func NewAdmin(user models.User) PlayerData {
	return PlayerData{
		UserId:  user.UserId,
		User:    user,
		IsHost:  true,
		IsReady: false,
		Score:   0,
	}
}

func (player *PlayerData) SetReady(ready bool) {
	player.IsReady = ready
}

func (player *PlayerData) AddScore(score int) {
	player.Score += score
}

func (player *PlayerData) Reset() {
	player.IsReady = false
	player.Score = 0
}
