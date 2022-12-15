package dataloaders

import (
	"context"
	"encoding/json"
	"graphql/graph/model"
	"strconv"

	"github.com/graph-gophers/dataloader"
)

func (d *DataLoader) GetReviewsByProduct(ctx context.Context, productID string) ([]*model.Review, error) {
	if d.redis.Ping(ctx).Err() == nil {
		var reviews []*model.Review
		res, err := d.redis.Do(ctx, "GET", "reviews:"+productID).Result()
		if err == nil {
			err := json.Unmarshal([]byte(res.(string)), &reviews)
			if err == nil {
				return reviews, nil
			}
		}
	}
	thunk := d.reviewLoaderByProduct.Load(ctx, dataloader.StringKey(productID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	reviews := result.([]*model.Review)
	if d.redis.Ping(ctx).Err() == nil {
		if d.redis.Do(ctx, "GET", "reviews:"+productID).Err() != nil {
			json, _ := json.Marshal(reviews)
			d.redis.Do(ctx, "SET", "reviews:"+productID, json)
		}
	}
	return reviews, nil
}

func (r *dbBatcher) getByProduct(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	productKeys := make([]any, len(keys))
	for i, key := range keys {
		productKeys[i] = key.String()
	}
	query := "SELECT * FROM \"Reviews\" WHERE \"productId\" IN ($1"
	for i := 1; i < len(productKeys); i++ {
		query += ",$" + strconv.Itoa(i+1)
	}
	query += ")"
	print(query + ", ")
	for i := 0; i < len(productKeys); i++ {
		print(productKeys[i].(string) + ", ")
	}
	print("\n")
	rows, err := r.db.Query(query, productKeys...)
	if err != nil {
		return []*dataloader.Result{
			{
				Error: err,
			},
		}
	}
	defer rows.Close()
	reviews := map[string][]*model.Review{}
	for rows.Next() {
		review := &model.Review{}
		err := rows.Scan(&review.ID, &review.ProductID, &review.UserID, &review.Name, &review.Description, &review.Rating)
		if err != nil {
			return []*dataloader.Result{
				{
					Error: err,
				},
			}
		}
		reviews[review.ProductID] = append(reviews[review.ProductID], review)
	}
	results := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		if reviews[key.String()] != nil {
			results[i] = &dataloader.Result{
				Data: reviews[key.String()],
			}
		} else {
			results[i] = &dataloader.Result{
				Data: []*model.Review{},
			}
		}
	}
	return results
}

func (d *DataLoader) GetReviewsByUser(ctx context.Context, userId string) ([]*model.Review, error) {
	if d.redis.Ping(ctx).Err() == nil {
		var reviews []*model.Review
		res, err := d.redis.Do(ctx, "GET", "reviews:"+userId).Result()
		if err == nil {
			err := json.Unmarshal([]byte(res.(string)), &reviews)
			if err == nil {
				return reviews, nil
			}
		}
	}
	thunk := d.reviewLoaderByProduct.Load(ctx, dataloader.StringKey(userId))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	reviews := result.([]*model.Review)
	if d.redis.Ping(ctx).Err() == nil {
		if d.redis.Do(ctx, "GET", "reviews:"+userId).Err() != nil {
			json, _ := json.Marshal(reviews)
			d.redis.Do(ctx, "SET", "reviews:"+userId, json)
		}
	}
	return reviews, nil
}

func (r *dbBatcher) getByUser(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	productKeys := make([]any, len(keys))
	for i, key := range keys {
		productKeys[i] = key.String()
	}
	query := "SELECT * FROM \"Reviews\" WHERE \"profileId\" IN ($1"
	for i := 1; i < len(productKeys); i++ {
		query += ",$" + strconv.Itoa(i+1)
	}
	query += ")"
	print(query + ", ")
	for i := 0; i < len(productKeys); i++ {
		print(productKeys[i].(string) + ", ")
	}
	print("\n")
	rows, err := r.db.Query(query, productKeys...)
	if err != nil {
		return []*dataloader.Result{
			{
				Error: err,
			},
		}
	}
	defer rows.Close()
	reviews := map[string][]*model.Review{}
	for rows.Next() {
		review := &model.Review{}
		err := rows.Scan(&review.ID, &review.ProductID, &review.UserID, &review.Name, &review.Description, &review.Rating)
		if err != nil {
			return []*dataloader.Result{
				{
					Error: err,
				},
			}
		}
		reviews[review.UserID] = append(reviews[review.UserID], review)
	}
	results := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		if reviews[key.String()] != nil {
			results[i] = &dataloader.Result{
				Data: reviews[key.String()],
			}
		} else {
			results[i] = &dataloader.Result{
				Data: []*model.Review{},
			}
		}
	}
	return results
}
