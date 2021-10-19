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
)

func GetBookInstances(c *gin.Context) {
	var res []models.BookInstance
	pipeline := qmgo.Pipeline{
		bson.D{{Key: operator.Lookup, Value: bson.M{
			"from":         "books",
			"foreignField": "_id",
			"localField":   "book",
			"as":           "books"}},
		}}

	err := database.BookInstanceCollection.Aggregate(context.Background(), pipeline).All(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}
