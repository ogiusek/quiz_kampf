package users

import (
	"lib/app/users/models"
	"lib/app/users/usecases"
	"lib/common/httpapi"
	"lib/common/wraps"
	"lib/services"
	"log"
	"net/http"
	"os"

	"github.com/ogiusek/hw/src/hw"
)

func AddUsers(api httpapi.IApi) {
	db, freeDb := services.Db()
	defer freeDb()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Could not migrate relations: %v", err)
		os.Exit(1)
	}
	log.Println("Migrated user relations")

	api.Map("v1/users/login", http.MethodPost, wraps.Wrap(hw.Run(usecases.Login)))
	api.Map("v1/users/register", http.MethodPost, wraps.Wrap(hw.Run(usecases.Register)))
	api.Map("v1/users/refresh", http.MethodPost, wraps.Wrap(hw.Run(usecases.Refresh)))
	api.Map("v1/account/rename", http.MethodPost, wraps.Wrap(hw.Run(usecases.Rename)))
	api.Map("v1/account/profile", http.MethodGet, wraps.Wrap(hw.Run(usecases.Profile)))
}
