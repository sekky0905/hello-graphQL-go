package gqlapi

import (
	"context"
	"fmt"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
)

type chatRoomResolver struct{ *Resolver }

func (r *chatRoomResolver) ID(ctx context.Context, obj *model.ChatRoom) (string, error) {
	if obj == nil {
		return "", nil
	}

	return fmt.Sprintf("ChatRoom:%s", obj.ID), nil
}
func (r *chatRoomResolver) Comments(ctx context.Context, obj *model.ChatRoom) ([]*model.Comment, error) {
	return r.CommentApplicationService.GetCommentListByChatRoomID(obj.ID), nil
}
func (r *chatRoomResolver) CreatedBy(ctx context.Context, obj *model.ChatRoom) (*model.User, error) {
	return r.UserApplicationService.GetUserByID(obj.UserID), nil
}
