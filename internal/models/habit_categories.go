package models

import (
	"time"

	"github.com/google/uuid"
)

type HabitCategory struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	IconUrl   string    `json:"icon_url"`
	CreatedBy uuid.UUID `json:"created_by"`
	UpdatedBy uuid.UUID `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
