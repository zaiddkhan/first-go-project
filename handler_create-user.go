package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/zaiddkhan/first-go-project/internal/database"
	"net/http"
	"time"
)

func (apiConfig *apiConfig) handlerCreateUser(
	w http.ResponseWriter,
	r *http.Request) {
	type paramter struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramter{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	user, err := apiConfig.DB.CreateUser(r.Context(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			Name:      params.Name,
		})
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}
	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiConfig *apiConfig) handlerGetUser(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	respondWithJSON(w, 200, databaseUserToUser(user))
}
