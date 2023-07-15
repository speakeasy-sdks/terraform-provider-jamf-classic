// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindMobileDeviceAppsByIDSubsetSubset - Subset to filter by
type FindMobileDeviceAppsByIDSubsetSubset string

const (
	FindMobileDeviceAppsByIDSubsetSubsetGeneral          FindMobileDeviceAppsByIDSubsetSubset = "General"
	FindMobileDeviceAppsByIDSubsetSubsetScope            FindMobileDeviceAppsByIDSubsetSubset = "Scope"
	FindMobileDeviceAppsByIDSubsetSubsetSelfService      FindMobileDeviceAppsByIDSubsetSubset = "SelfService"
	FindMobileDeviceAppsByIDSubsetSubsetVppCodes         FindMobileDeviceAppsByIDSubsetSubset = "VPPCodes"
	FindMobileDeviceAppsByIDSubsetSubsetVpp              FindMobileDeviceAppsByIDSubsetSubset = "VPP"
	FindMobileDeviceAppsByIDSubsetSubsetAppConfiguration FindMobileDeviceAppsByIDSubsetSubset = "AppConfiguration"
)

func (e FindMobileDeviceAppsByIDSubsetSubset) ToPointer() *FindMobileDeviceAppsByIDSubsetSubset {
	return &e
}

func (e *FindMobileDeviceAppsByIDSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "Scope":
		fallthrough
	case "SelfService":
		fallthrough
	case "VPPCodes":
		fallthrough
	case "VPP":
		fallthrough
	case "AppConfiguration":
		*e = FindMobileDeviceAppsByIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceAppsByIDSubsetSubset: %v", v)
	}
}

type FindMobileDeviceAppsByIDSubsetRequest struct {
	// ID to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Subset to filter by
	Subset FindMobileDeviceAppsByIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindMobileDeviceAppsByIDSubsetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
