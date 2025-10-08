## Meeting Scheduler Backend (Go)

This repository is a backend API scaffold for a Meeting Scheduler service using:

- Go 1.20+
- Gin (HTTP router)
- GORM (Postgres)
- JWT for authentication
- SMTP for transactional emails (e.g. SendPulse)

The project follows a clean architecture layout (handlers/controllers, services, repositories, models, utils, config, middleware) so you can extend and harden it for production.

Contents
- `cmd/` — application entrypoint
- `config/` — environment loader
- `internal/models` — GORM models (User, Meeting, AuditLog)
- `internal/repository` — repository interfaces and GORM implementation
- `internal/services` — business logic
- `internal/handlers` — Gin HTTP handlers
- `internal/middleware` — JWT and role middleware
- `internal/utils` — password hashing, JWT helpers, email helper
- `internal/scheduler` — reminder scheduler skeleton

Quick start (development)
1. Copy `.env.example` to `.env` and fill your values.
2. Create the Postgres database referenced by `DATABASE_URL`.
3. Fetch Go modules:

```powershell
go mod tidy
```

4. Run the server:

```powershell
go run ./cmd
```

Configuration / Environment variables
Set these (or similar) in your environment or `.env` file. Example values are in `.env.example`.

- DATABASE_URL: Postgres connection string. Example: `postgres://user:pass@localhost:5432/meetingdb`
- JWT_SECRET: long random secret for signing tokens
- SERVER_PORT: address to bind (recommended `:8080`)
- SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASS, SMTP_FROM: SMTP settings for transactional emails

Database migrations
- The scaffold uses GORM AutoMigrate for development. For production, use a proper migration tool (golang-migrate or similar).

API (implemented / scaffold)
Authentication
- POST /auth/login
  - Body: { "email": "user@example.com", "password": "secret" }
  - Response: { "token": "<jwt>" }

Users
- POST /users (admin-only) — create a new user (admin creates accounts)
- GET /users/me — get current user profile

Meetings
- POST /meetings — create a meeting and invite participants
  - Protected: Bearer JWT required
  - Body (example):
    {
      "title": "Weekly Sync",
      "description": "Discuss progress",
      "start_at": "2025-10-10T15:00:00Z",
      "link": "https://meet.example/abcd",
      "participants": [2,3] // user IDs
    }
  - Response: meeting object (with participants)

- PUT /meetings/:id — update / reschedule (organizer or admin)
- POST /meetings/:id/invite — add participants after creation
- POST /meetings/:id/cancel — cancel meeting and notify participants
- GET /meetings — user's meetings dashboard (organizer or participant)
- GET /meetings/:id — meeting details

Audit
- (Optional) GET /audit-logs — admin endpoint to fetch audit records

Email notifications (behavior)
- On creation: email all participants with details
- When participants are added later: email added participants
- On reschedule or cancel: email all participants
- Reminder: 10 minutes before start, send a reminder (includes meeting link)

Authentication & Authorization
- All protected routes require `Authorization: Bearer <token>` header.
- JWT tokens are signed with `JWT_SECRET`.
- Roles: users have `IsAdmin` flag. Admin-only endpoints are protected by middleware.

Example cURL requests
1) Login

```bash
curl -s -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"password"}'
```

2) Create meeting (authenticated)

```bash
TOKEN=<JWT_FROM_LOGIN>
curl -s -X POST http://localhost:8080/api/meetings \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"Team Sync","description":"Weekly","start_at":"2025-10-10T15:00:00Z","link":"https://meet.example/abc","participants":[2,3]}'
```

Seeding example data (dev)
- You can create a small Go program or add a script to insert an admin user and a few test users and meetings. The project contains a seed skeleton — add hashed passwords (bcrypt) and set `IsAdmin` for your admin.

Testing
- Add unit tests under `internal/services` and `internal/repository`.
- For integration tests use a test Postgres instance (Docker) and run the server in test mode.

Notes & Next steps (recommended)
- Implement remaining handlers and repository methods fully (participants, invites, cancel, reschedule).
- Ensure transactional emails are sent only after DB commit. Use GORM transactions and send emails after commit.
- Improve scheduler: run periodic worker (every 1 minute) that enquires meetings starting in ~10 minutes and enqueues reminders.
- Add input validation (`go-playground/validator`) to all request structs.
- Add rate-limiting on invitation endpoints to prevent abuse.
- Migrate from AutoMigrate to versioned migrations for prod.
- Add Swagger/OpenAPI for API documentation.
