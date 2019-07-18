package resolver

import (
	"context"
	"time"

	"github.com/crypto-papers/api/generated"
	"github.com/crypto-papers/api/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver creates collections of items
type Resolver struct {
	authors    []*model.Author
	currencies []*model.Currency
	files      []*model.File
	papers     []*model.Paper
	users      []*model.User
}

// Mutation executes GraphQL mutations
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query executes GraphQL queries
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAuthor(ctx context.Context, data model.CreateAuthorInput) (*model.Author, error) {
	author := &model.Author{
		Name:      data.Name,
		Psuedonym: data.Psuedonym,
		CreateAt:  time.Now(),
	}
	r.authors = append(r.authors, author)
	return author, nil
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id string) (*model.Author, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateAuthor(context.Context, model.UpdateAuthorInput) (*model.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateCurrency(ctx context.Context, data model.CreateCurrencyInput) (*model.Currency, error) {
	currency := &model.Currency{
		Name:     data.Name,
		Ticker:   data.Ticker,
		CreateAt: time.Now(),
	}
	r.currencies = append(r.currencies, currency)
	return currency, nil
}

func (r *mutationResolver) DeleteCurrency(ctx context.Context, id string) (*model.Currency, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateCurrency(context.Context, model.UpdateCurrencyInput) (*model.Currency, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateFile(ctx context.Context, data model.CreateFileInput) (*model.File, error) {
	file := &model.File{
		CoverImage: data.CoverImage,
		Source:     data.Source,
		URL:        data.URL,
		Version:    data.Version,
		CreateAt:   time.Now(),
	}
	r.files = append(r.files, file)
	return file, nil
}

func (r *mutationResolver) DeleteFile(ctx context.Context, id string) (*model.File, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateFile(context.Context, model.UpdateFileInput) (*model.File, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreatePaper(ctx context.Context, data model.CreatePaperInput) (*model.Paper, error) {
	paper := &model.Paper{
		Title:       data.Title,
		Description: data.Description,
		Excerpt:     data.Excerpt,
		PageNum:     data.PageNum,
		CreateAt:    time.Now(),
	}
	r.papers = append(r.papers, paper)
	return paper, nil
}

func (r *mutationResolver) DeletePaper(ctx context.Context, id string) (*model.Paper, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdatePaper(context.Context, model.UpdatePaperInput) (*model.Paper, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateUser(ctx context.Context, data model.CreateUserInput) (*model.User, error) {
	user := &model.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		CreateAt: time.Now(),
	}
	r.users = append(r.users, user)
	return user, nil
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
	return r.authors, nil
}
func (r *queryResolver) Currency(ctx context.Context, id string) (*model.Currency, error) {
	panic("not implemented")
}
func (r *queryResolver) Currencies(ctx context.Context) ([]*model.Currency, error) {
	return r.currencies, nil
}
func (r *queryResolver) File(ctx context.Context, id string) (*model.File, error) {
	panic("not implemented")
}
func (r *queryResolver) Files(ctx context.Context) ([]*model.File, error) {
	return r.files, nil
}
func (r *queryResolver) Paper(ctx context.Context, id string) (*model.Paper, error) {
	panic("not implemented")
}
func (r *queryResolver) Papers(ctx context.Context) ([]*model.Paper, error) {
	return r.papers, nil
}
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}
