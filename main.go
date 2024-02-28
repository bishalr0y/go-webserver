package main

import (
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

func main() {
	db = config.ConnectToDb()
	db.AutoMigrate(&Todo{})

	r := gin.Default()
	controller.HelloWorld()
	r.POST("/todos", controller.CreateTodo)
	r.GET("/todos", controller.FetchAllTodos)
	r.GET("/todos/:id", controller.FetchSingleTodo)
	r.PUT("/todos/:id", controller.UpdateTodo)
	r.DELETE("/todos/:id", controller.DeleteTodo)

	r.POST("/todos/multiple", controller.CreateTodos)

	r.Run(":8888")
}
