package tasks

import (
	"net/http"
	"strings"
	"todoProject/config"
	"todoProject/pkg/jsonconv"
	"todoProject/pkg/middleware"
	"todoProject/pkg/req"
)

type TaskHandler struct {
	TaskRepo *TaskRepository
}

type TaskHandlerDeps struct {
	*config.Config
	TaskRepo *TaskRepository
}

func NewTaskHandler(router *http.ServeMux, deps TaskHandlerDeps) {
	handler := &TaskHandler{
		TaskRepo: deps.TaskRepo,
	}

	router.HandleFunc("POST /tasks/add", middleware.IsAuthed(handler.Create(), deps.Config))
	router.HandleFunc("PATCH /tasks/name/{name}", middleware.IsAuthed(handler.ChangeTaskName(), deps.Config))
	router.HandleFunc("PATCH /tasks/mark/{name}", middleware.IsAuthed(handler.Mark(), deps.Config))
	router.HandleFunc("GET /tasks/all", middleware.IsAuthed(handler.GetAllTasks(), deps.Config))
	router.HandleFunc("GET /tasks/done", middleware.IsAuthed(handler.GetDoneTasks(), deps.Config))
	router.HandleFunc("GET /tasks/{name}", middleware.IsAuthed(handler.GetTaskByName(), deps.Config))

}

func (th *TaskHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CreateRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		uid := r.Context().Value(middleware.ContextUidKey)
		foundTask, _ := th.TaskRepo.GetByName(uid.(uint), body.Name)

		if foundTask != nil {
			http.Error(w, "task already exists", http.StatusBadRequest)
			return
		}

		task := NewTask(body.Name, uid.(uint))

		createdTask, err := th.TaskRepo.Create(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonconv.Json(w, createdTask, http.StatusCreated)
	}
}

func (th *TaskHandler) GetAllTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.Context().Value(middleware.ContextUidKey)
		tasks, err := th.TaskRepo.GetAll(uid.(uint))

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		jsonconv.Json(w, tasks, http.StatusOK)
	}
}

func (th *TaskHandler) GetDoneTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.Context().Value(middleware.ContextUidKey)
		task, err := th.TaskRepo.GetDone(uid.(uint))

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		jsonconv.Json(w, task, http.StatusOK)
	}
}

func (th *TaskHandler) GetTaskByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := r.Context().Value(middleware.ContextUidKey)
		path := r.PathValue("name")
		task, err := th.TaskRepo.GetByName(uid.(uint), path)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		jsonconv.Json(w, task, http.StatusOK)
	}
}

func (th *TaskHandler) Mark() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[MarkRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		pValue := r.PathValue("name")
		taskName := strings.Join(strings.Split(pValue, "+"), " ")
		uid := r.Context().Value(middleware.ContextUidKey)

		foundTask, _ := th.TaskRepo.GetByName(uid.(uint), taskName)

		err = th.TaskRepo.Mark(foundTask, body.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		jsonconv.Json(w, "OK", http.StatusOK)
	}
}

func (th *TaskHandler) ChangeTaskName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ChangeNameRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		uid := r.Context().Value(middleware.ContextUidKey)
		path := r.PathValue("name")
		taskName := strings.Join(strings.Split(path, "+"), " ")
		foundTask, _ := th.TaskRepo.GetByName(uid.(uint), taskName)
		if foundTask == nil {
			http.Error(w, "task does not exist", http.StatusNotFound)
			return
		}

		err = th.TaskRepo.ChangeName(foundTask, body.NewName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonconv.Json(w, "Name changed", http.StatusOK)
	}
}
