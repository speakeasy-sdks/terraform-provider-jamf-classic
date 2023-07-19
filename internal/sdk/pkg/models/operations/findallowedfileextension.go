// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindAllowedFileExtension200ApplicationXMLAllowedFileExtension struct {
	// File extension
	Extension string
	ID        *int64
}

type FindAllowedFileExtension200ApplicationXML struct {
	AllowedFileExtension *FindAllowedFileExtension200ApplicationXMLAllowedFileExtension
	Size                 *int64
}

type FindAllowedFileExtension200ApplicationJSONAllowedFileExtension struct {
	// File extension
	Extension string `json:"extension"`
	ID        *int64 `json:"id,omitempty"`
}

type FindAllowedFileExtension200ApplicationJSON struct {
	AllowedFileExtension *FindAllowedFileExtension200ApplicationJSONAllowedFileExtension `json:"allowed_file_extension,omitempty"`
	Size                 *int64                                                          `json:"size,omitempty"`
}

type FindAllowedFileExtensionResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindAllowedFileExtension200ApplicationJSONObjects []FindAllowedFileExtension200ApplicationJSON
}
