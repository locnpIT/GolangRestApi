package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"loc.com/hocgolang/models"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
	}

	context.JSON(http.StatusOK, event)

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again latter"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again latter"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})

}

// func updateEvent(context *gin.Context) {
// 	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
// 		return
// 	}
// 	_, err = models.GetEventByID(eventId)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not found event."})
// 		return
// 	}

// 	var updatedEvent models.Event
// 	err = context.ShouldBindJSON(&updatedEvent)

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request event."})
// 		return
// 	}

// 	updatedEvent.ID = eventId

// 	err = updatedEvent.Update()
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update event."})
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})

// }

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}
	var updatedEvent *models.Event
	updatedEvent, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Cannot fetch user"})
		return
	}
	var requestData models.Event
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
		return
	}

	if requestData.Name != "" {
		updatedEvent.Name = requestData.Name
	}
	if requestData.Description != "" {
		updatedEvent.Description = requestData.Description
	}
	if requestData.Location != "" {
		updatedEvent.Location = requestData.Location
	}
	if !requestData.DateTime.IsZero() {
		updatedEvent.DateTime = requestData.DateTime
	}

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})

}
