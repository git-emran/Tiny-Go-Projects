# Tiny Go Projects

A collection of small, focused Go projects—each folder is its own “tiny lab” exploring one idea (HTTP routing, cookies, sessions, DB access, WebSockets, templates, etc.).

Most folders are standalone Go modules (they include a `go.mod`). A few are single-file examples without modules.

## Quick Start

```bash
# pick a folder and run it
cd <project-folder>
go run .
```

If a folder doesn’t have a `main.go` (or doesn’t compile as-is), run the specific file called out in that section.

## Projects (by folder)

- [`crud-go-api`](#crud-go-api)
A Basic CRUD Application application api.

- [`form-login-logout`](#form-login-logout)
A UI template for a login FLow.

- [`go-cache`](#go-cache)
A tiny REST CRUD api for a Movie resource using Gorilla Mux.

- [`go-fiber-crm-basic`](#go-fiber-crm-basic)
CRM style API skeleton with Fiber+GORM+SQLite

- [`http-client`](#http-client)
An example of building an HTTP “client façade” using **Resty** to call another service, with Gorilla Mux routes defined for local endpoints.

- [`http-error`](#http-error)
Centralized HTTP error handling by wrapping handlers that can return errors.

- [`http-image-generator-api`](#http-image-generator-api)
A small gin API showcasing route groups, JSON binding/validation, and file upload handling. 

- [`http-request-gorilla-mux`](#http-request-gorilla-mux)
A gorilla mux server that does Basic routing, Request logging middleware, Path variables.

- [`http-rest`](#http-rest)
Compact rest API for an Employee resource, including versioned endpoints (`/v1`, `/v2`)

- [`http-server`](#http-server)
A "Hello World" HTTP server protected by **Basic Auth** with constant time comparision

- [`http-session`](#http-session)
A minimal authentication session example using **Redis-backed sessions** via `redistore` (Gorilla sessions).

- [`mysql-book-manager`](#mysql-book-manager)
A small REST API for managing `Book` records using **Gorilla Mux** + **GORM** + **MySQL**.

- [`mysql-db-go`](#mysql-db-go)
A direct MySQL CRUD demo using the standard library `database/sql` (plus Gorilla Mux for routing).

- [`secure-cookie`](#secure-cookie)
A minimal webserver demonstrating how to securely encode, decode, set and read HTTP cookies. 

- [`serving-templates`](#serving-templates)
A small server that serves HTML templates and static assets, plus a login form with validation.

- [`tcp-server`](#tcp-server)
A minimal TCP echo server using `net`—a great contrast to the HTTP examples.

- [`uploading-file`](#uploading-file)
A minimal HTML + Go server that accepts file uploads using multipart form data.

- [`websockets-go`](#websockets-go)

A broadcast-style WebSocket server using `gorilla/websocket` where messages from one client are fanned out to all connected clients.

- [`cli-markdown-preview`](#cli-markdown-preview)
A Minimalistic command line based Markdown preview tool. Using `bluemonday` package from golang which is an html sanitizer, `blackfriday/v2` which is a markdown processor and provides full UTF-8 support. 

- [`cli-todo`](#cli-todo)
A simple, fast, and lightweight Command-Line Interface (CLI) task manager written in Go. It persists tasks to a local JSON file, supports interactive input, and accepts environment variable configuration.





