package resolvers

import (
	"context"
	"graphql/graph/dataloaders"
	"graphql/graph/generated"
	"graphql/graph/model"
)

type userResolver struct{ *Resolver }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// Reviews is the resolver for the reviews field.
func (r *userResolver) Reviews(ctx context.Context, obj *model.User) ([]*model.Review, error) {
	return dataloaders.For(ctx).GetReviewsByUser(ctx, obj.ID)
}
