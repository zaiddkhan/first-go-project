package main

import (
	"github.com/zaiddkhan/first-go-project/internal/auth"
	"github.com/zaiddkhan/first-go-project/internal/database"
	"net/http"
)

type authedHandler func(
	http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, err.Error())
		}
		user, err := cfg.DB.GetUserByApiKey(
			r.Context(), apiKey,
		)
		if err != nil {
			respondWithError(w, 403, err.Error())
			return
		}
		handler(w, r, user)
	}
}
