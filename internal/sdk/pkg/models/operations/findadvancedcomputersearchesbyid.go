// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FindAdvancedComputerSearchesByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type FindAdvancedComputerSearchesByID200ApplicationXMLComputersComputer struct {
	ComputerName *string
	ID           *int64
	// Name of the computer
	Name *string
	Udid *string
}

type FindAdvancedComputerSearchesByID200ApplicationXMLComputers struct {
	Computer *FindAdvancedComputerSearchesByID200ApplicationXMLComputersComputer
	Size     *int64
}

type FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterionAndOr string

const (
	FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterionAndOrAnd FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterionAndOr = "and"
	FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterionAndOrOr  FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterionAndOr = "or"
)

func (e FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterionAndOr) ToPointer() *FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterionAndOr {
	return &e
}

func (e *FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterionAndOr) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "and":
		fallthrough
	case "or":
		*e = FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterionAndOr(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterionAndOr: %v", v)
	}
}

type FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterion struct {
	AndOr        *FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterionAndOr
	ClosingParen *bool
	// Name of the criteria
	Name         *string
	OpeningParen *bool
	Priority     *int64
	// Operator
	SearchType *string
	Value      *string
}

type FindAdvancedComputerSearchesByID200ApplicationXMLCriteria struct {
	Criterion *FindAdvancedComputerSearchesByID200ApplicationXMLCriteriaCriterion
	Size      *int64
}

type FindAdvancedComputerSearchesByID200ApplicationXMLDisplayFieldsDisplayField struct {
	// Name of the display field
	Name *string
}

type FindAdvancedComputerSearchesByID200ApplicationXMLDisplayFields struct {
	DisplayField *FindAdvancedComputerSearchesByID200ApplicationXMLDisplayFieldsDisplayField
	Size         *int64
}

type FindAdvancedComputerSearchesByID200ApplicationXMLSite struct {
	ID *int64
	// Name of the site
	Name string
}

// FindAdvancedComputerSearchesByID200ApplicationXML - OK
type FindAdvancedComputerSearchesByID200ApplicationXML struct {
	Computers     []FindAdvancedComputerSearchesByID200ApplicationXMLComputers
	Criteria      []FindAdvancedComputerSearchesByID200ApplicationXMLCriteria
	DisplayFields []FindAdvancedComputerSearchesByID200ApplicationXMLDisplayFields
	ID            *int64
	// Name of the advanced computer search
	Name   string
	Site   *FindAdvancedComputerSearchesByID200ApplicationXMLSite
	Sort1  *string
	Sort2  *string
	Sort3  *string
	ViewAs *string
}

type FindAdvancedComputerSearchesByID200ApplicationJSONComputersComputer struct {
	ComputerName *string `json:"Computer_Name,omitempty"`
	ID           *int64  `json:"id,omitempty"`
	// Name of the computer
	Name *string `json:"name,omitempty"`
	Udid *string `json:"udid,omitempty"`
}

type FindAdvancedComputerSearchesByID200ApplicationJSONComputers struct {
	Computer *FindAdvancedComputerSearchesByID200ApplicationJSONComputersComputer `json:"computer,omitempty"`
	Size     *int64                                                               `json:"size,omitempty"`
}

type FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterionAndOr string

const (
	FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterionAndOrAnd FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterionAndOr = "and"
	FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterionAndOrOr  FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterionAndOr = "or"
)

func (e FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterionAndOr) ToPointer() *FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterionAndOr {
	return &e
}

func (e *FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterionAndOr) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "and":
		fallthrough
	case "or":
		*e = FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterionAndOr(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterionAndOr: %v", v)
	}
}

type FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterion struct {
	AndOr        *FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterionAndOr `json:"and_or,omitempty"`
	ClosingParen *bool                                                                     `json:"closing_paren,omitempty"`
	// Name of the criteria
	Name         *string `json:"name,omitempty"`
	OpeningParen *bool   `json:"opening_paren,omitempty"`
	Priority     *int64  `json:"priority,omitempty"`
	// Operator
	SearchType *string `json:"search_type,omitempty"`
	Value      *string `json:"value,omitempty"`
}

type FindAdvancedComputerSearchesByID200ApplicationJSONCriteria struct {
	Criterion *FindAdvancedComputerSearchesByID200ApplicationJSONCriteriaCriterion `json:"criterion,omitempty"`
	Size      *int64                                                               `json:"size,omitempty"`
}

type FindAdvancedComputerSearchesByID200ApplicationJSONDisplayFieldsDisplayField struct {
	// Name of the display field
	Name *string `json:"name,omitempty"`
}

type FindAdvancedComputerSearchesByID200ApplicationJSONDisplayFields struct {
	DisplayField *FindAdvancedComputerSearchesByID200ApplicationJSONDisplayFieldsDisplayField `json:"display_field,omitempty"`
	Size         *int64                                                                       `json:"size,omitempty"`
}

type FindAdvancedComputerSearchesByID200ApplicationJSONSite struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the site
	Name string `json:"name"`
}

// FindAdvancedComputerSearchesByID200ApplicationJSON - OK
type FindAdvancedComputerSearchesByID200ApplicationJSON struct {
	Computers     []FindAdvancedComputerSearchesByID200ApplicationJSONComputers     `json:"computers,omitempty"`
	Criteria      []FindAdvancedComputerSearchesByID200ApplicationJSONCriteria      `json:"criteria,omitempty"`
	DisplayFields []FindAdvancedComputerSearchesByID200ApplicationJSONDisplayFields `json:"display_fields,omitempty"`
	ID            *int64                                                            `json:"id,omitempty"`
	// Name of the advanced computer search
	Name   string                                                  `json:"name"`
	Site   *FindAdvancedComputerSearchesByID200ApplicationJSONSite `json:"site,omitempty"`
	Sort1  *string                                                 `json:"sort_1,omitempty"`
	Sort2  *string                                                 `json:"sort_2,omitempty"`
	Sort3  *string                                                 `json:"sort_3,omitempty"`
	ViewAs *string                                                 `json:"view_as,omitempty"`
}

type FindAdvancedComputerSearchesByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindAdvancedComputerSearchesByID200ApplicationJSONObject *FindAdvancedComputerSearchesByID200ApplicationJSON
}
