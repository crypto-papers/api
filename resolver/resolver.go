package resolver

import (
	"context"

	"github.com/crypto-papers/api/generated"
	"github.com/crypto-papers/api/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAuthor(ctx context.Context, name string) (*model.Author, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model.Author, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateAuthor(ctx context.Context, name string, psuedonym *bool) (*model.Author, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateCurrency(ctx context.Context, name string, ticker string) (*model.Currency, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteCurrency(ctx context.Context, id string) (*model.Currency, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateCurrency(ctx context.Context, name string, ticker string) (*model.Currency, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateFile(ctx context.Context, url string) (*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteFile(ctx context.Context, id string) (*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateFile(ctx context.Context, coverImage *string, source *string, url *string, version *float64) (*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePaper(ctx context.Context, title string, description *string, excerpt *string, pageNum *int) (*model.Paper, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePaper(ctx context.Context, id string) (*model.Paper, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePaper(ctx context.Context, title *string, description *string, excerpt *string, pageNum *int) (*model.Paper, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string, password string) (*model.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(ctx context.Context, name *string, email *string, password *string) (*model.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Author(ctx context.Context, id string) (*model.Author, error) {
	panic("not implemented")
}
func (r *queryResolver) Currency(ctx context.Context, id string) (*model.Currency, error) {
	panic("not implemented")
}
func (r *queryResolver) File(ctx context.Context, id string) (*model.File, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Paper(ctx context.Context, id string) (*model.Paper, error) {
	panic("not implemented")
}
