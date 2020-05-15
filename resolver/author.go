package resolver

import (
	"context"

	"github.com/crypto-papers/api/db"
	"github.com/crypto-papers/api/errors"
	"github.com/crypto-papers/api/model"
)

// Author query resolvers

// Author returns data on an individual author based on a specified author id
func (r *queryResolver) Author(ctx context.Context, id string) (*model.Author, error) {
	var author = new(model.Author)

	sql := `
		SELECT
			id,
			author_name,
			bio,
			photo,
			pseudonym,
			created_at
		FROM public.authors
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
			&author.Pseudonym,
			&author.CreateAt,
		)
		if err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return author, nil
}

// Authors returns a list of authors with their data
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	var authors []*model.Author

	sql := `
		SELECT
			id,
			author_name,
			bio,
			photo,
			pseudonym,
			created_at
		FROM public.authors;
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
			&author.Pseudonym,
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
