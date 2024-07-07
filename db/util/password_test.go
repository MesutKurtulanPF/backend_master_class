package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashed_password1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed_password1)

	err = CheckPassword(password, hashed_password1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashed_password1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashed_password2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed_password2)
	require.NotEqual(t, hashed_password1, hashed_password2)
}
