// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindMobileDeviceApplications200ApplicationXMLMobileDeviceApplication struct {
	BundleID    *string
	DisplayName *string
	ID          *int64
	InternalApp *bool
	Name        *string
	Version     *string
}

type FindMobileDeviceApplications200ApplicationXML struct {
	MobileDeviceApplication *FindMobileDeviceApplications200ApplicationXMLMobileDeviceApplication
}

type FindMobileDeviceApplications200ApplicationJSONMobileDeviceApplication struct {
	BundleID    *string `json:"bundle_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	ID          *int64  `json:"id,omitempty"`
	InternalApp *bool   `json:"internal_app,omitempty"`
	Name        *string `json:"name,omitempty"`
	Version     *string `json:"version,omitempty"`
}

type FindMobileDeviceApplications200ApplicationJSON struct {
	MobileDeviceApplication *FindMobileDeviceApplications200ApplicationJSONMobileDeviceApplication `json:"mobile_device_application,omitempty"`
}

type FindMobileDeviceApplicationsResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDeviceApplications200ApplicationJSONObjects []FindMobileDeviceApplications200ApplicationJSON
}
