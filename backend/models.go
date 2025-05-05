package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// GET /employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, name, role FROM employees")
	if err != nil {
		http.Error(w, "Failed to fetch employees: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Role); err != nil {
			http.Error(w, "Failed to scan employee: "+err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, emp)
	}

	json.NewEncoder(w).Encode(employees)
}

// POST /employees
func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if emp.Name == "" || emp.Role == "" {
		http.Error(w, "Name and Role are required fields", http.StatusBadRequest)
		return
	}

	err := db.QueryRow(
		"INSERT INTO employees (name, role) VALUES ($1, $2) RETURNING id",
		emp.Name, emp.Role,
	).Scan(&emp.ID)

	if err != nil {
		http.Error(w, "Failed to insert employee: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}
