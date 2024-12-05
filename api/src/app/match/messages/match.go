package messages

import (
	matchmodels "lib/app/match/models"
	metadatarepo "lib/app/ws/repo"
	"lib/services"
	socketstorage "lib/services/socketStorage"
)

var MatchMessageTopic = "match"

func SendMatch(match matchmodels.Match) {
	metadataRepo := metadatarepo.GetMetadataRepo()
	socketStorage, freeSocketStorage := services.SocketStoage()
	dto := match // missing question data dto
	defer freeSocketStorage()
	for _, player := range match.Players {
		data, err := metadataRepo.GetByUserId(player.UserId)
		if err != nil {
			continue // if users exited when sending message than ignore it
		}
		socketStorage.
			Get(data.SocketId).
			Send(socketstorage.NewMessage(MatchMessageTopic, dto))
	}
}
