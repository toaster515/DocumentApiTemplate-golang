package document

import "time"

type Document struct {
	ID         string
	FileName   string
	Url        string
	UploadedAt time.Time
}
