-- +goose Up
CREATE TABLE refresh_tokens(
	token TEXT primary KEY,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
	expires_at TIMESTAMP,
	revoked_at TIMESTAMP
);

-- +goose Down
DROP TABLE refresh_tokens;

