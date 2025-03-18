package user_test

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcrypto5Hash(t *testing.T) {
	b, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	t.Log(string(b))

	err := bcrypt.CompareHashAndPassword(b, []byte("admin123"))
	if err != nil {
		t.Log(err)
	}
}
