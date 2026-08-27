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

