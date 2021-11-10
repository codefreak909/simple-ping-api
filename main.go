package main

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	r.GET("/ping", pingHandler)
	r.GET("/employee", employeeHandler)

	r.Run(":8000")
}

func pingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"response": "pong!"})
}

func employeeHandler(ctx *gin.Context) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:2000)/test")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	id := ctx.Params.ByName("id")

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
		return
	}

	empBytes, err := json.Marshal(emp)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.JSON(200, empBytes)
}
