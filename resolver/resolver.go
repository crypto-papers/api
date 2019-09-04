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
	var author = new(model.Author)

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
	}

	return author, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	var authors []*model.Author

	rows, err := db.LogAndQuery(r.db, "SELECT id, name, psuedonym, created_at FROM authors")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var author = new(model.Author)
		if err := rows.Scan(&author.ID, &author.Name, &author.Psuedonym, &author.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
		authors = append(authors, author)
	}

	return authors, nil
}

func (r *queryResolver) Currency(ctx context.Context, id string) (*model.Currency, error) {
	var currency = new(model.Currency)

	rows, err := db.LogAndQuery(r.db, "SELECT id, name, ticker, created_at FROM currencies")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		if err := rows.Scan(&currency.ID, &currency.Name, &currency.Ticker, &currency.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return currency, nil
}

func (r *queryResolver) Currencies(ctx context.Context) ([]*model.Currency, error) {
	var currencies []*model.Currency

	rows, err := db.LogAndQuery(r.db, "SELECT id, name, ticker, created_at FROM currencies")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var currency = new(model.Currency)
		if err := rows.Scan(&currency.ID, &currency.Name, &currency.Ticker, &currency.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
		currencies = append(currencies, currency)
	}

	return currencies, nil
}

func (r *queryResolver) File(ctx context.Context, id string) (*model.File, error) {
	var file = new(model.File)

	rows, err := db.LogAndQuery(r.db, "SELECT id, coverimage, source, url, created_at FROM files")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		if err := rows.Scan(&file.ID, &file.CoverImage, &file.Source, &file.URL, &file.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return file, nil
}

func (r *queryResolver) Files(ctx context.Context) ([]*model.File, error) {
	var files []*model.File

	rows, err := db.LogAndQuery(r.db, "SELECT id, coverimage, source, url, created_at FROM files")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var file = new(model.File)
		if err := rows.Scan(&file.ID, &file.CoverImage, &file.Source, &file.URL, &file.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
		files = append(files, file)
	}

	return files, nil
}

func (r *queryResolver) Paper(ctx context.Context, id string) (*model.Paper, error) {
	var paper = new(model.Paper)

	rows, err := db.LogAndQuery(r.db, "SELECT id, excerpt, page_num, title, created_at FROM papers")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		if err := rows.Scan(&paper.ID, &paper.Excerpt, &paper.PageNum, &paper.Title, &paper.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return paper, nil
}

func (r *queryResolver) Papers(ctx context.Context) ([]*model.Paper, error) {
	var papers []*model.Paper

	rows, err := db.LogAndQuery(r.db, "SELECT id, excerpt, page_num, title, created_at FROM papers")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var paper = new(model.Paper)
		if err := rows.Scan(&paper.ID, &paper.Excerpt, &paper.PageNum, &paper.Title, &paper.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
		papers = append(papers, paper)
	}

	return papers, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var user = new(model.User)

	rows, err := db.LogAndQuery(r.db, "SELECT id, name, email, created_at FROM users")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	rows, err := db.LogAndQuery(r.db, "SELECT id, name, email, created_at FROM users")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var user = new(model.User)
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
		users = append(users, user)
	}

	return users, nil
}
