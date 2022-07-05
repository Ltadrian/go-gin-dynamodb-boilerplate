package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ltadrian/test-dynamo-db-api/forms"
	"github.com/ltadrian/test-dynamo-db-api/models"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) FindUserByID(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := userModel.FindUserByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve user", "error": err})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User found!", "user": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

func (u UserController) AddUser(c *gin.Context) {
	var user forms.AddUser
	if c.Bind(&user) == nil {
		dbOutput, err := userModel.AddUser(user)
		if err != nil {
			fmt.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to add user"})
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "ok", "user": dbOutput})
	}
}
