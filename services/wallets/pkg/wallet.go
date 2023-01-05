package wallet

import (
	"encoding/json"
	"net/http"

	"github.com/formancehq/go-libs/metadata"
	"github.com/google/uuid"
)

type ListWallets struct {
	Metadata metadata.Metadata
	Name     string
}

type PatchRequest struct {
	Metadata metadata.Metadata `json:"metadata"`
}

func (c *PatchRequest) Bind(r *http.Request) error {
	return nil
}

type CreateRequest struct {
	PatchRequest
	Name string `json:"name"`
}

func (c *CreateRequest) Bind(r *http.Request) error {
	return nil
}

type Wallet struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Metadata metadata.Metadata `json:"metadata"`
}

type WithBalances struct {
	Wallet
	Balances map[string]int32 `json:"balances"`
}

func (w *WithBalances) UnmarshalJSON(data []byte) error {
	type view struct {
		Wallet
		Balances struct {
			Main ExpandedBalance `json:"main"`
		} `json:"balances"`
	}
	v := view{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*w = WithBalances{
		Wallet:   v.Wallet,
		Balances: v.Balances.Main.Assets,
	}
	return nil
}

func (w WithBalances) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Wallet
		Balances struct {
			Main ExpandedBalance `json:"main"`
		} `json:"balances"`
	}{
		Wallet: w.Wallet,
		Balances: struct {
			Main ExpandedBalance `json:"main"`
		}{
			Main: ExpandedBalance{
				Assets: w.Balances,
			},
		},
	})
}

func (w Wallet) LedgerMetadata() metadata.Metadata {
	return metadata.Metadata{
		MetadataKeyWalletSpecType:   PrimaryWallet,
		MetadataKeyWalletName:       w.Name,
		MetadataKeyWalletCustomData: map[string]any(w.Metadata),
		MetadataKeyWalletID:         w.ID,
		MetadataKeyWalletBalance:    TrueValue,
		MetadataKeyBalanceName:      MainBalance,
	}
}

func NewWallet(name string, m metadata.Metadata) Wallet {
	if m == nil {
		m = metadata.Metadata{}
	}
	return Wallet{
		ID:       uuid.NewString(),
		Metadata: m,
		Name:     name,
	}
}

func FromAccount(account metadata.Owner) Wallet {
	return Wallet{
		ID:       GetMetadata(account, MetadataKeyWalletID).(string),
		Name:     GetMetadata(account, MetadataKeyWalletName).(string),
		Metadata: GetMetadata(account, MetadataKeyWalletCustomData).(map[string]any),
	}
}

func WithBalancesFromAccount(account interface {
	metadata.Owner
	GetBalances() map[string]int32
},
) WithBalances {
	return WithBalances{
		Wallet:   FromAccount(account),
		Balances: account.GetBalances(),
	}
}
