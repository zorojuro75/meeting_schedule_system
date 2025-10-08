package models

import "time"

// User represents an application user.
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	IsAdmin      bool      `gorm:"default:false" json:"is_admin"`
	Disabled     bool      `gorm:"default:false" json:"disabled"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// MeetingStatus values
const (
	MeetingScheduled   = "scheduled"
	MeetingRescheduled = "rescheduled"
	MeetingCancelled   = "cancelled"
)

// Meeting represents a scheduled meeting.
type Meeting struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Title        string    `gorm:"not null" json:"title"`
	Description  string    `json:"description"`
	StartAt      time.Time `gorm:"not null" json:"start_at"`
	Link         string    `json:"link"`
	OrganizerID  uint      `gorm:"not null" json:"organizer_id"`
	Organizer    User      `gorm:"foreignKey:OrganizerID" json:"organizer"`
	Status       string    `gorm:"not null;default:'scheduled'" json:"status"`
	Participants []*User   `gorm:"many2many:meeting_participants;" json:"participants"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// MeetingParticipant holds extra fields for many2many relation.
type MeetingParticipant struct {
	ID        uint `gorm:"primaryKey"`
	MeetingID uint `gorm:"index"`
	UserID    uint `gorm:"index"`
	AddedBy   uint
	AddedAt   time.Time
}

// AuditLog stores actions for audit purposes.
type AuditLog struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ActorID    uint      `json:"actor_id"`
	Action     string    `json:"action"`
	EntityType string    `json:"entity_type"`
	EntityID   uint      `json:"entity_id"`
	Detail     string    `json:"detail"`
	CreatedAt  time.Time `json:"created_at"`
}
