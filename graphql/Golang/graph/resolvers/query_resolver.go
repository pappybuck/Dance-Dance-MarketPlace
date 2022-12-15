package resolvers

import (
	"context"
	"encoding/json"
	"graphql/graph/generated"
	"graphql/graph/model"
)

type queryResolver struct{ *Resolver }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Product is the resolver for the Product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	println("Product: " + id)
	if r.Redis.Ping(ctx).Err() == nil {
		var product []*model.Product
		res, err := r.Redis.Do(ctx, "GET", id).Result()
		if err == nil {
			err := json.Unmarshal([]byte(res.(string)), &product)
			if err == nil {
				return product[0], nil
			}
		}
	}
	query := "SELECT * FROM \"Products\" WHERE \"id\" = $1"
	println(query + ", " + id)
	row, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	product := &model.Product{}
	for row.Next() {
		err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			return nil, err
		}
	}
	if r.Redis.Ping(ctx).Err() == nil {
		json, _ := json.Marshal(product)
		r.Redis.Do(ctx, "SET", id, json)
	}
	return product, nil
}

// Products is the resolver for the Products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	query := "SELECT * FROM \"Products\""
	println(query)
	row, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	products := make([]*model.Product, 0)
	for row.Next() {
		product := &model.Product{}
		err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	for _, product := range products {
		if r.Redis.Ping(ctx).Err() == nil {
			if r.Redis.Do(ctx, "GET", product.ID).Err() != nil {
				json, _ := json.Marshal(product)
				r.Redis.Do(ctx, "SET", product.ID, json)
			}
		}
	}
	return products, nil
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	if r.Redis.Ping(ctx).Err() == nil {
		var user *model.User
		res, err := r.Redis.Do(ctx, "GET", id).Result()
		if err == nil {
			err := json.Unmarshal([]byte(res.(string)), &user)
			if err == nil {
				return user, nil
			}
		}
	}
	query := "SELECT \"ProfileId\", \"FirstName\", \"LastName\", \"Email\", \"Phone\" FROM \"Profiles\" WHERE \"ProfileId\" = $1"
	println(query + ", " + id)
	row, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	user := &model.User{}
	for row.Next() {
		err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}
	}
	if r.Redis.Ping(ctx).Err() == nil {
		json, _ := json.Marshal(user)
		r.Redis.Do(ctx, "SET", id, json)
	}
	return user, nil
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	query := "SELECT \"ProfileId\", \"FirstName\", \"LastName\", \"Email\", \"Phone\" FROM \"Profiles\""
	println(query)
	row, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	users := make([]*model.User, 0)
	for row.Next() {
		user := &model.User{}
		err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	for _, user := range users {
		if r.Redis.Ping(ctx).Err() == nil {
			if r.Redis.Do(ctx, "GET", user.ID).Err() != nil {
				json, _ := json.Marshal(user)
				r.Redis.Do(ctx, "SET", user.ID, json)
			}
		}
	}
	return users, nil
}
