package controllers

import (
	"fmt"
	"net/http"

	"github.com/bishalr0y/go_webserver/config"
	"github.com/bishalr0y/go_webserver/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectToDb()
var validate = validator.New(validator.WithRequiredStructEnabled())

func HelloWorld() {
	fmt.Println("Hello world from controller")
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)

	if err := validate.Struct(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	// validating the todos
	for _, todo := range todos {
		if err := validate.Struct(todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

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

	if err := validate.Struct(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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
