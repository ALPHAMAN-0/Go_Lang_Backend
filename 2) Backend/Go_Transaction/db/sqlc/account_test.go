package db

import (
	"context"
	"testing"
	"math/rand"
	"fmt"

	"github.com/stretchr/testify/require"
)

// helper to generate random strings
func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// helper to generate random owner name
func randomOwner() string {
	return randomString(6)
}

// helper to generate random balance
func randomBalance() int64 {
	return rand.Int63n(1000)
}

// helper to generate random currency
func randomCurrency() string {
	currencies := []string{"USD", "EUR", "BDT"}
	return currencies[rand.Intn(len(currencies))]
}

// createRandomAccount creates a test account in the database
func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    randomOwner(),
		Balance:  randomBalance(),
		Currency: randomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	fmt.Printf("Created account: ID=%d Owner=%s Balance=%d\n",
		account.ID, account.Owner, account.Balance)

	return account
}