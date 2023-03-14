package ledger_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/formancehq/ledger/pkg/core"
	"github.com/formancehq/ledger/pkg/ledger"
	"github.com/formancehq/ledger/pkg/ledgertesting"
	"github.com/formancehq/ledger/pkg/storage"
	"github.com/formancehq/ledger/pkg/storage/sqlstorage"
	ledgerstore "github.com/formancehq/ledger/pkg/storage/sqlstorage/ledger"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestStore(t *testing.T) {
	type testingFunction struct {
		name string
		fn   func(t *testing.T, store *ledgerstore.Store)
	}

	for _, tf := range []testingFunction{
		{name: "Commit", fn: testCommit},
		{name: "UpdateTransactionMetadata", fn: testUpdateTransactionMetadata},
		{name: "UpdateAccountMetadata", fn: testUpdateAccountMetadata},
		{name: "GetLastLog", fn: testGetLastLog},
		{name: "GetLogs", fn: testGetLogs},
		{name: "CountAccounts", fn: testCountAccounts},
		{name: "GetAssetsVolumes", fn: testGetAssetsVolumes},
		{name: "GetAccounts", fn: testGetAccounts},
		{name: "Transactions", fn: testTransactions},
		{name: "GetTransaction", fn: testGetTransaction},
		{name: "GetBalances", fn: testGetBalances},
		{name: "GetBalancesAggregated", fn: testGetBalancesAggregated},
	} {
		t.Run(fmt.Sprintf("postgres/%s", tf.name), func(t *testing.T) {
			done := make(chan struct{})
			app := fx.New(
				ledgertesting.ProvideStorageDriver(t),
				fx.NopLogger,
				fx.Invoke(func(driver *sqlstorage.Driver, lc fx.Lifecycle) {
					lc.Append(fx.Hook{
						OnStart: func(ctx context.Context) error {
							defer func() {
								close(done)
							}()
							store, _, err := driver.GetLedgerStore(ctx, uuid.NewString(), true)
							if err != nil {
								return err
							}
							defer store.Close(ctx)

							if _, err = store.Initialize(context.Background()); err != nil {
								return err
							}

							tf.fn(t, store)
							return nil
						},
					})
				}),
			)
			go func() {
				require.NoError(t, app.Start(context.Background()))
			}()
			defer func(app *fx.App, ctx context.Context) {
				require.NoError(t, app.Stop(ctx))
			}(app, context.Background())

			select {
			case <-time.After(5 * time.Second):
				t.Fatal("timeout")
			case <-done:
			}
		})
	}
}

