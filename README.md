# Voidside

## Docker

From the root folder:

```bash
docker-compose up --build
```

To rebuild just one of the services:

```bash
docker-compose up --build core
```

## Curl testing

It's a wonderful thing, use it.

```bash
curl -X POST http://localhost:8080/chat \
-H "Content-Type: application/json" \
-d '{"user":"test", "message":"are you functional?"}'
```

## Modules

### Core (API)

- Gin framework implementation.
- RESTful API HTTP server.
- ORM database library.
- SQLite driver.
- Keycloak JWKS.

#### *Core Key Packages*

- [Runtime](core/main.go)
- [Post controller](core/controllers/postController.go)
- [User controller](core/controllers/userController.go)
- [Routes](core/routes/router.go)

### Aibot (LLM bridge microservice)

- Gin framework implementation.
- Bridge JSON payload with LLM services.
- Temporary context window memory of messages and users.
- Static Key Auth.

#### *Aibot Key Packages*

- [Runtime](aibot/main.go)
- [Handler](aibot/handlers/chat.go)
- [Bridge](aibot/llm/client.go)
- [Memory](aibot/internal/store.go)

### Ggbot (Discord microservice)

- Gin framework implementation.
- Bridge JSON payload with Discord.
- Discord API endpoint connection.
- Static Key Auth.

#### *Ggbot Key Packages*

- [Runtime](ggbot/main.go)
- [Handler](ggbot/handlers/chat.go)

## Dependencies

- <https://github.com/gin-gonic/gin>
- <https://gorm.io/gorm>
- <https://github.com/glebarez/sqlite>
- <https://github.com/joho/godotenv>
- <https://github.com/MicahParks/keyfunc>
- <https://github.com/golang-jwt/jwt/v5>
