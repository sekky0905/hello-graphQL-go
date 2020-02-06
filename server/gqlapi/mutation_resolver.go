package gqlapi

import (
	"context"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*model.User, error) {
	return r.UserApplicationService.CreateUser(input.Name)
}
func (r *mutationResolver) CreateChatRoom(ctx context.Context, input NewChatRoom) (*model.ChatRoom, error) {
	return r.ChatRoomApplicationService.CreateChatRoom(input.UserID, input.Title)
}
func (r *mutationResolver) CreateComment(ctx context.Context, input NewComment) (*model.Comment, error) {
	return r.CommentApplicationService.CreateComment(input.UserID, input.ChatRoomID, input.Content)
}
