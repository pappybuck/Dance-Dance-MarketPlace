package resolvers

import (
	"context"
	"graphql/graph/dataloaders"
	"graphql/graph/generated"
	"graphql/graph/model"
)

type reviewResolver struct{ *Resolver }

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

// Product is the resolver for the product field.
func (r *reviewResolver) Product(ctx context.Context, obj *model.Review) (*model.Product, error) {
	return dataloaders.For(ctx).GetProductByReview(ctx, obj.ProductID)
}

// User is the resolver for the user field.
func (r *reviewResolver) User(ctx context.Context, obj *model.Review) (*model.User, error) {
	return dataloaders.For(ctx).GetUserByReview(ctx, obj.UserID)
}
