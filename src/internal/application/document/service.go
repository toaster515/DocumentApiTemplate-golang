package document

import (
	"time"

	"github.com/google/uuid"
	"github.com/toaster515/DocumentApiTemplate-golang/internal/domain/document"
)

type Service struct {
	Storage document.StorageProvider
	Repo    document.MetadataRepository
}

func (s *Service) Upload(name string, data []byte) (string, error) {
	id := uuid.NewString()
	url, err := s.Storage.UploadFile(id, data)
	if err != nil {
		return "", err
	}

	doc := document.Document{
		ID:         id,
		FileName:   name,
		Url:        url,
		UploadedAt: time.Now(),
	}

	if err := s.Repo.SaveMetadata(doc); err != nil {
		return "", err
	}

	return id, nil
}

func (s *Service) Download(id string) ([]byte, error) {
	return s.Storage.DownloadFile(id)
}

func (s *Service) GetMetadata(id string) (document.Document, error) {
	return s.Repo.GetMetadata(id)
}
