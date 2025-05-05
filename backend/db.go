package main

import (
"database/sql"
\_ "github.com/lib/pq"
"log"
"os"
)

var db \*sql.DB

func init() {
dbHost := os.Getenv("DB\_HOST")
dbUser := os.Getenv("DB\_USER")
dbPassword := os.Getenv("DB\_PASSWORD")
dbName := os.Getenv("DB\_NAME")
dbPort := os.Getenv("DB\_PORT")

```
dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	dbHost, dbPort, dbUser, dbPassword, dbName)
var err error
db, err = sql.Open("postgres", dsn)
if err != nil {
	log.Fatalf("Error opening database: %s", err)
}
```

}
