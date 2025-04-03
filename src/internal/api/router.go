package api

import (
	"github.com/gorilla/mux"
	"github.com/toaster515/DocumentApiTemplate-golang/internal/api/handler"
)

func NewRouter(h *handler.DocumentHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/documents", h.Upload).Methods("POST")
	r.HandleFunc("/documents/{id}", h.Download).Methods("GET")
	r.HandleFunc("/documents/{id}/meta", h.GetMetadata).Methods("GET")

	return r
}
