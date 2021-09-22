package store

import (
	"github.com/spolyakovs/racing-backend-ifmo/internal/app/model"
)

type UserRepository struct {
	DefaultRepository
}

func (userRepository *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	return userRepository.CreateDefault(user, "users", "username", "email", "encrypted_password")
}

func (userRepository *UserRepository) Find(id int) (*model.User, error) {
	mod, err := userRepository.FindByFieldDefault("id", id, "users", "id", "username", "email", "encrypted_password")

	if err != nil {
		return nil, err
	}

	switch m := mod.(type) {
	case *model.User:
		return m, nil
	default:
		return nil, ErrWrongModel
	}
}

func (userRepository *UserRepository) FindByUsername(username string) (*model.User, error) {
	mod, err := userRepository.FindByFieldDefault("username", username, "users", "id", "username", "email", "encrypted_password")

	if err != nil {
		return nil, err
	}

	switch m := mod.(type) {
	case *model.User:
		return m, nil
	default:
		return nil, ErrWrongModel
	}
}

func (userRepository *UserRepository) FindByEmail(email string) (*model.User, error) {
	mod, err := userRepository.FindByFieldDefault("email", email, "users", "id", "username", "email", "encrypted_password")

	if err != nil {
		return nil, err
	}

	switch m := mod.(type) {
	case *model.User:
		return m, nil
	default:
		return nil, ErrWrongModel
	}
}

func (userRepository *UserRepository) Update(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	return userRepository.UpdateDefault(user, "users", "username", "email", "encrypted_password")
}

func (userRepository *UserRepository) Delete(id int) error {
	return userRepository.DeleteDefault("users", id)
}
