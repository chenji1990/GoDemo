package controllers

import (
	"context"
	"log"
	"net/http"

	"root/database"
	"root/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(c *gin.Context) {
	var res []models.User
	err := database.UserCollection.Find(context.Background(), bson.D{}).All(&res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func InsertUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	ctx := context.Background()
	res, err := database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	log.Println("insert success:", res.InsertedID)
	c.JSON(http.StatusOK, res)
}

// func GetUsers(c *gin.Context) {
// 	sizeParam := c.Param("size")
// 	size, err := strconv.ParseInt(sizeParam, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	ctx := context.Background()
// 	var users []models.User

// 	findOptions := options.Find()
// 	findOptions.SetLimit(size)
// 	searchResult, err := database.UserCollection.Find(ctx, bson.D{}, findOptions)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
// 		return
// 	}

// 	defer searchResult.Close(ctx)

// 	if err = searchResult.All(ctx, &users); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, users)
// }

func DeleteUser(c *gin.Context) {
	userUid := c.Param("uid")

	ctx := context.Background()
	err := database.UserCollection.RemoveId(ctx, userUid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotModified, gin.H{"ERROR": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userUid)
}

// func UpdateUser(c *gin.Context) {
// 	var user models.User
// 	userUid := c.Param("uid")

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
// 		return
// 	}

// 	byteUser, err := bson.Marshal(user)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
// 		return
// 	}

// 	var bUser bson.M
// 	if err = bson.Unmarshal(byteUser, &bUser); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
// 		return
// 	}

// 	ctx := context.Background()
// 	var updatedUser models.User

// 	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
// 	filter := bson.M{"_uid": userUid}
// 	update := bson.D{{Key: "$set", Value: user}}
// 	err = database.UserCollection.Remove(ctx, filter, update, opts).Decode(&updatedUser)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			c.JSON(http.StatusNotModified, gin.H{"ERROR": err.Error()})
// 			return

// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, updatedUser)
// }
