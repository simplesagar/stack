package ledgertesting

import (
	"context"

	"github.com/formancehq/ledger/pkg/storage"
	"github.com/formancehq/ledger/pkg/storage/sqlstorage"
	"github.com/formancehq/ledger/pkg/storage/sqlstorage/schema"
	"github.com/formancehq/stack/libs/go-libs/pgtesting"
	"go.uber.org/fx"
)

// TODO(gfyrag): Remove second return
func StorageDriver(t pgtesting.TestingT) (*sqlstorage.Driver, func(), error) {
	pgServer := pgtesting.NewPostgresDatabase(t)

	db, err := sqlstorage.OpenSQLDB(pgServer.ConnString())
	if err != nil {
		return nil, nil, err
	}
	return sqlstorage.NewDriver(
		"postgres",
		schema.NewPostgresDB(db),
	), func() {}, nil
}

func ProvideStorageDriver(t pgtesting.TestingT) fx.Option {
	return fx.Provide(func(lc fx.Lifecycle) (storage.Driver, error) {
		driver, stopFn, err := StorageDriver(t)
		if err != nil {
			return nil, err
		}
		lc.Append(fx.Hook{
			OnStart: driver.Initialize,
			OnStop: func(ctx context.Context) error {
				stopFn()
				return driver.Close(ctx)
			},
		})
		return driver, nil
	})
}

func ProvideLedgerStorageDriver(t pgtesting.TestingT) fx.Option {
	return fx.Options(
		ProvideStorageDriver(t),
	)
}
