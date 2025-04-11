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

type AuthHandlerDeps struct {
	*config.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
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

		data, err := jwt.NewJWT(ah.Config.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token := RegisterResponse{
			JWT: data,
		}

		jsonconv.Json(w, token, http.StatusCreated)
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
		_, err = ah.AuthService.Login(email, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		data, err := jwt.NewJWT(ah.Config.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token := LoginResponse{
			JWT: data,
		}
		jsonconv.Json(w, token, http.StatusOK)

	}
}

// func (ah *AuthHandler) Delete() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		body, err := req.HandleBody[DeleteRequest](&w, r)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		email, password := body.Email, body.Password
// 		err = ah.AuthService.Delete(email, password)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		jsonconv.Json(w, "Account deleted", http.StatusOK)
// 	}
// }
