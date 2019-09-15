package resolver

import (
	"database/sql"

	"github.com/crypto-papers/api/generated"
)

// Resolver creates collections of items
type Resolver struct {
	db *sql.DB
}

// NewRootResolvers applies the Config struct to the resolvers
func NewRootResolvers(db *sql.DB) generated.Config {
	c := generated.Config{
		Resolvers: &Resolver{
			db: db,
		},
	}

	return c
}

// Mutation executes GraphQL mutations
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query executes GraphQL queries
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
