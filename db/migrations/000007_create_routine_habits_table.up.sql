CREATE TABLE IF NOT EXISTS routine_habits(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    daily_routine_id UUID NOT NULL,
    habit_id UUID NOT NULL,
    routine_status_id BIGINT NOT NULL,
    notes TEXT,
    planned_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (daily_routine_id) REFERENCES daily_routines(id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (habit_id) REFERENCES habits(id)
        ON DELETE CASCADE ON UPDATE CASCADE
);