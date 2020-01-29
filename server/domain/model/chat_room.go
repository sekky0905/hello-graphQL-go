package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

// ChatRoom は、チャットルームを表す。
type ChatRoom struct {
	ID        string
	UserID    string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewChatRoom は、ChatRoom を生成し、返す。
func NewChatRoom(userID, title string) (*ChatRoom, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, xerrors.Errorf("failed to generate uuid: %w", err)
	}

	return &ChatRoom{
		ID:        fmt.Sprintf("chat_room_%s", uid.String()),
		UserID:    userID,
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
