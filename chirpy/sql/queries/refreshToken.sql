-- name: RefreshToken :exec
UPDATE refresh_tokens
SET expires_at = $1
where token = $2;
