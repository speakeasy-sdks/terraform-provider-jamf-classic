// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindLDAPServersByNameRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type FindLDAPServersByNameResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
