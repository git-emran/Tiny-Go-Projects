## `http-client`

**What it is:** An example of building an HTTP “client façade” using **Resty** to call another service, with Gorilla Mux routes defined for local endpoints.

**Key ideas:**
- Calling a backing service at `http://localhost:8080` using Resty:
  - `GET /employees` proxies to `GET http://localhost:8080/employees`
  - `POST /employee/add` proxies to `POST http://localhost:8080/employee/add`
- Decoding JSON request bodies into Go structs.

**Main file:** `http-client/http-rest-client.go`

**Run (as-is):**
- The router is set up, but `http.ListenAndServe(...)` isn’t called in `main()`, so it won’t start a server until that’s added.

**When completed, you’d expect:**
- Local server on `localhost:8090` forwarding requests to a “real” service on `localhost:8080`.

