package session

import (
	"context"

	"github.com/go-pg/pg"
)

type contextValueKey int

const (
	keyRequest  contextValueKey = 0
	keyDatabase contextValueKey = 1
)

func WithDatabase(ctx context.Context, database *pg.DB) context.Context {
	return context.WithValue(ctx, keyDatabase, database)
}

func Database(ctx context.Context) *pg.DB {
	return ctx.Value(keyDatabase).(*pg.DB)
}
