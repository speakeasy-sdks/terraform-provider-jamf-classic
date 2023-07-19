// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindDirectoryBindings200ApplicationXMLDirectoryBinding struct {
	ID   *int64
	Name *string
}

type FindDirectoryBindings200ApplicationXML struct {
	DirectoryBinding *FindDirectoryBindings200ApplicationXMLDirectoryBinding
	Size             *int64
}

type FindDirectoryBindings200ApplicationJSONDirectoryBinding struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindDirectoryBindings200ApplicationJSON struct {
	DirectoryBinding *FindDirectoryBindings200ApplicationJSONDirectoryBinding `json:"directory_binding,omitempty"`
	Size             *int64                                                   `json:"size,omitempty"`
}

type FindDirectoryBindingsResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindDirectoryBindings200ApplicationJSONObjects []FindDirectoryBindings200ApplicationJSON
}
