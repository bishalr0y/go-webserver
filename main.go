package main

import (
	"github.com/bishalr0y/go_webserver/config"
	"github.com/bishalr0y/go_webserver/controllers"
	"github.com/bishalr0y/go_webserver/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = config.ConnectToDb()
	db.AutoMigrate(&models.Todo{})

	r := gin.Default()
	controllers.HelloWorld()
	r.POST("/todos", controllers.CreateTodo)
	r.GET("/todos", controllers.FetchAllTodos)
	r.GET("/todos/:id", controllers.FetchSingleTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	r.POST("/todos/multiple", controllers.CreateTodos)

	r.Run(":8888")
}
