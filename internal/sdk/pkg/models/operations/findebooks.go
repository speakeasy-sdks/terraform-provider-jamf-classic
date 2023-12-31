// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindEBooks200ApplicationXMLEbook struct {
	ID   *int64
	Name *string
}

type FindEBooks200ApplicationXML struct {
	Ebook *FindEBooks200ApplicationXMLEbook
	Size  *int64
}

type FindEBooks200ApplicationJSONEbook struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindEBooks200ApplicationJSON struct {
	Ebook *FindEBooks200ApplicationJSONEbook `json:"ebook,omitempty"`
	Size  *int64                             `json:"size,omitempty"`
}

type FindEBooksResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindEBooks200ApplicationJSONObjects []FindEBooks200ApplicationJSON
}
