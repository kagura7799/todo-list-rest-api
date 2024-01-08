package main

import (
	"github.com/gin-gonic/gin"
	"todo/internal/app"
)

func main() {
	r := gin.Default()

	myApp := app.App{}
	r.POST("/todoList", myApp.HomeHandler)
	r.POST("/todoList/add", myApp.NewTask)
	r.POST("/todoList/delete", myApp.DeleteTask)

	r.Run(":8080")
}
