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
