package application

import (
	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
	"github.com/sekky0905/hello-graphQL-go/server/repository"
)

// CommentApplicationService は、Comment の ApplicationService。
type CommentApplicationService struct {
	Repo *repository.CommentRepository
}

// CreateComment は、Comment を作成する。
func (s CommentApplicationService) CreateComment(chatRoomID, id, name string, user model.User) *model.Comment {
	comment := model.NewComment(id, name, user)
	// 普通だと諸々の処理を行う。
	s.Repo.AddComment(chatRoomID, comment)
	return comment
}
