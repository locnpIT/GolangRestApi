package routes

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"loc.com/hocgolang/models"
	"loc.com/hocgolang/utils"
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
	authHeader := context.Request.Header.Get("Authorization") // Lấy header

	if authHeader == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized. Authorization header missing."})
		return
	}

	// Kiểm tra xem header có bắt đầu bằng "Bearer " không
	// và tách lấy phần token thực sự
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized. Invalid Authorization header format."})
		return
	}

	tokenString := parts[1] // Đây mới là chuỗi JWT cần verify

	// Gọi VerifyToken với chuỗi JWT đã được tách ra
	err := utils.VerifyToken(tokenString)

	if err != nil {
		// Có thể log lỗi err ở đây để debug rõ hơn
		// log.Printf("Token verification failed: %v\n", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized. Invalid token."})
		return
	}

	// --- Phần còn lại của hàm giữ nguyên ---
	var event models.Event
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	// !!! Quan trọng: Phần này bạn đang hardcode UserID.
	// !!! Bạn nên lấy UserID từ trong token sau khi verify thành công.
	// event.ID = 1 // ID thường do database tự tạo hoặc không cần set ở đây
	event.UserID = 1 // <<-- Nên lấy từ token

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later"})
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

func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Cannot found id event"})
		return
	}

	var requestEvent *models.Event
	requestEvent, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Cannot found event"})
		return
	}

	err = requestEvent.Delete()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Delete event successfully"})

}
