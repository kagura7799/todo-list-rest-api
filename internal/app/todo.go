package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo/internal/db"
)

type App struct {
	DB *db.DB
}

func (a *App) HomeHandler(c *gin.Context) {
	c.String(200, "Todo list")
}

func (a *App) NewTask(c *gin.Context) {

}

func (a *App) DeleteTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task id"})
		return
	}

	isDelete, err := a.DB.DeleteTaskFromDB(taskID)
	
	if isDelete && err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task deletion error"})
	} else {
		c.JSON(200, gin.H{"message": "Task deleted succesfully"})
	}
	
}
