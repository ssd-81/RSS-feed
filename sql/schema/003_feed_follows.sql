-- +goose Up
-- +goose StatementBegin
CREATE TABLE feed_follows(id UUID PRIMARY KEY, created_at timestamp, updated_at timestamp, user_id UUID,feed_id UUID, CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
CONSTRAINT fk_feed FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
CONSTRAINT unique_user_feed UNIQUE (user_id, feed_id));
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feeds_follows;
-- +goose StatementEnd

