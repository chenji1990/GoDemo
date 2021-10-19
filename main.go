package main

import (
	"context"
	"fmt"
	"log"
	"root/database"
	"root/router"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
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
	router.InitRouters(gin.Default()).Run(":3000")
}
