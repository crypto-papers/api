package loader

import (
	"context"
	"sync"

	"github.com/graph-gophers/dataloader"

	"github.com/crypto-papers/api/model"
)

type authorGetter interface {
	Author(ctx context.Context, key string) model.Author
}

type authorLoader struct{}

func newAuthorLoader() dataloader.BatchFunc {
	return authorLoader{}.loadBatch
}

func (ldr authorLoader) loadBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	var (
		num     = len(keys)
		results = make([]*dataloader.Result, num)
		wg      sync.WaitGroup
	)

	wg.Add(num)

	for i, key := range keys {
		go func(i int, key dataloader.Key) {
			defer wg.Done()

			// resp, err := ctx.Value()
			// results[i] = &dataloader.Result{Data: resp, Error: err}
		}(i, key)
	}

	wg.Wait()

	return results
}
