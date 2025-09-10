-- +goose Up
-- +goose StatementBegin
ALTER TABLE feeds ADD last_fetched_at timestamp;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE feeds DROP COLUMN last_fetched_at;
-- +goose StatementEnd

