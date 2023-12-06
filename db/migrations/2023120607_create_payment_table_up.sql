CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    transaction_id INT NOT NULL,
    amount NUMERIC NULL,
    method TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
);

COMMIT;