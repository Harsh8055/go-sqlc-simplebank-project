package db

import (
	"context"
	"database/sql"

	"testing"

	"github.com/Harsh8055/simplebank/util"
	"github.com/stretchr/testify/require"
)



func TestInsertEntry(t *testing.T) {
    
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: sql.NullInt64{Int64: account.ID, Valid: true},
		Amount: util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

}



