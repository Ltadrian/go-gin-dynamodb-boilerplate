package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ltadrian/go-gin-dynamodb-boilerplate/config"
)

type TokenController struct{}

func (h TokenController) GetToken(c *gin.Context) {
	// generate jwt token and return as body
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": time.Now().UnixMilli(),
		"exp": time.Now().Add(time.Duration(time.Hour) * 1).UnixMilli(),
	})
	secret := config.GetConfig().GetString("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating token", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
