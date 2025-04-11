package middleware

import (
	"context"
	"net/http"
	"strings"
	"todoProject/config"
	"todoProject/pkg/jwt"
)

type key string

const (
	AutorizationHeader     = "Authorization"
	BearerPrefix           = "Bearer "
	ContextEmailKey    key = "contextEmailKey"
)

func unauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *config.Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(AutorizationHeader)

		if !strings.HasPrefix(authHeader, BearerPrefix) {
			unauthed(w)
			return
		}

		token := strings.TrimPrefix(authHeader, BearerPrefix)

		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			unauthed(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
