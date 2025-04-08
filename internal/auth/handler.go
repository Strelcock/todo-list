package auth

import (
	"net/http"
	"todoProject/config"
	"todoProject/pkg/jsonconv"
	"todoProject/pkg/jwt"
	"todoProject/pkg/req"
)

type AuthHandler struct {
	*config.Config
	*AuthService
}

type AuthhandlerDeps struct {
	*config.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthhandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (ah *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		name, email, password := body.Name, body.Email, body.Password

		_, err = ah.AuthService.Register(name, email, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(ah.Config.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

	}
}

func (ah *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		email, password := body.Email, body.Password
		result, err := ah.AuthService.Login(email, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		jsonconv.Json(w, result, 200)
	}
}
