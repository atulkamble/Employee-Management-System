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
