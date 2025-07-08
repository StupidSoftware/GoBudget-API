# GoBudget - Personal Budget Management API

## Overview

GoBudget is a REST API developed in Go that allows users to manage their personal budget through financial categories and transactions control. The application uses the Gin framework for API creation and implements authentication via Bearer Token.

## Architecture

The application follows a layered architecture with clear separation of responsibilities:

- **Controller**: Layer responsible for receiving HTTP requests and returning responses
- **Service**: Business logic layer
- **Repository**: Data access layer (Data Access Layer)
- **Model**: Data models definition layer
- **Middleware**: Intermediate processing layer (authentication, logging, etc.)
- **Router**: API routes configuration and definition

### Request Flow

```
HTTP Request → Router → Middleware → Controller → Service → Repository → Database
                                        ↓
HTTP Response ← Router ← Middleware ← Controller ← Service ← Repository ← Database
```

## Main Features

### 🔐 Authentication and Users
- New user creation
- Login system with JWT token generation
- User deletion

### 📂 Category Management
- Custom category creation
- Category listing by user
- Access control through authentication

### 💰 Transaction Control
- Financial transaction recording
- Transaction listing by user
- Transaction association with categories

## API Endpoints

### Users

#### `POST /users`
**Description**: Creates a new user in the system

**Headers**:
```
Content-Type: application/json
```

**Body**:
```json
{
  "name": "string",
  "email": "string",
  "password": "string"
}
```

**Responses**:
- `201`: User created successfully
- `400`: Invalid data in request body
- `409`: User already exists
- `500`: Internal server error

---

#### `GET /users/login`
**Description**: Performs login and returns authentication token

**Headers**:
```
Content-Type: application/json
```

**Body**:
```json
{
  "email": "string",
  "password": "string"
}
```

**Responses**:
- `200`: Login successful
  ```json
  {
    "token": "jwt_token_string"
  }
  ```
- `400`: Invalid data
- `404`: User not found
- `500`: Internal server error

### Categories

#### `POST /categories`
**Description**: Creates a new category for the authenticated user

**Headers**:
```
Content-Type: application/json
Authorization: Bearer {token}
```

**Body**:
```json
{
  "name": "string",
  "description": "string",
  "type": "string"
}
```

**Responses**:
- `201`: Category created successfully
- `400`: Invalid data
- `409`: Category already exists
- `500`: Internal server error

---

#### `GET /categories`
**Description**: Lists all categories for the authenticated user

**Headers**:
```
Authorization: Bearer {token}
```

**Responses**:
- `200`: Category list
  ```json
  [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "type": "string",
      "user_id": "string"
    }
  ]
  ```
- `400`: Request error
- `500`: Internal server error

### Transactions

#### `POST /transactions`
**Description**: Creates a new transaction for the authenticated user

**Headers**:
```
Content-Type: application/json
Authorization: Bearer {token}
```

**Body**:
```json
{
  "amount": 0.0,
  "description": "string",
  "category_id": "string",
  "type": "string",
  "date": "2024-01-01T00:00:00Z"
}
```

**Responses**:
- `201`: Transaction created successfully
- `400`: Invalid data
- `500`: Internal server error

---

#### `GET /transactions`
**Description**: Lists all transactions for the authenticated user

**Headers**:
```
Authorization: Bearer {token}
```

**Responses**:
- `200`: Transaction list
  ```json
  [
    {
      "id": "string",
      "amount": 0.0,
      "description": "string",
      "category_id": "string",
      "type": "string",
      "date": "2024-01-01T00:00:00Z",
      "user_id": "string"
    }
  ]
  ```
- `400`: Request error
- `500`: Internal server error

## Authentication

The API uses Bearer Token (JWT) authentication. After logging in through the `/users/login` endpoint, the returned token must be included in the `Authorization` header of all requests that require authentication.

**Header Format**:
```
Authorization: Bearer {your_jwt_token}
```

## Project Structure

