package domain

import (
	"context"
	"time"
)

type RoomType string

const (
	RoomTypeDirect  RoomType = "direct"
	RoomTypeGroup   RoomType = "group"
	RoomTypeChannel RoomType = "channel"
)

type Room struct {
	ID          string
	Name        string
	Description string
	Type        RoomType
	CreatorID   string
	MemberIDs   []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RoomMember struct {
	RoomID   string
	UserID   string
	JoinedAt time.Time
	Role     string
}

type RoomRepository interface {
	Create(ctx context.Context, room *Room) error
	GetById(ctx context.Context, id string) (*Room, error)
	ListByUser(ctx context.Context, userID string, limit, offset int) ([]*Room, int, error)
	Update(ctx context.Context, room *Room) error
	Delete(ctx context.Context, id string) error

	AddMember(ctx context.Context, roomID, userID string) error
	RemoveMember(ctx context.Context, roomID, userID string) error
	GetMembers(ctx context.Context, roomID string) ([]string, error)
	isMember(ctx context.Context, roomID, userID string) (bool, error)
}

type RoomUseCase interface {
	CreateRoom(ctx context.Context, name, description string, roomType RoomType, creatorID string, memberIDs []string) (*Room, error)
	GetRoom(ctx context.Context, roomID string) (*Room, error)
	ListRooms(ctx context.Context, userID string, limit, offset int) ([]*Room, int, error)
	JoinRoom(ctx context.Context, roomID, userID string) error
	LeaveRoom(ctx context.Context, roomID, userID string) error
	AddMember(ctx context.Context, roomID, userID, memberID string) error
	RemoveMember(ctx context.Context, roomID, userID, memberID string) error
}
