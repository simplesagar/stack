// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountsCursorCursor struct {
	Data     []PaymentsAccount `json:"data"`
	HasMore  bool              `json:"hasMore"`
	Next     *string           `json:"next,omitempty"`
	PageSize int64             `json:"pageSize"`
	Previous *string           `json:"previous,omitempty"`
}

// AccountsCursor - OK
type AccountsCursor struct {
	Cursor AccountsCursorCursor `json:"cursor"`
}
