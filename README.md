# Event Management Backend

A Go-based event management API built with Fiber framework and PostgreSQL.

## Features

- User authentication with JWT
- Event management (CRUD operations)
- Ticket booking system
- RESTful API design
- PostgreSQL database integration

## Local Development

### Prerequisites

- Go 1.23+
- PostgreSQL
- Make (optional)

### Setup

1. Clone the repository
2. Copy `.env.example` to `.env` and configure your environment variables
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## Deployment to Render

This application is configured for easy deployment on Render.

### Automatic Deployment

1. Push your code to a Git repository (GitHub, GitLab, etc.)
2. Connect your repository to Render
3. Render will automatically detect the `render.yaml` configuration
4. The deployment will create:
   - A PostgreSQL database
   - A web service running your Go application
   - All necessary environment variables will be auto-configured

### Environment Variables

The following environment variables are automatically configured by Render:

- `PORT` - Server port (automatically set by Render)
- `DATABASE_URL` - Complete database connection string
- `DB_HOST` - Database host
- `DB_NAME` - Database name
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password
- `DB_SSLMODE` - Set to "require" for production
- `JWT_SECRET` - Auto-generated secure secret
- `SERVER_PORT` - Fallback port (3000)

### Manual Deployment Steps

If you prefer manual setup:

1. Create a new Web Service on Render
2. Connect your repository
3. Configure build and start commands:
   - **Build Command**: `go build -o main ./cmd/api`
   - **Start Command**: `./main`
4. Add a PostgreSQL database
5. Set up the environment variables listed above

## API Endpoints

### Authentication
- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User login

### Events (Protected)
- `GET /api/event` - Get all events
- `POST /api/event` - Create new event
- `GET /api/event/:id` - Get event by ID
- `PUT /api/event/:id` - Update event
- `DELETE /api/event/:id` - Delete event

### Tickets (Protected)
- `GET /api/ticket` - Get all tickets
- `POST /api/ticket` - Book a ticket
- `GET /api/ticket/:id` - Get ticket by ID
- `PUT /api/ticket/:id` - Update ticket
- `DELETE /api/ticket/:id` - Cancel ticket

## Database Schema

The application uses GORM for database operations and includes auto-migration on startup.

## Security

- JWT-based authentication
- Password hashing with bcrypt
- CORS enabled for cross-origin requests
- SSL required in production

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request
