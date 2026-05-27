-- name: CreatePhone :one
INSERT INTO phones (user_id, phone_number)
VALUES ($1, $2)
RETURNING *;

-- name: GetPhone :one
SELECT * FROM phones
WHERE user_id = $1 LIMIT 1;

-- name: UpdatePhone :one
UPDATE phones
SET phone_number = $2
WHERE user_id = $1
RETURNING *;

-- name: DeletePhone :exec
DELETE FROM phones
WHERE user_id = $1;