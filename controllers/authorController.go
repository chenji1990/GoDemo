package controllers

import (
	"context"
	"net/http"
	"root/database"
	"root/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAuthors(c *gin.Context) {
	var res []models.Author
	err := database.AuthorCollection.Find(context.Background(), bson.D{}).All(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func GetAuthorDetail(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	var author models.Author
	err = database.AuthorCollection.Find(context.Background(), bson.D{{Key: "_id", Value: id}}).One(&author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, author)
}
