package dataloaders

import (
	"context"
	"graphql/graph/model"
	"strconv"

	"github.com/graph-gophers/dataloader"
)

func (d *DataLoader) GetProductByReview(ctx context.Context, reviewID string) (*model.Product, error) {
	thunk := d.productLoaderByReview.Load(ctx, dataloader.StringKey(reviewID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	product := result.(*model.Product)
	return product, nil
}

func (d *dbBatcher) getProductByReview(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	reviewKeys := make([]any, len(keys))
	for i, key := range keys {
		reviewKeys[i] = key.String()
	}
	query := "SELECT * FROM \"Products\" WHERE id IN ($1"
	for i := 1; i < len(reviewKeys); i++ {
		query += ",$" + strconv.Itoa(i+1)
	}
	query += ")"
	println(query + ", ")
	for i := 0; i < len(reviewKeys); i++ {
		print(reviewKeys[i].(string) + ", ")
	}
	println("\n")
	rows, err := d.db.Query(query, reviewKeys...)
	if err != nil {
		return []*dataloader.Result{
			{
				Error: err,
			},
		}
	}
	defer rows.Close()
	products := map[string]*model.Product{}
	for rows.Next() {
		product := &model.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity)
		if err != nil {
			return []*dataloader.Result{
				{
					Error: err,
				},
			}
		}
		products[product.ID] = product
	}
	results := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		if products[key.String()] != nil {
			results[i] = &dataloader.Result{
				Data: products[key.String()],
			}
		} else {
			results[i] = &dataloader.Result{
				Data: &model.Product{},
			}
		}
	}
	return results
}
