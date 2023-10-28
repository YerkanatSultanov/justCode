package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"justCode/lecture8/auth/model"
)

var users []model.User

func CreateUser(c *gin.Context) {
	var newItem model.User
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem.ID = len(users) + 1
	users = append(users, newItem)

	c.JSON(http.StatusCreated, newItem)
}

func OutUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var foundItem model.User
	for _, user := range users {
		if user.ID == id {
			foundItem = user
			break
		}
	}

	if foundItem.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, foundItem)
}
