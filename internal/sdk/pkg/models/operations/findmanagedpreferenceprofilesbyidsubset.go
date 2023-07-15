// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindManagedPreferenceProfilesByIDSubsetSubset - Subset to filter by
type FindManagedPreferenceProfilesByIDSubsetSubset string

const (
	FindManagedPreferenceProfilesByIDSubsetSubsetGeneral  FindManagedPreferenceProfilesByIDSubsetSubset = "General"
	FindManagedPreferenceProfilesByIDSubsetSubsetScope    FindManagedPreferenceProfilesByIDSubsetSubset = "Scope"
	FindManagedPreferenceProfilesByIDSubsetSubsetSettings FindManagedPreferenceProfilesByIDSubsetSubset = "Settings"
)

func (e FindManagedPreferenceProfilesByIDSubsetSubset) ToPointer() *FindManagedPreferenceProfilesByIDSubsetSubset {
	return &e
}

func (e *FindManagedPreferenceProfilesByIDSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "Scope":
		fallthrough
	case "Settings":
		*e = FindManagedPreferenceProfilesByIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindManagedPreferenceProfilesByIDSubsetSubset: %v", v)
	}
}

type FindManagedPreferenceProfilesByIDSubsetRequest struct {
	// ID to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Subset to filter by
	Subset FindManagedPreferenceProfilesByIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindManagedPreferenceProfilesByIDSubsetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
