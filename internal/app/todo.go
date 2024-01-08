package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type App struct {
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

	err = DeleteTaskFromDB(taskID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(200, gin.H{"message": "Task deleted succesfully"})
}
