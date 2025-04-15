package main

import (
	"fmt"
	"log"
	"net/http"
	"todoProject/config"
	"todoProject/internal/auth"
	"todoProject/internal/tasks"
	"todoProject/internal/user"
	"todoProject/pkg/db"
)

func main() {
	router := http.NewServeMux()

	conf := config.LoadConfig()
	database := db.NewDb(conf)

	//repo
	userRepo := user.NewUserRepositiry(database)
	taskRepo := tasks.NewTaskRepository(database)

	//service
	authService := auth.NewAuthService(userRepo)

	//handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	tasks.NewTaskHandler(router, tasks.TaskHandlerDeps{
		Config:   conf,
		TaskRepo: taskRepo,
	})

	user.NewAccHandler(router, user.AccHandlerDeps{
		Config:   conf,
		UserRepo: userRepo,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
