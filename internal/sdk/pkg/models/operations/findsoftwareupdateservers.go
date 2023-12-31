// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindSoftwareUpdateServers200ApplicationXMLSoftwareUpdateServer struct {
	ID   *int64
	Name *string
}

type FindSoftwareUpdateServers200ApplicationXML struct {
	Size                 *int64
	SoftwareUpdateServer *FindSoftwareUpdateServers200ApplicationXMLSoftwareUpdateServer
}

type FindSoftwareUpdateServers200ApplicationJSONSoftwareUpdateServer struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindSoftwareUpdateServers200ApplicationJSON struct {
	Size                 *int64                                                           `json:"size,omitempty"`
	SoftwareUpdateServer *FindSoftwareUpdateServers200ApplicationJSONSoftwareUpdateServer `json:"software_update_server,omitempty"`
}

type FindSoftwareUpdateServersResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindSoftwareUpdateServers200ApplicationJSONObjects []FindSoftwareUpdateServers200ApplicationJSON
}
