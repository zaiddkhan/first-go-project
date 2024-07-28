-- +goose Up

CREATE table feeds (
    id UUID PRIMARY KEY ,
    created_at TIMESTAMP not null,
    updated_at timestamp not null ,
    name TEXT NOT NULL ,
    url Text UNIQUE not null,
    user_id uuid REFERENCES users(id) ON Delete cascade
);

-- +goose Down
DROP TABLE feeds;