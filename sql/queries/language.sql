-- name: CreateLanguage :one
INSERT INTO languages (name, label)
VALUES ($1, $2)
RETURNING *;

-- name: RemoveLanguage :exec
DELETE FROM languages
WHERE id = @id;

-- name: GetLanguages :many
SELECT *
FROM languages;

-- name: GetLanguagebyID :one
SELECT *
FROM languages
WHERE id = @id;
