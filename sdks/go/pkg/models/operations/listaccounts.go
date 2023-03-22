// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/stack/sdks/go/pkg/models/shared"
	"net/http"
)

// ListAccountsBalanceOperatorEnum - Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not.
type ListAccountsBalanceOperatorEnum string

const (
	ListAccountsBalanceOperatorEnumGte ListAccountsBalanceOperatorEnum = "gte"
	ListAccountsBalanceOperatorEnumLte ListAccountsBalanceOperatorEnum = "lte"
	ListAccountsBalanceOperatorEnumGt  ListAccountsBalanceOperatorEnum = "gt"
	ListAccountsBalanceOperatorEnumLt  ListAccountsBalanceOperatorEnum = "lt"
	ListAccountsBalanceOperatorEnumE   ListAccountsBalanceOperatorEnum = "e"
	ListAccountsBalanceOperatorEnumNe  ListAccountsBalanceOperatorEnum = "ne"
)

type ListAccountsRequest struct {
	// Filter accounts by address pattern (regular expression placed between ^ and $).
	Address *string `queryParam:"style=form,explode=true,name=address"`
	// Pagination cursor, will return accounts after given address, in descending order.
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Filter accounts by their balance (default operator is gte)
	Balance *int64 `queryParam:"style=form,explode=true,name=balance"`
	// Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not.
	//
	BalanceOperator *ListAccountsBalanceOperatorEnum `queryParam:"style=form,explode=true,name=balanceOperator"`
	// Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not.
	// Deprecated, please use `balanceOperator` instead.
	//
	BalanceOperatorDeprecated *ListAccountsBalanceOperatorEnum `queryParam:"style=form,explode=true,name=balance_operator"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below.
	Metadata map[string]interface{} `queryParam:"style=deepObject,explode=true,name=metadata"`
	// The maximum number of results to return per page.
	//
	PageSize *int64 `queryParam:"style=form,explode=true,name=pageSize"`
	// The maximum number of results to return per page.
	// Deprecated, please use `pageSize` instead.
	//
	PageSizeDeprecated *int64 `queryParam:"style=form,explode=true,name=page_size"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	// Deprecated, please use `cursor` instead.
	//
	PaginationToken *string `queryParam:"style=form,explode=true,name=pagination_token"`
}

type ListAccountsResponse struct {
	// OK
	AccountsCursorResponse *shared.AccountsCursorResponse
	ContentType            string
	// Error
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
}