var now = time.Now().UTC().Truncate(time.Second)
var tx1 = core.ExpandedTransaction{
	Transaction: core.Transaction{
		TransactionData: core.TransactionData{
			Postings: []core.Posting{
				{
					Source:      "world",
					Destination: "central_bank",
					Amount:      core.NewMonetaryInt(100),
					Asset:       "USD",
				},
			},
			Reference: "tx1",
			Timestamp: now.Add(-3 * time.Hour),
		},
	},
	PostCommitVolumes: core.AccountsAssetsVolumes{
		"world": {
			"USD": {
				Input:  core.NewMonetaryInt(0),
				Output: core.NewMonetaryInt(100),
			},
		},
		"central_bank": {
			"USD": {
				Input:  core.NewMonetaryInt(100),
				Output: core.NewMonetaryInt(0),
			},
		},
	},
	PreCommitVolumes: core.AccountsAssetsVolumes{
		"world": {
			"USD": {
				Input:  core.NewMonetaryInt(0),
				Output: core.NewMonetaryInt(0),
			},
		},
		"central_bank": {
			"USD": {
				Input:  core.NewMonetaryInt(0),
				Output: core.NewMonetaryInt(0),
			},
		},
	},
}
var tx2 = core.ExpandedTransaction{
	Transaction: core.Transaction{
		ID: 1,
		TransactionData: core.TransactionData{
			Postings: []core.Posting{
				{
					Source:      "world",
					Destination: "central_bank",
					Amount:      core.NewMonetaryInt(100),
					Asset:       "USD",
				},
			},
			Reference: "tx2",
			Timestamp: now.Add(-2 * time.Hour),
		},
	},
	PostCommitVolumes: core.AccountsAssetsVolumes{
		"world": {
			"USD": {
				Input:  core.NewMonetaryInt(0),
				Output: core.NewMonetaryInt(200),
			},
		},
		"central_bank": {
			"USD": {
				Input:  core.NewMonetaryInt(200),
				Output: core.NewMonetaryInt(0),
			},
		},
	},
	PreCommitVolumes: core.AccountsAssetsVolumes{
		"world": {
			"USD": {
				Input:  core.NewMonetaryInt(0),
				Output: core.NewMonetaryInt(100),
			},
		},
		"central_bank": {
			"USD": {
				Input:  core.NewMonetaryInt(100),
				Output: core.NewMonetaryInt(0),
			},
		},
	},
}
var tx3 = core.ExpandedTransaction{
	Transaction: core.Transaction{
		ID: 2,
		TransactionData: core.TransactionData{
			Postings: []core.Posting{
				{
					Source:      "central_bank",
					Destination: "users:1",
					Amount:      core.NewMonetaryInt(1),
					Asset:       "USD",
				},
			},
			Reference: "tx3",
			Metadata: core.Metadata{
				"priority": json.RawMessage(`"high"`),
			},
			Timestamp: now.Add(-1 * time.Hour),
		},
	},
	PreCommitVolumes: core.AccountsAssetsVolumes{
		"central_bank": {
			"USD": {
				Input:  core.NewMonetaryInt(200),
				Output: core.NewMonetaryInt(0),
			},
		},
		"users:1": {
			"USD": {
				Input:  core.NewMonetaryInt(0),
				Output: core.NewMonetaryInt(0),
			},
		},
	},
	PostCommitVolumes: core.AccountsAssetsVolumes{
		"central_bank": {
			"USD": {
				Input:  core.NewMonetaryInt(200),
				Output: core.NewMonetaryInt(1),
			},
		},
		"users:1": {
			"USD": {
				Input:  core.NewMonetaryInt(1),
				Output: core.NewMonetaryInt(0),
			},
		},
	},
}

func testCommit(t *testing.T, store *ledgerstore.Store) {
	tx := core.ExpandedTransaction{
		Transaction: core.Transaction{
			ID: 0,
			TransactionData: core.TransactionData{
				Postings: []core.Posting{
					{
						Source:      "world",
						Destination: "central_bank",
						Amount:      core.NewMonetaryInt(100),
						Asset:       "USD",
					},
				},
				Reference: "foo",
				Timestamp: time.Now().Round(time.Second),
			},
		},
	}
	err := store.Commit(context.Background(), tx)
	require.NoError(t, err)
	logs := make([]core.Log, 0)
	logs = append(logs, core.NewTransactionLog(tx.Transaction))
	errChan := store.AppendLogs(context.Background(), logs...)
	require.NoError(t, <-errChan)

	err = store.Commit(context.Background(), tx)
	require.Error(t, err)
	require.True(t, storage.IsErrorCode(err, storage.ConstraintFailed))

	cursor, err := store.GetLogs(context.Background(), ledger.NewLogsQuery())
	require.NoError(t, err)
	require.Len(t, cursor.Data, 1)
}

func testUpdateTransactionMetadata(t *testing.T, store *ledgerstore.Store) {
	tx := core.ExpandedTransaction{
		Transaction: core.Transaction{
			ID: 0,
			TransactionData: core.TransactionData{
				Postings: []core.Posting{
					{
						Source:      "world",
						Destination: "central_bank",
						Amount:      core.NewMonetaryInt(100),
						Asset:       "USD",
					},
				},
				Reference: "foo",
				Timestamp: time.Now().Round(time.Second),
			},
		},
	}
	err := store.Commit(context.Background(), tx)
	require.NoError(t, err)
	logs := make([]core.Log, 0)
	logs = append(logs, core.NewTransactionLog(tx.Transaction))
	errChan := store.AppendLogs(context.Background(), logs...)
	require.NoError(t, <-errChan)

	at := time.Now()
	err = store.UpdateTransactionMetadata(context.Background(), tx.ID, core.Metadata{
		"foo": "bar",
	})
	require.NoError(t, err)
	logs = nil
	logs = append(logs, core.NewSetMetadataLog(at, core.SetMetadata{
		TargetType: core.MetaTargetTypeTransaction,
		TargetID:   tx.ID,
		Metadata: core.Metadata{
			"foo": "bar",
		},
	}))
	errChan = store.AppendLogs(context.Background(), logs...)
	require.NoError(t, <-errChan)

	retrievedTransaction, err := store.GetTransaction(context.Background(), tx.ID)
	require.NoError(t, err)
	require.EqualValues(t, "bar", retrievedTransaction.Metadata["foo"])

	cursor, err := store.GetLogs(context.Background(), ledger.NewLogsQuery())
	require.NoError(t, err)
	require.Len(t, cursor.Data, 2)
}

