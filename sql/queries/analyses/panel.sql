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
SELECT panel.id,
       panel.name,
       panel.description,
       array_agg(
        json_build_object('id', component.id,
                          'name', component.name,
                          'description', component.description)) as components
FROM panel
JOIN panel_component ON panel.id = panel_component.panel_id
JOIN component ON panel_component.component_id = component.id
WHERE panel.id = $1
GROUP BY panel.id, panel.name, panel.description;
