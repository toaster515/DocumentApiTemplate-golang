// @title Document API
// @version 1.0
// @description A simple document management API.
// @host localhost:8080
// @BasePath /
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	router "github.com/toaster515/DocumentApiTemplate-golang/internal/api"
	handler "github.com/toaster515/DocumentApiTemplate-golang/internal/api/handler"
	appdoc "github.com/toaster515/DocumentApiTemplate-golang/internal/application/document"
	"github.com/toaster515/DocumentApiTemplate-golang/internal/infrastructure/document"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/toaster515/DocumentApiTemplate-golang/cmd/api/docs"
)

func main() {
	_ = godotenv.Load(".env")

	bucket := os.Getenv("S3_BUCKET_NAME")
	storage, err := document.NewS3Storage(bucket)
	if err != nil {
		log.Fatal("Failed to create S3 repo: ", err)
	}

	dsn := os.Getenv("PG_CONN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to Postgres:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping Postgres:", err)
	}

	repo := &document.PostgresRepo{DB: db}

	service := &appdoc.Service{
		Storage: storage,
		Repo:    repo,
	}
	docHandler := &handler.DocumentHandler{Service: service}
	r := router.NewRouter(docHandler)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Running on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
