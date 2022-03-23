package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employee = []Employee{
	{ID: 1, Name: "Raditeo", Age: 23, Division: "test1"},
	{ID: 2, Name: "Warma", Age: 23, Division: "test2"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

// func getEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	if r.Method == "GET" {
// 		json.NewEncoder(w).Encode(employee)
// 		return
// 	}

// 	http.Error(w, "Invalid method", http.StatusBadRequest)
// }

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// if r.Method == "GET" {
	// 	json.NewEncoder(w).Encode(employee)
	// 	return
	// }

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee{
			ID:       len(employee) + 1,
			Name:     name,
			Age:      convertAge,
			Division: division,
		}

		employee = append(employee, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employee)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
