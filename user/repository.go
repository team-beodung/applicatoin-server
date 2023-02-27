package user

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	GetUserByID(id int64) (*User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB

	logger *logrus.Logger
}

func NewUserRepository(db *sql.DB, logger *logrus.Logger) UserRepository {
	return &UserRepositoryImpl{db, logger}
}

func (r *UserRepositoryImpl) GetUserByID(id int64) (*User, error) {
	row := r.db.QueryRow("SELECT * FROM users where id = ?", id)
	var user *User
	row.Scan(&user)

	return user, nil
}

type User struct {
	ID       int64
	Name     string
	Email    string
	password string
}