```
GOBUDGET-API-MAIN/
├── assets/                     # Static resources
├── cmd/
│   └── server/
│       └── main.go            # Application entry point
├── config/                    # Application configuration
│   ├── config.go             # Main configuration
│   ├── env.go                # Environment variables
│   ├── jwt.go                # JWT configuration
│   ├── logger.go             # Logger configuration
│   ├── postgres.go           # PostgreSQL configuration
│   └── validator.go          # Custom validators
├── db/
│   └── migrations/           # Database migrations
│       ├── 000001_init_schema.down.sql
│       └── 000001_init_schema.up.sql
├── internal/
│   ├── controller/           # HTTP controllers layer
│   │   ├── category.go      # Category controller
│   │   ├── transaction.go   # Transaction controller
│   │   └── user.go          # User controller
│   ├── docs/                # Swagger documentation
│   │   ├── docs.go
│   │   ├── swagger.json
│   │   └── swagger.yaml
│   ├── middleware/          # Application middlewares
│   │   └── auth.go          # Authentication middleware
│   ├── model/               # Data models
│   │   ├── category.go      # Category model
│   │   ├── transaction.go   # Transaction model
│   │   └── user.go          # User model
│   ├── repository/          # Data access layer
│   │   ├── category.go      # Category repository
│   │   ├── transaction.go   # Transaction repository
│   │   └── user.go          # User repository
│   ├── router/              # Route configuration
│   │   ├── router.go        # Main router
│   │   └── routes.go        # Route definitions
│   ├── service/             # Business logic layer
│   │   ├── category.go      # Category service
│   │   ├── transaction.go   # Transaction service
│   │   └── user.go          # User service
│   └── utils/               # Utilities
├── scripts/                 # Helper scripts
├── tests/                   # Automated tests
├── .air.toml               # Air configuration (hot reload)
├── .gitignore              # Git ignored files
├── docker-compose.yml      # Docker Compose configuration
├── Dockerfile              # Application Docker image
├── go.mod                  # Go dependencies
├── go.sum                  # Dependencies checksums
├── LICENSE                 # Project license
├── main.go                 # Alternative main file
├── makefile                # Make commands
├── mise.toml               # Mise configuration
└── README.md               # Project documentation
```

## Error Handling

The API returns standardized errors with appropriate HTTP codes and descriptive messages. All errors follow the format:

```json
{
  "error": "Error description",
  "code": 400
}
```

## Logging

The system implements structured logging to facilitate debugging and monitoring:

- **Debug**: Detailed information for development
- **Error**: Critical errors that require attention

## Technologies Used

- **Go**: Programming language
- **Gin**: Web framework for Go
- **PostgreSQL**: Relational database
- **JWT**: Token-based authentication
- **Swagger**: Automatic API documentation
- **Docker**: Application containerization
- **Air**: Hot reload for development
- **Mise**: Development tools manager

## Swagger/OpenAPI

The API is documented using Swagger annotations, allowing automatic generation of interactive documentation. Annotations include:

- Endpoint descriptions
- Model schemas
- Response codes
- Authentication requirements

## Database

The application uses PostgreSQL as the main database. Migrations are managed through SQL files in the `db/migrations/` folder.

### Migrations

- **000001_init_schema.up.sql**: Initial table creation
- **000001_init_schema.down.sql**: Initial table rollback

### Table Structure

Based on the models, the main tables are:

- **users**: User storage
- **categories**: Transaction categories by user
- **transactions**: User financial transactions

## Development

### Requirements

- Go 1.21+
- PostgreSQL 13+
- Docker (optional)
- Air (for hot reload)

### Environment Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/breno5g/GoBudget.git
   cd GoBudget
   ```

2. **Configure environment variables**:
   ```bash
   # Copy the example file and edit the configurations
   cp .env.example .env
   ```

3. **Run migrations**:
   ```bash
   make migrate-up
   ```

4. **Run the application**:
   ```bash
   # Development (with hot reload)
   make dev
   
   # Production
   make build
   make run
   ```

### Docker

To run with Docker:

```bash
# Start services
docker-compose up -d

# Access the application at http://localhost:8080
```

### Make Commands

The project includes a Makefile with useful commands:

```bash
make build       # Compile the application
make run         # Run the application
make dev         # Run in development mode
make test        # Run tests
make migrate-up  # Run migrations
make migrate-down # Rollback migrations
make clean       # Clean temporary files
```

### Swagger UI

Interactive API documentation is available at:
- **Development**: http://localhost:8080/swagger/index.html
- **Production**: https://your-api.com/swagger/index.html

## Next Steps

To use this API, you can:

1. Implement a web or mobile client
2. Add features like financial reports
3. Implement advanced transaction filters
4. Add multi-currency support
5. Create data visualization dashboards

## Security Considerations

- All passwords must be hashed before storage
- JWT tokens should have appropriate expiration time
- Strict input data validation through middlewares
- Rate limiting to prevent API abuse
- HTTPS mandatory in production
- Authentication middleware protects sensitive routes
- Input data validation with custom validators

## Contributing

1. Fork the project
2. Create a branch for your feature (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Open a Pull Request

## License

This project is under the license specified in the `LICENSE` file.
