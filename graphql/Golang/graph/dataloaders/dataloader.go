package dataloaders

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/graph-gophers/dataloader"
)

type ctxKey string

const (
	loadersKey = ctxKey("loaders")
)

type DataLoader struct {
	reviewLoaderByProduct *dataloader.Loader
	reviewLoaderByUser    *dataloader.Loader
	userLoaderByReview    *dataloader.Loader
	productLoaderByReview *dataloader.Loader
	redis                 *redis.Client
}

type dbBatcher struct {
	db *sql.DB
}

func NewLoaders(db *sql.DB, redis *redis.Client) *DataLoader {
	dbReader := &dbBatcher{
		db: db,
	}
	loaders := &DataLoader{
		reviewLoaderByProduct: dataloader.NewBatchedLoader(
			dbReader.getByProduct,
			dataloader.WithBatchCapacity(500),
			dataloader.WithWait(1*time.Millisecond),
			dataloader.WithCache(&dataloader.NoCache{}),
		),
		reviewLoaderByUser: dataloader.NewBatchedLoader(
			dbReader.getByUser,
			dataloader.WithBatchCapacity(500),
			dataloader.WithWait(1*time.Millisecond),
			dataloader.WithCache(&dataloader.NoCache{}),
		),
		userLoaderByReview: dataloader.NewBatchedLoader(
			dbReader.getUserByReview,
			dataloader.WithBatchCapacity(500),
			dataloader.WithWait(1*time.Millisecond),
			dataloader.WithCache(&dataloader.NoCache{}),
		),
		productLoaderByReview: dataloader.NewBatchedLoader(
			dbReader.getProductByReview,
			dataloader.WithBatchCapacity(500),
			dataloader.WithWait(1*time.Millisecond),
			dataloader.WithCache(&dataloader.NoCache{}),
		),
		redis: redis,
	}
	return loaders
}

func Middleware(loader *DataLoader, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loader)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *DataLoader {
	return ctx.Value(loadersKey).(*DataLoader)
}
