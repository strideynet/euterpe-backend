package domain

import (
	"golang.org/x/oauth2"
	"time"
)

// User represents a single instance of the Emoji entity
type User struct {
	ID           string
	RefreshToken string
	AuthToken    string
	TokenExpiry  time.Time
}

func (u User) Token() oauth2.Token {
	return oauth2.Token{
		AccessToken:  u.AuthToken,
		RefreshToken: u.RefreshToken,
		Expiry:       u.TokenExpiry,
		TokenType:    "Bearer",
	}
}
