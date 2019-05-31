package loader

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

type contextKey string

// Set context key variables
const (
	authorLoaderKey   = contextKey("author")
	currencyLoaderKey = contextKey("currency")
	fileLoaderKey     = contextKey("file")
	paperLoaderKey    = contextKey("paper")
	userLoaderKey     = contextKey("user")
)

// Client provides getters for all the model types
type Client interface {
	authorGetter
	// currencyGetter
	// fileGetter
	// paperGetter
	// userGetter
}

// Collection holds an internal lookup of initialized batch data load functions.
type Collection struct {
	lookup map[contextKey]dataloader.BatchFunc
}

// Initialize a lookup map of context keys to batch functions.
//
// When Attach is called on the Collection, the batch functions are used to create new dataloader
// instances. The instances are attached to the request context at the provided keys.
//
// The keys are then used to extract the dataloader instances from the request context.
func Initialize() Collection {
	return Collection{
		lookup: map[contextKey]dataloader.BatchFunc{
			authorLoaderKey: newAuthorLoader(),
			// currencyLoaderKey: newCurrencyLoader(client),
			// fileLoaderKey:     newFileLoader(client),
			// paperLoaderKey:    newPaperLoader(client),
			// userLoaderKey:     newUserLoader(client),
		},
	}
}

// Attach creates new instances of dataloader.Loader and attaches the instances on the request context.
// We do this because the dataloader has an in-memory cache that is scoped to the request.
func (c Collection) Attach(ctx context.Context) context.Context {
	for key, batchFn := range c.lookup {
		ctx = context.WithValue(ctx, key, dataloader.NewBatchedLoader(batchFn))
	}

	return ctx
}

// extract is a helper function to make common get-value, assert-type, return-error-or-value
// operations easier.
func extract(ctx context.Context, key contextKey) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(key).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("unable to find %s loader on the request context", key)
	}

	return ldr, nil
}

// String returns the context key as a string
func (key contextKey) String() string {
	return string(key)
}
