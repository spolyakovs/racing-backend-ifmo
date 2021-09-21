package model_test

import (
	"testing"

	"github.com/spolyakovs/racing-backend-ifmo/internal/app/model"
	"github.com/stretchr/testify/assert"
)

// TODO: add tests for username
func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			user: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "with encrypted password",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Password = ""
				user.EncryptedPassword = "encryptedpassword"

				return user
			},
			isValid: true,
		},
		{
			name: "empty email",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Email = ""

				return user
			},
			isValid: false,
		},
		{
			name: "invalid email",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Email = "invalid"

				return user
			},
			isValid: false,
		},
		{
			name: "empty password",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Password = ""

				return user
			},
			isValid: false,
		},
		{
			name: "short password",
			user: func() *model.User {
				user := model.TestUser(t)
				user.Password = "short"

				return user
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.user().Validate())
			} else {
				assert.Error(t, tc.user().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	user := model.TestUser(t)
	assert.NoError(t, user.BeforeCreate())
	assert.NotEmpty(t, user.EncryptedPassword)
}
