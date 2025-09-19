-- name: GetUsrEmail :one
Select * from users 
WHERE email = $1;
