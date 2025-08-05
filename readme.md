# Terminal E-Ticketing System - Setup and Running Guide

## 📌 Project Overview

Terminal E-Ticketing System is a public transportation e-ticketing platform with JWT-based authentication. The system is built using:

- **Go 1.21+**
- **Fiber** web framework
- **GORM** ORM
- **PostgreSQL** database
- **Docker** (optional, for containerized environments)

---

## ⚙️ Prerequisites

### ✅ Required Tools

- [Go 1.21+](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/)
  - Migrate CLI  
    Install using:
    ```bash
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

    ```
    🚀 Setup Instructions
    1. Clone the Repository
    ```
    git clone <repository-url>
    cd terminal-workspace
    ```
    2. Initialize Go Workspace and Modules
    
    ```aiignore
        make configure
    ```
    ▶️ Running the Application
    Option 1: Run Directly (Recommended for Development)
  - # Run database migrations


    make migrate
    
    make run
    
    
# Build and run the application
Option 2: Run with Docker (Recommended for Production)
```aiignore
# Build and start all services
docker-compose up -d --build

# View logs
docker-compose logs -f api

```