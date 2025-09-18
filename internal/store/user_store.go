package store

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(plaintextPass string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPass), 12)
	if err != nil {
		return err
	}

	p.text = &plaintextPass
	p.hash = hash
	return nil
}

func (p *password) Matches(plaintextPass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPass))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
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
	GetUserByUsername(username string) (*User, error)
	UpdateUser(*User) error
}

func (s *PostgresUserStore) CreateUser(u *User) error {

	tx, err := s.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `INSERT INTO USERS (username, email, password_hash, bio) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	err = tx.QueryRow(query, u.Username, u.Email, u.PasswordHash.hash, u.Bio).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresUserStore) GetUserByUsername(username string) (*User, error) {

	user := &User{
		PasswordHash: password{},
	}

	query := `SELECT id, username, email, password_hash, bio, created_at, updated_at FROM users WHERE username = $1`
	err := s.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash.hash,
		&user.Bio,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *PostgresUserStore) UpdateUser(u *User) error {

	tx, err := s.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `
	UPDATE USERS 
	SET username = $1, email = $2, bio = $3, updated_at = CURRENT_TIMESTAMP
	WHERE id = $4;
	`
	//ejecuta la query sin devolver filas
	result, err := tx.Exec(query, u.Username, u.Email, u.Bio, u.ID)

	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return sql.ErrNoRows
	}
	err = tx.Commit()

	if err != nil {
		return err
	}
	return nil

}