func testUpdateAccountMetadata(t *testing.T, store *ledgerstore.Store) {
	tx := core.ExpandedTransaction{
		Transaction: core.Transaction{
			ID: 0,
			TransactionData: core.TransactionData{
				Postings: []core.Posting{
					{
						Source:      "world",
						Destination: "central_bank",
						Amount:      core.NewMonetaryInt(100),
						Asset:       "USD",
					},
				},
				Reference: "foo",
				Timestamp: time.Now().Round(time.Second),
			},
		},
	}
	err := store.Commit(context.Background(), tx)
	require.NoError(t, err)
	logs := make([]core.Log, 0)
	logs = append(logs, core.NewTransactionLog(tx.Transaction))
	errChan := store.AppendLogs(context.Background(), logs...)
	require.NoError(t, <-errChan)

	at := time.Now()
	err = store.UpdateAccountMetadata(context.Background(), "central_bank", core.Metadata{
		"foo": "bar",
	})
	require.NoError(t, err)
	logs = nil
	logs = append(logs, core.NewSetMetadataLog(at, core.SetMetadata{
		TargetType: core.MetaTargetTypeAccount,
		TargetID:   "central_bank",
		Metadata: core.Metadata{
			"foo": "bar",
		},
	}))
	errChan = store.AppendLogs(context.Background(), logs...)
	require.NoError(t, <-errChan)

	account, err := store.GetAccount(context.Background(), "central_bank")
	require.NoError(t, err)
	require.EqualValues(t, "bar", account.Metadata["foo"])

	cursor, err := store.GetLogs(context.Background(), ledger.NewLogsQuery())
	require.NoError(t, err)
	require.Len(t, cursor.Data, 2)
}

func testCountAccounts(t *testing.T, store *ledgerstore.Store) {
	tx := core.ExpandedTransaction{
		Transaction: core.Transaction{
			ID: 0,
			TransactionData: core.TransactionData{
				Postings: []core.Posting{
					{
						Source:      "world",
						Destination: "central_bank",
						Amount:      core.NewMonetaryInt(100),
						Asset:       "USD",
					},
				},
				Timestamp: time.Now().Round(time.Second),
			},
		},
	}
	err := store.Commit(context.Background(), tx)
	require.NoError(t, err)

	countAccounts, err := store.CountAccounts(context.Background(), ledger.AccountsQuery{})
	require.NoError(t, err)
	require.EqualValues(t, 2, countAccounts) // world + central_bank
}

func testGetAssetsVolumes(t *testing.T, store *ledgerstore.Store) {
	tx := core.ExpandedTransaction{
		Transaction: core.Transaction{
			TransactionData: core.TransactionData{
				Postings: []core.Posting{
					{
						Source:      "world",
						Destination: "central_bank",
						Amount:      core.NewMonetaryInt(100),
						Asset:       "USD",
					},
				},
				Timestamp: time.Now().Round(time.Second),
			},
		},
		PostCommitVolumes: core.AccountsAssetsVolumes{
			"central_bank": core.AssetsVolumes{
				"USD": {
					Input:  core.NewMonetaryInt(100),
					Output: core.NewMonetaryInt(0),
				},
			},
		},
		PreCommitVolumes: core.AccountsAssetsVolumes{
			"central_bank": core.AssetsVolumes{
				"USD": {
					Input:  core.NewMonetaryInt(100),
					Output: core.NewMonetaryInt(0),
				},
			},
		},
	}
	err := store.Commit(context.Background(), tx)
	require.NoError(t, err)

	volumes, err := store.GetAssetsVolumes(context.Background(), "central_bank")
	require.NoError(t, err)
	require.Len(t, volumes, 1)
	require.EqualValues(t, core.NewMonetaryInt(100), volumes["USD"].Input)
	require.EqualValues(t, core.NewMonetaryInt(0), volumes["USD"].Output)
}

