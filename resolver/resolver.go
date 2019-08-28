package resolver

import (
	"context"
	"database/sql"
	"time"

	"github.com/crypto-papers/api/db"
	"github.com/crypto-papers/api/errors"
	"github.com/crypto-papers/api/generated"
	"github.com/crypto-papers/api/model"
)

// Resolver creates collections of items
type Resolver struct {
	// authors    []*model.Author
	currencies []*model.Currency
	files      []*model.File
	papers     []*model.Paper
	users      []*model.User
	db         *sql.DB
}

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

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAuthor(ctx context.Context, data model.CreateAuthorInput) (*model.Author, error) {
	author := &model.Author{
		Name:      data.Name,
		Psuedonym: data.Psuedonym,
		CreateAt:  time.Now(),
	}

	rows, err := db.LogAndQuery(
		r.db,
		"INSERT INTO authors (name, psuedonym, created_at) VALUES ($1, $2, $3) RETURNING id",
		data.Name, data.Psuedonym, author.CreateAt,
	)

	if err != nil || !rows.Next() {
		return author, err
	}

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

	rows, err := db.LogAndQuery(
		r.db,
		"INSERT INTO currencies (name, ticker, created_at) VALUES ($1, $2, $3) RETURNING id",
		data.Name, data.Ticker, currency.CreateAt,
	)

	if err != nil || !rows.Next() {
		return currency, err
	}

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

	rows, err := db.LogAndQuery(
		r.db,
		"INSERT INTO files (coverimage, source, url, version, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		data.CoverImage, data.Source, data.URL, data.Version, file.CreateAt,
	)

	if err != nil || !rows.Next() {
		return file, err
	}

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

	rows, err := db.LogAndQuery(
		r.db,
		"INSERT INTO papers (title, description, excerpt, page_num, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		data.Title, data.Description, data.Excerpt, data.PageNum, paper.CreateAt,
	)

	if err != nil || !rows.Next() {
		return paper, err
	}

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

	rows, err := db.LogAndQuery(
		r.db,
		"INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
		data.Name, data.Email, data.Password, user.CreateAt,
	)

	if err != nil || !rows.Next() {
		return user, err
	}

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
	var author *model.Author
	var authors []*model.Author

	rows, err := db.LogAndQuery(r.db, "SELECT id, name, psuedonym, created_at FROM authors")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		if err := rows.Scan(&author.ID, &author.Name, &author.Psuedonym, &author.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
		authors = append(authors, author)
	}

	return authors, nil
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
