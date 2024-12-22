-- name: CreateMeasurement :copyfrom
INSERT INTO panel_measurements (component_id, date_measured, pet_id, amount)
VALUES ($1, $2, $3, $4);
