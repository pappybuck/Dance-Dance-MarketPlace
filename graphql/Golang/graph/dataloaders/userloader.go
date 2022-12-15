package dataloaders

import (
	"context"
	"graphql/graph/model"
	"strconv"

	"github.com/graph-gophers/dataloader"
)

func (d *DataLoader) GetUserByReview(ctx context.Context, reviewID string) (*model.User, error) {
	thunk := d.userLoaderByReview.Load(ctx, dataloader.StringKey(reviewID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	user := result.(*model.User)
	return user, nil
}

func (d *dbBatcher) getUserByReview(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	reviewKeys := make([]any, len(keys))
	for i, key := range keys {
		reviewKeys[i] = key.String()
	}
	query := "SELECT \"ProfileId\", \"FirstName\", \"LastName\", \"Email\", \"Phone\" FROM \"Profiles\" WHERE \"ProfileId\" IN ($1"
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
	users := map[string]*model.User{}
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
		if err != nil {
			return []*dataloader.Result{
				{
					Error: err,
				},
			}
		}
		users[user.ID] = user
	}
	results := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		if users[key.String()] != nil {
			results[i] = &dataloader.Result{
				Data: users[key.String()],
			}
		} else {
			results[i] = &dataloader.Result{
				Data: &model.User{},
			}
		}
	}
	return results
}
