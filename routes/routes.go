package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start(router *gin.Engine) {
	//default router
	router.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "default router")
	})

}
