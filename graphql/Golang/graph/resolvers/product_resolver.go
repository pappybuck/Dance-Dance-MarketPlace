package resolvers

import (
	"context"
	"graphql/graph/dataloaders"
	"graphql/graph/generated"
	"graphql/graph/model"
)

type productResolver struct{ *Resolver }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Reviews is the resolver for the reviews field.
func (r *productResolver) Reviews(ctx context.Context, obj *model.Product) ([]*model.Review, error) {
	return dataloaders.For(ctx).GetReviewsByProduct(ctx, obj.ID)
}
