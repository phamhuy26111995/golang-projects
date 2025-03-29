package main

import (
	"net/http"

	"example.com/rest-api/database"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	server := gin.Default()

	server.GET("/events",getEvents )
	server.POST("/events", createEvent)

	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	event, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	event.ID = 1
	event.UserId = 1

	err := event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	context.JSON(http.StatusCreated, event)
}