package main

import (
	"fmt"
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
		UserRepo: userRepo,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
