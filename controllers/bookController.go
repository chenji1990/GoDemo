package controllers

import (
	"context"
	"net/http"
	"root/database"
	"root/models"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(c *gin.Context) {
	var res []models.Book
	pipline := qmgo.Pipeline{
		bson.D{{Key: operator.Lookup, Value: bson.M{
			"from":         "authors",
			"foreignField": "_id",
			"localField":   "author",
			"as":           "authors"}},
		},
		bson.D{{Key: operator.Lookup, Value: bson.M{
			"from":         "genres",
			"foreignField": "_id",
			"localField":   "genre",
			"as":           "genres"},
		}},
	}
	err := database.BookCollection.Aggregate(context.Background(), pipline).All(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func BookIndex(c *gin.Context) {
	bookCount, _ := database.BookCollection.Find(context.Background(), bson.D{}).Count()
	bookInstanceCount, _ := database.BookInstanceCollection.Find(context.Background(), bson.D{}).Count()
	bookInstanceAvailableCount, _ := database.BookInstanceCollection.Find(context.Background(), bson.D{{Key: "status", Value: "Available"}}).Count()
	authorCount, _ := database.AuthorCollection.Find(context.Background(), bson.D{}).Count()
	genreCount, _ := database.GenreCollection.Find(context.Background(), bson.D{}).Count()

	c.JSON(http.StatusOK, bson.M{
		"bookCount":                  bookCount,
		"bookInstanceCount":          bookInstanceCount,
		"bookInstanceAvailableCount": bookInstanceAvailableCount,
		"authorCount":                authorCount,
		"genreCount":                 genreCount,
	})
}

func GetBookDetail(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	var book models.Book
	err = database.BookCollection.Find(context.Background(), bson.D{{Key: "_id", Value: id}}).One(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	var bookInstances []models.BookInstance
	err = database.BookInstanceCollection.Find(context.Background(), bson.D{{Key: "book", Value: id}}).All(&bookInstances)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bson.M{"book": book, "bookInstances": bookInstances})
}
