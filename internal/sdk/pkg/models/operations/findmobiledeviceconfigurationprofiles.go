// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindMobileDeviceConfigurationProfiles200ApplicationXMLConfigurationProfile struct {
	ID *int64
	// Name of the configuration profile
	Name *string
}

type FindMobileDeviceConfigurationProfiles200ApplicationXML struct {
	ConfigurationProfile *FindMobileDeviceConfigurationProfiles200ApplicationXMLConfigurationProfile
}

type FindMobileDeviceConfigurationProfiles200ApplicationJSONConfigurationProfile struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the configuration profile
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfiles200ApplicationJSON struct {
	ConfigurationProfile *FindMobileDeviceConfigurationProfiles200ApplicationJSONConfigurationProfile `json:"configuration_profile,omitempty"`
}

type FindMobileDeviceConfigurationProfilesResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDeviceConfigurationProfiles200ApplicationJSONObjects []FindMobileDeviceConfigurationProfiles200ApplicationJSON
}
