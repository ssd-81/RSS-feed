-- +goose Up
-- +goose StatementBegin
CREATE TABLE feeds(id UUID PRIMARY KEY, created_at timestamp, updated_at timestamp, name varchar(255), url varchar(255) unique, user_id UUID,CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feeds;
-- +goose StatementEnd

