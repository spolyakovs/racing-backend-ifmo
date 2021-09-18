package store_test

import (
	"testing"

	"github.com/spolyakovs/racing-backend-ifmo/internal/model"
	"github.com/spolyakovs/racing-backend-ifmo/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	store, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	user, err := store.User().Create(&model.User{
		Username: "usernameExample",
		Email:    "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	store, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	username := "usernameExample"
	_, err := store.User().FindByUsername(username)
	assert.Error(t, err)

	store.User().Create(&model.User{
		Username: username,
	})
	user, err := store.User().FindByUsername(username)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
