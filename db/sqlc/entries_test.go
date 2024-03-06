package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func CreateRandomEntries(t *testing.T, account Account) Entry {
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQuerries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntries(t *testing.T) {
	account := CreateRandomAccount(t)
	CreateRandomEntries(t, account)
}

func TestGetEntries(t *testing.T) {
	account := CreateRandomAccount(t)
	entry1 := CreateRandomEntries(t, account)

	entry2, err := testQuerries.GetEntries(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {

	account := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		CreateRandomEntries(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQuerries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, transfer := range entries {
		require.NotEmpty(t, transfer)
		// require.Equal(t, arg.AccountID, account.ID)
	}
}