func testGetAccounts(t *testing.T, store *ledgerstore.Store) {
	logs := make([]core.Log, 0)
	require.NoError(t, store.UpdateAccountMetadata(context.Background(), "world", core.Metadata{
		"foo": json.RawMessage(`"bar"`),
	}))
	logs = append(logs, core.NewSetMetadataLog(time.Now(), core.SetMetadata{
		TargetType: core.MetaTargetTypeAccount,
		TargetID:   "world",
		Metadata: core.Metadata{
			"foo": json.RawMessage(`"bar"`),
		},
	}))
	require.NoError(t, store.UpdateAccountMetadata(context.Background(), "bank", core.Metadata{
		"hello": json.RawMessage(`"world"`),
	}))
	logs = append(logs, core.NewSetMetadataLog(time.Now(), core.SetMetadata{
		TargetType: core.MetaTargetTypeAccount,
		TargetID:   "bank",
		Metadata: core.Metadata{
			"hello": json.RawMessage(`"world"`),
		},
	}))
	require.NoError(t, store.UpdateAccountMetadata(context.Background(), "order:1", core.Metadata{
		"hello": json.RawMessage(`"world"`),
	}))
	logs = append(logs, core.NewSetMetadataLog(time.Now(), core.SetMetadata{
		TargetType: core.MetaTargetTypeAccount,
		TargetID:   "order:1",
		Metadata: core.Metadata{
			"hello": json.RawMessage(`"world"`),
		},
	}))
	require.NoError(t, store.UpdateAccountMetadata(context.Background(), "order:2", core.Metadata{
		"number":  json.RawMessage(`3`),
		"boolean": json.RawMessage(`true`),
		"a":       json.RawMessage(`{"super": {"nested": {"key": "hello"}}}`),
	}))
	logs = append(logs, core.NewSetMetadataLog(time.Now(), core.SetMetadata{
		TargetType: core.MetaTargetTypeAccount,
		TargetID:   "order:2",
		Metadata: core.Metadata{
			"number":  json.RawMessage(`3`),
			"boolean": json.RawMessage(`true`),
			"a":       json.RawMessage(`{"super": {"nested": {"key": "hello"}}}`),
		},
	}))
	errChan := store.AppendLogs(context.Background(), logs...)
	require.NoError(t, <-errChan)

	accounts, err := store.GetAccounts(context.Background(), ledger.AccountsQuery{
		PageSize: 1,
	})
	require.NoError(t, err)
	require.Equal(t, 1, accounts.PageSize)
	require.Len(t, accounts.Data, 1)

	accounts, err = store.GetAccounts(context.Background(), ledger.AccountsQuery{
		PageSize:     1,
		AfterAddress: string(accounts.Data[0].Address),
	})
	require.NoError(t, err)
	require.Equal(t, 1, accounts.PageSize)

	accounts, err = store.GetAccounts(context.Background(), ledger.AccountsQuery{
		PageSize: 10,
		Filters: ledger.AccountsQueryFilters{
			Address: ".*der.*",
		},
	})
	require.NoError(t, err)
	require.Len(t, accounts.Data, 2)
	require.Equal(t, 10, accounts.PageSize)

	accounts, err = store.GetAccounts(context.Background(), ledger.AccountsQuery{
		PageSize: 10,
		Filters: ledger.AccountsQueryFilters{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, accounts.Data, 1)

	accounts, err = store.GetAccounts(context.Background(), ledger.AccountsQuery{
		PageSize: 10,
		Filters: ledger.AccountsQueryFilters{
			Metadata: map[string]string{
				"number": "3",
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, accounts.Data, 1)

	accounts, err = store.GetAccounts(context.Background(), ledger.AccountsQuery{
		PageSize: 10,
		Filters: ledger.AccountsQueryFilters{
			Metadata: map[string]string{
				"boolean": "true",
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, accounts.Data, 1)

	accounts, err = store.GetAccounts(context.Background(), ledger.AccountsQuery{
		PageSize: 10,
		Filters: ledger.AccountsQueryFilters{
			Metadata: map[string]string{
				"a.super.nested.key": "hello",
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, accounts.Data, 1)
}

func testTransactions(t *testing.T, store *ledgerstore.Store) {
	err := store.Commit(context.Background(), tx1, tx2, tx3)
	require.NoError(t, err)

	t.Run("Count", func(t *testing.T) {
		count, err := store.CountTransactions(context.Background(), ledger.TransactionsQuery{})
		require.NoError(t, err)
		// Should get all the transactions
		require.EqualValues(t, 3, count)

		count, err = store.CountTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				Account: "world",
			},
		})
		require.NoError(t, err)
		// Should get the two first transactions involving the 'world' account.
		require.EqualValues(t, 2, count)

		count, err = store.CountTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				Account:   "world",
				StartTime: now.Add(-2 * time.Hour),
				EndTime:   now.Add(-1 * time.Hour),
			},
		})
		require.NoError(t, err)
		// Should get only tx2, as StartTime is inclusive and EndTime exclusive.
		require.EqualValues(t, 1, count)

		count, err = store.CountTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				Metadata: map[string]string{
					"priority": "high",
				},
			},
		})
		require.NoError(t, err)
		require.EqualValues(t, 1, count)
	})

	t.Run("Get", func(t *testing.T) {
		cursor, err := store.GetTransactions(context.Background(), ledger.TransactionsQuery{
			PageSize: 1,
		})
		require.NoError(t, err)
		// Should get only the first transaction.
		require.Equal(t, 1, cursor.PageSize)

		cursor, err = store.GetTransactions(context.Background(), ledger.TransactionsQuery{
			AfterTxID: cursor.Data[0].ID,
			PageSize:  1,
		})
		require.NoError(t, err)
		// Should get only the second transaction.
		require.Equal(t, 1, cursor.PageSize)

		cursor, err = store.GetTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				Account:   "world",
				Reference: "tx1",
			},
			PageSize: 1,
		})
		require.NoError(t, err)
		require.Equal(t, 1, cursor.PageSize)
		// Should get only the first transaction.
		require.Len(t, cursor.Data, 1)

		cursor, err = store.GetTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				Account: "users:.*",
			},
			PageSize: 10,
		})
		require.NoError(t, err)
		require.Equal(t, 10, cursor.PageSize)
		require.Len(t, cursor.Data, 1)

		cursor, err = store.GetTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				Source: "central_bank",
			},
			PageSize: 10,
		})
		require.NoError(t, err)
		require.Equal(t, 10, cursor.PageSize)
		// Should get only the third transaction.
		require.Len(t, cursor.Data, 1)

		cursor, err = store.GetTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				Destination: "users:1",
			},
			PageSize: 10,
		})
		require.NoError(t, err)
		require.Equal(t, 10, cursor.PageSize)
		// Should get only the third transaction.
		require.Len(t, cursor.Data, 1)

		cursor, err = store.GetTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				Destination: "users:.*", // Use regex
			},
			PageSize: 10,
		})
		assert.NoError(t, err)
		assert.Equal(t, 10, cursor.PageSize)
		// Should get only the third transaction.
		assert.Len(t, cursor.Data, 1)

		cursor, err = store.GetTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				Destination: ".*:1", // Use regex
			},
			PageSize: 10,
		})
		assert.NoError(t, err)
		assert.Equal(t, 10, cursor.PageSize)
		// Should get only the third transaction.
		assert.Len(t, cursor.Data, 1)

		cursor, err = store.GetTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				Source: ".*bank", // Use regex
			},
			PageSize: 10,
		})
		assert.NoError(t, err)
		assert.Equal(t, 10, cursor.PageSize)
		// Should get only the third transaction.
		assert.Len(t, cursor.Data, 1)

		cursor, err = store.GetTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				StartTime: now.Add(-2 * time.Hour),
				EndTime:   now.Add(-1 * time.Hour),
			},
			PageSize: 10,
		})
		require.NoError(t, err)
		require.Equal(t, 10, cursor.PageSize)
		// Should get only tx2, as StartTime is inclusive and EndTime exclusive.
		require.Len(t, cursor.Data, 1)

		cursor, err = store.GetTransactions(context.Background(), ledger.TransactionsQuery{
			Filters: ledger.TransactionsQueryFilters{
				Metadata: map[string]string{
					"priority": "high",
				},
			},
			PageSize: 10,
		})
		require.NoError(t, err)
		require.Equal(t, 10, cursor.PageSize)
		// Should get only the third transaction.
		require.Len(t, cursor.Data, 1)
	})
}

