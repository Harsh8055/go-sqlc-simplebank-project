-- name: SelectAllTransfers :many
SELECT * FROM transfers
ORDER BY id;


-- name: InsertTransfer :one
INSERT INTO transfers (from_accout_id, to_account_id, amount, created_at)
VALUES
($1, $2, $3, $4)
RETURNING id;