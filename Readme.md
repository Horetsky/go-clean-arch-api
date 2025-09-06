# Job Search API (Go)

## Description

This project is an API for a job search platform, implemented in **Go** using the **Clean Architecture** approach.

The API provides:

* User authentication and authorization.
* CRUD operations for **talents**, **recruiters**, **jobs**, and **applications**.
* Email notifications on registration and job status updates.

---

## Tech Stack

* **Go** — main programming language.
* **Clean Architecture** — layered separation (Domain, Use Cases, Infrastructure, Transport).
* **PostgreSQL** — database.
* **JWT** — authorization.
* **SMTP / Mailgun (or other)** — email sending.
* **Makefile** — build and run automation.

---

## Project Structure

```
/ cmd               # Application entry point (main.go)
/ internal          # Core business logic
   / app            # Application setup, dependency injection
   / domain         # Entities (Talent, Recruiter, Job, Application)
   / infrastructure # Database, email
   / migrations     # Database migrations
   / transport      # HTTP handlers, gRPC (if needed)
/ pkg               # Shared utilities (middleware, helpers)
```

---

## Setup & Run

### 1. Clone repository

```bash
git clone https://github.com/Horetsky/go-clean-arch-api.git
cd job-search-api
```

### 2. Using Makefile

Build and run the application with:

```bash
make build   # Build binary into ./bin/api
make run     # Build and run
make start   # Run already built binary
```

---

## Roadmap

* Job search with filtering.
* Multi-language support.
* Integration with external APIs (LinkedIn, Indeed, etc.).
* Real-time notifications via WebSockets.