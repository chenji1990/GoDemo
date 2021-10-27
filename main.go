package main

import (
	"context"
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

	url := "mongodb+srv://admin:admin@locallibrary.l8n9d.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	// url := "mongodb://localhost:27017"
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: url})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Close(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	database.InitDatabase(client)
	gin.ForceConsoleColor()
	controllers.InitRouters(gin.Default()).Run(":8080")

	// router := controllers.InitRouters(gin.Default())
	// srv := &http.Server{Addr: ":80", Handler: router}
	// go func() {
	// 	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
	// 		log.Printf("listen: %s\n", err)
	// 	}
	// }()

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit

	// log.Println("Shutting down server...")

	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server forced to shutdown: ", err)
	// }
	// log.Println("Server exiting")
}
