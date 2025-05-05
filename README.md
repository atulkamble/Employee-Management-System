📁 employee-management-system/
├── backend/
│   ├── main.go
│   ├── db.go
│   ├── models.go
│   └── go.mod
├── frontend/
│   ├── package.json
│   ├── public/
│   │   └── index.html
│   └── src/
│       ├── App.js
│       ├── index.js
│       └── components/
│           └── EmployeeList.js
└── README.md

// ---------------------- backend/main.go ----------------------
package main

import (
"fmt"
"log"
"net/http"
"os"
"github.com/gorilla/mux"
)

func main() {
r := mux.NewRouter()
r.HandleFunc("/employees", getEmployees).Methods("GET")
r.HandleFunc("/employees", createEmployee).Methods("POST")

```
port := "8080"
fmt.Printf("Server running on port %s\n", port)
log.Fatal(http.ListenAndServe(":"+port, r))
```

}

// ---------------------- backend/db.go ----------------------
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

// ---------------------- backend/models.go ----------------------
package main

type Employee struct {
ID   int    `json:"id"`
Name string `json:"name"`
}

func getEmployees(w http.ResponseWriter, r \*http.Request) {
rows, err := db.Query("SELECT id, name FROM employees")
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
defer rows.Close()

```
var employees []Employee
for rows.Next() {
	var emp Employee
	if err := rows.Scan(&emp.ID, &emp.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	employees = append(employees, emp)
}
json.NewEncoder(w).Encode(employees)
```

}

func createEmployee(w http.ResponseWriter, r \*http.Request) {
var emp Employee
json.NewDecoder(r.Body).Decode(\&emp)
\_, err := db.Exec("INSERT INTO employees (name) VALUES (\$1)", emp.Name)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
w\.WriteHeader(http.StatusCreated)
}

// ---------------------- backend/go.mod ----------------------
module backend

go 1.19

require (
github.com/gorilla/mux v1.8.0
github.com/lib/pq v1.10.4
)

// ---------------------- frontend/src/App.js ----------------------
import React, { useEffect, useState } from 'react';
import axios from 'axios';

function App() {
const \[employees, setEmployees] = useState(\[]);
const \[name, setName] = useState("");

useEffect(() => {
axios.get("[http://localhost:8080/employees](http://localhost:8080/employees)")
.then(res => setEmployees(res.data));
}, \[]);

const addEmployee = () => {
axios.post("[http://localhost:8080/employees](http://localhost:8080/employees)", { name })
.then(() => {
setEmployees(\[...employees, { name }]);
setName("");
});
};

return ( <div> <h1>Employee Management System</h1>
\<input value={name} onChange={(e) => setName(e.target.value)} /> <button onClick={addEmployee}>Add</button> <ul>
{employees.map((emp, i) => ( <li key={i}>{emp.name}</li>
))} </ul> </div>
);
}

export default App;

// ---------------------- frontend/package.json ----------------------
{
"name": "employee-frontend",
"version": "1.0.0",
"dependencies": {
"axios": "^1.4.0",
"react": "^17.0.2",
"react-dom": "^17.0.2",
"react-scripts": "4.0.3"
},
"scripts": {
"start": "react-scripts start",
"build": "react-scripts build"
}
}

// ---------------------- README.md ----------------------

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
