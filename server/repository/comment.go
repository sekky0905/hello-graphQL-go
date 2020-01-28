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
		UserID:     comment.PostedBy.ID,
		ChatRoomID: chatRoomID,
		Content:    comment.Content,
		CreatedAt:  comment.CreatedAt,
		UpdatedAt:  comment.UpdatedAt,
	}
}

// CommentRepository は、Comment の Repository。
// 練習用の Project なので、Interface 等は定義しない。
type CommentRepository struct {
	MockDB []*CommentDTO
}

// AddComment は、Comment を追加する。
func (r *CommentRepository) AddComment(chatRoomID string, comment *model.Comment) {
	c := newCommentDTOFromComment(chatRoomID, comment)
	r.MockDB = append(r.MockDB, c)
}
