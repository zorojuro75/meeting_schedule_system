package handlers

import (
	"meeting_scheduler/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MeetingHandler struct {
	svc *services.MeetingService
}

func NewMeetingHandler(svc *services.MeetingService) *MeetingHandler {
	return &MeetingHandler{svc: svc}
}

type createMeetingRequest struct {
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description"`
	StartAt      string `json:"start_at" binding:"required"`
	Link         string `json:"link"`
	Participants []uint `json:"participants"`
}

func (h *MeetingHandler) Create(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented yet"})
}
