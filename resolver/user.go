package resolver

import (
	"context"

	"github.com/crypto-papers/api/db"
	"github.com/crypto-papers/api/errors"
	"github.com/crypto-papers/api/model"
)

// User query resolvers

// User returns data on an individual user based on a specified user id
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	var user = new(model.User)

	sql := `
		SELECT
			id,
			user_name,
			email,
			created_at
		FROM public.users
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

// Users returns a list of users with their data
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	sql := `
		SELECT
			id,
			user_name,
			email,
			created_at
		FROM public.users;
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
