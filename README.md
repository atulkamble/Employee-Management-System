# Employee Management System

## Backend (Go 1.19)

```bash
cd backend
go get ./...
DB_HOST=localhost DB_USER=postgres DB_PASSWORD=admin123 DB_NAME=employees DB_PORT=5432 go run main.go
```

## Frontend (React 14.17.0)

```bash
cd frontend
npm install
npm start
```

## PostgreSQL Table Setup

```sql
CREATE TABLE employees (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100)
);
```
