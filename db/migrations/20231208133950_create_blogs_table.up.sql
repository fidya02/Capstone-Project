CREATE TABLE blogs (
    Id SERIAL PRIMARY KEY,
    Image TEXT,
	Title VARCHAR(255) NOT NULL,
	Description TEXT,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
);

COMMIT;