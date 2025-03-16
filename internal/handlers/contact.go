package handlers

import (
	"net/http"
	"strconv"

	"cmd/main.go/internal/models"
	"cmd/main.go/internal/services"

	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
	contacts := services.GetContacts()
	c.JSON(http.StatusOK, contacts)
}

func PostContact(c *gin.Context) {
	var newContact models.Contact
	if err := c.BindJSON(&newContact); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}
	services.CreateContact(newContact)
	c.IndentedJSON(http.StatusCreated, newContact)
}

func UpdateContact(c *gin.Context) {
	id := c.Param("id")
	key := c.Param("key")
	value := c.Param("value")

	contactID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid contact ID"})
		return
	}

	updatedContact, err := services.UpdateContact(contactID, key, value)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if updatedContact == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid key"})
		return
	}
	c.JSON(http.StatusOK, updatedContact)
}

func DeleteContact(c *gin.Context) {
	id := c.Param("id")
	contactID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid contact ID"})
		return
	}

	if services.DeleteContact(contactID) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Delete successful"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Contact not found"})
	}
}