package document_test

import (
	"testing"

	app "github.com/toaster515/DocumentApiTemplate-golang/internal/application/document"
	dom "github.com/toaster515/DocumentApiTemplate-golang/internal/domain/document"
)

type mockStorage struct{}
type mockRepo struct {
	saved dom.Document
}

func (m *mockStorage) UploadFile(id string, data []byte) (string, error) {
	return "https://fake-url.com/" + id, nil
}

func (m *mockStorage) DownloadFile(id string) ([]byte, error) {
	return nil, nil
}

func (m *mockRepo) SaveMetadata(doc dom.Document) error {
	m.saved = doc
	return nil
}
func (m *mockRepo) GetMetadata(id string) (dom.Document, error) {
	return m.saved, nil
}

func TestUpload(t *testing.T) {
	storage := &mockStorage{}
	repo := &mockRepo{}
	service := app.Service{Storage: storage, Repo: repo}

	id, err := service.Upload("test.txt", "testfile", []byte("dummy content"))
	if err != nil {
		t.Fatalf("Upload failed: %v", err)
	}

	if repo.saved.Description != "testfile" {
		t.Errorf("Expected description 'testfile', got %s", repo.saved.Description)
	}

	if repo.saved.ID != id {
		t.Errorf("Expected saved ID %s, got %s", id, repo.saved.ID)
	}
}
