package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, h handler) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"test": "test",
		})
	})
	router.POST("/auth/login", h.LogIn)
	router.POST("/auth/refresh-token", h.RefreshToken)
}
