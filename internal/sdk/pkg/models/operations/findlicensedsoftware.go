// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindLicensedSoftware200ApplicationXMLLicensedSoftware struct {
	ID *int64
	// Name of the licensed software
	Name *string
}

type FindLicensedSoftware200ApplicationXML struct {
	LicensedSoftware *FindLicensedSoftware200ApplicationXMLLicensedSoftware
}

type FindLicensedSoftware200ApplicationJSONLicensedSoftware struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the licensed software
	Name *string `json:"name,omitempty"`
}

type FindLicensedSoftware200ApplicationJSON struct {
	LicensedSoftware *FindLicensedSoftware200ApplicationJSONLicensedSoftware `json:"licensed_software,omitempty"`
}

type FindLicensedSoftwareResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindLicensedSoftware200ApplicationJSONObjects []FindLicensedSoftware200ApplicationJSON
}
