// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FlushComputerExtenstionAttributeDataRequest struct {
	// ID of the computer extension attribute for which data will be deleted
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type FlushComputerExtenstionAttributeDataResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}