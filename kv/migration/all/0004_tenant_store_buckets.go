package all

import "github.com/influxdata/influxdb/v2/tenant"

// Migration0004_TenantStoreBuckets creates all the bucket needed for the tenant store
var Migration0004_TenantStoreBuckets = tenant.InitialMigration()
