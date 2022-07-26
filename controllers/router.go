package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 鉴权中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := c.Request.Header.Get("Authorization"); checkSessionToken(token) {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
	}
}

// 检查token
func checkSessionToken(token string) bool {
	return true
}

func InitRouters(router *gin.Engine) *gin.Engine {
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.StaticFS("/public", http.Dir("./public"))

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	InitAuthorRouter(router)
	InitBookRouter(router)
	InitBookInstanceRouter(router)
	InitGenreRouter(router)
	InitSettingRouter(router)
	InitUserRouter(router)
	InitLoginRouter(router)

	// api := router.Group("/api")

	// users := api.Group(("/users"))
	// users.GET("/:size", controllers.GetUsers)
	// users.GET("/insertUser", controllers.InsertUser)
	// users.DELETE("/:uid/delete", controllers.DeleteUser)
	// users.PUT("/:uid/update", controllers.UpdateUser)

	return router
}
