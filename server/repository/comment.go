package repository

import (
	"time"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
)

// CommentDTO は、Comment を表す。
type CommentDTO struct {
	ID         string
	UserID     string
	ChatRoomID string
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// newCommentDTOFromComment は、domain model の Comment から CommentDTO を作成する。
func newCommentDTOFromComment(chatRoomID string, comment *model.Comment) *CommentDTO {
	return &CommentDTO{
		ID:         comment.ID,
		UserID:     comment.UserID,
		ChatRoomID: chatRoomID,
		Content:    comment.Content,
		CreatedAt:  comment.CreatedAt,
		UpdatedAt:  comment.UpdatedAt,
	}
}

// newCommentFromCommentDTO は、CommentDTO から domain model の Comment を作成する。
func newCommentFromCommentDTO(dto *CommentDTO) *model.Comment {
	return &model.Comment{
		ID:         dto.ID,
		UserID:     dto.UserID,
		ChatRoomID: dto.ChatRoomID,
		Content:    dto.Content,
		CreatedAt:  dto.CreatedAt,
		UpdatedAt:  dto.UpdatedAt,
	}
}

// CommentRepository は、Comment の Repository。
// 練習用の Project なので、Interface 等は定義しない。
type CommentRepository struct {
	MockDB []*CommentDTO
}

// GetCommentListByUserID は、指定された UserID を持つ Comment の一覧を取得する。
func (r *CommentRepository) GetCommentListByUserID(userID string) []*model.Comment {
	var comments []*model.Comment
	for _, dto := range r.MockDB {
		if dto.UserID == userID {
			comments = append(comments, newCommentFromCommentDTO(dto))
		}
	}

	return comments
}

// AddComment は、Comment を追加する。
func (r *CommentRepository) AddComment(chatRoomID string, comment *model.Comment) {
	c := newCommentDTOFromComment(chatRoomID, comment)
	r.MockDB = append(r.MockDB, c)
}
