// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

// FindComputerHistoryByIDSubsetSubset - Subset to filter by
type FindComputerHistoryByIDSubsetSubset string

const (
	FindComputerHistoryByIDSubsetSubsetGeneral                 FindComputerHistoryByIDSubsetSubset = "General"
	FindComputerHistoryByIDSubsetSubsetComputerUsageLogs       FindComputerHistoryByIDSubsetSubset = "ComputerUsageLogs"
	FindComputerHistoryByIDSubsetSubsetAudits                  FindComputerHistoryByIDSubsetSubset = "Audits"
	FindComputerHistoryByIDSubsetSubsetPolicyLogs              FindComputerHistoryByIDSubsetSubset = "PolicyLogs"
	FindComputerHistoryByIDSubsetSubsetCasperRemoteLogs        FindComputerHistoryByIDSubsetSubset = "CasperRemoteLogs"
	FindComputerHistoryByIDSubsetSubsetScreenSharingLogs       FindComputerHistoryByIDSubsetSubset = "ScreenSharingLogs"
	FindComputerHistoryByIDSubsetSubsetCasperImagingLogs       FindComputerHistoryByIDSubsetSubset = "CasperImagingLogs"
	FindComputerHistoryByIDSubsetSubsetCommands                FindComputerHistoryByIDSubsetSubset = "Commands"
	FindComputerHistoryByIDSubsetSubsetUserLocation            FindComputerHistoryByIDSubsetSubset = "UserLocation"
	FindComputerHistoryByIDSubsetSubsetMacAppStoreApplications FindComputerHistoryByIDSubsetSubset = "MacAppStoreApplications"
)

func (e FindComputerHistoryByIDSubsetSubset) ToPointer() *FindComputerHistoryByIDSubsetSubset {
	return &e
}

func (e *FindComputerHistoryByIDSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "ComputerUsageLogs":
		fallthrough
	case "Audits":
		fallthrough
	case "PolicyLogs":
		fallthrough
	case "CasperRemoteLogs":
		fallthrough
	case "ScreenSharingLogs":
		fallthrough
	case "CasperImagingLogs":
		fallthrough
	case "Commands":
		fallthrough
	case "UserLocation":
		fallthrough
	case "MacAppStoreApplications":
		*e = FindComputerHistoryByIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHistoryByIDSubsetSubset: %v", v)
	}
}

type FindComputerHistoryByIDSubsetRequest struct {
	// Computer ID to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Subset to filter by
	Subset FindComputerHistoryByIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindComputerHistoryByIDSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerHistory *shared.ComputerHistory
}
