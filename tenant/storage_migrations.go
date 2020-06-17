package tenant

import (
	"context"

	"github.com/influxdata/influxdb/v2/kv"
)

// Migration is a type which implements migration.Spec
// It is used to create tenant specific migrations
type Migration struct {
	name string
	up   func(context.Context, kv.SchemaStore) error
}

// MigrationName returns the name of the migration
func (m *Migration) MigrationName() string {
	return m.name
}

// Up calls the embedded up function on the migration
func (m *Migration) Up(ctx context.Context, store kv.SchemaStore) error {
	return m.up(ctx, store)
}

// Down is a no operation and just returns nil
func (*Migration) Down(context.Context, kv.SchemaStore) error { return nil }

// InitialMigration returns the initial migration spec used to
// get the supplied store into the correct state for use with
// the tenant service.
func InitialMigration() *Migration {
	return &Migration{
		name: "initial tenant migration",
		up: func(ctx context.Context, store kv.SchemaStore) error {
			if err := store.CreateBucket(ctx, userBucket); err != nil {
				return err
			}

			if err := store.CreateBucket(ctx, userIndex); err != nil {
				return err
			}

			if err := store.CreateBucket(ctx, userpasswordBucket); err != nil {
				return err
			}

			if err := store.CreateBucket(ctx, urmBucket); err != nil {
				return err
			}

			if err := store.CreateBucket(ctx, urmByUserIndexBucket); err != nil {
				return err
			}

			if err := store.CreateBucket(ctx, organizationBucket); err != nil {
				return err
			}

			if err := store.CreateBucket(ctx, organizationIndex); err != nil {
				return err
			}

			if err := store.CreateBucket(ctx, bucketBucket); err != nil {
				return err
			}

			if err := store.CreateBucket(ctx, bucketIndex); err != nil {
				return err
			}

			return nil
		},
	}
}
