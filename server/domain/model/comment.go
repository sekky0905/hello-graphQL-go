package model

import "time"

// Comment は、コメントを表す。
type Comment struct {
	ID        string
	Content   string
	PostedBy  User
	CreatedAt time.Time
	UpdatedAt time.Time
}
