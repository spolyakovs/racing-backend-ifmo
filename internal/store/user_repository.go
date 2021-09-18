package store

import "github.com/spolyakovs/racing-backend-ifmo/internal/model"

type UserRepository struct {
	store *Store
}

func (userRepository *UserRepository) Create(user *model.User) (*model.User, error) {
	create_user_query := "INSERT INTO users (username, email, encrypted_password) VALUES ($1, $2, $3) RETURNING id"
	if err := userRepository.store.db.QueryRow(
		create_user_query,
		user.Username, user.Email, user.EncryptedPassword,
	).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	if err := userRepository.store.db.QueryRow(
		"SELECT id, username, email, encrypted_password FROM users WHERE username = $1",
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) Update(user *model.User) (*model.User, error) {
	return nil, nil
}
