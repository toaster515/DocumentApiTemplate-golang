package document

type StorageProvider interface {
	UploadFile(id string, data []byte) (string, error)
	DownloadFile(id string) ([]byte, error)
}

type MetadataRepository interface {
	SaveMetadata(doc Document) error
	GetMetadata(id string) (Document, error)
}
