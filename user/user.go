package user

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class string `json:"class"`
}

var db *sql.DB

func Setup(database *sql.DB) {
	db = database
}

// func ListUsers(c *gin.Context) {
// 	var users []User
// 	rows, err := db.Query("SELECT id, name, age, class FROM users")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var u User
// 		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Class); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		users = append(users, u)
// 	}
// 	c.JSON(http.StatusOK, users)
// }

func ListUsers(c *gin.Context) {
	// Parse query parameters for pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "10"))

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * perPage

	// Query to count total items
	var total int
	countQuery := "SELECT COUNT(*) FROM users"
	err := db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error retrieving resources.", "error": err.Error()})
		return
	}

	// Adjust totalPages to ensure it always has a value
	totalPages := total / perPage
	if total%perPage != 0 {
		totalPages++
	}

	// Query to fetch paginated items
	var users []User
	query := "SELECT id, name, age, class FROM users LIMIT ? OFFSET ?"
	rows, err := db.Query(query, perPage, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error retrieving resources.", "error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Class); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error retrieving resources.", "error": err.Error()})
			return
		}
		users = append(users, u)
	}

	// Construct the response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Resources retrieved successfully.",
		"data": gin.H{
			"items":       users,
			"total":       total,
			"perPage":     perPage,
			"currentPage": page,
			"totalPages":  totalPages,
		},
	})
}

func AddUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO users (name, age, class) VALUES (?, ?, ?)", newUser.Name, newUser.Age, newUser.Class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	newUser.ID = int(id)
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Resource created successfully.", "data": newUser})
}

func EditUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid user ID.", "error": err.Error()})
		return
	}

	var updatedUser User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error updating resource.", "error": err.Error()})
		return
	}

	_, err = db.Exec("UPDATE users SET name = ?, age = ?, class = ? WHERE id = ?", updatedUser.Name, updatedUser.Age, updatedUser.Class, id)
	if err != nil { // Assuming 'err' holds the error from your update operation
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error updating resource.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Resource updated successfully.", "data": updatedUser})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid user ID.", "error": err.Error()})
		return
	}

	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil { // Assuming 'err' holds the error from your delete operation
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error deleting resource.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Resource deleted successfully."})
}

// CreateTable creates the users table if it does not exist
func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        age INTEGER NOT NULL,
        class TEXT NOT NULL
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
}
