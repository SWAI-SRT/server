package dto

type CreateReportDto struct {
  Type      string  `json:"type" validate:"required"`
  Title     string  `json:"title" validate:"required"`
	Content   string  `json:"content" validate:"required"`
  Latitude  float64 `json:"latitude" validate:"required"`
  Longitude float64 `json:"longitude" validate:"required"`
  ImageUri  string  `json:"imageUri"`
}