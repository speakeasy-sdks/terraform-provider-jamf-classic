// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteMobiledeviceextensionattributeByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteMobiledeviceextensionattributeByIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
