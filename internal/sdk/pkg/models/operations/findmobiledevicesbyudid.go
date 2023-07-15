// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindMobileDevicesByUDIDRequest struct {
	// UDID to filter by
	Udid string `pathParam:"style=simple,explode=false,name=udid"`
}

type FindMobileDevicesByUDIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MobileDevice *shared.MobileDevice
}