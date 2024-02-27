package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var db *gorm.DB
var err error

func main() {
	// dsn := "postgresql://admin:admin@localhost:5432/my_db?schema=public"
	dsn := "user=admin password=admin dbname=my_db port=5432 sslmode=disable TimeZone=Asia/Kolkata host=localhost"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Todo{})

	r := gin.Default()
	r.POST("/todos", createTodo)
	r.GET("/todos", fetchAllTodos)
	r.GET("/todos/:id", fetchSingleTodo)
	r.PUT("/todos/:id", updateTodo)
	r.DELETE("/todos/:id", deleteTodo)

	r.POST("/todos/multiple", createTodos)

	r.Run(":8888")
}

func createTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	db.Create(&todo)
	c.JSON(200, todo)
}

func createTodos(c *gin.Context) {
	var todos []Todo
	c.BindJSON(&todos)
	fmt.Println(todos)

	for _, todo := range todos {
		err := db.Create(&todo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the todos", "todo": todo})
			return
		}
	}

	c.JSON(200, todos)
}

func fetchAllTodos(c *gin.Context) {
	var todos []Todo
	db.Find(&todos)
	c.JSON(200, todos)
}

func fetchSingleTodo(c *gin.Context) {
	var todo Todo
	id := c.Params.ByName("id")
	db.First(&todo, id)
	c.JSON(200, todo)
}

func updateTodo(c *gin.Context) {
	var todo Todo
	id := c.Params.ByName("id")
	db.First(&todo, id)
	c.BindJSON(&todo)
	db.Save(&todo)
	c.JSON(200, todo)
}

func deleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	db.Where("id = ?", id).Delete(&todo)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
