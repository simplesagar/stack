// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/stack/sdks/go/pkg/models/shared"
	"net/http"
)

type ListClientsResponse struct {
	ContentType string
	// List of clients
	ListClientsResponse *shared.ListClientsResponse
	StatusCode          int
	RawResponse         *http.Response
}
