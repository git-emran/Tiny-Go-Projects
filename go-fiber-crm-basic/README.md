
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
