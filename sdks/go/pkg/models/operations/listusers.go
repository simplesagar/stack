// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance/sdks/go/pkg/models/shared"
	"net/http"
)

type ListUsersResponse struct {
	ContentType string
	// List of users
	ListUsersResponse *shared.ListUsersResponse
	StatusCode        int
	RawResponse       *http.Response
}
