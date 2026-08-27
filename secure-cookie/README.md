# Go Secure Cookie Example

A minimal Go web server demonstrating how to securely encode, decode, set, and read HTTP cookies using the [`github.com/gorilla/securecookie`](https://github.com/gorilla/securecookie) package.

## Features

- **Authenticated & Encrypted Cookies:** Prevents client-side cookie tampering using HMAC authentication and AES encryption.
- **Dynamic Key Generation:** Automatically generates cryptographically strong 64-byte hash and 32-byte block keys at startup.
- **REST Endpoints:** Straightforward HTTP handlers to create and read secure cookie data.

## Prerequisites

- [Go](https://go.dev/dl/) 1.18 or higher installed.

## Getting Started

### 1. Clone & Install Dependencies

Clone the repository and install the Gorilla `securecookie` package:

```bash
git clone [https://github.com/your-username/your-repo-name.git](https://github.com/your-username/your-repo-name.git)
cd your-repo-name
go get [github.com/gorilla/securecookie](https://github.com/gorilla/securecookie)
