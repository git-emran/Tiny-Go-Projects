# Tiny Go Projects

A collection of small, focused Go projects—each folder is its own “tiny lab” exploring one idea (HTTP routing, cookies, sessions, DB access, WebSockets, templates, etc.).

Most folders are standalone Go modules (they include a `go.mod`). A few are single-file examples without modules.

## Quick Start

```bash
# pick a folder and run it
cd <project-folder>
go run .
```

If a folder doesn’t have a `main.go` (or doesn’t compile as-is), run the specific file called out in that section.

## Projects (by folder)

- [`crud-go-api`](#crud-go-api)
- [`form-login-logout`](#form-login-logout)
- [`go-cache`](#go-cache)
- [`go-fiber-crm-basic`](#go-fiber-crm-basic)
- [`http-client`](#http-client)
- [`http-error`](#http-error)
- [`http-image-generator-api`](#http-image-generator-api)
- [`http-request-gorilla-mux`](#http-request-gorilla-mux)
- [`http-rest`](#http-rest)
- [`http-server`](#http-server)
- [`http-session`](#http-session)
- [`mysql-book-manager`](#mysql-book-manager)
- [`mysql-db-go`](#mysql-db-go)
- [`secure-cookie`](#secure-cookie)
- [`serving-templates`](#serving-templates)
- [`tcp-server`](#tcp-server)
- [`uploading-file`](#uploading-file)
- [`websockets-go`](#websockets-go)
- [`cli-markdown-preview`](#cli-markdown-preview)

---

## `crud-go-api`

**What it is:** A tiny REST-ish CRUD API for a `Movie` resource using **Gorilla Mux** and in-memory storage.

**Key ideas:**
- Routing with path params (`/movies/{id}`) and HTTP methods.
- Encoding/decoding JSON with `encoding/json`.
- In-memory slice as a “database” (great for learning flows without persistence complexity).

**Main file:** `crud-go-api/main.go`

**Endpoints:**
- `GET /movies` → list all movies
- `GET /movies/{id}` → fetch one movie
- `POST /movies/` → create a movie (reads JSON body)
- `PUT /movies/{id}` → update a movie (reads JSON body)
- `DELETE /movies/{id}` → delete a movie

**Run:**
```bash
cd crud-go-api
go run .
# server: http://localhost:8000
```


---

## `form-login-logout`

**What it is:** UI templates for a login flow (HTML only right now). The Go server file (`html-form-auth.go`) exists but is empty.

**Key ideas:**
- Template structure for authentication UI:
  - `templates/login-form.html` (basic login form)
  - `templates/home.html` (home screen with logout form)

**Run:**
- Not runnable yet (no server implementation in `html-form-auth.go`).

**Suggested next step (if you want to evolve it):**
- Add handlers for `GET /` (show login), `POST /login` (validate), `POST /logout` (clear session/cookie) and render these templates.

---

## `go-cache`

**What it is:** A tiny HTTP server that demonstrates **in-memory caching** using `github.com/patrickmn/go-cache`.

**Key ideas:**
- Application-level cache with TTL:
  - cache is created with a default expiration (5 minutes)
  - a sample key `foo=bar` is stored in `init()`
- Simple “read-through” behavior: on request, attempt cache read and return a response.

**Main file:** `go-cache/go-cache.go`

**Run:**
```bash
cd go-cache
go run .
# server: http://localhost:8080
```

**Try:**
- `GET /` → returns `Hello bar` (as long as the cached key exists and hasn’t expired)

---

## `go-fiber-crm-basic`

**What it is:** A minimal “CRM-style” API skeleton using **Fiber** + **GORM** with **SQLite** (`leads.db`).

**Key ideas:**
- Fiber routing under `/api/v1/...`
- SQLite persistence via GORM
- Auto-migration for the `Lead` model

**Main files:**
- `go-fiber-crm-basic/main.go` (server + routes + DB init)
- `go-fiber-crm-basic/lead/lead.go` (model + handlers)
- `go-fiber-crm-basic/database/database.go` (global DB connection)

**Routes (intended):**
- `GET /api/v1/lead` → list leads
- `GET /api/v1/lead/:id` → fetch a lead
- `POST /api/v1/lead` → create a lead
- `DELETE /api/v1/lead/:id` → delete a lead

**Run:**
```bash
cd go-fiber-crm-basic
go run .
# server: http://localhost:3000
```

---

## `http-client`

**What it is:** An example of building an HTTP “client façade” using **Resty** to call another service, with Gorilla Mux routes defined for local endpoints.

**Key ideas:**
- Calling a backing service at `http://localhost:8080` using Resty:
  - `GET /employees` proxies to `GET http://localhost:8080/employees`
  - `POST /employee/add` proxies to `POST http://localhost:8080/employee/add`
- Decoding JSON request bodies into Go structs.

**Main file:** `http-client/http-rest-client.go`

**Run (as-is):**
- The router is set up, but `http.ListenAndServe(...)` isn’t called in `main()`, so it won’t start a server until that’s added.

**When completed, you’d expect:**
- Local server on `localhost:8090` forwarding requests to a “real” service on `localhost:8080`.

---

## `http-error`

**What it is:** A pattern for **centralized HTTP error handling** by wrapping handlers that can return errors.

**Key ideas:**
- Define `type WrapperHandler func(http.ResponseWriter, *http.Request) error`
- Implement `ServeHTTP` on that function type to handle errors consistently.
- Use a custom error type (`NameNotFoundError`) to map business errors to HTTP status codes.

**Main file:** `http-error/http-error.go`

**Run:**
```bash
cd http-error
go run .
# server: http://localhost:8080
```

**Try:**
- `GET /employee/get/foo` → `Hello foo`
- `GET /employee/get/anything-else` → error response

---

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

## `http-request-gorilla-mux`

**What it is:** A Gorilla Mux server that demonstrates:
- Basic routing
- Request logging middleware (`gorilla/handlers`)
- Path variables

**Main file:** `http-request-gorilla-mux/gorilla-mux-routing.go`

**Run:**
```bash
cd http-request-gorilla-mux
go run .
# server: http://localhost:8080
```

**Endpoints:**
- `GET /` → `Hello Go` (logged to stdout)
- `POST /post` → `I am posting` (logged to `server.log`)
- `GET /hello/{name}` → `Hi{name}` (combined logging to `server.log`)

---

## `http-rest`

**What it is:** A compact REST API for an `Employee` resource, including **versioned endpoints** (`/v1`, `/v2`).

**Key ideas:**
- Route table pattern (`Routes` slice) + helper (`AddRoutes`)
- Basic CRUD-ish flows (in-memory slice)
- Versioning via subrouters:
  - `/v1/employees`
  - `/v2/employees`

**Main file:** `http-rest/http-rest-api.go`

**Run:**
```bash
cd http-rest
go run .
# server: http://localhost:8080
```

**Endpoints:**
- `GET /employees` → base dataset
- `GET /employee/{id}` → fetch by ID
- `POST /employee/add` → add
- `PUT /employee/update` → update (upsert behavior)
- `DELETE /employee/delete` → delete (expects JSON body)
- `GET /v1/employees` → v1 dataset
- `GET /v2/employees` → v2 dataset

---

## `http-server`

**What it is:** A “hello world” HTTP server protected by **Basic Auth**, using constant-time comparisons.

**Key ideas:**
- `r.BasicAuth()` extraction
- `crypto/subtle` for constant-time string comparison
- Wrapper auth middleware returning `http.HandlerFunc`

**Main file:** `http-server/http-server.go`

**Run:**
```bash
cd http-server
go run .
# server: http://localhost:8080
```

**Try:**
- `GET /` → requires basic auth (`admin` / `admin`)
- `GET /login` and `GET /logout` → simple text pages

---

## `http-session`

**What it is:** A minimal authentication session example using **Redis-backed sessions** via `redistore` (Gorilla sessions).

**Key ideas:**
- Redis-backed session store:
  - connects to `localhost:6379`
  - session key: `"session-name"`
  - session value: `authenticated=true/false`
- Route protection: `/home` checks the session value before allowing access.

**Main file:** `http-session/http-session.go`

**Run:**
```bash
cd http-session
go run .
# server: http://localhost:8080
```

**Prerequisites:**
- Redis running on `localhost:6379`

**Endpoints:**
- `GET /login` → sets `authenticated=true`
- `GET /home` → requires `authenticated=true`
- `GET /logout` → sets `authenticated=false`

---

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


---

## `mysql-db-go`

**What it is:** A direct MySQL CRUD demo using the standard library `database/sql` (plus Gorilla Mux for routing).

**Key ideas:**
- `sql.Open` + `db.Query` + `db.Exec`
- Prepared statements for update/delete
- Mix of path params and query params

**Main file:** `mysql-db-go/connect-mysql.go`

**Run:**
```bash
cd mysql-db-go
go run .
# server: http://localhost:8080
```

**Prerequisites:**
- MySQL running and a database/table matching:
  - DSN: `root:@/mydb`
  - table: `employee` (columns expected: `id`, `name`)

**Endpoints:**
- `GET /` → prints current DB name
- `GET /employees` → list records
- `POST /employee/create?name=Alice` → insert
- `PUT /employee/update/{id}?name=Bob` → update
- `DELETE /employee/delete?name=Bob` → delete by name

---

## `secure-cookie`

**What it is:** A demonstration of signed/encoded cookies using `gorilla/securecookie`.

**Key ideas:**
- Create a secure cookie with generated hash + block keys.
- Encode a map payload into a cookie value.
- Decode and read that payload back on later requests.

**Main file:** `secure-cookie/http-cookie.go`

**Run:**
```bash
cd secure-cookie
go run .
# server: http://localhost:8080
```

**Endpoints:**
- `GET /create` → sets cookie `first-cookie`
- `GET /read` → reads and decodes `first-cookie`

---

## `serving-templates`

**What it is:** A small server that serves HTML templates and static assets, plus a login form with validation.

**Key ideas:**
- Template rendering with `html/template`
- Static assets with `http.FileServer` mounted under `/static/`
- Form parsing + schema decoding
- Validation using `govalidator` (alpha + required)

**Main files:**
- `serving-templates/first-template.go`
- `serving-templates/templates/first-template.html`
- `serving-templates/templates/login-form.html`
- `serving-templates/static/main.css`

**Run:**
```bash
cd serving-templates
go run .
# server: http://localhost:8080
```

**Flow:**
- `GET /` → shows login form template
- `POST /` → validates fields and responds with a greeting or validation message
- `/static/*` → serves CSS

---

## `tcp-server`

**What it is:** A minimal TCP echo server using `net`—a great contrast to the HTTP examples.

**Key ideas:**
- `net.Listen("tcp", ...)`
- Accept loop + per-connection goroutine
- Read until newline and echo back

**Main file:** `tcp-server/tcp-server.go`

**Run:**
```bash
go run tcp-server/tcp-server.go
# server: tcp://localhost:8080
```

**Try:**
```bash
nc localhost 8080
hello there
```

---

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


---

## `websockets-go`

**What it is:** A broadcast-style WebSocket server using `gorilla/websocket` where messages from one client are fanned out to all connected clients.

**Key ideas:**
- WebSocket upgrade + connection tracking in a `clients` map
- Broadcast channel to decouple read/write loops
- Server-side fan-out to all active clients

**Main file:** `websockets-go/websocket-server.go`

**Run:**
```bash
cd websockets-go
go run .
# server: http://localhost:8080
```

**Endpoints:**
- `GET /` → serves `index.html` (client page)
- `GET /echo` → WebSocket endpoint (JSON messages like `{ "message": "hi" }`)


## `cli-markdown-preview`

**What it is:** A Minimalistic command line based Markdown preview tool. Using `bluemonday` package from golang which is an html sanitizer, `blackfriday/v2` which is a markdown processor and provides full UTF-8 support. 

### Features
     1	markdown → html conversion with blackfriday
     2	html sanitization with bluemonday
     3	custom html templates
     4	temporary html file generation
     5	automatic browser preview
     6	cross-platform support: macos, linux, and windows

### Usage
go run . -file README.md

Skip opening the browser:
`
go run . -file README.md -s
`

Use a custom HTML template:
`
go run . -file README.md -t template.html
`

The generated HTML file path is printed to stdout.

Build
`
go build -o mdp .
`

Then:
`
./mdp -file README.md
`
