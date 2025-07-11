CREATE TABLE IF NOT EXISTS habits(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    habit_category_id BIGINT NOT NULL,
    user_id UUID NOT NULL,
    name VARCHAR(255),
    description TEXT,
    thumbnail_url TEXT,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (habit_category_id) REFERENCES habit_categories(id)
        ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE ON UPDATE CASCADE

);