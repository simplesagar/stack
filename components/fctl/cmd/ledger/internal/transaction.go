package internal

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func TransactionIDOrLastN(ctx context.Context, ledgerClient *formance.Formance, ledger, id string) (int64, error) {
	if strings.HasPrefix(id, "last") {
		id = strings.TrimPrefix(id, "last")
		sub := int64(0)
		if id != "" {
			var err error
			sub, err = strconv.ParseInt(id, 10, 64)
			if err != nil {
				return 0, err
			}
		}
		pageSize := int64(1)
		request := operations.ListTransactionsRequest{
			Ledger:   ledger,
			PageSize: &pageSize,
		}
		response, err := ledgerClient.Ledger.ListTransactions(ctx, request)
		if err != nil {
			return 0, err
		}
		if len(response.TransactionsCursorResponse.Cursor.Data) == 0 {
			return 0, errors.New("no transaction found")
		}
		return response.TransactionsCursorResponse.Cursor.Data[0].Txid + sub, nil
	}

	return strconv.ParseInt(id, 10, 64)
}
