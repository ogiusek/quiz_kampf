package messages

import (
	"lib/app/match/models"
	"lib/app/ws/repo"
	"lib/services"
	socketstorage "lib/services/socketStorage"
)

var StartedTopic = "started"

func SendStarted(match models.Match) {
	metadataRepo := repo.GetMetadataRepo()
	socketStorage, freeSocketStorage := services.SocketStoage()
	defer freeSocketStorage()
	dto := struct{}{} // no data on match start
	for _, player := range match.Players {
		metadata, err := metadataRepo.GetByUserId(player.UserId)
		if err != nil {
			continue
		}
		socketStorage.
			Get(metadata.SocketId).
			Send(socketstorage.NewMessage(StartedTopic, dto))
	}
}
