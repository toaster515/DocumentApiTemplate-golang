# File Storage API Template (Golang)

A containerized file storage API template built in Go, using Domain-Driven Design. This project stores files in AWS S3 and metadata in Postgres.

---
## Tech Stack

- **Golang 1.24**
- **Docker + Docker Compose**
- **PostgreSQL**
- **AWS S3 (via SDK v2)**
- **Swagger (via swaggo)**
### To Do

- **JWT Authentication**
- **Validation (go-playground/validator)**

---
## Architecture

```bash
src/ 
├── cmd/api/ # Entry point 
├── internal/ 
│ ├── api/ # Handlers + Router 
│ ├── application/ # Service layer (use cases) 
│ ├── domain/ # Core models + interfaces 
│ └── infrastructure/ # Postgres + S3 implementations 
├── go.mod / go.sum 
├── Dockerfile 
└── .env
```

---
## Getting Started

### 1. Prerequisites

- [Docker](https://www.docker.com/)
- [AWS credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html)
- AWS S3 bucket created
- Go 1.24 if running outside Docker

---
### 2. Clone + Setup

```bash
git clone https://github.com/your-username/document-api-golang.git
cd document-api-golang/src
```

---
### 3. Environment Variables

Keeping this simple as it's a template

Create a `.env` file in `/src/`:

See `.env.example` for local setup

```bash
AWS_ACCESS_KEY_ID=your-key
AWS_SECRET_ACCESS_KEY=your-secret
AWS_REGION=us-east-1
S3_BUCKET_NAME=your-bucket-name

PG_CONN=postgres://postgres:postgres@postgres:5432/recordsdb?sslmode=disable
```

---
### 4. Run with Docker

```bash
docker-compose up --build
```
The API will be live at: http://localhost:8080

---
## Swagger UI

Once running, access docs at:

`http://localhost:8080/swagger/index.html`

To regenerate:

`cd src swag init --generalInfo cmd/api/main.go --output cmd/api/docs`

You can test swagger locally at `http://localhost:8080/swagger/index.html`

However, I set this up parse a multipart request for file uploads with metadata in one call, so testing it via the swagger is difficult. 

Here is an example python call to test:

```python
import requests
import json

url = "http://localhost:8080/documents"

metadata = {
    "name": "example.txt",
    "description":"test description"
}

files = {
    'file': ('test.txt', open('test.txt', 'rb'), 'text/plain'),
    'metadata': (None, json.dumps(metadata), "application/json")
}

response = requests.post(url, files=files)

print("Status Code:", response.status_code)
print("Response:", response.json)
```