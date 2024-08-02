# Agnos Backend

This project is a backend service built with Go, using PostgreSQL as the database and Nginx as a reverse proxy.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

1. Clone the repository:
git clone https://github.com/santhureinoo/santhureinoo_agnos_backend.git cd santhureinoo_agnos_backend


2. Set up environment variables:
Create a `.env` file in the root directory by using .env.example file

3. Start the services:
docker-compose up -d

This command will start the following services:
- App: The main Go application
- Nginx: Reverse proxy
- DB: PostgreSQL database
- Migrate: Runs database migrations
- PgAdmin: PostgreSQL administration tool

4. Access the application:
The application will be available at `http://localhost`

5. Access PgAdmin:
PgAdmin will be available at `http://localhost:5050`
- Email: admin@example.com
- Password: admin

## Development

To make changes to the application:

1. Modify the Go code in the project directory
2. Rebuild and restart the services:

docker-compose down docker-compose up --build -d

## Database Migrations

Migrations are automatically run when the services start. To add new migrations:

1. Add new SQL files to the `./migrations` directory
2. Rebuild and restart the services

## Troubleshooting

If you encounter any issues, check the logs of the services:

docker-compose logs app docker-compose logs db docker-compose logs migrate
