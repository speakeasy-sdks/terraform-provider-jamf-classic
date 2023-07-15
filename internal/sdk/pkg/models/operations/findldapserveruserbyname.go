// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindLDAPServerUserByNameRequest struct {
	// Server name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// User to filter by
	User string `pathParam:"style=simple,explode=false,name=user"`
}

type FindLDAPServerUserByNameResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}