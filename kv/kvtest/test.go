package kvtest

import (
	"context"
	"testing"

	"github.com/influxdata/influxdb/v2/kv"
	"github.com/influxdata/influxdb/v2/kv/migration"
	"github.com/influxdata/influxdb/v2/kv/migration/all"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

// NewService constructs a new kv service suitable for use in tests.
// It applies all migrations to the supplied store and also returns
// a test logger.
func NewService(t *testing.T, ctx context.Context, store kv.SchemaStore) (logger *zap.Logger, service *kv.Service) {
	t.Helper()

	logger = zaptest.NewLogger(t)
	migrator, err := migration.NewMigrator(logger, store, all.Migrations[:]...)
	if err != nil {
		t.Fatal(err)
	}

	if err := migrator.Up(ctx); err != nil {
		t.Fatal(err)
	}

	service = kv.NewService(logger, store)

	return
}
