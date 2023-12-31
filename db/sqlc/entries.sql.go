// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: entries.sql

package db

import (
	"context"
	"database/sql"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (account_id, amount)
VALUES ($1, $2)
RETURNING id, account_id, amount, created_at
`

type CreateEntryParams struct {
	AccountID sql.NullInt64 `json:"account_id"`
	Amount    int64         `json:"amount"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.AccountID, arg.Amount)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const selectAllEntries = `-- name: SelectAllEntries :many
SELECT id, account_id, amount, created_at FROM entries 
ORDER BY id
`

func (q *Queries) SelectAllEntries(ctx context.Context) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, selectAllEntries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectEntriesByAccountID = `-- name: SelectEntriesByAccountID :many
SELECT id, account_id, amount, created_at FROM entries WHERE account_id = $1
ORDER BY id
`

func (q *Queries) SelectEntriesByAccountID(ctx context.Context, accountID sql.NullInt64) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, selectEntriesByAccountID, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
