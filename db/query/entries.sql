-- name: SelectAllEntries :many
SELECT * FROM entries 
ORDER BY id;

-- name: CreateEntry :one
INSERT INTO entries (account_id, amount)
VALUES ($1, $2)
RETURNING *;


-- name: SelectEntriesByAccountID :many
SELECT * FROM entries WHERE account_id = $1
ORDER BY id;
