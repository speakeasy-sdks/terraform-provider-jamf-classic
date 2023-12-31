// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PatchsoftwaretitlesIDByIDGetRequest struct {
	// ID value to filter by
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type PatchsoftwaretitlesIDByIDGet200ApplicationJSONCategory struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the category
	Name string `json:"name"`
}

type PatchsoftwaretitlesIDByIDGet200ApplicationJSONNotifications struct {
	EmailNotification *bool `json:"email_notification,omitempty"`
	WebNotification   *bool `json:"web_notification,omitempty"`
}

type PatchsoftwaretitlesIDByIDGet200ApplicationJSONVersionsVersionPackage struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type PatchsoftwaretitlesIDByIDGet200ApplicationJSONVersionsVersion struct {
	Package         *PatchsoftwaretitlesIDByIDGet200ApplicationJSONVersionsVersionPackage `json:"package,omitempty"`
	SoftwareVersion *string                                                               `json:"software_version,omitempty"`
}

type PatchsoftwaretitlesIDByIDGet200ApplicationJSONVersions struct {
	Version *PatchsoftwaretitlesIDByIDGet200ApplicationJSONVersionsVersion `json:"version,omitempty"`
}

// PatchsoftwaretitlesIDByIDGet200ApplicationJSON - OK
type PatchsoftwaretitlesIDByIDGet200ApplicationJSON struct {
	Categories    *PatchsoftwaretitlesIDByIDGet200ApplicationJSONCategory      `json:"categories,omitempty"`
	ID            *int64                                                       `json:"id,omitempty"`
	Name          *string                                                      `json:"name,omitempty"`
	NameID        *string                                                      `json:"name_id,omitempty"`
	Notifications *PatchsoftwaretitlesIDByIDGet200ApplicationJSONNotifications `json:"notifications,omitempty"`
	SourceID      *int64                                                       `json:"source_id,omitempty"`
	Versions      []PatchsoftwaretitlesIDByIDGet200ApplicationJSONVersions     `json:"versions,omitempty"`
}

type PatchsoftwaretitlesIDByIDGetResponse struct {
	ContentType string
	// OK
	PatchsoftwaretitlesIDByIDGet200ApplicationJSONObject *PatchsoftwaretitlesIDByIDGet200ApplicationJSON
	StatusCode                                           int
	RawResponse                                          *http.Response
}
