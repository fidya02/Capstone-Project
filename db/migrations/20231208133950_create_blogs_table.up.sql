BEGIN;

CREATE TABLE IF NOT EXISTS blogs (
    id SERIAL PRIMARY KEY,
    image TEXT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

COMMIT;
