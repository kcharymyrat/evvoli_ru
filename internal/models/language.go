package models

type Language struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
