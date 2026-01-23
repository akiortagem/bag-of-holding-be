# Clean Architecture Layout

This project follows a clean-architecture-inspired layout.

## Directories

- `config/`: Configuration loading and environment setup.
- `domain/entities/`: Core business entities.
- `domain/valueobjects/`: Domain value objects.
- `handlers/`: HTTP handlers/controllers.
- `infrastructure/`: External integrations (databases, messaging, etc.).
- `middleware/`: HTTP middleware.
- `repository/`: Data access interfaces and implementations.
- `routes/`: HTTP route registration.
- `usecase/`: Application use cases.
