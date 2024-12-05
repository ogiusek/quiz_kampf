package ws

import (
	"lib/app/users/dto"
	"lib/app/ws/models"
	"lib/app/ws/repo"
	"lib/common/errs"
	"lib/common/httpapi"
	"lib/common/id"
	"lib/services"
	socketstorage "lib/services/socketStorage"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func migrate() {
	db, freeDb := services.Db()
	db.AutoMigrate(&models.Metadata{})
	freeDb()
}

func UseWs(app httpapi.IApi) {
	migrate()
	metadataRepo := repo.GetMetadataRepo()
	metadataRepo.DeleteAll()

	storage, freeStorage := services.SocketStoage()
	defer freeStorage()
	storage.OnConnect(func(i id.ID, w socketstorage.Ws) {
		w.OnClose(func() {
			metadataRepo := repo.GetMetadataRepo()
			metadataRepo.Delete(i)
		})
	})

	app.Map("v1/ws", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		sessionTokenParam := "session_token"
		sessionToken := dto.SessionToken(r.URL.Query().Get(sessionTokenParam))
		session, err := sessionToken.DecodeSession()

		if err != nil {
			errs.ToHttp(w, err)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error during connection upgrade:", err)
			return
		}

		ws := socketstorage.NewWebSocket(conn)
		storage, freeStorage := services.SocketStoage()
		id := storage.Add(ws)
		freeStorage()

		metadataRepo := repo.GetMetadataRepo()
		metadataRepo.Add(models.Metadata{
			SocketId: id,
			Session:  session,
		})
	})
}
