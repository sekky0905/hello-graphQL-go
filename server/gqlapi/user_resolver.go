package gqlapi

import (
	"context"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
)

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	panic("not implemented")
}
func (r *userResolver) ChatRooms(ctx context.Context, obj *model.User) ([]*model.ChatRoom, error) {
	panic("not implemented")
}
func (r *userResolver) Comments(ctx context.Context, obj *model.User) ([]*model.Comment, error) {
	panic("not implemented")
}
