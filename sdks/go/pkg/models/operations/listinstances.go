// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type ListInstancesRequest struct {
	// Filter running instances
	Running *bool `queryParam:"style=form,explode=true,name=running"`
	// A workflow id
	WorkflowID *string `queryParam:"style=form,explode=true,name=workflowID"`
}

type ListInstancesResponse struct {
	ContentType string
	// General error
	Error *shared.Error
	// List of workflow instances
	ListRunsResponse interface{}
	StatusCode       int
	RawResponse      *http.Response
}