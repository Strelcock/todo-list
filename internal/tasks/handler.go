package tasks

import (
	"net/http"
	"todoProject/config"
	"todoProject/internal/user"
	"todoProject/pkg/jsonconv"
	"todoProject/pkg/middleware"
	"todoProject/pkg/req"
)

type TaskHandler struct {
	TaskRepo *TaskRepository
	UserRepo *user.UserRepository
}

type TaskHandlerDeps struct {
	*config.Config
	TaskRepo *TaskRepository
	UserRepo *user.UserRepository
}

func NewTaskHandler(router *http.ServeMux, deps TaskHandlerDeps) {
	handler := &TaskHandler{
		TaskRepo: deps.TaskRepo,
		UserRepo: deps.UserRepo,
	}

	router.HandleFunc("POST /tasks", middleware.IsAuthed(handler.Create(), deps.Config))
	router.HandleFunc("GET /tasks/{id}", middleware.IsAuthed(handler.GetTasks(), deps.Config))
}

func (th *TaskHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CreateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		email := r.Context().Value(middleware.ContextEmailKey)
		foundUser, err := th.UserRepo.FindByEmail(email.(string))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		body.UserID = foundUser.ID
		task := NewTask(body.Name, body.UserID)

		createdTask, err := th.TaskRepo.Create(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonconv.Json(w, createdTask, http.StatusCreated)
	}
}

func (th *TaskHandler) GetTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
