package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zaiddkhan/first-go-project/internal/database"
	"log"
	"net/http"
	"os"
	"time"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")
	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}
	portString := os.Getenv("PORT")

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
		AllowCredentials: false,
	}))
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("exit")
	}
	conn, error := sql.Open("postgres", dbUrl)
	queries := database.New(conn)

	apiConfig := apiConfig{
		DB: queries,
	}

	go startScraping(
		queries,
		10,
		time.Minute,
	)
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handleReadiness)
	v1Router.Get("/err", handleErr)
	v1Router.Post("/user", apiConfig.handlerCreateUser)
	v1Router.Get("/user", apiConfig.middlewareAuth(apiConfig.handlerGetUser))
	router.Mount("/v1", v1Router)

	log.Printf("Server starting on port %v", portString)

	if error != nil {
		log.Fatal(error)
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	println(portString)
}
