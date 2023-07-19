// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindSitesByNameRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

// FindSitesByName200ApplicationXML - OK
type FindSitesByName200ApplicationXML struct {
	ID *int64
	// Name of the site
	Name string
}

// FindSitesByName200ApplicationJSON - OK
type FindSitesByName200ApplicationJSON struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the site
	Name string `json:"name"`
}

type FindSitesByNameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindSitesByName200ApplicationJSONObject *FindSitesByName200ApplicationJSON
}
