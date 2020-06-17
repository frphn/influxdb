package all

import (
	"context"

	"github.com/influxdata/influxdb/v2/kv"
	"github.com/influxdata/influxdb/v2/kv/migration"
)

// Migrations contains all the migrations required for the entire of the
// kv store backing influxdb's metadata.
var Migrations = [...]migration.Spec{
	// initial migrations
	Migration0001_InitialMigration,
	// add index user resource mappings by user id
	Migration0002_AddURMByUserIndex,
	// add index for tasks with missing owner IDs
	Migration0003_TaskOwnerIDUpMigration,
	// initial tenant migrations
	Migration0004_TenantStoreBuckets,
}

// MigrationFunc is a function which can be used as either an up or down operation.
type MigrationFunc func(context.Context, kv.SchemaStore) error

func noopMigration(context.Context, kv.SchemaStore) error {
	return nil
}

// Migration is a type which implements the migration packages Spec interface
// It can be used to conveniently create migration specs for the all package
type Migration struct {
	name string
	up   MigrationFunc
	down MigrationFunc
}

// UpOnlyMigration is a migration with an up function and a noop down function
func UpOnlyMigration(name string, up MigrationFunc) *Migration {
	return &Migration{name, up, noopMigration}
}

// MigrationName returns the underlying name of the migation
func (m *Migration) MigrationName() string {
	return m.name
}

// Up delegates to the underlying anonymous up migration function
func (m *Migration) Up(ctx context.Context, store kv.SchemaStore) error {
	return m.up(ctx, store)
}

// Down delegates to the underlying anonymous down migration function
func (m *Migration) Down(ctx context.Context, store kv.SchemaStore) error {
	return m.down(ctx, store)
}
