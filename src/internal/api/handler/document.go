package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/toaster515/DocumentApiTemplate-golang/internal/application/document"
)

type DocumentHandler struct {
	Service *document.Service
}

// Upload godoc
// @Summary Upload a document
// @Accept  multipart/form-data
// @Produce json
// @Param file formData file true "Document file"
// @Success 201 {object} map[string]string
// @Router /documents [post]
func (h *DocumentHandler) Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File upload error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// metadata JSON
	metaJson := r.FormValue("metadata")
	if metaJson == "" {
		http.Error(w, "Missing metadata", http.StatusBadRequest)
		return
	}

	// metadata struct
	var meta struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.Unmarshal([]byte(metaJson), &meta); err != nil {
		http.Error(w, "Invalid metadata", http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	id, err := h.Service.Upload(meta.Name, meta.Description, data)
	if err != nil {
		http.Error(w, "Upload failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

// Download godoc
// @Summary Download a document
// @Description Streams a document file by ID
// @Tags documents
// @Param id path string true "Document ID"
// @Produce application/octet-stream
// @Success 200 {file} file
// @Failure 404 {object} map[string]string
// @Router /documents/{id} [get]
func (h *DocumentHandler) Download(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	doc, err := h.Service.GetMetadata(id)
	if err != nil {
		http.Error(w, "Document not found", http.StatusNotFound)
		return
	}

	data, err := h.Service.Download(id)
	if err != nil {
		http.Error(w, "Failed to download file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", doc.FileName))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// GetMetadata godoc
// @Summary Get document metadata
// @Description Retrieves metadata for a document
// @Tags documents
// @Param id path string true "Document ID"
// @Produce json
// @Success 200 {object} document.Document
// @Failure 404 {object} map[string]string
// @Router /documents/{id}/meta [get]
func (h *DocumentHandler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	doc, err := h.Service.GetMetadata(id)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(doc)
}