func testGetTransaction(t *testing.T, store *ledgerstore.Store) {
	err := store.Commit(context.Background(), tx1, tx2)
	require.NoError(t, err)

	tx, err := store.GetTransaction(context.Background(), tx1.ID)
	require.NoError(t, err)
	require.Equal(t, tx1.Postings, tx.Postings)
	require.Equal(t, tx1.Reference, tx.Reference)
	require.Equal(t, tx1.Timestamp, tx.Timestamp)
}

func TestInitializeStore(t *testing.T) {
	driver, stopFn, err := ledgertesting.StorageDriver(t)
	require.NoError(t, err)
	defer stopFn()
	defer func(driver storage.Driver[*ledgerstore.Store], ctx context.Context) {
		require.NoError(t, driver.Close(ctx))
	}(driver, context.Background())

	err = driver.Initialize(context.Background())
	require.NoError(t, err)

	store, _, err := driver.GetLedgerStore(context.Background(), uuid.NewString(), true)
	require.NoError(t, err)

	modified, err := store.Initialize(context.Background())
	require.NoError(t, err)
	require.True(t, modified)

	modified, err = store.Initialize(context.Background())
	require.NoError(t, err)
	require.False(t, modified)
}

func testGetLastLog(t *testing.T, store *ledgerstore.Store) {
	err := store.Commit(context.Background(), tx1)
	require.NoError(t, err)
	logs := make([]core.Log, 0)
	logs = append(logs, core.NewTransactionLog(tx1.Transaction))
	errChan := store.AppendLogs(context.Background(), logs...)
	require.NoError(t, <-errChan)

	lastLog, err := store.GetLastLog(context.Background())
	require.NoError(t, err)
	require.NotNil(t, lastLog)

	require.Equal(t, tx1.Postings, lastLog.Data.(core.Transaction).Postings)
	require.Equal(t, tx1.Reference, lastLog.Data.(core.Transaction).Reference)
	require.Equal(t, tx1.Timestamp, lastLog.Data.(core.Transaction).Timestamp)
}

