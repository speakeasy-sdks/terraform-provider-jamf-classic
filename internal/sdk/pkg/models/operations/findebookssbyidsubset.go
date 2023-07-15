// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindEbookssByIDSubsetSubset - Subset to filter by
type FindEbookssByIDSubsetSubset string

const (
	FindEbookssByIDSubsetSubsetGeneral     FindEbookssByIDSubsetSubset = "General"
	FindEbookssByIDSubsetSubsetScope       FindEbookssByIDSubsetSubset = "Scope"
	FindEbookssByIDSubsetSubsetSelfService FindEbookssByIDSubsetSubset = "SelfService"
)

func (e FindEbookssByIDSubsetSubset) ToPointer() *FindEbookssByIDSubsetSubset {
	return &e
}

func (e *FindEbookssByIDSubsetSubset) UnmarshalJSON(data []byte) error {
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
		*e = FindEbookssByIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindEbookssByIDSubsetSubset: %v", v)
	}
}

type FindEbookssByIDSubsetRequest struct {
	// ID to filter by
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Subset to filter by
	Subset FindEbookssByIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindEbookssByIDSubsetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}