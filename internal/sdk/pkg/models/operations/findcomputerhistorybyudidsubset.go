// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"jamf/internal/sdk/pkg/models/shared"
	"net/http"
)

// FindComputerHistoryByUDIDSubsetSubset - Subset to filter by
type FindComputerHistoryByUDIDSubsetSubset string

const (
	FindComputerHistoryByUDIDSubsetSubsetGeneral                 FindComputerHistoryByUDIDSubsetSubset = "General"
	FindComputerHistoryByUDIDSubsetSubsetComputerUsageLogs       FindComputerHistoryByUDIDSubsetSubset = "ComputerUsageLogs"
	FindComputerHistoryByUDIDSubsetSubsetAudits                  FindComputerHistoryByUDIDSubsetSubset = "Audits"
	FindComputerHistoryByUDIDSubsetSubsetPolicyLogs              FindComputerHistoryByUDIDSubsetSubset = "PolicyLogs"
	FindComputerHistoryByUDIDSubsetSubsetCasperRemoteLogs        FindComputerHistoryByUDIDSubsetSubset = "CasperRemoteLogs"
	FindComputerHistoryByUDIDSubsetSubsetScreenSharingLogs       FindComputerHistoryByUDIDSubsetSubset = "ScreenSharingLogs"
	FindComputerHistoryByUDIDSubsetSubsetCasperImagingLogs       FindComputerHistoryByUDIDSubsetSubset = "CasperImagingLogs"
	FindComputerHistoryByUDIDSubsetSubsetCommands                FindComputerHistoryByUDIDSubsetSubset = "Commands"
	FindComputerHistoryByUDIDSubsetSubsetUserLocation            FindComputerHistoryByUDIDSubsetSubset = "UserLocation"
	FindComputerHistoryByUDIDSubsetSubsetMacAppStoreApplications FindComputerHistoryByUDIDSubsetSubset = "MacAppStoreApplications"
)

func (e FindComputerHistoryByUDIDSubsetSubset) ToPointer() *FindComputerHistoryByUDIDSubsetSubset {
	return &e
}

func (e *FindComputerHistoryByUDIDSubsetSubset) UnmarshalJSON(data []byte) error {
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
		*e = FindComputerHistoryByUDIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHistoryByUDIDSubsetSubset: %v", v)
	}
}

type FindComputerHistoryByUDIDSubsetRequest struct {
	// Subset to filter by
	Subset FindComputerHistoryByUDIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
	// Computer UDID to filter by
	Udid string `pathParam:"style=simple,explode=false,name=udid"`
}

type FindComputerHistoryByUDIDSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ComputerHistory *shared.ComputerHistory
}
