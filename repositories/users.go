package repositories

import (
	"gocred/ports"
)

type UserRepository struct {
	db ports.Database
}

func NewUserRepository(db ports.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(id int) (string, error) {
	var username string
	err := r.db.QueryRow("SELECT username FROM users WHERE id = $1", id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
