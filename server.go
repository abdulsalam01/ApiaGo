package main

import (
	"fmt"

	"github.com/abdulsalam01/go_gorm/config"
	"github.com/abdulsalam01/go_gorm/routes"
	"github.com/gorilla/mux"
)

func InitServer() {
	config.InitEnvironment()
	port := fmt.Sprintf(":%s", config.InitPort())
	db, err := config.InitConnection()
	routeInit := routes.InitializeRoutes(db, &mux.Router{})

	if err != nil {
		panic(fmt.Sprintf("Connection refused %v", err))
	}

	fmt.Println("It's seems great!")
	fmt.Printf("Listening to port %s", port)

	config.InitSchema(db)
	routes.Run(port, routeInit.Router)
}

func RunServer() {
	InitServer()
}
