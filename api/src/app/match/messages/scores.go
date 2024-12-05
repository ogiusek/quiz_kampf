package messages

import (
	"lib/app/match/models"
	metadatarepo "lib/app/ws/repo"
	"lib/services"
	socketstorage "lib/services/socketStorage"
)

var ScoresMessageTopic = "scores"

func SendScores(match models.Match) {
	metadataRepo := metadatarepo.GetMetadataRepo()
	socketStorage, freeSocketStorage := services.SocketStoage()
	dto := struct {
		Players []models.PlayerData `json:"players"`
	}{
		Players: match.Players,
	}
	defer freeSocketStorage()
	for _, player := range match.Players {
		data, err := metadataRepo.GetByUserId(player.UserId)
		if err != nil {
			continue // if users exited when sending message than ignore it
		}
		socketStorage.
			Get(data.SocketId).
			Send(socketstorage.NewMessage(ScoresMessageTopic, dto))
	}
}
