package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func socketHandler(c *gin.Context) {
	//upgrade get request to websocket protocol
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	for {
		//Read Message from client
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		msg := string(message)
		//If client message is ping will return pong
		if msg == "ping" {
			msg = "pong"
		}
		msg = "hello: " + msg
		//Response message to client
		err = ws.WriteMessage(mt, []byte(msg))
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func InitRouters(router *gin.Engine) *gin.Engine {
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.StaticFS("/public", http.Dir("./public"))

	router.GET("/ws", socketHandler)

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
