package resolver

import (
	"context"

	"github.com/crypto-papers/api/db"
	"github.com/crypto-papers/api/errors"
	"github.com/crypto-papers/api/model"
)

// Asset query resolvers

// Asset returns data on an individual asset based on a specified asset id
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

// Assets returns a list of assets with their data
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
