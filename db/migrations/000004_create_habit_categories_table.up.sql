CREATE TABLE IF NOT EXISTS habit_categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    icon_url TEXT,
    created_by UUID,
    updated_by UUID,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);