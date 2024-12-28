package db

import (
	"context"
	"testing"
	"time"

	util "github.com/practice/simple_bank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
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
	Account := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), Account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, Account.ID, account2.ID)
	require.Equal(t, Account.Owner, account2.Owner)
	require.Equal(t, Account.Balance, account2.Balance)
	require.Equal(t, Account.Currency, account2.Currency)
	require.WithinDuration(t, Account.CreatedAt, account2.CreatedAt, time.Second)

}

func TestListAccount(t *testing.T) {
	
	for i := 0; i < 10; i++ {
		 CreateRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestUpdateAccount(t *testing.T) {
	Account := CreateRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      Account.ID,
		Balance: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, Account.ID, account2.ID)
	require.Equal(t, Account.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, Account.Currency, account2.Currency)
	require.WithinDuration(t, Account.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	Account := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), Account.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), Account.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}
