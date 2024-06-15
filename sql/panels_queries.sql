-- name: CreateMeasurementUnit :exec
INSERT INTO panel_units (name, language_id)
VALUES ($1, $2);

-- name: CreatePanelComponent :one
INSERT INTO panel_components (name, unit_id, average_amount, comment, language_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: CreateMeasurement :exec
INSERT INTO panel_measurements (component_id, date_measured, pet_id, amount)
VALUES ($1, $2, $3, $4);

-- name: GetLanguageID :one
SELECT id FROM languages WHERE label = $1;