package main

import (
	"context"
	"fmt"
	"log"
	"root/controllers"
	"root/database"
	"time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"

	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
	"github.com/mochi-co/mqtt/server/listeners/auth"

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

	// url2 := "mongodb+srv://admin:admin@locallibrary.l8n9d.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	// client2, err2 := qmgo.NewClient(ctx, &qmgo.Config{Uri: url2})
	// if err2 != nil {
	// 	log.Fatal(err)
	// }

	// DB := client.Database("my_database")
	// DB2 := client2.Database("my_database")
	// // AuthorCollection := DB.Collection("authors")
	// // BookCollection := DB.Collection("books")
	// // BookInstanceCollection := DB.Collection("bookinstances")
	// // GenreCollection := DB.Collection("genres")
	// // SettingCollection := DB.Collection("settings")
	// // UserCollection := DB.Collection("users")

	// var res []bson.D
	// DB.Collection("books").Find(context.Background(), bson.D{}).All(&res)
	// DB2.Collection("books").InsertMany(context.Background(), res)

	// DB.Collection("bookinstances").Find(context.Background(), bson.D{}).All(&res)
	// DB2.Collection("bookinstances").InsertMany(context.Background(), res)

	// DB.Collection("genres").Find(context.Background(), bson.D{}).All(&res)
	// DB2.Collection("genres").InsertMany(context.Background(), res)

	// DB.Collection("settings").Find(context.Background(), bson.D{}).All(&res)
	// DB2.Collection("settings").InsertMany(context.Background(), res)

	// DB.Collection("users").Find(context.Background(), bson.D{}).All(&res)
	// DB2.Collection("users").InsertMany(context.Background(), res)

	defer func() {
		if err = client.Close(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	database.InitDatabase(client)

	gin.ForceConsoleColor()

	server := mqtt.NewServer(nil)

	ws := listeners.NewWebsocket("ws1", ":1882")
	err = server.AddListener(ws, &listeners.Config{
		Auth: &auth.Allow{},
	})
	if err != nil {
		log.Fatal(err)
	}

	tcp := listeners.NewTCP("t1", ":1883")
	err = server.AddListener(tcp, &listeners.Config{
		Auth: &auth.Allow{},
	})
	if err != nil {
		log.Fatal(err)
	}

	go server.Serve()
	fmt.Println("MQTT   Started! ")

	defer func() {
		server.Close()
		fmt.Println("MQTT   Finished")
	}()

	controllers.InitRouters(gin.Default()).Run(":3000")

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
