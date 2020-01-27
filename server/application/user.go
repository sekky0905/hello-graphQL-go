package application

import (
	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
	"github.com/sekky0905/hello-graphQL-go/server/repository"
)

// UserApplicationService は、User の ApplicationService。
type UserApplicationService struct {
	Repo *repository.UserRepository
}

// CreateUser は、User を作成する。
func (s UserApplicationService) CreateUser(id, name string) *model.User {
	user := model.NewUser(id, name)
	// 普通だと諸々の処理を行う。
	s.Repo.AddUser(user)
	return user
}
