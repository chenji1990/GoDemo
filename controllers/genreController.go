package controllers

import (
	"context"
	"net/http"
	"root/database"
	"root/models"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/locales/id"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitGenreRouter(router *gin.Engine) {
	group := router.Group("/genres")
	group.GET("/", GetGenres)
	group.GET("/:id", GetGenreDetail)
}

func GetGenres(c *gin.Context) {
	var res []models.Genre
	err := database.GenreCollection.Find(context.Background(), bson.D{}).All(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetGenreDetail(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	var genre []models.Genre
	err = database.GenreCollection.Find(context.Background(), bson.D{{Key: "_id", Value: id}}).All(&genre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	var genreBooks []models.Book
	err = database.BookCollection.Find(context.Background(), bson.D{{Key: "genre", Value: id}}).All(&genreBooks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bson.M{"genre": genre, "genreBooks": genreBooks})
}
