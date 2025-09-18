package tokens

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"
)

const (
	ScopeAuth = "authentication"
)

type Token struct {
	Plaintext string    `json:"token"`
	UserID    int       `json:"-"`
	Hash      []byte    `json:"-"`
	Scope     string    `json:"-"`
	Expiry    time.Time `json:"expiry"`
}

func GenerateToken(userID int, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	emptyBytes := make([]byte, 32)

	_, err := rand.Read(emptyBytes) //pone en emptyBytes randoms bytes criptograficamente seguros

	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString(emptyBytes)
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	return token, nil
}
