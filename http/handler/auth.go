package handler

import (
	"net/http"
	"todoapp/config"
	jwt "todoapp/infrastructure/JWT"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}
	for _, user := range config.Users {
		if user.UserName == input.UserName && user.Password == input.Password {
			token, err := jwt.GenerateToken(user.ID, int(user.Type))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Token error"})
			}
			c.JSON(http.StatusOK, gin.H{"token": token})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Credentials"})
}
