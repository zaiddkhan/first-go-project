-- +goose Up

ALTER TABLE users add column  api_key VARCHAR(64) UNIQUE  NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea),'hex')
 );

-- +goose Down
ALTER TABLE users drop column api_key;