package gqlapi

import (
	"context"
	"fmt"

	"github.com/sekky0905/hello-graphQL-go/server/domain/model"
)

type commentResolver struct{ *Resolver }

func (r *commentResolver) ID(ctx context.Context, obj *model.Comment) (string, error) {
	if obj == nil {
		return "", nil
	}

	return fmt.Sprintf("Comment:%s", obj.ID), nil
}
func (r *commentResolver) PostedBy(ctx context.Context, obj *model.Comment) (*model.User, error) {
	return r.UserApplicationService.GetUserByID(obj.UserID), nil
}
