## `serving-templates`

**What it is:** A small server that serves HTML templates and static assets, plus a login form with validation.

**Key ideas:**
- Template rendering with `html/template`
- Static assets with `http.FileServer` mounted under `/static/`
- Form parsing + schema decoding
- Validation using `govalidator` (alpha + required)

**Main files:**
- `serving-templates/first-template.go`
- `serving-templates/templates/first-template.html`
- `serving-templates/templates/login-form.html`
- `serving-templates/static/main.css`

**Run:**
```bash
cd serving-templates
go run .
# server: http://localhost:8080
```

**Flow:**
- `GET /` → shows login form template
- `POST /` → validates fields and responds with a greeting or validation message
- `/static/*` → serves CSS

