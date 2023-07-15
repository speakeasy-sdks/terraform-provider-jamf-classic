// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

type FindComputerApplicationUsageBySerialRequest struct {
	// End date (e.g. yyyy-mm-dd)
	EndDate string `pathParam:"style=simple,explode=false,name=end_date"`
	// Serial number to filter by
	Serialnumber string `pathParam:"style=simple,explode=false,name=serialnumber"`
	// Start date (e.g. yyyy-mm-dd)
	StartDate string `pathParam:"style=simple,explode=false,name=start_date"`
}

type FindComputerApplicationUsageBySerialResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerApplicationUsage []shared.ComputerApplicationUsage
}
