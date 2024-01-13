package main

import (
	"todo/internal/app"
	"todo/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database, err := db.NewDB()

	if err != nil {
		panic(err)
	}

	myApp := app.App{
		DB: database,
	}

	defer myApp.DB.Close()

	r.POST("/todoList", myApp.HomeHandler)
	r.POST("/todoList/add", myApp.NewTask)
	r.POST("/todoList/delete", myApp.DeleteTask)
	r.GET("/todoList/get", myApp.ReadTasks)

	r.Run(":8080")
}
