## `form-login-logout`

**What it is:** UI templates for a login flow (HTML only right now). The Go server file (`html-form-auth.go`) exists but is empty.

**Key ideas:**
- Template structure for authentication UI:
  - `templates/login-form.html` (basic login form)
  - `templates/home.html` (home screen with logout form)

**Run:**
- Not runnable yet (no server implementation in `html-form-auth.go`).

**Suggested next step (if you want to evolve it):**
- Add handlers for `GET /` (show login), `POST /login` (validate), `POST /logout` (clear session/cookie) and render these templates.

