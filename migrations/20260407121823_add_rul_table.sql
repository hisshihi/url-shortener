-- +goose Up
CREATE TABLE urls (
    id BIGSERIAL PRIMARY KEY ,
    alias VARCHAR(16) NOT NULL UNIQUE ,
    long_url TEXT NOT NULL ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_urls_alias ON urls (alias);

-- +goose Down
DROP TABLE IF EXISTS urls;
