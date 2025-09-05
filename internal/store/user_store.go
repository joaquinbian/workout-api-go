package store

import (
	"database/sql"
	"time"
)

type password struct {
	text *string
	hash []byte
}

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash password  `json:"-"` //- means ignore the value
	Bio          string    `json:"bio"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db: db,
	}
}

type UserStore interface {
	CreateUser(*User) error
	GetUserByID(id int) (*User, error)
	UpdateUser(*User) error
}
