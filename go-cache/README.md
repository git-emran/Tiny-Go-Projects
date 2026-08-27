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
