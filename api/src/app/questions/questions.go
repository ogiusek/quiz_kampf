package questions

import (
	"lib/app/questions/models"
	"lib/app/questions/usecases"
	"lib/common/httpapi"
	"lib/common/wraps"
	"lib/services"
	"log"
	"net/http"
	"os"

	"github.com/ogiusek/hw/src/hw"
)

func AddQuestions(app httpapi.IApi) {
	db, freeDb := services.Db()
	defer freeDb()
	err := db.AutoMigrate(&models.Question{})
	if err != nil {
		log.Fatalf("Could not migrate relations: %v", err)
		os.Exit(1)
	}
	log.Println("migrated questions")

	app.Map("v1/questions", http.MethodGet, wraps.Wrap(hw.Run(usecases.SelectQuestions)))
	app.Map("v1/questions", http.MethodPost, wraps.Wrap(hw.Run(usecases.AddQuestion)))
	app.Map("v1/questions", http.MethodPatch, wraps.Wrap(hw.Run(usecases.UpdateQuestion)))
	app.Map("v1/questions", http.MethodDelete, wraps.Wrap(hw.Run(usecases.RemoveQuestion)))
}
