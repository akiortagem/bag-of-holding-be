# Bag of Holding Backend

Backend of the Bag of Holding app written in Go

# Directory Structure

```text
.
├── cmd/
│   └── main.go
├── internal/
│   ├── core/
│   │   └── some_core_feature/
│   └── shared/
│   └── features/
│       └── foo_feature/
│           ├── domain/
│           ├── interfaces/
│           |   └── http/
│           |       └── handlers/
│           ├── repositories/
│           └── usecases/
├── schemas/
│   └── 001_users.sql
├── go.mod
├── go.sum
└── README.md
```

- `cmd/`: application entrypoint and top-level wiring.
- `core/`: features and utils that are shared app-wide.
- `features/`: discrete features that do not depend on each other.
- `shared/`: a feature that is used by two or more other features
- `schemas/`: database migration SQL.

Each feature has:
- `domain/`: models and interfaces
- `interfaces/`: how this app is accessed from the outside world (HTTP, Discord bot (planned), etc)
- `repositories/`: service implementation that connects to some data source
- `usecases/`: business logic goes here
