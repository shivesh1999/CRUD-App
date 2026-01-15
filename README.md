# CRUD-App

A full-stack web application for managing contacts with a **Go backend** and **React frontend**, containerized with Docker.

## 📋 Overview

CRUD-App is a boilerplate project demonstrating a complete contact management system with:
- **Backend**: REST API built with Go using Gin web framework
- **Frontend**: Modern React application with Tailwind CSS styling
- **Database**: PostgreSQL for persistent data storage
- **Infrastructure**: Docker & Docker Compose for easy deployment

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────────┐
│                   Frontend (React)                       │
│  - Contact List View                                    │
│  - Add/Edit/Delete Contacts                            │
│  - Tailwind CSS Styling                                │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓ HTTP (REST API)
┌─────────────────────────────────────────────────────────┐
│            Backend (Go + Gin)                            │
│  - API Routes & Controllers                            │
│  - Contact CRUD Operations                             │
│  - Data Validation                                      │
│  - CORS Support                                         │
   - API Routes & Controllers                            │
   - Contact CRUD Operations                             │
   - Data Validation                                     │
   - CORS Support                                        │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓ SQL
┌─────────────────────────────────────────────────────────┐
│            Database (PostgreSQL)                        │
│  - Contact Records Storage                             │
│  - Automatic Migrations                                │
└─────────────────────────────────────────────────────────┘
```

## 🚀 Getting Started

### Prerequisites
- Docker & Docker Compose
- OR
- Go 1.20+
- Node.js 14+
- PostgreSQL 13+

### Quick Start with Docker

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd CRUD-App
   ```

2. **Create environment file**
   ```bash
   cp .env.example .env
   ```

3. **Start the application**
   ```bash
   docker-compose up --build
   ```

4. **Access the application**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - Database: localhost:5432

### Local Development Setup

#### Backend Setup

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Configure environment variables**
   ```bash
   cp .env.example .env
   ```

3. **Run the backend server**
   ```bash
   go run main.go
   ```
   Server will start on `http://localhost:8080` (now powered by Gin)

#### Frontend Setup

1. **Navigate to client directory**
   ```bash
   cd client
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Start development server**
   ```bash
   npm start
   ```
   Application will open at `http://localhost:3000`

4. **Build for production**
   ```bash
   npm run build
   ```

## 📁 Project Structure

```
CRUD-App/
├── main.go                      # Application entry point
├── go.mod                       # Go module dependencies
├── docker-compose.yml           # Multi-container orchestration
├── Dockerfile                   # Backend container configuration
├── .env.example                 # Environment variables template
│
├── bootstrap/
│   └── app.go                   # Application initialization & setup
│
├── database/
│   ├── models/
│   │   └── contact.go          # Contact data model with validation
│   ├── storage/
│   │   └── postgres.go         # Database connection & configuration
│   └── migrations/
│       └── migrations.go        # Database schema migrations
│
├── repository/
│   ├── controller.go           # API endpoints & business logic
│   ├── repository.go           # Database operations interface
│   └── routes.go               # Route definitions
│
└── client/                      # React frontend application
    ├── Dockerfile             # Frontend container configuration
    ├── package.json           # Node.js dependencies
    ├── tailwind.config.js      # Tailwind CSS configuration
    ├── public/                # Static assets
    │   ├── index.html         # Main HTML file
    │   ├── manifest.json      # PWA manifest
    │   └── robots.txt         # SEO robots file
    └── src/                   # React components & styles
        ├── App.js             # Main React component
        ├── App.css            # Application styles
        ├── index.js           # React root entry
        ├── config.js          # Application configuration
        ├── index.css          # Global styles
        └── contacts/          # Contact management components
            ├── List.js        # Contact list view
            ├── EachContact.js # Individual contact component
            └── ViewContact.js # Contact detail view
```

## 🔗 API Endpoints

All API endpoints are prefixed with `/api`

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/contacts` | Get all contacts |
| POST | `/api/contacts` | Create a new contact |
| GET | `/api/contacts/:id` | Get contact by ID |
| PATCH | `/api/contacts/:id` | Update contact |
| DELETE | `/api/contacts/:id` | Delete contact |

### Contact Model

```json
{
  "name": "string (5-40 chars, required)",
  "contactNumber": "string (exactly 10 digits, required)",
  "city": "string (max 40 chars, required)",
  "country": "string (max 40 chars, required)",
  "email": "string (valid email, 10-40 chars, required)"
}
```

## 🛠️ Technology Stack

### Backend
- **Go 1.20** - Programming language
- **Fiber v2** - Web framework for Go (fast HTTP server)
- **PostgreSQL** - Database
- **pgx** - PostgreSQL driver
- **godotenv** - Environment variable management

### Frontend
- **React 18** - UI library
- **React Router v6** - Client-side routing
- **Axios** - HTTP client
- **Tailwind CSS** - Utility-first CSS framework
- **Moment.js** - Date/time handling

### DevOps
- **Docker** - Containerization
- **Docker Compose** - Multi-container orchestration

## 📝 Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
DB_HOST=localhost          # PostgreSQL host
DB_PORT=5432              # PostgreSQL port
DB_USER=postgres          # Database user
DB_PASS=postgres          # Database password
DB_NAME=postgres          # Database name
DB_SSLMODE=disable        # SSL mode for database connection
APP_ENV=development       # Application environment (optional)
```

For Docker deployment, these are automatically configured in `docker-compose.yml`.

## 🗄️ Database

### Automatic Migrations
The application automatically creates and manages the database schema on startup through the migration system in `database/migrations/migrations.go`.

### Connection Details
- **Host**: `db` (Docker) / `localhost` (local)
- **Port**: `5432`
- **Default User**: `postgres`
- **Default Password**: `postgres`
- **Default Database**: `postgres`

## 🔄 Development Workflow

1. **Backend changes**: Modify Go files and restart the server
2. **Frontend changes**: Changes auto-reload in development mode
3. **Database schema changes**: Update migration files and restart backend

## 🐳 Docker Commands

```bash
# Build and start all services
docker-compose up --build

# Run in background
docker-compose up -d

# View logs
docker-compose logs -f

# Stop all services
docker-compose down

# Remove volumes (clears database)
docker-compose down -v
```

## 📚 Key Features

- ✅ Complete CRUD operations for contacts
- ✅ Input validation on both frontend and backend
- ✅ PostgreSQL database with automatic migrations
- ✅ CORS enabled for cross-origin requests
- ✅ Responsive React UI with Tailwind CSS
- ✅ RESTful API design
- ✅ Docker containerization
- ✅ Hot reload in development

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is open source and available under the MIT License.

## 📞 Support

For issues and questions, please open an issue in the repository.

---

**Happy coding!** 🎉
