package repository

import (
	"time"

	"github.com/sekky0905/hello-graphQL-go/server/gqlapi/domain/model"
)

// UserDTO は、User を表す。
type UserDTO struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// newUserDTOFromUser は、domain model の User から UserDTO を作成する。
func newUserDTOFromUser(user model.User) *UserDTO {
	return &UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// UserRepository は、User の Repository。
// 練習用の Project なので、Interface 等は定義しない。
type UserRepository struct {
	MockDB []*UserDTO
}

// AddUser は、User を追加する。
func (r *UserRepository) AddUser(user model.User) {
	u := newUserDTOFromUser(user)
	r.MockDB = append(r.MockDB, u)
}
