// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteMobileDeviceGroupByNameRequest struct {
	// Name value to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type DeleteMobileDeviceGroupByNameResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
