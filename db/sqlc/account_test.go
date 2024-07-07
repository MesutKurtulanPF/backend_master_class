package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"backend_master_class/db/util"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg) // CreateAccount functios returns Account and error
	require.NoError(t, err)                                              // NoError asserts that a function returned no error (i.e. `nil`).
	require.NotEmpty(t, account)                                         // NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either

	require.Equal(t, arg.Owner, account.Owner)       // Equal asserts that two objects are equal.
	require.Equal(t, arg.Balance, account.Balance)   // Equal asserts that two objects are equal.
	require.Equal(t, arg.Currency, account.Currency) // Equal asserts that two objects are equal.

	require.NotZero(t, account.ID)        // NotZero asserts that i is not the zero value for its type.
	require.NotZero(t, account.CreatedAt) // NotZero asserts that i is not the zero value for its type.

	return account
}

func TestCreateAccount(t *testing.T) { // Every unit test function in Go must start with the Test prefix (with uppercase letter T) and takes a testing.T object as input. We will use this T object to manage the test state.
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second) // WithinDuration asserts that the two times are within duration delta of each other.
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)                             // Error asserts that a function returned an error (i.e. not `nil`).
	require.EqualError(t, err, sql.ErrNoRows.Error()) // EqualError asserts that a function returned an error (i.e. not `nil`) and that it is equal to the provided error.
	require.Empty(t, account2)                        // Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
