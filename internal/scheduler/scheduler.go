package scheduler

import (
	"fmt"
	"log"
	"meeting_scheduler/config"
	"meeting_scheduler/internal/models"
	"meeting_scheduler/internal/utils"
	"time"

	"gorm.io/gorm"
)

// Start runs the scheduler loop once to find meetings starting in ~10 minutes and sends reminders.
func Start(db *gorm.DB, cfg *config.Config) error {
	// This is a simple implementation for demo: run once at startup.
	now := time.Now().UTC()
	after := now.Add(9 * time.Minute)
	before := now.Add(11 * time.Minute)

	var meetings []models.Meeting
	if err := db.Preload("Participants").Where("start_at BETWEEN ? AND ?", after, before).Find(&meetings).Error; err != nil {
		return fmt.Errorf("query upcoming: %w", err)
	}

	for _, m := range meetings {
		to := []string{}
		for _, p := range m.Participants {
			to = append(to, p.Email)
		}
		if len(to) == 0 {
			continue
		}
		subject := fmt.Sprintf("Reminder: %s starts at %s UTC", m.Title, m.StartAt.Format(time.RFC3339))
		body := fmt.Sprintf("Meeting %s will start at %s. Link: %s", m.Title, m.StartAt.Format(time.RFC3339), m.Link)
		go func(to []string, subject, body string) {
			if err := utils.SendEmail(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPFrom, to, subject, body); err != nil {
				log.Printf("failed to send reminder: %v", err)
			}
		}(to, subject, body)
	}

	return nil
}
