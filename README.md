# Bangkit Capstone Backend
## Project Setup
Follow these steps to set up and run the application:

### Setup Environment
#### 1. Copy the `.env` structure to create your development environment file:
```bash
cp .env-sample .env.dev
```
Update the values in `.env.dev` as needed.

#### 2. Run the Application
Use the `Makefile` to start the application in development mode:
```bash
make start-dev
```

### Database Migrations
#### Manage database migrations with the following commands:

* Migrate Up

To apply all pending migrations and update the database schema:
```bash
make migrate-up
```

* Migrate Down

To roll back the most recent migration:
```bash
make migrate-down
```