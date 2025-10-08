package repository

import (
	"meeting_scheduler/internal/models"
	"time"
)

// Ensure Repos implements the repo interfaces via methods below.

// Create a new user record
func (r *Repos) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *Repos) GetByEmail(email string) (*models.User, error) {
	var u models.User
	if err := r.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repos) GetByID(id uint) (*models.User, error) {
	var u models.User
	if err := r.DB.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// Meeting methods
func (r *Repos) CreateMeeting(m *models.Meeting) error {
	return r.DB.Create(m).Error
}

func (r *Repos) GetMeetingByID(id uint) (*models.Meeting, error) {
	var m models.Meeting
	if err := r.DB.Preload("Participants").Preload("Organizer").First(&m, id).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *Repos) UpdateMeeting(m *models.Meeting) error {
	return r.DB.Save(m).Error
}

func (r *Repos) ListMeetingsByUser(userID uint) ([]models.Meeting, error) {
	var meetings []models.Meeting
	// find meetings where user is organizer or participant
	if err := r.DB.Joins("JOIN meeting_participants mp ON mp.meeting_id = meetings.id").Where("mp.user_id = ? OR organizer_id = ?", userID, userID).Find(&meetings).Error; err != nil {
		return nil, err
	}
	return meetings, nil
}

func (r *Repos) UpcomingBetween(after, before time.Time) ([]models.Meeting, error) {
	var meetings []models.Meeting
	if err := r.DB.Preload("Participants").Where("start_at BETWEEN ? AND ?", after, before).Find(&meetings).Error; err != nil {
		return nil, err
	}
	return meetings, nil
}
