-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash(255) NOT NULL,
    bio TEXT,
    creted_at TIMESTAMP WITH TIMEZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIMEZONE DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
