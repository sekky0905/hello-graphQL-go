package gqlapi

import (
	"context"

	"github.com/sekky0905/hello-graphQL-go/server/application"
	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
	"github.com/sekky0905/hello-graphQL-go/server/repository"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	UserApplicationService     *application.UserApplicationService
	ChatRoomApplicationService *application.ChatRoomApplicationService
	CommentApplicationService  *application.CommentApplicationService
}

// NewResolver は、Resolver を生成し、返す。
func NewResolver() *Resolver {
	return &Resolver{
		UserApplicationService: &application.UserApplicationService{
			Repo: &repository.UserRepository{
				MockDB: make([]*repository.UserDTO, 0),
			},
		},
		ChatRoomApplicationService: &application.ChatRoomApplicationService{
			Repo: &repository.ChatRoomRepository{
				MockDB: make([]*repository.ChatRoomDTO, 0),
			},
		},
		CommentApplicationService: &application.CommentApplicationService{
			Repo: &repository.CommentRepository{
				MockDB: make([]*repository.CommentDTO, 0),
			},
		},
	}
}

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

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.UserApplicationService.GetUserList(), nil
}

func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	return r.UserApplicationService.GetUserByID(userID), nil
}

func (r *queryResolver) ChatRooms(ctx context.Context) ([]*model.ChatRoom, error) {
	return r.ChatRoomApplicationService.GetChatRoomList(), nil
}

func (r *queryResolver) ChatRoomsByUserID(ctx context.Context, userID string) ([]*model.ChatRoom, error) {
	return r.ChatRoomApplicationService.GetChatRoomListByUserID(userID), nil
}
func (r *queryResolver) ChatRoom(ctx context.Context, chatRoomID string) (*model.ChatRoom, error) {
	return r.ChatRoomApplicationService.GetChatRoomByID(chatRoomID), nil
}
func (r *queryResolver) Comments(ctx context.Context, chatRoomID string) ([]*model.Comment, error) {
	return r.CommentApplicationService.GetCommentListByChatRoomID(chatRoomID), nil
}
func (r *queryResolver) CommentsByUserID(ctx context.Context, userID string) ([]*model.Comment, error) {
	return r.CommentApplicationService.GetCommentListByUserID(userID), nil
}
