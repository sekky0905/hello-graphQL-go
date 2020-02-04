package application

import (
	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
	"github.com/sekky0905/hello-graphQL-go/server/repository"
	"golang.org/x/xerrors"
)

// CommentApplicationService は、Comment の ApplicationService。
type CommentApplicationService struct {
	Repo *repository.CommentRepository
}

// GetCommentListByUserID は、指定された UserID を持つ Comment の一覧を取得する。
func (s CommentApplicationService) GetCommentListByUserID(userID string) []*model.Comment {
	return s.Repo.GetCommentListByUserID(userID)
}

// GetCommentListByChatRoomID は、指定された ChatRoomID を持つ Comment の一覧を取得する。
func (s CommentApplicationService) GetCommentListByChatRoomID(userID string) []*model.Comment {
	return s.Repo.GetCommentListByChatRoomID(userID)
}

// CreateComment は、Comment を作成する。
func (s CommentApplicationService) CreateComment(userID, chatRoomID, content string) (*model.Comment, error) {
	comment, err := model.NewComment(userID, chatRoomID, content)
	if err != nil {
		return nil, xerrors.Errorf("failed to generate comment: %w", err)
	}
	// 普通だと諸々の処理を行う。
	s.Repo.AddComment(chatRoomID, comment)
	return comment, nil
}
