// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CreateAdvancedComputerSearchgByIDRequest struct {
	// ID to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type CreateAdvancedComputerSearchgByIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
