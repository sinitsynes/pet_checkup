-- name: CreatePanelComponent :one
INSERT INTO panel_components (name, unit_id, min_amount, max_amount, comment, language_id)
VALUES ($1, $2, $3, $4, $5, (SELECT id FROM languages WHERE label=$6))
RETURNING id;

-- name: GetComponents :many
SELECT panel_components.id,
	   panel_components.name,
	   panel_measurement_units.name AS unit,
	   min_amount,
	   max_amount,
	   comment
FROM panel_components
JOIN panel_measurement_units ON panel_components.unit_id = panel_measurement_units.id AND panel_components.language_id = panel_measurement_units.language_id
WHERE panel_components.language_id = (SELECT id FROM languages WHERE label=$1);
