package model

import "time"

// User は、ユーザーを表す。
type User struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser は、User を生成し、返す。
func NewUser(id, name string) *User {
	return &User{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
