package controllers

import (
	"net/http"

	responses "github.com/abdulsalam01/go_gorm/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Db     *gorm.DB
	Router *mux.Router
	Model  *interface{}
}

func (server *Server) HelloApi(w http.ResponseWriter, r *http.Request) {
	responses.JsonResponse(w, http.StatusOK, "Welcome to gormAPI!")
}

func InitServer(db *gorm.DB, router *mux.Router) *Server {
	return &Server{Db: db, Router: router}
}
