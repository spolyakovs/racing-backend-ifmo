package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Username: "username",
		Email:    "user@example.org",
		Password: "password",
	}
}
