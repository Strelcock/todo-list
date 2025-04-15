package user

import (
	"net/http"
	"todoProject/config"
	"todoProject/internal/errs"
	"todoProject/pkg/middleware"
	"todoProject/pkg/req"
)

type AccHandler struct {
	UserRepo *UserRepository
}

type AccHandlerDeps struct {
	*config.Config
	UserRepo *UserRepository
}

func NewAccHandler(router *http.ServeMux, deps AccHandlerDeps) {
	handler := &AccHandler{
		UserRepo: deps.UserRepo,
	}

	router.HandleFunc("DELETE /user/delete", middleware.IsAuthed(handler.Delete(), deps.Config))
	router.HandleFunc("PATCH /user/name", middleware.IsAuthed(handler.Change(), deps.Config))
	router.HandleFunc("PATCH /user/password", middleware.IsAuthed(handler.Delete(), deps.Config))
}

func (uh *AccHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[DeleteRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		email := r.Context().Value(middleware.ContextEmailKey)

		if email.(string) != body.Email {
			http.Error(w, "you cannot delete another user", http.StatusBadRequest)
			return
		}

		foundUser, _ := uh.UserRepo.FindByEmail(body.Email)

		if foundUser.Model == nil {
			http.Error(w, errs.UserNotExists, http.StatusNotFound)
		}

		err = uh.UserRepo.Delete(foundUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}
}

func (uh *AccHandler) Change() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ChangeRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		email := r.Context().Value(middleware.ContextEmailKey)

		if email.(string) != body.Email {
			http.Error(w, "wrong credentials", http.StatusBadRequest)
			return
		}

		foundUser, _ := uh.UserRepo.FindByEmail(body.Email)
		if foundUser.Model == nil {
			http.Error(w, errs.UserNotExists, http.StatusNotFound)
		}

		err = uh.UserRepo.ChangeName(foundUser, body.NewName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
