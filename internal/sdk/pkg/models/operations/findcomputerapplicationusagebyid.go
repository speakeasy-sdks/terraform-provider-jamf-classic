// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"jamf/internal/sdk/pkg/models/shared"
	"jamf/internal/sdk/pkg/types"
	"net/http"
)

type FindComputerApplicationUsageByIDRequest struct {
	// End date (e.g. yyyy-mm-dd)
	EndDate types.Date `pathParam:"style=simple,explode=false,name=end_date"`
	// ID value to filter by
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Start date (e.g. yyyy-mm-dd)
	StartDate types.Date `pathParam:"style=simple,explode=false,name=start_date"`
}

type FindComputerApplicationUsageByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerApplicationUsage []shared.ComputerApplicationUsage
}