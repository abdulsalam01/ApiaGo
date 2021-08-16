package routes

import (
	"github.com/abdulsalam01/go_gorm/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RouteServer(db *gorm.DB, router *mux.Router) *controllers.Server {
	return controllers.InitServer(db, router)
}
