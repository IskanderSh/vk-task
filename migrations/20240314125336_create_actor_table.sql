-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS actors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    sex VARCHAR(255) NOT NULL,
    birthday DATE NOT NULL DEFAULT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS actors;
-- +goose StatementEnd
