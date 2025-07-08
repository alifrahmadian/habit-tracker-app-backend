package models

import (
	"time"

	"github.com/google/uuid"
)

type Habit struct {
	Id              int64         `json:"id"`
	HabitCategoryId int64         `json:"habit_category_id"`
	HabitCategory   HabitCategory `json:"habit_category"`
	UserId          uuid.UUID     `json:"user_id"`
	User            User          `json:"user"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	ThumbnailUrl    string        `json:"thumbnail_url"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
