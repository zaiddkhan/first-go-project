-- +goose Up
ALTER TABLE feeds add column last_fetched_at TIMESTAMP;

-- +goose Down
ALTER table feeds drop column last_fetched_at