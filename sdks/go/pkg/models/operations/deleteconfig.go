// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteConfigRequest struct {
	// Config ID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteConfigResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}