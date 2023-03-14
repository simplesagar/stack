package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	sdk "github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/stack/libs/go-libs/metadata"
	wallet "github.com/formancehq/wallets/pkg"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestWalletSummary(t *testing.T) {
	t.Parallel()

	w := wallet.NewWallet(uuid.NewString(), "default", metadata.Metadata{})

	req := newRequest(t, http.MethodGet, "/wallets/"+w.ID+"/summary", nil)
	rec := httptest.NewRecorder()

	coupon1Balance := wallet.NewBalance("coupon1", ptr(time.Now().Add(-time.Minute).Round(time.Second).UTC()))
	coupon2Balance := wallet.NewBalance("coupon2", ptr(time.Now().Add(time.Minute).Round(time.Second).UTC()))
	hold1 := wallet.NewDebitHold(w.ID, wallet.NewLedgerAccountSubject("bank"), "USD", "", metadata.Metadata{})
	hold2 := wallet.NewDebitHold(w.ID, wallet.NewLedgerAccountSubject("bank"), "USD", "", metadata.Metadata{})

	var testEnv *testEnv
	testEnv = newTestEnv(
		WithGetAccount(func(ctx context.Context, ledger, account string) (*sdk.AccountWithVolumesAndBalances, error) {
			require.Equal(t, testEnv.LedgerName(), ledger)
			switch account {
			case testEnv.Chart().GetMainBalanceAccount(w.ID):
				return &sdk.AccountWithVolumesAndBalances{
					Address:  account,
					Metadata: w.LedgerMetadata(),
					Balances: ptr(map[string]int64{
						"USD": 100,
					}),
				}, nil
			case testEnv.Chart().GetBalanceAccount(w.ID, coupon1Balance.Name):
				return &sdk.AccountWithVolumesAndBalances{
					Address:  account,
					Metadata: coupon1Balance.LedgerMetadata(w.ID),
					Balances: ptr(map[string]int64{
						"USD": 10,
					}),
				}, nil
			case testEnv.Chart().GetBalanceAccount(w.ID, coupon2Balance.Name):
				return &sdk.AccountWithVolumesAndBalances{
					Address:  account,
					Metadata: coupon2Balance.LedgerMetadata(w.ID),
					Balances: ptr(map[string]int64{
						"USD": 20,
					}),
				}, nil
			case testEnv.Chart().GetHoldAccount(hold1.ID):
				return &sdk.AccountWithVolumesAndBalances{
					Address:  account,
					Metadata: hold1.LedgerMetadata(testEnv.Chart()),
					Balances: ptr(map[string]int64{
						"USD": 10,
					}),
				}, nil
			case testEnv.Chart().GetHoldAccount(hold2.ID):
				return &sdk.AccountWithVolumesAndBalances{
					Address:  account,
					Metadata: hold2.LedgerMetadata(testEnv.Chart()),
					Balances: ptr(map[string]int64{
						"USD": 20,
					}),
				}, nil
			default:
				require.Fail(t, "unexpected account query")
			}
			panic("should not happen")
		}),
		WithListAccounts(func(ctx context.Context, ledger string, query wallet.ListAccountsQuery) (*sdk.AccountsCursorResponseCursor, error) {
			switch {
			case query.Metadata[wallet.MetadataKeyWalletID] == w.ID:
				return &sdk.AccountsCursorResponseCursor{
					Data: []sdk.Account{
						{
							Address:  testEnv.Chart().GetMainBalanceAccount(w.ID),
							Metadata: w.LedgerMetadata(),
						},
						{
							Address:  testEnv.Chart().GetBalanceAccount(w.ID, "coupon1"),
							Metadata: coupon1Balance.LedgerMetadata(w.ID),
						},
						{
							Address:  testEnv.Chart().GetBalanceAccount(w.ID, "coupon2"),
							Metadata: coupon2Balance.LedgerMetadata(w.ID),
						},
					},
				}, nil
			case query.Metadata[wallet.MetadataKeyHoldWalletID] == w.ID:
				return &sdk.AccountsCursorResponseCursor{
					Data: []sdk.Account{
						{
							Address:  testEnv.Chart().GetHoldAccount(hold1.ID),
							Metadata: hold1.LedgerMetadata(testEnv.Chart()),
						},
						{
							Address:  testEnv.Chart().GetHoldAccount(hold2.ID),
							Metadata: hold2.LedgerMetadata(testEnv.Chart()),
						},
					},
				}, nil
			default:
				require.Fail(t, "unexpected list accounts query")
			}
			panic("should not happen")
		}),
	)
	testEnv.Router().ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Result().StatusCode)
	summary := wallet.Summary{}
	readResponse(t, rec, &summary)

	require.Equal(t, wallet.Summary{
		Balances: []wallet.ExpandedBalance{
			{
				Balance: wallet.Balance{
					Name: "main",
				},
				Assets: map[string]int64{
					"USD": 100,
				},
			},
			{
				Balance: coupon1Balance,
				Assets: map[string]int64{
					"USD": 10,
				},
			},
			{
				Balance: coupon2Balance,
				Assets: map[string]int64{
					"USD": 20,
				},
			},
		},
		AvailableFunds: map[string]int64{
			"USD": 120,
		},
		ExpiredFunds: map[string]int64{
			"USD": 10,
		},
		ExpirableFunds: map[string]int64{
			"USD": 20,
		},
		HoldFunds: map[string]int64{
			"USD": 30,
		},
	}, summary)
}
