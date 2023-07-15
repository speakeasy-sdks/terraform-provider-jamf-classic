// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteMobileDeviceGroupByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteMobileDeviceGroupByIDResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
