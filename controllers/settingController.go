package controllers

import (
	"context"
	"net/http"
	"root/database"
	"root/models"

	// . "root/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSettings(c *gin.Context) {
	var res []models.Setting
	err := database.SettingCollection.Find(context.Background(), bson.D{}).All(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}

func Query(c *gin.Context) {
	var res bson.D
	err := database.SettingCollection.Find(context.Background(), bson.D{}).All(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func Add(c *gin.Context) {
	res, err := database.SettingCollection.InsertOne(context.Background(), bson.M{"name": "pi",
		"value": 3.14159})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
