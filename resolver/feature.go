package resolver

import (
	"context"

	"github.com/crypto-papers/api/db"
	"github.com/crypto-papers/api/errors"
	"github.com/crypto-papers/api/model"
)

// Feature query resolvers

// Feature returns data on an individual featured paper based on a specified feature id
func (r *queryResolver) Feature(ctx context.Context, id string) (*model.Feature, error) {
	var feature = new(model.Feature)

	sql := `
		SELECT
			id,
			paper_id,
			promoted,
			sponsor,
			created_at
		FROM public.features
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
			&feature.ID,
			&feature.Paper,
			&feature.Promoted,
			&feature.Sponsor,
			&feature.CreateAt,
		)
		if err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
	}

	return feature, nil
}

// Features returns a list of featured items with their data
func (r *queryResolver) Features(ctx context.Context) ([]*model.Feature, error) {
	var features []*model.Feature

	sql := `
		SELECT
			id,
			paper_id,
			promoted,
			sponsor,
			created_at
		FROM public.features;
	`

	rows, err := db.LogAndQuery(r.db, sql)
	defer rows.Close()

	if err != nil {
		errors.DebugError(err)
		return nil, errors.InternalServerError
	}

	for rows.Next() {
		var feature = new(model.Feature)
		err := rows.Scan(
			&feature.ID,
			&feature.Paper,
			&feature.Promoted,
			&feature.Sponsor,
			&feature.CreateAt,
		)
		if err != nil {
			errors.DebugError(err)
			return nil, errors.InternalServerError
		}
		features = append(features, feature)
	}

	return features, nil
}
