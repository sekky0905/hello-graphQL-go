package gqlapi

import (
	"context"

	"github.com/sekky0905/hello-graphQL-go/server/repository"

	"github.com/sekky0905/hello-graphQL-go/server/application"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
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

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*model.User, error) {
	return r.UserApplicationService.CreateUser(input.ID, input.Name), nil
}

func (r *mutationResolver) CreateChatRoom(ctx context.Context, input NewChatRoom) (*model.ChatRoom, error) {
	return r.ChatRoomApplicationService.CreateChatRoom(input.UserID, input.Title)
}
func (r *mutationResolver) CreateComment(ctx context.Context, input NewComment) (*model.Comment, error) {
	return r.CommentApplicationService.CreateComment(input.UserID, input.ChatRoomID, input.Content)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ChatRooms(ctx context.Context) ([]*ChatRoom, error) {
	panic("not implemented")
}
func (r *queryResolver) ChatRoom(ctx context.Context, chatRoomID string) (*ChatRoom, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	panic("not implemented")
}
