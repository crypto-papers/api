package resolver

import (
	"context"
	"time"

	"github.com/crypto-papers/api/db"
	"github.com/crypto-papers/api/model"
)

// Creation resolvers

// CreateAsset accepts asset data values and creates an asset entry in the database
func (r *mutationResolver) CreateAsset(ctx context.Context, data model.AssetCreateInput) (*model.Asset, error) {
	asset := &model.Asset{
		Logo:     data.Logo,
		Name:     data.Name,
		Ticker:   data.Ticker,
		CreateAt: time.Now(),
	}

	sql := `
		INSERT INTO public.assets (
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

// CreateAuthor accepts author data values and creates an author entry in the database
func (r *mutationResolver) CreateAuthor(ctx context.Context, data model.AuthorCreateInput) (*model.Author, error) {
	author := &model.Author{
		Bio:       data.Bio,
		Name:      data.Name,
		Photo:     data.Photo,
		Pseudonym: data.Pseudonym,
		CreateAt:  time.Now(),
	}

	authorSQL := `
		INSERT INTO public.authors (
			author_name,
			bio,
			photo,
			pseudonym,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	// paperSQL := `
	// 	INSERT INTO public.author_paper (
	// 		author_id,
	// 		paper_id
	// 	)
	// 	VALUES ($1, $2);
	// `

	authorID, err := db.LogQueryAndScan(
		r.db,
		authorSQL,
		author.Name,
		author.Bio,
		author.Photo,
		author.Pseudonym,
		author.CreateAt,
	)
	if err != nil {
		return nil, err
	}

	author.ID = authorID

	// papers := data.Papers.Connect
	// for _, paper := range papers {
	// 	paperID := paper.ID

	// 	row, err := db.LogAndQuery(
	// 		r.db,
	// 		paperSQL,
	// 		authorID,
	// 		paperID,
	// 	)

	// 	if err != nil || !row.Next() {
	// 		return author, err
	// 	}
	// }

	return author, nil
}

func (r *mutationResolver) CreateFeature(ctx context.Context, data model.FeatureCreateInput) (*model.Feature, error) {
	feature := &model.Feature{
		Paper:    data.Paper,
		Promoted: data.Promoted,
		Sponsor:  data.Sponsor,
		CreateAt: time.Now(),
	}

	sql := `
		INSERT INTO public.features (
			paper_id,
			promoted,
			sponsor,
			created_at
		)
		VALUES ($1, $2, $3, $4)
	`

	featureID, err := db.LogQueryAndScan(
		r.db,
		sql,
		feature.Paper,
		feature.Promoted,
		feature.Sponsor,
		feature.CreateAt,
	)

	if err != nil {
		return nil, err
	}

	feature.ID = featureID

	return feature, nil
}

// CreateFile accepts file data values and creates a file entry in the database
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
		INSERT INTO public.files (
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

// CreatePaper accepts paper data values and creates a paper entry in the database
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
		INSERT INTO public.papers (
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

// CreateUser accepts user data values and creates a user entry in the database
func (r *mutationResolver) CreateUser(ctx context.Context, data model.UserCreateInput) (*model.User, error) {
	user := &model.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		CreateAt: time.Now(),
	}

	sql := `
		INSERT INTO public.users (
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
