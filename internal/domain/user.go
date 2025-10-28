package domain

import (
	"context"
	"time"
)

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	AvatarURL string
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
	LastSeen  time.Time
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetById(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Search(ctx context.Context, query string, limit, offset int) ([]*User, int, error)
	UpdateLastSeen(ctx context.Context, userID string) error
}

type UserUseCase interface {
	GetUser(ctx context.Context, userID string) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	SearchUsers(ctx context.Context, query string, limit, offset int) ([]*User, int, error)
}
