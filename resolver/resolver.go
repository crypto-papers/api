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

func (r *mutationResolver) CreateAuthor(context.Context, model.CreateAuthorInput) (*model.Author, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model.Author, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateAuthor(context.Context, model.UpdateAuthorInput) (*model.Author, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateCurrency(context.Context, model.CreateCurrencyInput) (*model.Currency, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteCurrency(ctx context.Context, id string) (*model.Currency, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateCurrency(context.Context, model.UpdateCurrencyInput) (*model.Currency, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateFile(context.Context, model.CreateFileInput) (*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteFile(ctx context.Context, id string) (*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateFile(context.Context, model.UpdateFileInput) (*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePaper(context.Context, model.CreatePaperInput) (*model.Paper, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeletePaper(ctx context.Context, id string) (*model.Paper, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePaper(context.Context, model.UpdatePaperInput) (*model.Paper, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUser(context.Context, model.CreateUserInput) (*model.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateUser(context.Context, model.UpdateUserInput) (*model.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Author(ctx context.Context, id string) (*model.Author, error) {
	panic("not implemented")
}
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	panic("not implemented")
}
func (r *queryResolver) Currency(ctx context.Context, id string) (*model.Currency, error) {
	panic("not implemented")
}
func (r *queryResolver) Currencies(ctx context.Context) ([]*model.Currency, error) {
	panic("not implemented")
}
func (r *queryResolver) File(ctx context.Context, id string) (*model.File, error) {
	panic("not implemented")
}
func (r *queryResolver) Files(ctx context.Context) ([]*model.File, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Paper(ctx context.Context, id string) (*model.Paper, error) {
	panic("not implemented")
}
func (r *queryResolver) Papers(ctx context.Context) ([]*model.Paper, error) {
	panic("not implemented")
}
