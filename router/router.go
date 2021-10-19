package router

import (
	"root/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) *gin.Engine {

	router.GET("/", func(c *gin.Context) {
		if c.Request.URL.Path != "/favicon.ico" {
			c.JSON(200, gin.H{"message": "Hello World"})
		}
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// api := router.Group("/api")

	// users := api.Group(("/users"))
	// users.GET("/:size", controllers.GetUsers)
	// users.GET("/insertUser", controllers.InsertUser)
	// users.DELETE("/:uid/delete", controllers.DeleteUser)
	// users.PUT("/:uid/update", controllers.UpdateUser)

	authors := router.Group("/authors")
	authors.GET("/", controllers.GetAuthors)
	authors.GET("/:id", controllers.GetAuthorDetail)

	books := router.Group("/books")
	books.GET("/", controllers.BookIndex)
	books.GET("/:id", controllers.GetBookDetail)
	books.GET("/list", controllers.GetBooks)

	bookInstances := router.Group("/bookinstances")
	bookInstances.GET("/", controllers.GetBookInstances)

	genres := router.Group("/genres")
	genres.GET("/", controllers.GetGenres)
	genres.GET("/:id", controllers.GetGenreDetail)

	settings := router.Group("/settings")
	settings.GET("/", controllers.GetSettings)
	settings.GET("/add", controllers.Add)

	users := router.Group(("/users"))
	users.GET("/", controllers.GetUsers)

	return router
}
