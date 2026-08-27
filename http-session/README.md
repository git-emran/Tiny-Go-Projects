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

