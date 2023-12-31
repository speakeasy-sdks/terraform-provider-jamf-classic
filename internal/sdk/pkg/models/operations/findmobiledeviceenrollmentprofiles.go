// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindMobileDeviceEnrollmentProfiles200ApplicationXMLMobileDeviceEnrollmentProfile struct {
	ID         *int64
	Invitation *string
	// Name of the enrollment profile
	Name *string
}

type FindMobileDeviceEnrollmentProfiles200ApplicationXML struct {
	MobileDeviceEnrollmentProfile *FindMobileDeviceEnrollmentProfiles200ApplicationXMLMobileDeviceEnrollmentProfile
	Size                          *int64
}

type FindMobileDeviceEnrollmentProfiles200ApplicationJSONMobileDeviceEnrollmentProfile struct {
	ID         *int64  `json:"id,omitempty"`
	Invitation *string `json:"invitation,omitempty"`
	// Name of the enrollment profile
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceEnrollmentProfiles200ApplicationJSON struct {
	MobileDeviceEnrollmentProfile *FindMobileDeviceEnrollmentProfiles200ApplicationJSONMobileDeviceEnrollmentProfile `json:"mobile_device_enrollment_profile,omitempty"`
	Size                          *int64                                                                             `json:"size,omitempty"`
}

type FindMobileDeviceEnrollmentProfilesResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDeviceEnrollmentProfiles200ApplicationJSONObjects []FindMobileDeviceEnrollmentProfiles200ApplicationJSON
}
