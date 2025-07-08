package models

import (
	"time"

	"github.com/google/uuid"
)

type RoutineHabit struct {
	Id              uuid.UUID     `json:"id"`
	DailyRoutineId  uuid.UUID     `json:"daily_routine_id"`
	DailyRoutine    DailyRoutine  `json:"daily_routine"`
	HabitId         uuid.UUID     `json:"habit_id"`
	Habit           Habit         `json:"habit"`
	RoutineStatusId int64         `json:"routine_status_id"`
	RoutineStatus   RoutineStatus `json:"routine_status"`
	Notes           string        `json:"notes"`
	PlannedTime     time.Time     `json:"planned_time"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
