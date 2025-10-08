package repository

import (
	"meeting_scheduler/internal/models"
	"time"

	"gorm.io/gorm"
)

type Repos struct {
	DB *gorm.DB
}

// User repository operations
type UserRepo interface {
	CreateUser(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
}

// Meeting repository operations
type MeetingRepo interface {
	CreateMeeting(meeting *models.Meeting) error
	GetMeetingByID(id uint) (*models.Meeting, error)
	UpdateMeeting(meeting *models.Meeting) error
	ListMeetingsByUser(userID uint) ([]models.Meeting, error)
	UpcomingBetween(after, before time.Time) ([]models.Meeting, error)
}

// NewRepos returns simple Repos wrapper around gorm DB.
func NewRepos(db *gorm.DB) *Repos {
	return &Repos{DB: db}
}
