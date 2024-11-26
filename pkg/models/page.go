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
	Slug         string    `gorm:"type:varchar(255);unique_index;not null"`
	Title        string    `gorm:"type:varchar(255)"`
	Author       string    `gorm:"type:varchar(255)"`
	Text         string    `gorm:"type:text"`
	AccessSecret string    `gorm:"type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PageCreateRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	HTML   string `json:"html"`
}

type PageUpdateRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	HTML   string `json:"html"`
	Slug   string `json:"author"`
	Secret string `json:"secret"`
}

type PageCreateResponse struct {
	Slug   string `json:"url"`
	Secret string `json:"secret"`
}
type PageUpdateResponse struct {
	Slug string `json:"url"`
}
