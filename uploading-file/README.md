## `uploading-file`

**What it is:** A minimal HTML + Go server that accepts file uploads using multipart form data.

**Key ideas:**
- `r.FormFile("file")` to fetch uploaded file
- Persist uploaded bytes using `io.Copy`
- Render a simple upload HTML form

**Main files:**
- `uploading-file/upload-file.go`
- `uploading-file/templates/upload-file.html`

**Run:**
```bash
go run uploading-file/upload-file.go
# server: http://localhost:8080
```

**Endpoints:**
- `GET /` → upload form
- `POST /upload` → saves file to `tmp/uploadedFile`


