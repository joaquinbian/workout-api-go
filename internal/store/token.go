package store

import (
	"database/sql"
	"time"

	"github.com/joaquinbian/workout-api-go/internal/tokens"
)

type PostgresTokenStore struct {
	db *sql.DB
}

func NewPostgresTokenStore(db *sql.DB) *PostgresTokenStore {
	return &PostgresTokenStore{
		db: db,
	}
}

type TokenStore interface {
	Insert(token *tokens.Token) error
	CreateNewToken(userID int, ttl time.Duration, scope string) (*tokens.Token, error)
	DeleteAllTokensForUser(userID int, scope string) error
}

func (ts *PostgresTokenStore) CreateNewToken(userID int, ttl time.Duration, scope string) (*tokens.Token, error) {
	token, err := tokens.GenerateToken(userID, ttl, scope)

	if err != nil {
		return nil, err
	}

	err = ts.Insert(token)

	return token, nil
}

func (ts *PostgresTokenStore) Insert(token *tokens.Token) error {

	tx, err := ts.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Commit()

	query := `INSERT INTO tokens(hash, user_id, expiry, scope) VALUES($1, $2, $3, $4);`

	_, err = tx.Exec(query, token.Hash, token.UserID, token.Expiry, token.Scope)

	if err != nil {
		return err
	}

	return nil
}

func (ts *PostgresTokenStore) DeleteAllTokensForUser(userID int, scope string) error {
	tx, err := ts.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Commit()
	query := `DELETE FROM hash WHERE user_id = $1 AND scope = $2;`

	res, err := tx.Exec(query, userID, scope)

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
