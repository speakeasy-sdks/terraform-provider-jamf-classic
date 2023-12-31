// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UpdateSoftwareTitlesByNameRequest struct {
	// Name value to update by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type UpdateSoftwareTitlesByNameResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