func testGetLogs(t *testing.T, store *ledgerstore.Store) {
	require.NoError(t, store.Commit(context.Background(), tx1, tx2, tx3))
	logs := make([]core.Log, 0)
	for _, tx := range []core.ExpandedTransaction{tx1, tx2, tx3} {
		logs = append(logs, core.NewTransactionLog(tx.Transaction))
	}
	errChan := store.AppendLogs(context.Background(), logs...)
	require.NoError(t, <-errChan)

	cursor, err := store.GetLogs(context.Background(), ledger.NewLogsQuery())
	require.NoError(t, err)
	require.Equal(t, ledger.QueryDefaultPageSize, cursor.PageSize)

	require.Equal(t, 3, len(cursor.Data))
	require.Equal(t, uint64(2), cursor.Data[0].ID)
	require.Equal(t, tx3.Postings, cursor.Data[0].Data.(core.Transaction).Postings)
	require.Equal(t, tx3.Reference, cursor.Data[0].Data.(core.Transaction).Reference)
	require.Equal(t, tx3.Timestamp, cursor.Data[0].Data.(core.Transaction).Timestamp)

	cursor, err = store.GetLogs(context.Background(), &ledger.LogsQuery{
		PageSize: 1,
	})
	require.NoError(t, err)
	// Should get only the first log.
	require.Equal(t, 1, cursor.PageSize)
	require.Equal(t, uint64(2), cursor.Data[0].ID)

	cursor, err = store.GetLogs(context.Background(), &ledger.LogsQuery{
		AfterID:  cursor.Data[0].ID,
		PageSize: 1,
	})
	require.NoError(t, err)
	// Should get only the second log.
	require.Equal(t, 1, cursor.PageSize)
	require.Equal(t, uint64(1), cursor.Data[0].ID)

	cursor, err = store.GetLogs(context.Background(), &ledger.LogsQuery{
		Filters: ledger.LogsQueryFilters{
			StartTime: now.Add(-2 * time.Hour),
			EndTime:   now.Add(-1 * time.Hour),
		},
		PageSize: 10,
	})
	require.NoError(t, err)
	require.Equal(t, 10, cursor.PageSize)
	// Should get only the second log, as StartTime is inclusive and EndTime exclusive.
	require.Len(t, cursor.Data, 1)
	require.Equal(t, uint64(1), cursor.Data[0].ID)
}