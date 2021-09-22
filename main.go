package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/employee", employeeHandler)

	http.ListenAndServe(":8000", nil)
}

func pingHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong")
}

func employeeHandler(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:2000)/test")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	id := req.URL.Query().Get("id")

	row := db.QueryRow("SELECT id, name, age FROM employee where id=?", id)

	type employee struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	emp := employee{}

	err = row.Scan(&emp.ID, &emp.Name, &emp.Age)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "OOPS!")
		return
	}

	empBytes, err := json.Marshal(emp)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "OOOPS!")
		return
	}

	fmt.Fprintf(w, string(empBytes))
}
