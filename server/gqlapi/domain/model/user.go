package model

import "time"

// User は、ユーザーを表す。
type User struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
