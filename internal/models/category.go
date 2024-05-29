package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID `json:"id"`
	Slug        string    `json:"slug"`
	Image       string    `json:"image"`
	Thumbnail   string    `json:"thumbnail"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	CreatedByID uuid.UUID `json:"created_by_id"`
}

type CategoryTranslation struct {
	ID          int64     `json:"id"`
	CategoryID  uuid.UUID `json:"category_id"`
	LanguageID  Language  `json:"language"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
