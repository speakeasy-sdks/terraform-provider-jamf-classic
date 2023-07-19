// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindPatchExternalSources200ApplicationXMLPatchExternalSource struct {
	ID   *int64
	Name *string
}

type FindPatchExternalSources200ApplicationXML struct {
	PatchExternalSource *FindPatchExternalSources200ApplicationXMLPatchExternalSource
	Size                *int64
}

type FindPatchExternalSources200ApplicationJSONPatchExternalSource struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchExternalSources200ApplicationJSON struct {
	PatchExternalSource *FindPatchExternalSources200ApplicationJSONPatchExternalSource `json:"patch_external_source,omitempty"`
	Size                *int64                                                         `json:"size,omitempty"`
}

type FindPatchExternalSourcesResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindPatchExternalSources200ApplicationJSONObjects []FindPatchExternalSources200ApplicationJSON
}
