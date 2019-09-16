package resolver

import (
	"context"

	"github.com/crypto-papers/api/model"
)

// Update resolvers

func (r *mutationResolver) UpdateAsset(context.Context, model.AssetUpdateInput) (*model.Asset, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateAuthor(context.Context, model.AuthorUpdateInput) (*model.Author, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateFeature(context.Context, model.FeatureUpdateInput) (*model.Feature, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateFile(context.Context, model.FileUpdateInput) (*model.File, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdatePaper(context.Context, model.PaperUpdateInput) (*model.Paper, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(context.Context, model.UserUpdateInput) (*model.User, error) {
	panic("not implemented")
}
