package gqlapi

import (
	"context"
	"fmt"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
)

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	if obj == nil {
		return "", nil
	}

	return fmt.Sprintf("User:%s", obj.ID), nil
}
func (r *userResolver) ChatRooms(ctx context.Context, obj *model.User) ([]*model.ChatRoom, error) {
	return r.ChatRoomApplicationService.GetChatRoomListByUserID(obj.ID), nil
}
func (r *userResolver) Comments(ctx context.Context, obj *model.User) ([]*model.Comment, error) {
	return r.CommentApplicationService.GetCommentListByUserID(obj.ID), nil
}
