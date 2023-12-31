// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindRemovableMacAddressesByNameRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

// FindRemovableMacAddressesByName200ApplicationXML - OK
type FindRemovableMacAddressesByName200ApplicationXML struct {
	ID *int64
	// MAC address to ignore when identifying computers
	Name *string
}

// FindRemovableMacAddressesByName200ApplicationJSON - OK
type FindRemovableMacAddressesByName200ApplicationJSON struct {
	ID *int64 `json:"id,omitempty"`
	// MAC address to ignore when identifying computers
	Name *string `json:"name,omitempty"`
}

type FindRemovableMacAddressesByNameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindRemovableMacAddressesByName200ApplicationJSONObject *FindRemovableMacAddressesByName200ApplicationJSON
}
