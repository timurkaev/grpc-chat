package domain

import (
	"context"
	"time"
)

type MessageType string

const (
	MessaegTypeText   MessageType = "text"
	MessageTypeImage  MessageType = "image"
	MessaegTypeFile   MessageType = "file"
	MessaegTypeSystem MessageType = "system"
)

type Message struct {
	ID             string
	RoomID         string
	SenderID       string
	Content        string
	Type           MessageType
	ReplyToID      string
	AttachmentURLs []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsEdited       bool
	ReadyBy        []string
}

type MessageRepository interface {
	Create(ctx context.Context, msg *Message) error
	GetByID(ctx context.Context, id string) (*Message, error)
	GetByRoom(ctx context.Context, roomID string, limit, offset int, before time.Time) ([]*Message, int, error)
	Update(ctx context.Context, msg *Message) error
	Delete(ctx context.Context, id string) error
	MarkAsRead(ctx context.Context, messageID, userID string) error
}

type MessageUseCase interface {
	SendMessage(ctx context.Context, msg *Message) (*Message, error)
	GetMessage(ctx context.Context, roomID string, limit, offset int, before time.Time) ([]*Message, int, bool, error)
	EditMessage(ctx context.Context, messageID, userID, content string) (*Message, error)
	DeleteMessage(ctx context.Context, messageID, userID string) error
	MarkAsRead(ctx context.Context, roomID, userID, messageID string) error
}
