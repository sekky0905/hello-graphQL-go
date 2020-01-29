package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

// Comment は、コメントを表す。
type Comment struct {
	ID         string
	UserID     string
	ChatRoomID string
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// NewComment は、コメントを表す。
func NewComment(userID, chatRoomID, content string) (*Comment, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, xerrors.Errorf("failed to generate uuid: %w", err)
	}

	return &Comment{
		ID:         fmt.Sprintf("comment_%s", uid.String()),
		UserID:     userID,
		ChatRoomID: chatRoomID,
		Content:    content,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
