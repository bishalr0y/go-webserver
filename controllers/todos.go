package controllers

import (
	"fmt"
	"net/http"

	"github.com/bishalr0y/go_webserver/config"
	"github.com/bishalr0y/go_webserver/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectToDb()

func HelloWorld() {
	fmt.Println("Hello world from controller")
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := db.Create(&todo)
	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create the todo"})
		return
	}
	c.JSON(200, todo)
}

func CreateTodos(c *gin.Context) {
	var todos []models.Todo
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

func FetchAllTodos(c *gin.Context) {
	var todos []models.Todo
	err := db.Find(&todos)
	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch all todos"})
		return
	}
	c.JSON(200, todos)
}

func FetchSingleTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := db.First(&todo, id)
	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the todo"})
	}
	c.JSON(200, todo)
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
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

func DeleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id paramter not found"})
		return
	}

	var todo models.Todo
	err := db.Where("id = ?", id).Delete(&todo)
	if err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the todo"})
	}
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
