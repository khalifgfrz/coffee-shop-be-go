package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "testpassword"

	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err, "Error hashing password")
	assert.NotEmpty(t, hashedPassword, "Hashed password should not be empty")

	err = VerifyPassword(hashedPassword, password)
	assert.NoError(t, err, "Error verifying password")
}

func TestVerifyPassword(t *testing.T) {
	password := "testpassword"
	wrongPassword := "wrongpassword"

	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err, "Error hashing password")

	err = VerifyPassword(hashedPassword, password)
	assert.NoError(t, err, "Password should be verified successfully")

	err = VerifyPassword(hashedPassword, wrongPassword)
	assert.Error(t, err, "Verification should fail with the wrong password")
}