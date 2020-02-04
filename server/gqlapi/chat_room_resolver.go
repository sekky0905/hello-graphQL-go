package gqlapi

import (
	"context"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
)

type chatRoomResolver struct{ *Resolver }

func (r *chatRoomResolver) ID(ctx context.Context, obj *model.ChatRoom) (string, error) {
	panic("not implemented")
}
func (r *chatRoomResolver) Comments(ctx context.Context, obj *model.ChatRoom) ([]*model.Comment, error) {
	panic("not implemented")
}
func (r *chatRoomResolver) CreatedBy(ctx context.Context, obj *model.ChatRoom) (*model.User, error) {
	panic("not implemented")
}
