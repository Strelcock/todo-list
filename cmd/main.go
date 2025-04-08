package main

import (
	"fmt"
	"net/http"
	"todoProject/config"
	"todoProject/internal/auth"
	"todoProject/internal/user"
	"todoProject/pkg/db"
)

func main() {
	router := http.NewServeMux()

	conf := config.LoadConfig()
	database := db.NewDb(conf)

	//repo
	userRepo := user.NewUserRepositiry(database)

	//service
	authService := auth.NewAuthService(userRepo)

	//handler
	auth.NewAuthHandler(router, auth.AuthhandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
