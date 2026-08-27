## `mysql-db-go`

**What it is:** A direct MySQL CRUD demo using the standard library `database/sql` (plus Gorilla Mux for routing).

**Key ideas:**
- `sql.Open` + `db.Query` + `db.Exec`
- Prepared statements for update/delete
- Mix of path params and query params

**Main file:** `mysql-db-go/connect-mysql.go`

**Run:**
```bash
cd mysql-db-go
go run .
# server: http://localhost:8080
```

**Prerequisites:**
- MySQL running and a database/table matching:
  - DSN: `root:@/mydb`
  - table: `employee` (columns expected: `id`, `name`)

**Endpoints:**
- `GET /` → prints current DB name
- `GET /employees` → list records
- `POST /employee/create?name=Alice` → insert
- `PUT /employee/update/{id}?name=Bob` → update
- `DELETE /employee/delete?name=Bob` → delete by name

