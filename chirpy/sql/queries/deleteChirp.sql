-- name: DeleteChirp :exec
DELETE from chirps where id = $1;
