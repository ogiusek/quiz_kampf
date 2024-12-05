package match

import (
	"encoding/json"
	"lib/app/match/messages"
	matchrepo "lib/app/match/repo"
	"lib/app/match/usecases"
	metadatarepo "lib/app/ws/repo"
	"lib/common/httpapi"
	"lib/common/id"
	"lib/services"
	socketstorage "lib/services/socketStorage"
	"log"
	"time"
)

func LogState() {
	time.Sleep(5 * time.Second)
	repo := matchrepo.GetMatchRepo()
	matches := repo.GetAll()
	json, _ := json.Marshal(matches)
	log.Printf("all current matches %s\n", json)
	go LogState()
}

func AddMatches(app httpapi.IApi) {
	socketStorage, freeSocketStorage := services.SocketStoage()
	defer freeSocketStorage()
	go LogState()
	socketStorage.OnConnect(func(id id.ID, ws socketstorage.Ws) {
		ws.OnClose(func() {
			metadataRepo := metadatarepo.GetMetadataRepo()
			data, err := metadataRepo.GetBySocket(id)
			if err != nil {
				log.Fatalf("there is no socket %v", id)
				return
			}
			usecases.Quit(usecases.QuitArgs{Session: data.Session})
		})

		// messages from server
		messages.MatchMessageTopic = "match_state"
		messages.QuestionMessageTopic = "question"
		messages.ScoresMessageTopic = "scores"
		messages.StartedTopic = "start"

		// messages to server
		ws.OnMessage("host", Parse(id, ws, usecases.Host))
		ws.OnMessage("join", Parse(id, ws, usecases.Join))
		ws.OnMessage("quit", Parse(id, ws, usecases.Quit))

		ws.OnMessage("add_question", Parse(id, ws, usecases.AddQuestion))
		ws.OnMessage("remove_question", Parse(id, ws, usecases.RemoveQuestion))
		ws.OnMessage("set_answer_time", Parse(id, ws, usecases.SetAnswerTime))

		ws.OnMessage("start", Parse(id, ws, usecases.Start))
		ws.OnMessage("answer", Parse(id, ws, usecases.Answer))
	})
}
