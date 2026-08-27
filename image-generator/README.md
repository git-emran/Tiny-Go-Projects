## `http-image-generator-api`

**What it is:** A small **Gin** API showcasing route groups, JSON binding/validation, and file upload handling.

**Key ideas:**
- `r.Group("/user")` to group routes
- `BindJSON` with validation tags (example: `email` required)
- Multipart upload with `FormFile` and `SaveUploadedFile`

**Main file:** `http-image-generator-api/main.go`

**Run:**
```bash
cd http-image-generator-api
go run .
# server: http://localhost:8080 (Gin default)
```

**Endpoints:**
- `GET /user/hello/:name` → text greeting
- `POST /user/post` → validates and echoes JSON payload
- `POST /user/upload` → uploads a file and saves it to `/tmp/tempfile`

---
