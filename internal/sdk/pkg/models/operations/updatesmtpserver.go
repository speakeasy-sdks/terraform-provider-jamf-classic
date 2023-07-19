// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UpdateSMTPServerRequestBody struct {
	AuthorizationRequired *bool
	Enabled               *bool
	// Hostname or IP address of the SMTP server
	Host          *string
	Password      *string
	Port          *int64
	SendFromEmail *string
	SendFromName  *string
	Ssl           *bool
	// Measured in seconds
	Timeout  *int64
	TLS      *bool
	Username *string
}

type UpdateSMTPServerResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
