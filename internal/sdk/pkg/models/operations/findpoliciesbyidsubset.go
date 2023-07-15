// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindPoliciesByIDSubsetSubset - Subset to filter by
type FindPoliciesByIDSubsetSubset string

const (
	FindPoliciesByIDSubsetSubsetGeneral              FindPoliciesByIDSubsetSubset = "General"
	FindPoliciesByIDSubsetSubsetScope                FindPoliciesByIDSubsetSubset = "Scope"
	FindPoliciesByIDSubsetSubsetSelfService          FindPoliciesByIDSubsetSubset = "SelfService"
	FindPoliciesByIDSubsetSubsetPackageConfiguration FindPoliciesByIDSubsetSubset = "PackageConfiguration"
	FindPoliciesByIDSubsetSubsetScripts              FindPoliciesByIDSubsetSubset = "Scripts"
	FindPoliciesByIDSubsetSubsetPrinters             FindPoliciesByIDSubsetSubset = "Printers"
	FindPoliciesByIDSubsetSubsetDockItems            FindPoliciesByIDSubsetSubset = "DockItems"
	FindPoliciesByIDSubsetSubsetAccountMaintenance   FindPoliciesByIDSubsetSubset = "AccountMaintenance"
	FindPoliciesByIDSubsetSubsetReboot               FindPoliciesByIDSubsetSubset = "Reboot"
	FindPoliciesByIDSubsetSubsetMaintenance          FindPoliciesByIDSubsetSubset = "Maintenance"
	FindPoliciesByIDSubsetSubsetFilesProcesses       FindPoliciesByIDSubsetSubset = "FilesProcesses"
	FindPoliciesByIDSubsetSubsetUserInteraction      FindPoliciesByIDSubsetSubset = "UserInteraction"
	FindPoliciesByIDSubsetSubsetDiskEncryption       FindPoliciesByIDSubsetSubset = "DiskEncryption"
)

func (e FindPoliciesByIDSubsetSubset) ToPointer() *FindPoliciesByIDSubsetSubset {
	return &e
}

func (e *FindPoliciesByIDSubsetSubset) UnmarshalJSON(data []byte) error {
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
	case "PackageConfiguration":
		fallthrough
	case "Scripts":
		fallthrough
	case "Printers":
		fallthrough
	case "DockItems":
		fallthrough
	case "AccountMaintenance":
		fallthrough
	case "Reboot":
		fallthrough
	case "Maintenance":
		fallthrough
	case "FilesProcesses":
		fallthrough
	case "UserInteraction":
		fallthrough
	case "DiskEncryption":
		*e = FindPoliciesByIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindPoliciesByIDSubsetSubset: %v", v)
	}
}

type FindPoliciesByIDSubsetRequest struct {
	// ID to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Subset to filter by
	Subset FindPoliciesByIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindPoliciesByIDSubsetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}