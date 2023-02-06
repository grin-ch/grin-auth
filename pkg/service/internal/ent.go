package internal

import (
	"context"

	"github.com/grin-ch/grin-auth/pkg/model"
)

func Transaction[T any](ctx context.Context, client *model.Client, fn func(*model.Tx) (T, error)) (model T, err error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return
	}
	model, err = fn(tx)
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}
