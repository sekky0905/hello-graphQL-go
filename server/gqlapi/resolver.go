package gqlapi

import (
	"context"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) ChatRoom() ChatRoomResolver {
	return &chatRoomResolver{r}
}
func (r *Resolver) Comment() CommentResolver {
	return &commentResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

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

type commentResolver struct{ *Resolver }

func (r *commentResolver) ID(ctx context.Context, obj *model.Comment) (string, error) {
	panic("not implemented")
}
func (r *commentResolver) PostedBy(ctx context.Context, obj *model.Comment) (*model.User, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*model.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateChatRoom(ctx context.Context, input NewChatRoom) (*model.ChatRoom, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateComment(ctx context.Context, input NewComment) (*model.Comment, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ChatRooms(ctx context.Context) ([]*model.ChatRoom, error) {
	panic("not implemented")
}
func (r *queryResolver) ChatRoom(ctx context.Context, chatRoomID string) (*model.ChatRoom, error) {
	panic("not implemented")
}

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
