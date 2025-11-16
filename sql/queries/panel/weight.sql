-- name: CreateWeightMeasure :exec
INSERT INTO pet_weight (weight, date_measured, pet_id)
VALUES ($1, $2, $3);


-- name: RemoveWeightMeasure :exec
DELETE FROM pet_weight 
WHERE id = @id;


-- name: GetWeightMeasures :many
SELECT * FROM pet_weight
WHERE pet_id = @pet_id;