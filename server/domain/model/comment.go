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

// NewComment は、コメントを表す。
func NewComment(ID string, content string, postedBy User) *Comment {
	return &Comment{
		ID:        ID,
		Content:   content,
		PostedBy:  postedBy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
