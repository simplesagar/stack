// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type GetInstanceStageHistoryRequest struct {
	// The instance id
	InstanceID string `pathParam:"style=simple,explode=false,name=instanceID"`
	// The stage number
	Number int64 `pathParam:"style=simple,explode=false,name=number"`
}

type GetInstanceStageHistoryResponse struct {
	ContentType string
	// General error
	Error *shared.Error
	// The workflow instance stage history
	GetWorkflowInstanceHistoryStageResponse *shared.GetWorkflowInstanceHistoryStageResponse
	StatusCode                              int
	RawResponse                             *http.Response
}
