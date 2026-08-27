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
