-- name: CreateTransfer :one
INSERT INTO transfer ( 
    from_account_id,
    to_account_id,
    amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfer
WHERE id = $1 LIMIT 1;

-- name: ListTransfer :many
SELECT * FROM transfers
WHERE 
    from_account_id = $1 OR
    to_account_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;

-- name: UpdateTransfer :one
UPDATE transfer 
SET from_account_id = $2, 
to_account_id = $3,
amount = $4,
created_at = $5
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfer WHERE id = $1;