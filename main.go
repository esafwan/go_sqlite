package main

import (
	"database/sql"
	"log"

	"github.com/esafwan/gosqlite/user"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to SQLite Database
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup user package and create table
	user.Setup(db)
	user.CreateTable()

	// Set up Gin
	r := gin.Default()

	// Routes
	r.GET("/users", user.ListUsers)
	r.POST("/users", user.AddUser)
	r.PUT("/users/:id", user.EditUser)
	r.DELETE("/users/:id", user.DeleteUser)

	// Start server
	r.Run() // listen and serve on 0.0.0.0:8080
}
