package repository

import (
	"time"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
)

// UserDTO は、User を表す。
type UserDTO struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// newUserDTOFromUser は、domain model の User から UserDTO を作成する。
func newUserDTOFromUser(user *model.User) *UserDTO {
	return &UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// newUserFromUserDTO は、UserDTO から domain model の User を作成する。
func newUserFromUserDTO(dto *UserDTO) *model.User {
	return &model.User{
		ID:        dto.ID,
		Name:      dto.Name,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

// UserRepository は、User の Repository。
// 練習用の Project なので、Interface 等は定義しない。
type UserRepository struct {
	MockDB []*UserDTO
}

// GetUserByID は、指定した ID を持っている User を取得する。
func (r *UserRepository) GetUserByID(id string) *model.User {
	for _, dto := range r.MockDB {
		if dto.ID == id {
			return newUserFromUserDTO(dto)
		}
	}

	return nil
}

// GetUserList は、User の一覧を取得する。
func (r *UserRepository) GetUserList() []*model.User {
	n := len(r.MockDB)
	users := make([]*model.User, n, n)
	for i, dto := range r.MockDB {
		users[i] = newUserFromUserDTO(dto)
	}

	return users
}

// AddUser は、User を追加する。
func (r *UserRepository) AddUser(user *model.User) {
	u := newUserDTOFromUser(user)
	r.MockDB = append(r.MockDB, u)
}
