// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindUserextensionattributes200ApplicationXMLUserExtensionAttribute struct {
	ID   *int64
	Name *string
}

type FindUserextensionattributes200ApplicationXML struct {
	Size                   *int64
	UserExtensionAttribute *FindUserextensionattributes200ApplicationXMLUserExtensionAttribute
}

type FindUserextensionattributes200ApplicationJSONUserExtensionAttribute struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindUserextensionattributes200ApplicationJSON struct {
	Size                   *int64                                                               `json:"size,omitempty"`
	UserExtensionAttribute *FindUserextensionattributes200ApplicationJSONUserExtensionAttribute `json:"user_extension_attribute,omitempty"`
}

type FindUserextensionattributesResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindUserextensionattributes200ApplicationJSONObjects []FindUserextensionattributes200ApplicationJSON
}
