## `mysql-book-manager`

**What it is:** A small REST API for managing `Book` records using **Gorilla Mux** + **GORM** + **MySQL**.

**Key ideas:**
- “Package layering” (config → models → controllers → routes)
- GORM auto-migration of the `Book` model on init
- CRUD endpoints with path params (`/book/{bookId}`)

**Main files:**
- `mysql-book-manager/cmd/main/main.go` (server entrypoint)
- `mysql-book-manager/pkg/config/app.go` (MySQL connection)
- `mysql-book-manager/pkg/models/book.go` (GORM model + queries)
- `mysql-book-manager/pkg/controllers/book-controller.go` (HTTP handlers)
- `mysql-book-manager/pkg/routes/bookstore-routes.go` (route registration)
- `mysql-book-manager/pkg/utils/utils.go` (body parsing helper)

**Run:**
```bash
cd mysql-book-manager
go run ./cmd/main
# server: http://localhost:9010
```

**Prerequisites:**
- A MySQL instance reachable with the DSN in `pkg/config/app.go`

**Endpoints:**
- `POST /book/` → create book (JSON body)
- `GET /book/` → list books
- `GET /book/{bookId}` → fetch by ID
- `PUT /book/{bookId}` → update
- `DELETE /book/{bookId}` → delete (note: route path in code is missing a leading `/`)

