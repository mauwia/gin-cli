# Gin CLI

A powerful command-line interface tool for quickly scaffolding and setting up Go web applications using the Gin framework.

## Features

- 🚀 Quick project scaffolding with clean architecture
- 📦 Database setup (PostgreSQL and MongoDB)
- 🔄 CRUD operation generation
- 🛠️ Service layer generation
- 📝 Environment configuration management
- 🏗️ Standardized project structure

## Installation

```bash
go install github.com/mauwia/gin-cli@latest
```

## Usage

### Create a New Project

```bash
gin-cli new my-project
```

This will create a new project with the following structure:
```
my-project/
├── config/           # Configuration management
├── internal/         # Private application code
│   ├── server/      # Server setup and routing
│   ├── controllers/ # HTTP request handlers
│   ├── models/      # Data structures
│   ├── services/    # Business logic
│   ├── repositories/# Data access layer
│   ├── middlewares/ # HTTP middleware
│   └── utils/       # Utility functions
├── migrations/      # Database migrations
└── .env            # Environment configuration
```

### Database Setup

#### PostgreSQL Setup
```bash
gin-cli setup postgres
```
This will:
- Install required dependencies (gorm.io/gorm, gorm.io/driver/postgres)
- Configure environment variables
- Set up database configuration
- Initialize database connection

#### MongoDB Setup
```bash
gin-cli setup mongodb
```
This will:
- Install required dependencies (go.mongodb.org/mongo-driver/mongo)
- Configure environment variables
- Set up MongoDB configuration
- Initialize database connection

### Code Generation

#### Generate CRUD Operations
```bash
gin-cli generate crud user
```
This will generate:
- Controller
- Service
- Repository
- Model
for the specified entity (in this case, "user")

#### Generate Service
```bash
gin-cli generate service auth
```
This will generate a new service file for the specified service name.

### Environment Configuration
```bash
gin-cli update env
```
Updates the environment configuration based on your .env file.

## Project Structure

The generated project follows clean architecture principles:

- **Controllers**: Handle HTTP requests and responses
- **Services**: Contain business logic
- **Repositories**: Handle data access
- **Models**: Define data structures
- **Middlewares**: Handle cross-cutting concerns
- **Utils**: Contains utility functions

## Database Configuration

### PostgreSQL
Default environment variables:
```
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=postgres
DB_PORT=5432
DB_HOST=localhost
```

### MongoDB
Default environment variables:
```
MONGODB_URI=
MONGODB_DATABASE=
```

## Requirements

- Go 1.x or higher
- Git

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) 