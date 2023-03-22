// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance/sdks/go/pkg/models/shared"
	"net/http"
)

type ConnectorsTransferRequest struct {
	TransferRequest shared.TransferRequest `request:"mediaType=application/json"`
	// The name of the connector.
	Connector shared.ConnectorEnum `pathParam:"style=simple,explode=false,name=connector"`
}

type ConnectorsTransferResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	TransferResponse *shared.TransferResponse
}
