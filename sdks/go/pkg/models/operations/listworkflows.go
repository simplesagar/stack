// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance/sdks/go/pkg/models/shared"
	"net/http"
)

type ListWorkflowsResponse struct {
	ContentType string
	// General error
	Error *shared.Error
	// List of workflows
	ListWorkflowsResponse *shared.ListWorkflowsResponse
	StatusCode            int
	RawResponse           *http.Response
}
