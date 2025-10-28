package domain

import (
	"context"
	"time"
)

type AuthTokens struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

type RefreshToken struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

type AuthRepository interface {
	CreateRefreshToken(ctx context.Context, token *RefreshToken) error
	GetRefreshToken(ctx context.Context, token string) (*RefreshToken, error)
	DeleteRefreshToken(ctx context.Context, token string) error
	DeleteUserRefreshTokens(ctx context.Context, userID string) error
}

type AuthUseCase interface {
	Register(ctx context.Context, username, email, password string) (*User, *AuthTokens, error)
	Login(ctx context.Context, email, password string) (*User, *AuthTokens, error)
	RefreshToken(ctx context.Context, token string) (*AuthTokens, error)
	ValidateToken(ctx context.Context, token string) (string, error) // returns userID
}
