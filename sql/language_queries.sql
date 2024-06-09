-- name: CreateLanguage :one
INSERT INTO languages (name)
VALUES ($1)
RETURNING *;

-- name: RemoveLanguage :exec
DELETE FROM languages
WHERE id = sqlc.arg('id');

-- name: GetLanguages :many
SELECT name
FROM languages;