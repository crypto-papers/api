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

// Mutation resolvers
type mutationResolver struct{ *Resolver }

// Asset mutation resolvers
func (r *mutationResolver) CreateAsset(ctx context.Context, data model.AssetCreateInput) (*model.Asset, error) {
	asset := &model.Asset{
		Name:     data.Name,
		Ticker:   data.Ticker,
		CreateAt: time.Now(),
	}

	rows, err := db.LogAndQuery(
		r.db,
		"INSERT INTO assets (name, ticker, created_at) VALUES ($1, $2, $3) RETURNING id",
		asset.Name, asset.Ticker, asset.CreateAt,
	)

	if err != nil || !rows.Next() {
		return asset, err
	}

	return asset, nil
}

func (r *mutationResolver) DeleteAsset(ctx context.Context, id model.AssetWhereUniqueInput) (*model.Asset, error) {
	rows, err := db.LogAndQuery(
		r.db,
		"DELETE FROM assets WHERE id = $1",
		id.ID,
	)

	if err != nil || !rows.Next() {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) UpdateAsset(context.Context, model.AssetUpdateInput) (*model.Asset, error) {
	panic("not implemented")
}

// Author mutation resolvers
func (r *mutationResolver) CreateAuthor(ctx context.Context, data model.AuthorCreateInput) (*model.Author, error) {
	author := &model.Author{
		Bio:  data.Bio,
		Name: data.Name,
		// Papers:    data.Papers,
		Photo:     data.Photo,
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

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id model.AuthorWhereUniqueInput) (*model.Author, error) {
	rows, err := db.LogAndQuery(
		r.db,
		"DELETE FROM autors WHERE id = $1",
		id.ID,
	)

	if err != nil || !rows.Next() {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) UpdateAuthor(context.Context, model.AuthorUpdateInput) (*model.Author, error) {
	panic("not implemented")
}

// File mutation resolvers
func (r *mutationResolver) CreateFile(ctx context.Context, data model.FileCreateInput) (*model.File, error) {
	file := &model.File{
		CoverImage: data.CoverImage,
		Source:     data.Source,
		URL:        data.URL,
		Version:    data.Version,
		CreateAt:   time.Now(),
	}

	rows, err := db.LogAndQuery(
		r.db,
		"INSERT INTO files (cover_image, latest, source, url, version, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		file.CoverImage, file.Source, file.URL, file.Version, file.CreateAt,
	)

	if err != nil || !rows.Next() {
		return file, err
	}

	return file, nil
}

func (r *mutationResolver) DeleteFile(ctx context.Context, id model.FileWhereUniqueInput) (*model.File, error) {
	rows, err := db.LogAndQuery(
		r.db,
		"DELETE FROM files WHERE id = $1",
		id.ID,
	)

	if err != nil || !rows.Next() {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) UpdateFile(context.Context, model.FileUpdateInput) (*model.File, error) {
	panic("not implemented")
}

// Paper mutation resolvers
func (r *mutationResolver) CreatePaper(ctx context.Context, data model.PaperCreateInput) (*model.Paper, error) {
	paper := &model.Paper{
		Description: data.Description,
		Excerpt:     data.Excerpt,
		PageNum:     data.PageNum,
		Title:       data.Title,
		CreateAt:    time.Now(),
	}

	rows, err := db.LogAndQuery(
		r.db,
		"INSERT INTO papers (title, description, excerpt, page_num, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		paper.Title, paper.Description, paper.Excerpt, paper.PageNum, paper.CreateAt,
	)

	if err != nil || !rows.Next() {
		return paper, err
	}

	return paper, nil
}

func (r *mutationResolver) DeletePaper(ctx context.Context, id model.PaperWhereUniqueInput) (*model.Paper, error) {
	rows, err := db.LogAndQuery(
		r.db,
		"DELETE FROM papers WHERE id = $1",
		id.ID,
	)

	if err != nil || !rows.Next() {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) UpdatePaper(context.Context, model.PaperUpdateInput) (*model.Paper, error) {
	panic("not implemented")
}

// User mutation resolvers
func (r *mutationResolver) CreateUser(ctx context.Context, data model.UserCreateInput) (*model.User, error) {
	user := &model.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		CreateAt: time.Now(),
	}

	rows, err := db.LogAndQuery(
		r.db,
		"INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
		user.Name, user.Email, user.Password, user.CreateAt,
	)

	if err != nil || !rows.Next() {
		return user, err
	}

	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id model.UserWhereUniqueInput) (*model.User, error) {
	rows, err := db.LogAndQuery(
		r.db,
		"DELETE FROM users WHERE id = $1",
		id.ID,
	)

	if err != nil || !rows.Next() {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) UpdateUser(context.Context, model.UserUpdateInput) (*model.User, error) {
	panic("not implemented")
}

// Query resolvers
type queryResolver struct{ *Resolver }

// Asset query resolvers
func (r *queryResolver) Asset(ctx context.Context, id string) (*model.Asset, error) {
	var asset = new(model.Asset)

	rows, err := db.LogAndQuery(r.db, "SELECT id, name, ticker, created_at FROM assets")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		if err := rows.Scan(&asset.ID, &asset.Name, &asset.Ticker, &asset.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return asset, nil
}

func (r *queryResolver) Assets(ctx context.Context) ([]*model.Asset, error) {
	var assets []*model.Asset

	rows, err := db.LogAndQuery(r.db, "SELECT id, name, ticker, created_at FROM assets")
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var asset = new(model.Asset)
		if err := rows.Scan(&asset.ID, &asset.Name, &asset.Ticker, &asset.CreateAt); err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

// Author query resolvers
func (r *queryResolver) Author(ctx context.Context, id string) (*model.Author, error) {
	var author = new(model.Author)

	rows, err := db.LogAndQuery(r.db, "SELECT id, author_name, psuedonym, created_at FROM authors")
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

	rows, err := db.LogAndQuery(r.db, "SELECT id, author_name, psuedonym, created_at FROM authors")
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

// File query resolvers
func (r *queryResolver) File(ctx context.Context, id string) (*model.File, error) {
	var file = new(model.File)

	rows, err := db.LogAndQuery(r.db, "SELECT id, cover_image, latest, source, url, created_at FROM files")
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

	rows, err := db.LogAndQuery(r.db, "SELECT id, cover_image, latest, source, url, created_at FROM files")
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

// Paper query resolvers
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

// User query resolvers
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
