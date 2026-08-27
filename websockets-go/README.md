
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
