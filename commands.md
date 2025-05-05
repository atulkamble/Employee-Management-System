brew update
brew doctor
brew tap petere/postgresql

brew install postgresql@14

brew services start postgresql@14

sudo initdb /usr/local/var/postgresql@14 -E utf8
createuser -s postgres
psql -U postgres

\c employees

```
CREATE TABLE employees (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  role VARCHAR(100)
);
```

\q

```
DB_HOST=localhost \
DB_PORT=5432 \
DB_USER=postgres \
DB_PASSWORD=admin123 \
DB_NAME=employees \
go run *.go
```


```
cd backend
go mod tidy         # Clean and update module dependencies
go run *.go         # Compile and run all .go files in backend
```
