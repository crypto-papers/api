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
		Logo:     data.Logo,
		Name:     data.Name,
		Ticker:   data.Ticker,
		CreateAt: time.Now(),
	}

	sql := `
		INSERT INTO assets (
			asset_name,
			logo,
			ticker,
			created_at
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	assetID, err := db.LogQueryAndScan(
		r.db,
		sql,
		asset.Name,
		asset.Logo,
		asset.Ticker,
		asset.CreateAt,
	)
	if err != nil {
		return nil, err
	}

	asset.ID = assetID

	return asset, nil
}

func (r *mutationResolver) DeleteAsset(ctx context.Context, id model.AssetWhereUniqueInput) (string, error) {
	i := id.ID

	_, err := db.LogAndQuery(
		r.db,
		"DELETE FROM assets WHERE id = $1",
		i,
	)

	if err != nil {
		return i, err
	}

	return i, nil
}

func (r *mutationResolver) UpdateAsset(context.Context, model.AssetUpdateInput) (*model.Asset, error) {
	panic("not implemented")
}

// Author mutation resolvers
func (r *mutationResolver) CreateAuthor(ctx context.Context, data model.AuthorCreateInput) (*model.Author, error) {
	author := &model.Author{
		Bio:       data.Bio,
		Name:      data.Name,
		Photo:     data.Photo,
		Psuedonym: data.Psuedonym,
		CreateAt:  time.Now(),
	}

	authorSQL := `
		INSERT INTO authors (
			author_name,
			bio,
			photo,
			psuedonym,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	authorID, err := db.LogQueryAndScan(
		r.db,
		authorSQL,
		author.Bio,
		author.Name,
		author.Photo,
		author.Psuedonym,
		author.CreateAt,
	)
	if err != nil {
		return nil, err
	}

	author.ID = authorID

	papers := data.Papers.Connect
	for _, paper := range papers {
		paperID := paper.ID

		row, err := db.LogAndQuery(
			r.db,
			"INSERT INTO author_paper (author_id, paper_id) VALUES ($1, $2)",
			authorID,
			paperID,
		)

		if err != nil || !row.Next() {
			return author, err
		}
	}

	return author, nil
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id model.AuthorWhereUniqueInput) (string, error) {
	i := id.ID

	_, err := db.LogAndQuery(
		r.db,
		"DELETE FROM authors WHERE id = $1",
		i,
	)

	if err != nil {
		return i, err
	}

	return i, nil
}

func (r *mutationResolver) UpdateAuthor(context.Context, model.AuthorUpdateInput) (*model.Author, error) {
	panic("not implemented")
}

