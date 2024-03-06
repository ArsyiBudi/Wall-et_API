package db

import (
    "testing"
    "context"
    "github.com/techschool/simplebank/util"
    "github.com/stretchr/testify/require"
    "time"
    "database/sql"  
)

func CreateRandomAccount(t *testing.T) Account {
    arg := CreateAccountParams{
        Owner:    util.RandomOwner(),
        Balance:  util.RandomMoney(), //saldo
        Currency: util.RandomCurrency(),
    }

    account, err := testQuerries.CreateAccount(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, account)

    require.Equal(t, arg.Owner, account.Owner)
    require.Equal(t, arg.Balance, account.Balance)
    require.Equal(t, arg.Currency, account.Currency)

    require.NotZero(t, account.ID)
    require.NotZero(t, account.CreatedAt)

    return account
}

func TestCreateAccount(t *testing.T) {
    CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
    account1 := CreateRandomAccount(t)
    account2, err := testQuerries.GetAccount(context.Background(), account1.ID)
    require.NoError(t, err)
    require.NotEmpty(t, account2)

    require.Equal(t, account1.ID, account2.ID)
    require.Equal(t, account1.Owner, account2.Owner)
    require.Equal(t, account1.Balance, account2.Balance)
    require.Equal(t, account1.Currency, account2.Currency)
    require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
    account1 := CreateRandomAccount(t)

    updateArg := UpdateAccountParams{
        ID:      account1.ID,
        Balance: util.RandomMoney(),
    }

    account2, err := testQuerries.UpdateAccount(context.Background(), updateArg)
    require.NoError(t, err)
    require.NotEmpty(t, account2)

    require.Equal(t, account1.ID, account2.ID)
    require.Equal(t, account1.Owner, account2.Owner)
    require.Equal(t, updateArg.Balance, account2.Balance)
    require.NotEqual(t, account1.Balance, account2.Balance)
    require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// func TestDeleteAccount(t *testing.T) {
//     account1 := CreateRandomAccount(t)
//     err := testQuerries.DeleteAccount(context.Background(), account1.ID)
//     require.NoError(t, err)

//     account2, err := testQuerries.GetAccount(context.Background(), account1.ID)
//     require.NoError(t, err)
//     require.EqualError(t, err, sql.ErrNoRows.Error())
//     require.Empty(t, account2)
// }

func TestDeleteAccount(t *testing.T) {
    account1 := CreateRandomAccount(t)
    
    // Menghapus akun dengan ID yang valid
    err := testQuerries.DeleteAccount(context.Background(), account1.ID)
    require.NoError(t, err)

    // Memastikan akun telah dihapus
    _, err = testQuerries.GetAccount(context.Background(), account1.ID)
    require.Error(t, err)
    require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListAccount(t *testing.T) {
    for i := 0; i < 10; i++ {
        CreateRandomAccount(t)
    } 

    arg := ListAccountParams{
        Limit: 5,
        Offset: 5,
    }

    account, err := testQuerries.ListAccount(context.Background(), arg)
    require.NoError(t, err)
    require.Len(t, account, 5)

    for _, account := range account {
        require.NotEmpty(t, account)
    }
}