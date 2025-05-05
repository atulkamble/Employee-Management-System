package main

import (
	"encoding/json"
	"net/http"
)


type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, role FROM employees")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		rows.Scan(&emp.ID, &emp.Name, &emp.Role)
		employees = append(employees, emp)
	}
	json.NewEncoder(w).Encode(employees)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	err := db.QueryRow("INSERT INTO employees (name, role) VALUES ($1, $2) RETURNING id", emp.Name, emp.Role).Scan(&emp.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(emp)
}
