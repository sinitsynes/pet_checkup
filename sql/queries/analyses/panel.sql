-- name: Create :one
INSERT INTO panel (name, description)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateByID :one
UPDATE panel
SET name = coalesce(sqlc.narg('name'), name),
    description = coalesce(sqlc.narg('description'), description)
WHERE id = $1
RETURNING *;

-- name: DeleteByID :exec
DELETE FROM panel
WHERE id = $1;

-- name: GetByID :one
SELECT * FROM panel
JOIN component ON panel.id = component.panel_id
WHERE panel.id = $1;
