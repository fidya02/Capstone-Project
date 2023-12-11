CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    price INT,
    status TEXT DEFAULT 'available',
    image VARCHAR(255),
    location VARCHAR(255),
    quantity INT,
    category VARCHAR(255),
    date DATE,
    sold INT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);