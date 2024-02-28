package main

import (
	"fmt"
	"net/http"

	"github.com/bishalr0y/go_webserver/config"
	"github.com/bishalr0y/go_webserver/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var db *gorm.DB

// var err error

func main() {
	// dsn := "user=admin password=admin dbname=my_db port=5432 sslmode=disable TimeZone=Asia/Kolkata host=localhost"
	// db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	db = config.ConnectToDb()
	db.AutoMigrate(&Todo{})

	r := gin.Default()
	controller.HelloWorld()
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
	err := db.Create(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the todo"})
		return
	}
	c.JSON(200, todo)
}

func createTodos(c *gin.Context) {
	var todos []Todo
	c.BindJSON(&todos)
	fmt.Println(todos)

	for _, todo := range todos {
		err := db.Create(&todo)
		if err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the todo", "todo": todo})
			return
		}
	}

	c.JSON(200, todos)
}

func fetchAllTodos(c *gin.Context) {
	var todos []Todo
	err := db.Find(&todos)
	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all todos"})
		return
	}
	c.JSON(200, todos)
}

func fetchSingleTodo(c *gin.Context) {
	var todo Todo
	id := c.Params.ByName("id")
	err := db.First(&todo, id)
	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the todo"})
	}
	c.JSON(200, todo)
}

func updateTodo(c *gin.Context) {
	var todo Todo
	id := c.Params.ByName("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id paramter not found"})
		return
	}

	err := db.First(&todo, id)

	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to find todo with that id"})
		return
	}

	c.BindJSON(&todo)
	err = db.Save(&todo)
	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the todo"})
		return
	}
	c.JSON(200, todo)
}

func deleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id paramter not found"})
		return
	}

	var todo Todo
	err := db.Where("id = ?", id).Delete(&todo)
	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the todo"})
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
