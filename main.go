package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
}

func main() {
	router := gin.Default()
	db, err := sql.Open("mysql", "root:root@cp(localhst:330)/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	router.GET("/ping", func(c *gin.Context) {
		var user User
		err = db.QueryRow("SELECT id, name FROM users WHERE id = ?", 1).Scan(&user.ID, &user.Name)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("pong, user id: %d, user name: %s", user.ID, user.Name),
		})
	})
	router.Run(":8080")
}
