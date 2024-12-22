-- name: CreateLanguage :one
INSERT INTO languages (name, label)
VALUES ($1, $2)
RETURNING *;

-- name: RemoveLanguage :exec
DELETE FROM languages
WHERE id = @id;

-- name: GetLanguages :many
SELECT name
FROM languages;

-- name: GetLanguageID :one
SELECT id
FROM languages
WHERE label = @label;