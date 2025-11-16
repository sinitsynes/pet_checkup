-- name: CreateMeasurementUnit :exec
INSERT INTO panel_measurement_units (name, language_id)
VALUES ($1, (SELECT id FROM languages WHERE label=$2));

-- name: UpdateMeasurementUnit :exec
UPDATE panel_measurement_units
SET name = @name
WHERE id = @id;

-- name: GetMeasurementUnits :many
SELECT id, name FROM panel_measurement_units WHERE language_id = (SELECT languages.id FROM languages WHERE label = $1);

-- name: GetMeasurementUnit :one
SELECT id, name FROM panel_measurement_units WHERE language_id = (SELECT languages.id FROM languages WHERE label = $1) AND panel_measurement_units.id = @id;
