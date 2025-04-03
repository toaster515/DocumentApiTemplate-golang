package document

import (
	"database/sql"

	"github.com/toaster515/DocumentApiTemplate-golang/internal/domain/document"
)

type PostgresRepo struct {
	DB *sql.DB
}

func (r *PostgresRepo) SaveMetadata(doc document.Document) error {
	_, err := r.DB.Exec(
		`INSERT INTO file_records ("Id", "FileName", "Url", "UploadedAt") VALUES ($1, $2, $3, $4)`,
		doc.ID, doc.FileName, doc.Url, doc.UploadedAt,
	)
	return err
}

func (r *PostgresRepo) GetMetadata(id string) (document.Document, error) {
	var doc document.Document
	err := r.DB.QueryRow(
		`SELECT "Id", "FileName", "Url", "UploadedAt" FROM file_records WHERE "Id" = $1`,
		id,
	).Scan(&doc.ID, &doc.FileName, &doc.Url, &doc.UploadedAt)
	return doc, err
}
