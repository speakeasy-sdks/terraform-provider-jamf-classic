// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindComputerReports200ApplicationXMLComputerReport struct {
	ID   *int64
	Name *string
}

type FindComputerReports200ApplicationXML struct {
	ComputerReport *FindComputerReports200ApplicationXMLComputerReport
	Size           *int64
}

type FindComputerReports200ApplicationJSONComputerReport struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerReports200ApplicationJSON struct {
	ComputerReport *FindComputerReports200ApplicationJSONComputerReport `json:"computer_report,omitempty"`
	Size           *int64                                               `json:"size,omitempty"`
}

type FindComputerReportsResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindComputerReports200ApplicationJSONObjects []FindComputerReports200ApplicationJSON
}
