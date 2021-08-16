package routes

import (
	"log"
	"net/http"

	"github.com/abdulsalam01/go_gorm/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitializeRoutes(db *gorm.DB, router *mux.Router) *controllers.Server {
	route := RouteServer(db, router)
	// home routes
	route.Router.HandleFunc("/", route.HelloApi).Methods("GET")

	// user routes
	user := route.Router.PathPrefix("/user").Subrouter()
	user.HandleFunc("/", route.Create).Methods("POST")
	user.HandleFunc("/{id}", route.Update).Methods("PUT")
	user.HandleFunc("/{id}", route.Delete).Methods("DELETE")
	user.HandleFunc("/all", route.GetAll).Methods("GET")
	user.HandleFunc("/all/{id}", route.GetById).Methods("GET")

	// contact routes
	_ = route.Router.PathPrefix("/contact").Subrouter()

	return route
}

func Run(addr string, route *mux.Router) {
	log.Fatal(http.ListenAndServe(addr, route))
}
