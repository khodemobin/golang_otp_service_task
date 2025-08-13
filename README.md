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
git clone <repository-url>
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

## 📁 Project Structure

```
golang_otp_service_task/
├── cmd/                    # Command line applications
│   └── serve.go           # Server command
├── internal/              # Internal application code
│   ├── app/              # Application container & setup
│   ├── config/           # Configuration management
│   ├── model/            # Data models
│   ├── server/           # HTTP server & routing
│   │   ├── dto/          # Data Transfer Objects
│   │   ├── handler/      # HTTP handlers
│   │   ├── middleware/   # HTTP middleware
│   │   ├── route.go      # Route definitions
│   │   └── server.go     # Server setup
│   └── service/          # Business logic services
├── pkg/                   # Reusable packages
│   ├── apperror/         # Error handling
│   ├── cache/            # Cache implementations
│   ├── logger/           # Logging utilities
│   ├── pgsql/            # PostgreSQL connection
│   └── response/         # Response utilities
├── docs/                 # API documentation
├── docker/               # Docker configurations
├── docker-compose.yaml   # Docker services
├── Dockerfile           # Application container
├── go.mod              # Go modules
└── README.md           # This file
```

## 🔧 API Endpoints

### Authentication

#### Send OTP
```http
POST /api/auth/otp/send
Content-Type: application/json

{
  "phone": "09123456789"
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
  "phone": "09123456789",
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

## 🗄️ Database

### Models

#### User
```go
type User struct {
    ID        uint           `gorm:"primarykey"`
    Phone     string         `gorm:"uniqueIndex;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

### Auto Migration
The application automatically migrates the database schema on startup.

## 🔐 Authentication System

### OTP Features
- **6-digit OTP generation**: Secure random number generation
- **2-minute expiration**: OTP expires after 2 minutes
- **3-attempt limit**: Maximum 3 verification attempts per OTP
- **Console output**: OTP is printed to console (for development)
- **Automatic cleanup**: Expired OTPs are automatically removed

### JWT Tokens
- **24-hour validity**: Tokens expire after 24 hours
- **User claims**: Contains user ID and phone number
- **Protected routes**: Use for API authentication

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Run specific package tests
go test ./pkg/cache -v
```

## 📊 Logging

The project uses Zap logger for structured logging:

```go
import "github.com/khodemobin/golang_otp_service_task/internal/app"

// Usage
app.Log.Info("Application started")
app.Log.Error("An error occurred", "error", err)
```

## 🔧 Configuration

Configuration is managed through environment variables:

```go
import "github.com/khodemobin/golang_otp_service_task/internal/app"

// Access configuration
config := app.Config
db := app.DB
cache := app.Cache
```

## 🐳 Docker

### Build Image
```bash
docker build -t golang-otp-service .
```

### Run Container
```bash
docker run -p 3000:3000 --env-file .env golang-otp-service
```

### Docker Compose
```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f container

# Stop services
docker-compose down
```

## 🚀 CLI Commands

The application uses Cobra for CLI commands:

```bash
# Run the server
go run main.go serve

# Show help
go run main.go --help
```

## 🔍 Development

### Project Dependencies

Key dependencies include:
- **Fiber v2**: Web framework
- **GORM v1**: ORM for database operations
- **PostgreSQL**: Primary database
- **JWT v5**: JSON Web Token implementation
- **Zap**: Structured logging
- **Cobra**: CLI framework
- **Cleanenv**: Configuration management
- **Invopop Validation**: Request validation

### Code Organization

The project follows clean architecture principles:
- **DTOs**: Data Transfer Objects for API requests/responses
- **Services**: Business logic layer
- **Handlers**: HTTP request handling
- **Middleware**: Cross-cutting concerns
- **Models**: Database entities

## 🚀 Deployment

### Production Checklist
1. Set `APP_ENV=production` in `.env`
2. Configure strong `JWT_SECRET`
3. Set up proper database credentials
4. Configure reverse proxy (nginx)
5. Set up SSL/TLS certificates
6. Configure monitoring and logging
7. Set up backup strategies

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | Application port | 8000 |
| `APP_ENV` | Environment | local |
| `JWT_SECRET` | JWT signing secret | secret |
| `DB_HOST` | Database host | localhost |
| `DB_PORT` | Database port | 3306 |
| `DB_DATABASE` | Database name | test |
| `DB_USERNAME` | Database user | test |
| `DB_PASSWORD` | Database password | secret |

## 🤝 Contributing

1. Fork the project
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support and questions:
- Create an issue on GitHub
- Contact the development team

## 🔄 Changelog

### v1.0.0
- Initial project setup
- OTP-based authentication system
- JWT token support
- PostgreSQL integration with GORM
- Docker containerization
- Swagger documentation
- CLI interface with Cobra
- Structured logging with Zap
- Request validation
- Error handling system
