package handlers

import (
	"net/http"
	"strconv"

	"cmd/main.go/internal/models"
	"cmd/main.go/internal/services"

	"github.com/gin-gonic/gin"
)

func GetGroups(c *gin.Context) {
	groups := services.GetGroups()
	c.JSON(http.StatusOK, groups)
}

func PostGroup(c *gin.Context) {
	var newGroup models.Group
	if err := c.BindJSON(&newGroup); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}
	services.CreateGroup(newGroup)
	c.IndentedJSON(http.StatusCreated, newGroup)
}

func UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	key := c.Param("key")
	value := c.Param("value")

	groupID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid group ID"})
		return
	}

	updatedGroup, err := services.UpdateGroup(groupID, key, value)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if updatedGroup == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid key"})
		return
	}
	c.JSON(http.StatusOK, updatedGroup)
}

func DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	groupID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid group ID"})
		return
	}

	if services.DeleteGroup(groupID) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Delete successful"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Group not found"})
	}
}