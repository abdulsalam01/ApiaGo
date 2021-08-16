package routes

import (
	"fmt"
	"log"
	"net/http"
)

func (s *RouteServer) InitializeRoutes() {
	fmt.Println(s.BaseServer.Router, "adaaa")

	// home routes
	s.BaseServer.Router.HandleFunc("/", s.BaseServer.HelloApi).Methods("GET")

	// user routes
	user := s.BaseServer.Router.PathPrefix("/user").Subrouter()
	user.HandleFunc("/", s.BaseServer.Create).Methods("POST")
	user.HandleFunc("/{id}", s.BaseServer.Update).Methods("PUT")
	user.HandleFunc("/{id}", s.BaseServer.Delete).Methods("DELETE")
	user.HandleFunc("/all", s.BaseServer.GetAll).Methods("GET")
	user.HandleFunc("/all/{id}", s.BaseServer.GetById).Methods("GET")

	// contact routes
	_ = s.BaseServer.Router.PathPrefix("/contact").Subrouter()
}

func (s *RouteServer) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.BaseServer.Router))
}
