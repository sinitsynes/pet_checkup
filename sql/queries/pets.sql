-- name: CreatePet :exec
INSERT INTO pets (name, birth, passport, chip)
VALUES ($1, $2, $3, $4);

-- name: GetPet :one
SELECT * FROM pets WHERE id = $1;

-- name: UpdatePet :exec
UPDATE pets SET 
    name = coalesce(sqlc.arg('name'), name), 
    birth = coalesce(sqlc.narg('birth'), birth), 
    passport = coalesce(sqlc.narg('passport'), passport), 
    chip = coalesce(sqlc.narg('chip'), chip)
WHERE id = @id;

-- name: DeletePet :exec
DELETE FROM pets WHERE id = $1;

