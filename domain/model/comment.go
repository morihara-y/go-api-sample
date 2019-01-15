package model

import "time"

// Comment has comment data
type Comment struct {
	ID     int
	UserID int
	Date   time.Time
	Text   string
}
