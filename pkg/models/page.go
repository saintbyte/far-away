package models

import (
	"github.com/google/uuid"
	"time"
)

const (
	DirectionFromUser = 1
	DirectionToUser   = 2
)

type PageDBModel struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Title        string    `gorm:"type:varchar(255)"`
	Author       string    `gorm:"type:varchar(255)"`
	Text         string    `gorm:"type:text"`
	AccessSecret string    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
