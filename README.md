# Golang OTP Service

A project featuring OTP-based authentication, built with Fiber, GORM, PostgreSQL, and modern Go practices.

## 🚀 Features

- **Web Framework**: Fiber (high-performance HTTP framework)
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: OTP-based login/registration system
- **JWT**: JSON Web Token support for API authentication
- **Caching**: In-memory cache for OTP storage
- **Logging**: Structured logging with Zap
- **Validation**: Request validation with Invopop validation
- **Configuration**: Environment-based configuration with Cleanenv
- **Documentation**: Swagger/OpenAPI documentation
- **Containerization**: Docker & Docker Compose support
- **CLI**: Command-line interface with Cobra
- **Error Handling**: Structured error handling
- **Middleware**: JWT authentication middleware

## 📋 Prerequisites

- Go 1.24+
- Docker & Docker Compose
- PostgreSQL 16+

## 🛠️ Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/khodemobin/golang_otp_service_task
cd golang_otp_service_task
```

### 2. Environment Configuration

Create a `.env` file in the project root:

```env
# Application Configuration
APP_PORT=3000
APP_ENV=local
JWT_SECRET=your-secret-key-here

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=otp_service
DB_USERNAME=postgres
DB_PASSWORD=password

# Optional: Database Port Forwarding
FORWARD_DB_PORT=5432
```

### 3. Using Docker Compose

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f container
```

### 4. Local Development

```bash
# Install dependencies
go mod tidy

# Run the application
go run main.go serve
```

## 🔧 API Endpoints

### Authentication

#### Send OTP
```http
POST /api/auth/otp/send
Content-Type: application/json

{
  "phone": "09121111111"
}
```

Response:
```json
{
  "message": "OTP sent successfully",
  "phone": "09123456789"
}
```

#### Verify OTP and Login/Register
```http
POST /api/auth/otp/verify
Content-Type: application/json

{
  "phone": "09121111111",
  "otp": "123456"
}
```

Response:
```json
{
  "message": "Authentication successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "phone": "09123456789"
    }
  }
}
```

### User Management (Protected Routes)

#### List Users
```http
GET /api/users?page=1&limit=10
Authorization: Bearer <jwt-token>
```

#### Get User by ID
```http
GET /api/users/1
Authorization: Bearer <jwt-token>
```

### Swagger Documentation
```http
GET /swagger
```