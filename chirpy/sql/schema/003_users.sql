-- +goose Up
ALTER TABLE users
ADD column hashed_password TEXT NOT NULL DEFAULT 'unset';

-- +goose Down
ALTER TABLE users
drop column hashed_password;
