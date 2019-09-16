package resolver

import (
	"context"

	"github.com/crypto-papers/api/db"
	"github.com/crypto-papers/api/model"
)

// Deletion resolvers

// DeleteAsset removes an asset from the database based on the asset id
func (r *mutationResolver) DeleteAsset(ctx context.Context, id model.AssetWhereUniqueInput) (string, error) {
	i := id.ID

	_, err := db.LogAndQuery(
		r.db,
		"DELETE FROM public.assets WHERE id = $1",
		i,
	)

	if err != nil {
		return i, err
	}

	return i, nil
}

// DeleteAuthor removes an author from the database based on the author id
func (r *mutationResolver) DeleteAuthor(ctx context.Context, id model.AuthorWhereUniqueInput) (string, error) {
	i := id.ID

	_, err := db.LogAndQuery(
		r.db,
		"DELETE FROM public.authors WHERE id = $1",
		i,
	)

	if err != nil {
		return i, err
	}

	return i, nil
}

// DeleteFeature removes a featured item from the database based on the feature id
func (r *mutationResolver) DeleteFeature(ctx context.Context, id model.FeatureWhereUniqueInput) (string, error) {
	i := id.ID

	_, err := db.LogAndQuery(
		r.db,
		"DELETE FROM public.features WHERE id = $1",
		i,
	)

	if err != nil {
		return i, err
	}

	return i, nil
}

// DeleteFile removes a file from the database based on the file id
func (r *mutationResolver) DeleteFile(ctx context.Context, id model.FileWhereUniqueInput) (string, error) {
	i := id.ID

	_, err := db.LogAndQuery(
		r.db,
		"DELETE FROM public.files WHERE id = $1",
		i,
	)

	if err != nil {
		return i, err
	}

	return i, nil
}

// DeletePaper removes a paper from the database based on the paper id
func (r *mutationResolver) DeletePaper(ctx context.Context, id model.PaperWhereUniqueInput) (string, error) {
	i := id.ID

	_, err := db.LogAndQuery(
		r.db,
		"DELETE FROM public.papers WHERE id = $1",
		i,
	)

	if err != nil {
		return i, err
	}

	return i, nil
}

// DeleteUser removes a user from the database based on the user id
func (r *mutationResolver) DeleteUser(ctx context.Context, id model.UserWhereUniqueInput) (string, error) {
	i := id.ID

	_, err := db.LogAndQuery(
		r.db,
		"DELETE FROM public.users WHERE id = $1",
		i,
	)

	if err != nil {
		return i, err
	}

	return i, nil
}
