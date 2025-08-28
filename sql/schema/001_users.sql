-- +goose Up
CREATE TABLE users(id UUID PRIMARY KEY, created_at timestamp, updated_at timestamp, name varchar(255) unique not null);

-- +goose Down
DROP TABLE users;