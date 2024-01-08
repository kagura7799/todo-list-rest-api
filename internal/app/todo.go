package app

import (
	"github.com/gin-gonic/gin"
)

type App struct {
}

func (a *App) HomeHandler(c *gin.Context) {
	c.String(200, "Todo list")
}

func (a *App) NewTask(c *gin.Context) {

}

func (a *App) DeleteTask(c *gin.Context) {

}
