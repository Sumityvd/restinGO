package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)
// Registers a user for a specific event
//
// @Summary 			Register for an event
// @Description 		Register the authenticated user for a specific event
// @Id 				RegisterForEvent
// @Tags 			Events
// @Accept 			json
// @Produce 			json
// @Param 			id path int true "Event ID"
// @Success 			201 	{object} 	models.EventRegister 		"Registered!"
// @Failure 			500 	{object} 	models.EventRegisterError 	"Could not fetch event or register user for event."
// @Security			ApiKeyAuth
// @Router 			/events/{id}/register [post]
func RegisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
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

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}
// Xancels the registration of a user for a specific event
//
// @Summary 			Cancel event registration
// @Description 		Cancel the registration of the authenticated user for a specific event
// @Id 				CancelRegistration
// @Tags 			Events
// @Accept 			json
// @Produce 			json
// @Param 			id path int true "Event ID"
// @Success 			200	{object} 	models.EventCancalled 		  "Cancelled!"
// @Failure 			400 	{object} 	models.EventCancalledError 	  "Invalid event ID."
// @Security			ApiKeyAuth
// @Router 			/events/{id}/register [delete]
func CancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}
