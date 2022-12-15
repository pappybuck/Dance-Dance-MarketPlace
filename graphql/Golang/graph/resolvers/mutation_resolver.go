package resolvers

import (
	"context"
	"fmt"
	"graphql/graph/generated"
	"graphql/graph/model"
)

type mutationResolver struct{ *Resolver }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// AddProduct is the resolver for the AddProduct field.
func (r *mutationResolver) AddProduct(ctx context.Context, description string, name string, price float64, quantity int) (*model.Product, error) {
	query := `INSERT INTO "Products" (name, description, price, quantity) VALUES ($1, $2, $3, $4) RETURNING id`
	println(query)
	var id string
	err := r.DB.QueryRow(query, name, description, price, quantity).Scan(&id)
	if err == nil {
		product := &model.Product{
			ID:          id,
			Name:        name,
			Description: description,
			Price:       price,
			Quantity:    quantity,
		}

		// if r.Redis.Ping().Err() == nil {

		// }

		return product, nil
	}
	return nil, err
}

// RemoveProduct is the resolver for the RemoveProduct field.
func (r *mutationResolver) RemoveProduct(ctx context.Context, id string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: RemoveProduct - RemoveProduct"))
}

// UpdateProduct is the resolver for the UpdateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, description *string, name *string, price *float64, quantity *int) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: UpdateProduct - UpdateProduct"))
}

// CreateReview is the resolver for the CreateReview field.
func (r *mutationResolver) CreateReview(ctx context.Context, description string, name string, productID string, rating int, userID string) (*model.Review, error) {
	query := `INSERT INTO "Reviews" (name, description, rating, product_id, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	println(query)
	var id string
	err := r.DB.QueryRow(query, name, description, rating, productID, userID).Scan(&id)
	if err == nil {
		review := &model.Review{
			ID:          id,
			Name:        name,
			Description: description,
			Rating:      rating,
			ProductID:   productID,
			UserID:      userID,
		}
		return review, nil
	}
	return nil, err
}

// DeleteReview is the resolver for the DeleteReview field.
func (r *mutationResolver) DeleteReview(ctx context.Context, id string) (*model.Review, error) {
	panic(fmt.Errorf("not implemented: DeleteReview - DeleteReview"))
}

// UpdateReview is the resolver for the UpdateReview field.
func (r *mutationResolver) UpdateReview(ctx context.Context, id string, description *string, name *string, productID *string, rating *int, userID *string) (*model.Review, error) {
	panic(fmt.Errorf("not implemented: UpdateReview - UpdateReview"))
}
