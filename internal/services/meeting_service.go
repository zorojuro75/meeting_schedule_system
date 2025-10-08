package services

import (
	"errors"
	"meeting_scheduler/internal/models"
	"meeting_scheduler/internal/repository"
	"time"
)

type MeetingService struct {
	meetings repository.MeetingRepo
}

func NewMeetingService(meetings repository.MeetingRepo) *MeetingService {
	return &MeetingService{meetings: meetings}
}

// Create schedules a new meeting.
func (s *MeetingService) Create(m *models.Meeting) error {
	if m.StartAt.Before(time.Now().UTC()) {
		return errors.New("start time must be in the future")
	}
	m.Status = models.MeetingScheduled
	return s.meetings.CreateMeeting(m)
}
