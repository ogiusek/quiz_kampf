package main

import (
	"lib/app/match"
	"lib/app/ping"
	"lib/app/questions"
	"lib/app/users"
	usersconfig "lib/app/users/config"
	userdto "lib/app/users/dto"
	"lib/app/ws"
	"lib/common/httpapi"
	"lib/common/servicepool"
	"lib/config"
	"lib/services"
	"lib/services/filestorage"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func configDb() {
	pool := servicepool.NewPool(
		func() *gorm.DB {
			db, err := gorm.Open(postgres.Open(config.Config.ConnectionString), &gorm.Config{})
			if err != nil {
				log.Fatalf("Could not connect to the database: %v", err)
				os.Exit(1)
			}
			var maintainConnection func()
			maintainConnection = func() {
				time.Sleep(12 * time.Hour)
				db.Raw("SELECT 1", nil)
				go maintainConnection()
			}
			go maintainConnection()
			return db
		},
		25,
	)
	services.SetDb(pool)
	log.Println("Configured database")
}

//

func configFileStorage() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Could not get caller information")
	}
	dir := filepath.Dir(filepath.Dir(filename)) + "/storage/"
	filestorage.Config = filestorage.Configuration{
		StoragePath: dir,
		ApiUrl:      "http://localhost:5000/",
	}
	pool := servicepool.NewInfinitePool(
		func() filestorage.FileStorage { return filestorage.NewFileStorage() },
	)
	services.SetFileStorage(pool)
	log.Println("Configured file storage")
}

//

func configUserModule() {
	usersconfig.Config = usersconfig.Configuration{
		JwtSecret: []byte("64f6d320cbe67cba3d23802bebbea99190e4f1e49117c3c8f956b93ac68ab93c"),
	}
}

func main() {
	configDb()
	configFileStorage()
	configUserModule()

	app := httpapi.New()
	app.Prefix("api")

	ws.UseWs(&app)
	ping.UsePing(&app)
	users.AddUsers(&app)
	questions.AddQuestions(&app)
	match.AddMatches(&app)
	app.Use(func(w http.ResponseWriter, r *http.Request, next func()) {
		log.Printf("received request to %v", r.RequestURI)
		next()
	})

	port := strconv.Itoa(config.Config.Port)
	log.Println("listening on :" + port)

	corsObj := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Origin", "Content-Type", userdto.SessionHeader})

	err := http.ListenAndServe(":"+port, handlers.CORS(corsObj, methods, headers)(app))
	if err != nil {
		panic(err)
	}
}
