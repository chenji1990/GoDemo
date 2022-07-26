package controllers

import (
	"context"
	"fmt"
	"net/http"
	"root/database"
	"root/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

var jwtkey = []byte("www.zhuchengji.com")

func InitLoginRouter(router *gin.Engine) {

	router.GET("/getToken", getToken)
	router.GET("/setToken", setToken)

	group := router.Group("/user")
	group.POST("/signUp", signUp)
	group.POST("/signIn", signIn)
	group.POST("/signOut", signOut)
}

func signUp(c *gin.Context) {

}

func signIn(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")

	var user models.User
	err := database.UserCollection.Find(context.Background(), bson.M{
		"username": username,
	}).One(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}

	if user.Password != password {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "password is wrong"})
		return
	}

}

func signOut(c *gin.Context) {

}

type Claims struct {
	Id uint
	jwt.StandardClaims
}

func setToken(c *gin.Context) {

	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Id: 2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "sessionToken",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"sessionToken": tokenString})
}

func getToken(c *gin.Context) {

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "token is empty",
		})
		return
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "token is error",
		})
		return
	}
	fmt.Println(111)
	fmt.Println(claims.Id)
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, claims, err
}
