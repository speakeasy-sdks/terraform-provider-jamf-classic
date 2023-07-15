// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindPatchExternalSourcesByIDRequest struct {
	// ID value to filter by
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type FindPatchExternalSourcesByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	PatchExternalSource *shared.PatchExternalSource
}