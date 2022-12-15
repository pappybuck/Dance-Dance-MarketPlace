package main

import (
	"database/sql"
	"fmt"
	"graphql/graph/dataloaders"
	"graphql/graph/generated"
	"graphql/graph/resolvers"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	_ "github.com/lib/pq"
)

var db *sql.DB
var redisClient *redis.Client

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	loader := dataloaders.NewLoaders(db, redisClient)
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		DB:    db,
		Redis: redisClient,
	}}))

	dataloaderSrv := dataloaders.Middleware(loader, h)

	return func(c *gin.Context) {
		dataloaderSrv.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	//Setting up Gin
	// r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.Default())
	dbConnection := ""
	if gin.Mode() == gin.ReleaseMode {
		dbConnection = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require&options=--%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_OPTIONS"))
	} else {
		dbConnection = "postgresql://dance-dance:password@localhost:3000/dance-dance-db?sslmode=disable"
	}
	database, err := sql.Open("postgres", dbConnection)
	if err != nil {
		panic(err)
	}
	db = database
	if _, present := os.LookupEnv("REDIS_URL"); present {
		opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
			panic(err)
		}
		redisClient = redis.NewClient(opt)
	} else {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
	}
	defer db.Close()
	r.POST("/graphql/query", graphqlHandler())
	r.GET("/graphql", playgroundHandler())
	println("Server is running at http://localhost:8080/")
	r.Run()
}
