package model

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Token struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	TokenType string
	IssuedAt  time.Time
	ExpiredAt time.Time
	Duration  string

	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

func NewToken(signingKey []byte, userID uuid.UUID, tokenType string) (*Token, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	oneDay := time.Duration(24 * time.Hour)
	expiredAt := time.Now().Add(oneDay)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "nguji-auth",
		ExpiresAt: jwt.NewNumericDate(expiredAt),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        id.String(),
	})

	ss, err := token.SignedString(signingKey)
	if err != nil {
		return nil, err
	}

	return &Token{
		ID:        id,
		UserID:    userID,
		Token:     ss,
		TokenType: tokenType,
		IssuedAt:  now,
		ExpiredAt: expiredAt,
		Duration:  oneDay.String(),
	}, nil
}

func (t *Token) IsTokenValid(signingKey []byte) (uuid.UUID, error) {
	claim := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(t.Token, claim, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	switch {
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return uuid.Nil, errors.New("token expired")
	case token != nil && token.Valid:
		tokenID, err := uuid.Parse(claim.ID)
		if err != nil {
			return uuid.Nil, nil
		}
		return tokenID, nil
	default:
		return uuid.Nil, errors.New("not valid token")
	}
}

func (t Token) TableName() string {
	return "tokens"
}

func (t Token) Columns() []string {
	return []string{
		"id",
		"user_id",
		"token",
		"token_type",
		"issued_at",
		"expired_at",
		"duration",
	}
}

func (t Token) Data() map[string]any {
	return map[string]any{
		"id":         t.ID,
		"user_id":    t.UserID,
		"token":      t.Token,
		"token_type": t.TokenType,
		"issued_at":  t.IssuedAt,
		"expired_at": t.ExpiredAt,
		"duration":   t.Duration,
	}
}
