CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    message TEXT NOT NULL,
	Is_Read BOOLEAN DEFAULT FALSE,
    status TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
	deleted_at TIMESTAMP,
);

COMMIT;