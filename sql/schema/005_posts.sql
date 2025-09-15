-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts(id UUID primary key, created_at timestamp, updated_at timestamp, title varchar(255), url varchar(255) unique, description varchar(512z), published_at timestamp, feed_id UUID,
CONSTRAINT fk_feed FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd

