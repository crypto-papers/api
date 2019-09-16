package resolver

import (
	"context"

	"github.com/crypto-papers/api/db"
	"github.com/crypto-papers/api/errors"
	"github.com/crypto-papers/api/model"
)

// File query resolvers

// File returns data on an individual file based on a specified file id
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
		FROM public.files
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

// Files returns a list of files with their data
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
		FROM public.files;
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
