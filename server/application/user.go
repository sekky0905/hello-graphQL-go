package application

import (
	"github.com/sekky0905/hello-graphQL-go/server/gqlapi/domain/model"
	"github.com/sekky0905/hello-graphQL-go/server/gqlapi/repository"
)

// UserApplicationService は、User の ApplicationService。
type UserApplicationService struct {
	Repo repository.UserRepository
}

// CreateUser は、User を作成する。
func (s UserApplicationService) CreateUser(user model.User) {
	// 普通だと諸々の処理を行う。
	s.Repo.AddUser(user)
}
