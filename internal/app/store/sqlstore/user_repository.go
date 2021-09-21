package sqlstore

import (
	"database/sql"

	"github.com/spolyakovs/racing-backend-ifmo/internal/app/model"
	"github.com/spolyakovs/racing-backend-ifmo/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (userRepository *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	createUserQuery := "INSERT INTO users (username, email, encrypted_password) VALUES ($1, $2, $3) RETURNING id"

	return userRepository.store.db.QueryRow(
		createUserQuery,
		user.Username, user.Email, user.EncryptedPassword,
	).Scan(&user.ID)
}

func (userRepository *UserRepository) Find(id int) (*model.User, error) {
	user := &model.User{}

	findUserByIDQuery := "SELECT id, username, email, encrypted_password FROM users WHERE id = $1"
	if err := userRepository.store.db.QueryRow(
		findUserByIDQuery,
		id,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}

	findUserByUsernameQuery := "SELECT id, username, email, encrypted_password FROM users WHERE username = $1"
	if err := userRepository.store.db.QueryRow(
		findUserByUsernameQuery,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	findUserByEmailQuery := "SELECT id, username, email, encrypted_password FROM users WHERE email = $1"
	if err := userRepository.store.db.QueryRow(
		findUserByEmailQuery,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return user, nil
}

func (userRepository *UserRepository) Update(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := user.BeforeCreate(); err != nil {
		return err
	}

	updateUserQuery := "UPDATE users " +
		"SET username = $2, " +
		"email = $3, " +
		"encrypted_password = $4 " +
		"WHERE id = $1"
	countResult, countResultErr := userRepository.store.db.Exec(
		updateUserQuery,
		user.ID,
		user.Username,
		user.Email,
		user.EncryptedPassword,
	)

	if countResultErr != nil {
		return countResultErr
	}

	count, countErr := countResult.RowsAffected()

	if countErr != nil {
		return countErr
	}

	if count == 0 {
		return store.ErrRecordNotFound
	}

	return nil
}

func (userRepository *UserRepository) Delete(id int) error {
	deleteUserQuery := "DELETE FROM users WHERE id = $1"

	countResult, countResultErr := userRepository.store.db.Exec(
		deleteUserQuery,
		id,
	)

	if countResultErr != nil {
		return countResultErr
	}

	count, countErr := countResult.RowsAffected()

	if countErr != nil {
		return countErr
	}

	if count == 0 {
		return store.ErrRecordNotFound
	}

	return nil
}
