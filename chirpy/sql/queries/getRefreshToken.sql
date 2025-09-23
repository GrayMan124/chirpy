-- name: GetRefreshToken :one
SELECT * from refresh_tokens 
WHERE token = $1;
