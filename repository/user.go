package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
	FetchByID(id int) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	quer := "INSERT INTO users (username, password) VALUES ($1, $2)"

	_, err := u.db.Exec(quer, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *userRepository) CheckAvail(user model.User) error {
	row := u.db.QueryRow("SELECT id FROM users WHERE username = $1 and password = $2", user.Username, user.Password)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *userRepository) FetchByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT id, username, password FROM users WHERE id = $1", id)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
