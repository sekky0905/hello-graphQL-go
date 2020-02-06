package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

// User は、ユーザーを表す。
type User struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser は、User を生成し、返す。
func NewUser(name string) (*User, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, xerrors.Errorf("failed to generate uuid: %w", err)
	}

	return &User{
		ID:        fmt.Sprintf("user_%s", uid.String()),
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
