package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/ch3nnn/webstack-go/internal/graph/generated"
	"github.com/ch3nnn/webstack-go/internal/graph/model"
)

type Resolver struct{}

func (r *mutationResolver) UpdateUserMobile(ctx context.Context, data model.UpdateUserMobileInput) (*model.User, error) {
	panic("not implemented")
}

func (r *queryResolver) BySex(ctx context.Context, sex string) ([]*model.User, error) {
	panic("not implemented")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
