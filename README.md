# Voidside

## Modules

### Core

- Gin framework implementation.
- RESTful API HTTP server.
- ORM database library.
- SQLite driver.
- Keycloak JWKS.

#### **Key packages**

- [API entry](main.go)
- [Post controller](core/controllers/postController.go)
- [User controller](core/controllers/userController.go)
- [Post models](core/models/post.go)
- [User models](core/models/user.go)
- [Routes](core/routes/router.go)

## Dependencies

- <https://github.com/gin-gonic/gin>
- <https://gorm.io/gorm>
- <https://github.com/glebarez/sqlite>
- <https://github.com/joho/godotenv>
- <https://github.com/MicahParks/keyfunc>
- <https://github.com/golang-jwt/jwt/v5>