// File mutation resolvers
func (r *mutationResolver) CreateFile(ctx context.Context, data model.FileCreateInput) (*model.File, error) {
	file := &model.File{
		CoverImage: data.CoverImage,
		Filename:   data.Filename,
		Latest:     data.Latest,
		PageNum:    data.PageNum,
		Source:     data.Source,
		Version:    data.Version,
		CreateAt:   time.Now(),
	}

	sql := `
		INSERT INTO files (
			cover_image,
			filename,
			is_latest,
			page_num,
			source,
			version,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	fileID, err := db.LogQueryAndScan(
		r.db,
		sql,
		file.CoverImage,
		file.Filename,
		file.Latest,
		file.PageNum,
		file.Source,
		file.Version,
		file.CreateAt,
	)

	if err != nil {
		return nil, err
	}

	file.ID = fileID

	return file, nil
}

func (r *mutationResolver) DeleteFile(ctx context.Context, id model.FileWhereUniqueInput) (string, error) {
	i := id.ID

	_, err := db.LogAndQuery(
		r.db,
		"DELETE FROM files WHERE id = $1",
		i,
	)

	if err != nil {
		return i, err
	}

	return i, nil
}

func (r *mutationResolver) UpdateFile(context.Context, model.FileUpdateInput) (*model.File, error) {
	panic("not implemented")
}

// Paper mutation resolvers
func (r *mutationResolver) CreatePaper(ctx context.Context, data model.PaperCreateInput) (*model.Paper, error) {
	paper := &model.Paper{
		Description:   data.Description,
		Excerpt:       data.Excerpt,
		LatestVersion: data.LatestVersion,
		Title:         data.Title,
		SubTitle:      data.SubTitle,
		CreateAt:      time.Now(),
	}

	sql := `
		INSERT INTO papers (
			description,
			excerpt,
			latest_version,
			title_primary,
			title_secondary,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	paperID, err := db.LogQueryAndScan(
		r.db,
		sql,
		paper.Description,
		paper.Excerpt,
		paper.LatestVersion,
		paper.Title,
		paper.SubTitle,
		paper.CreateAt,
	)

	if err != nil {
		return nil, err
	}

	paper.ID = paperID

	return paper, nil
}

func (r *mutationResolver) DeletePaper(ctx context.Context, id model.PaperWhereUniqueInput) (string, error) {
	i := id.ID

	_, err := db.LogAndQuery(
		r.db,
		"DELETE FROM papers WHERE id = $1",
		i,
	)

	if err != nil {
		return i, err
	}

	return i, nil
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

	sql := `
		INSERT INTO users (
			user_name,
			email,
			password,
			created_at
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	userID, err := db.LogQueryAndScan(
		r.db,
		sql,
		user.Name,
		user.Email,
		user.Password,
		user.CreateAt,
	)
	if err != nil {
		return nil, err
	}

	user.ID = userID

	return user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id model.UserWhereUniqueInput) (string, error) {
	i := id.ID

	_, err := db.LogAndQuery(
		r.db,
		"DELETE FROM users WHERE id = $1",
		i,
	)

	if err != nil {
		return i, err
	}

	return i, nil
}

func (r *mutationResolver) UpdateUser(context.Context, model.UserUpdateInput) (*model.User, error) {
	panic("not implemented")
}

// Query resolvers
type queryResolver struct{ *Resolver }

// Asset query resolvers
func (r *queryResolver) Asset(ctx context.Context, id string) (*model.Asset, error) {
	var asset = new(model.Asset)

	sql := `
		SELECT
			id,
			asset_name,
			logo,
			ticker,
			created_at
		FROM assets
		WHERE id = $1
		LIMIT 1;
	`

	row, err := db.LogAndQuery(r.db, sql, id)
	defer row.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for row.Next() {
		err := row.Scan(
			&asset.ID,
			&asset.Name,
			&asset.Logo,
			&asset.Ticker,
			&asset.CreateAt,
		)
		if err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return asset, nil
}

func (r *queryResolver) Assets(ctx context.Context) ([]*model.Asset, error) {
	var assets []*model.Asset

	sql := `
		SELECT
			id,
			asset_name,
			logo,
			ticker,
			created_at
		FROM assets;
	`

	rows, err := db.LogAndQuery(r.db, sql)
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var asset = new(model.Asset)
		err := rows.Scan(
			&asset.ID,
			&asset.Name,
			&asset.Logo,
			&asset.Ticker,
			&asset.CreateAt,
		)
		if err != nil {
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

	sql := `
		SELECT
			id,
			author_name,
			bio,
			photo,
			psuedonym,
			created_at
		FROM authors
		WHERE id = $1
		LIMIT 1;
	`

	row, err := db.LogAndQuery(r.db, sql, id)
	defer row.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for row.Next() {
		err := row.Scan(
			&author.ID,
			&author.Name,
			&author.Bio,
			&author.Photo,
			&author.Psuedonym,
			&author.CreateAt,
		)
		if err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return author, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	var authors []*model.Author

	sql := `
		SELECT
			id,
			author_name,
			bio,
			photo,
			psuedonym,
			created_at
		FROM authors;
	`

	rows, err := db.LogAndQuery(r.db, sql)
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var author = new(model.Author)
		err := rows.Scan(
			&author.ID,
			&author.Name,
			&author.Bio,
			&author.Photo,
			&author.Psuedonym,
			&author.CreateAt,
		)
		if err != nil {
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

	sql := `
		SELECT
			id,
			cover_image,
			filename,
			is_latest,
			page_num,
			pub_date,
			source,
			version,
			created_at
		FROM files
		WHERE id = $1
		LIMIT 1;
	`

	row, err := db.LogAndQuery(r.db, sql, id)
	defer row.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for row.Next() {
		err := row.Scan(
			&file.ID,
			&file.CoverImage,
			&file.Filename,
			&file.Latest,
			&file.PageNum,
			&file.PubDate,
			&file.Source,
			&file.Version,
			&file.CreateAt,
		)
		if err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return file, nil
}

func (r *queryResolver) Files(ctx context.Context) ([]*model.File, error) {
	var files []*model.File

	sql := `
		SELECT
			id,
			cover_image,
			filename,
			is_latest,
			page_num,
			pub_date,
			source,
			version,
			created_at
		FROM files;
	`

	rows, err := db.LogAndQuery(r.db, sql)
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var file = new(model.File)
		err := rows.Scan(
			&file.ID,
			&file.CoverImage,
			&file.Filename,
			&file.Latest,
			&file.PageNum,
			&file.PubDate,
			&file.Source,
			&file.Version,
			&file.CreateAt,
		)
		if err != nil {
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

	sql := `
		SELECT
			id,
			description,
			excerpt,
			latest_version,
			pretty_id,
			title_primary,
			title_secondary,
			created_at 
		FROM papers
		WHERE id = $1
		LIMIT 1;
	`

	row, err := db.LogAndQuery(r.db, sql, id)
	defer row.Close()

	if err != nil {
		errors.DebugError(err)
	}

	for row.Next() {
		err = row.Scan(
			&paper.ID,
			&paper.Description,
			&paper.Excerpt,
			&paper.LatestVersion,
			&paper.PrettyID,
			&paper.Title,
			&paper.SubTitle,
			&paper.CreateAt,
		)
		if err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return paper, nil
}

func (r *queryResolver) PaperByPid(ctx context.Context, prettyID int) (*model.Paper, error) {
	var paper = new(model.Paper)

	sql := `
		SELECT
			id,
			description,
			excerpt,
			latest_version,
			pretty_id,
			title_primary,
			title_secondary,
			created_at
		FROM papers
		WHERE pretty_id = $1
		LIMIT 1;
	`

	row, err := db.LogAndQuery(r.db, sql, prettyID)
	defer row.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for row.Next() {
		err = row.Scan(
			&paper.ID,
			&paper.Description,
			&paper.Excerpt,
			&paper.LatestVersion,
			&paper.PrettyID,
			&paper.Title,
			&paper.SubTitle,
			&paper.CreateAt,
		)
		if err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return paper, nil
}

func (r *queryResolver) Papers(ctx context.Context) ([]*model.Paper, error) {
	var papers []*model.Paper

	sql := `
		SELECT
			id,
			description,
			excerpt,
			latest_version,
			pretty_id,
			title_primary,
			title_secondary,
			created_at
		FROM papers;
	`

	rows, err := db.LogAndQuery(r.db, sql)
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var paper = new(model.Paper)
		err := rows.Scan(
			&paper.ID,
			&paper.Description,
			&paper.Excerpt,
			&paper.LatestVersion,
			&paper.PrettyID,
			&paper.Title,
			&paper.SubTitle,
			&paper.CreateAt,
		)
		if err != nil {
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

	sql := `
		SELECT
			id,
			user_name,
			email,
			created_at
		FROM users
		WHERE id = $1
		LIMIT 1;
	`

	row, err := db.LogAndQuery(r.db, sql, id)
	defer row.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for row.Next() {
		err := row.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreateAt,
		)
		if err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	sql := `
		SELECT
			id,
			user_name,
			email,
			created_at
		FROM users;
	`
	rows, err := db.LogAndQuery(r.db, sql)
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var user = new(model.User)
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreateAt,
		)
		if err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
		users = append(users, user)
	}

	return users, nil
}
