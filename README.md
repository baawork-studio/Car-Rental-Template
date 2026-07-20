# Car Rental Platform

Monorepo starter for a LINE OA car-rental platform.

## Applications

- `apps/customer` — mobile-first LINE LIFF customer experience.
- `apps/admin` — responsive administration dashboard.
- `apps/api` — Gin API with PostgreSQL and Redis integrations.

## Start with Docker

```sh
docker compose up --build
```

Customer: http://localhost:5173 · Admin: http://localhost:5174 · API health: http://localhost:8080/health

Copy `.env.example` to `.env` before implementing LINE Messaging, Rich Menu, or LIFF features.
