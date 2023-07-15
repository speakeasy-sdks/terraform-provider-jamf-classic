// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateMobileDeviceProvisioningProfilesByIDRequest struct {
	// Id value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateMobileDeviceProvisioningProfilesByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	MobileDeviceProvisioningProfile *shared.MobileDeviceProvisioningProfile
}
