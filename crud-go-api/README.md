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
