package auth

import (
	"net/http"
	"stocksync_api/internal/middleware"
	"stocksync_api/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, h handler) {
	router.POST("/auth/login", h.LogIn)
	router.POST("/auth/refresh-token", h.RefreshToken)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("auth/get-claims", func(c *gin.Context) {
			claims := c.MustGet("claims").(*jwt.Claims)
			c.JSON(http.StatusOK, gin.H{
				"user_id": claims.UserID,
				"email":   claims.Email,
				"role_id": claims.RoleID,
			})
		})
	}

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"test": "test",
	// 	})
	// })
}
