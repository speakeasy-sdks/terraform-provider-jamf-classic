// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindManagedPreferenceProfilesByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type FindManagedPreferenceProfilesByID200ApplicationXMLGeneral struct {
	Enabled *bool
	ID      *int64
	Name    string
	Plist   *string
}

// FindManagedPreferenceProfilesByID200ApplicationXML - OK
type FindManagedPreferenceProfilesByID200ApplicationXML struct {
	General *FindManagedPreferenceProfilesByID200ApplicationXMLGeneral
}

type FindManagedPreferenceProfilesByID200ApplicationJSONGeneral struct {
	Enabled *bool   `json:"enabled,omitempty"`
	ID      *int64  `json:"id,omitempty"`
	Name    string  `json:"name"`
	Plist   *string `json:"plist,omitempty"`
}

// FindManagedPreferenceProfilesByID200ApplicationJSON - OK
type FindManagedPreferenceProfilesByID200ApplicationJSON struct {
	General *FindManagedPreferenceProfilesByID200ApplicationJSONGeneral `json:"general,omitempty"`
}

type FindManagedPreferenceProfilesByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindManagedPreferenceProfilesByID200ApplicationJSONObject *FindManagedPreferenceProfilesByID200ApplicationJSON
}
