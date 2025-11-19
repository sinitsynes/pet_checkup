-- name: Create :one
INSERT INTO panel_component (panel_id, pet_id, component_id, date_measured, value)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetByID :one
SELECT * FROM panel_component
WHERE id = $1;

-- name: UpdateByID :one
UPDATE panel_component
SET panel_id = coalesce(sqlc.narg('panel_id'), panel_id),
    pet_id = coalesce(sqlc.narg('pet_id'), pet_id),
    component_id = coalesce(sqlc.narg('component_id'), component_id),
    date_measured = coalesce(sqlc.narg('date_measured'), date_measured),
    value = coalesce(sqlc.narg('value'), value)
WHERE id = @id
RETURNING *;

-- name: DeleteByID :exec
DELETE FROM panel_component
WHERE id = @id;

-- name: ListWithFilters :many
SELECT * FROM panel_component
WHERE (sqlc.narg('pet_id')::bigint IS NULL OR pet_id = sqlc.narg('pet_id'))
  AND (sqlc.narg('component_id')::bigint IS NULL OR component_id = sqlc.narg('component_id'))
  AND (sqlc.narg('panel_id')::bigint IS NULL OR panel_id = sqlc.narg('panel_id'))
  AND (sqlc.narg('date_from')::date IS NULL OR date_measured >= sqlc.narg('date_from'))
  AND (sqlc.narg('date_to')::date IS NULL OR date_measured <= sqlc.narg('date_to'))
ORDER BY date_measured DESC, id DESC;
