package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)
// GetEvents retrieves all events
//
// @Summary 			Get all events
// @Description 		Retrieve a list of all events
// @Id 				GetEvents
// @Tags 			Events
// @Accept 			json
// @Produce 			json
// @Success 			200 	{array} 	models.Event 		"List of events"
// @Failure 			500 	{object} 	models.EventError 	"Could not fetch events"
// @Router 			/events [get]
func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}
// GetEvent retrieves a specific event by ID
//
// @Summary 			Get a specific event
// @Description 		Retrieve details of a specific event by ID
// @Id 				GetEvent
// @Tags 			Events
// @Accept 			json
// @Produce 			json
// @Param 			id path int true "Event ID"
// @Success 			200 	{object} 	models.Event 		"Event details"
// @Failure 			500 	{object} 	models.EventErrorId 	"Could not fetch event."
// @Router 			/events/{id} [get]
func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}
// CreateEvent creates a new event
//
// @Summary 			Create a new event
// @Description 		Create a new event with the provided data
// @Id 				CreateEvent
// @Tags 			Events
// @Accept 			json
// @Produce 			json
// @Param 			event body models.EventPlane true "Event data"
// @Success 			201 	{object} 	models.EventCreated 		"Event created!"
// @Failure 			500 	{object} 	models.EventCreatedError 	"Could not create event. Try again later."
// @Security			ApiKeyAuth
// @Router 			/events [post]
func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
// UpdateEvent updates an existing event
//
// @Summary 			Update an event
// @Description 		Update an existing event with the provided data
// @Id 				UpdateEvent
// @Tags 			Events
// @Accept 			json
// @Produce 			json
// @Param 			id path int true "Event ID"
// @Param 			event body models.EventPlane true "Updated event data"
// @Success 			201 	{object} 	models.EventUpdated 		"Event updated!"
// @Failure 			500 	{object} 	models.EventUpdatedError 	"Could not update event. Try again later."
// @Security			ApiKeyAuth
// @Router 			/events/{id} [put]
func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}
// DeleteEvent deletes an existing event
//
// @Summary 			Delete an event
// @Description 		Delete an existing event by ID
// @Id 				DeleteEvent
// @Tags 			Events
// @Accept 			json
// @Produce 			json
// @Param 			id path int true "Event ID"
// @Success 			200 {object} 	models.EventDeleted 		"Event deleted successfully!"
// @Failure 			401 {object} 	models.EventDeletedError 	"Not authorized to delete event."
// @Security			ApiKeyAuth
// @Router 			/events/{id} [delete]
func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
