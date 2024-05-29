package api

import "github.com/google/uuid"

type LanguageDTO struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CategoryRequestDTO struct {
	Slug         string                   `json:"slug"`
	Image        string                   `json:"image"`
	Thumbnail    string                   `json:"thumbnail,omitempty"`
	Translations []CategoryTranslationDTO `json:"translations"`
}

type CategoryResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	LanguageCode string    `json:"language_code"`
	Slug         string    `json:"slug"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Image        string    `json:"image"`
	Thumbnail    string    `json:"thumbnail"`
}

type CategoryTranslationDTO struct {
	LanguageCode string `json:"language_code"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
}
