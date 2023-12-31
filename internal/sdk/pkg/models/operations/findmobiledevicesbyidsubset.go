// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindMobileDevicesByIDSubsetSubset - Subset to filter by
type FindMobileDevicesByIDSubsetSubset string

const (
	FindMobileDevicesByIDSubsetSubsetGeneral               FindMobileDevicesByIDSubsetSubset = "General"
	FindMobileDevicesByIDSubsetSubsetLocation              FindMobileDevicesByIDSubsetSubset = "Location"
	FindMobileDevicesByIDSubsetSubsetPurchasing            FindMobileDevicesByIDSubsetSubset = "Purchasing"
	FindMobileDevicesByIDSubsetSubsetApplications          FindMobileDevicesByIDSubsetSubset = "Applications"
	FindMobileDevicesByIDSubsetSubsetSecurity              FindMobileDevicesByIDSubsetSubset = "Security"
	FindMobileDevicesByIDSubsetSubsetNetwork               FindMobileDevicesByIDSubsetSubset = "Network"
	FindMobileDevicesByIDSubsetSubsetCertificates          FindMobileDevicesByIDSubsetSubset = "Certificates"
	FindMobileDevicesByIDSubsetSubsetConfigurationProfiles FindMobileDevicesByIDSubsetSubset = "ConfigurationProfiles"
	FindMobileDevicesByIDSubsetSubsetProvisioningProfiles  FindMobileDevicesByIDSubsetSubset = "ProvisioningProfiles"
	FindMobileDevicesByIDSubsetSubsetMobileDeviceGroups    FindMobileDevicesByIDSubsetSubset = "MobileDeviceGroups"
	FindMobileDevicesByIDSubsetSubsetExtensionAttributes   FindMobileDevicesByIDSubsetSubset = "ExtensionAttributes"
)

func (e FindMobileDevicesByIDSubsetSubset) ToPointer() *FindMobileDevicesByIDSubsetSubset {
	return &e
}

func (e *FindMobileDevicesByIDSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "Location":
		fallthrough
	case "Purchasing":
		fallthrough
	case "Applications":
		fallthrough
	case "Security":
		fallthrough
	case "Network":
		fallthrough
	case "Certificates":
		fallthrough
	case "ConfigurationProfiles":
		fallthrough
	case "ProvisioningProfiles":
		fallthrough
	case "MobileDeviceGroups":
		fallthrough
	case "ExtensionAttributes":
		*e = FindMobileDevicesByIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDevicesByIDSubsetSubset: %v", v)
	}
}

type FindMobileDevicesByIDSubsetRequest struct {
	// ID to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Subset to filter by
	Subset FindMobileDevicesByIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindMobileDevicesByIDSubsetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
