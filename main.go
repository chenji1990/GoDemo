package main

import (
	"context"
	"fmt"
	"log"
	"root/controllers"
	"root/database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// url := "mongodb+srv://admin:admin@locallibrary.l8n9d.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	url := "mongodb://localhost:27017"
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: url})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("数据库连接成功！")
	defer func() {
		if err = client.Close(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	database.InitDatabase(client)
	gin.ForceConsoleColor()
	controllers.InitRouters(gin.Default()).Run(":80")
}
