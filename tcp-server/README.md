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

