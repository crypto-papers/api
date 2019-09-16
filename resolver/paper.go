package resolver

import (
	"context"

	"github.com/crypto-papers/api/db"
	"github.com/crypto-papers/api/errors"
	"github.com/crypto-papers/api/model"
)

// Paper query resolvers

// Paper returns data on an individual paper based on a specified paper id
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
		FROM public.papers
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

// PaperByPid returns data on an individual paper based on a specified paper pretty id
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
		FROM public.papers
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

// Papers returns a list of papers with their data
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
		FROM public.papers;
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
