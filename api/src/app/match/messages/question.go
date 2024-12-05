package messages

import (
	"lib/app/match/models"
	metadatarepo "lib/app/ws/repo"
	"lib/services"
	socketstorage "lib/services/socketStorage"
)

var QuestionMessageTopic = "new-question"

func SendQuestion(match models.Match) {
	metadataRepo := metadatarepo.GetMetadataRepo()
	socketStorage, freeSocketStorage := services.SocketStoage()
	question, _ := match.GetCurrentQuestion()
	dto := struct {
		Question models.QuestionData `json:"question"`
	}{
		Question: *question,
	}
	defer freeSocketStorage()
	for _, player := range match.Players {
		data, err := metadataRepo.GetByUserId(player.UserId)
		if err != nil {
			continue // if users exited when sending message than ignore him
		}
		socketStorage.
			Get(data.SocketId).
			Send(socketstorage.NewMessage(QuestionMessageTopic, dto))
	}
}
