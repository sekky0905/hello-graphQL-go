package gqlapi

import (
	"context"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
)

type commentResolver struct{ *Resolver }

func (r *commentResolver) ID(ctx context.Context, obj *model.Comment) (string, error) {
	panic("not implemented")
}
func (r *commentResolver) PostedBy(ctx context.Context, obj *model.Comment) (*model.User, error) {
	panic("not implemented")
}
