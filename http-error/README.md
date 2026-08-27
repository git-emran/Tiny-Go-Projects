
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

