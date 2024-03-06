-- name: CreateEntries :one
INSERT INTO entries ( 
    account_id,
    amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateEntries :one
UPDATE entries 
SET account_id = $1, 
amount = $2
RETURNING *;

-- name: DeleteEntries :exec
DELETE FROM entries WHERE id = $1;