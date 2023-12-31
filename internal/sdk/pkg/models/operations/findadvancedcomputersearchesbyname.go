// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FindAdvancedComputerSearchesByNameRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type FindAdvancedComputerSearchesByName200ApplicationXMLComputersComputer struct {
	ComputerName *string
	ID           *int64
	// Name of the computer
	Name *string
	Udid *string
}

type FindAdvancedComputerSearchesByName200ApplicationXMLComputers struct {
	Computer *FindAdvancedComputerSearchesByName200ApplicationXMLComputersComputer
	Size     *int64
}

type FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterionAndOr string

const (
	FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterionAndOrAnd FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterionAndOr = "and"
	FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterionAndOrOr  FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterionAndOr = "or"
)

func (e FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterionAndOr) ToPointer() *FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterionAndOr {
	return &e
}

func (e *FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterionAndOr) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "and":
		fallthrough
	case "or":
		*e = FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterionAndOr(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterionAndOr: %v", v)
	}
}

type FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterion struct {
	AndOr        *FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterionAndOr
	ClosingParen *bool
	// Name of the criteria
	Name         *string
	OpeningParen *bool
	Priority     *int64
	// Operator
	SearchType *string
	Value      *string
}

type FindAdvancedComputerSearchesByName200ApplicationXMLCriteria struct {
	Criterion *FindAdvancedComputerSearchesByName200ApplicationXMLCriteriaCriterion
	Size      *int64
}

type FindAdvancedComputerSearchesByName200ApplicationXMLDisplayFieldsDisplayField struct {
	// Name of the display field
	Name *string
}

type FindAdvancedComputerSearchesByName200ApplicationXMLDisplayFields struct {
	DisplayField *FindAdvancedComputerSearchesByName200ApplicationXMLDisplayFieldsDisplayField
	Size         *int64
}

type FindAdvancedComputerSearchesByName200ApplicationXMLSite struct {
	ID *int64
	// Name of the site
	Name string
}

// FindAdvancedComputerSearchesByName200ApplicationXML - OK
type FindAdvancedComputerSearchesByName200ApplicationXML struct {
	Computers     []FindAdvancedComputerSearchesByName200ApplicationXMLComputers
	Criteria      []FindAdvancedComputerSearchesByName200ApplicationXMLCriteria
	DisplayFields []FindAdvancedComputerSearchesByName200ApplicationXMLDisplayFields
	ID            *int64
	// Name of the advanced computer search
	Name   string
	Site   *FindAdvancedComputerSearchesByName200ApplicationXMLSite
	Sort1  *string
	Sort2  *string
	Sort3  *string
	ViewAs *string
}

type FindAdvancedComputerSearchesByName200ApplicationJSONComputersComputer struct {
	ComputerName *string `json:"Computer_Name,omitempty"`
	ID           *int64  `json:"id,omitempty"`
	// Name of the computer
	Name *string `json:"name,omitempty"`
	Udid *string `json:"udid,omitempty"`
}

type FindAdvancedComputerSearchesByName200ApplicationJSONComputers struct {
	Computer *FindAdvancedComputerSearchesByName200ApplicationJSONComputersComputer `json:"computer,omitempty"`
	Size     *int64                                                                 `json:"size,omitempty"`
}

type FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterionAndOr string

const (
	FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterionAndOrAnd FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterionAndOr = "and"
	FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterionAndOrOr  FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterionAndOr = "or"
)

func (e FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterionAndOr) ToPointer() *FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterionAndOr {
	return &e
}

func (e *FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterionAndOr) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "and":
		fallthrough
	case "or":
		*e = FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterionAndOr(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterionAndOr: %v", v)
	}
}

type FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterion struct {
	AndOr        *FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterionAndOr `json:"and_or,omitempty"`
	ClosingParen *bool                                                                       `json:"closing_paren,omitempty"`
	// Name of the criteria
	Name         *string `json:"name,omitempty"`
	OpeningParen *bool   `json:"opening_paren,omitempty"`
	Priority     *int64  `json:"priority,omitempty"`
	// Operator
	SearchType *string `json:"search_type,omitempty"`
	Value      *string `json:"value,omitempty"`
}

type FindAdvancedComputerSearchesByName200ApplicationJSONCriteria struct {
	Criterion *FindAdvancedComputerSearchesByName200ApplicationJSONCriteriaCriterion `json:"criterion,omitempty"`
	Size      *int64                                                                 `json:"size,omitempty"`
}

type FindAdvancedComputerSearchesByName200ApplicationJSONDisplayFieldsDisplayField struct {
	// Name of the display field
	Name *string `json:"name,omitempty"`
}

type FindAdvancedComputerSearchesByName200ApplicationJSONDisplayFields struct {
	DisplayField *FindAdvancedComputerSearchesByName200ApplicationJSONDisplayFieldsDisplayField `json:"display_field,omitempty"`
	Size         *int64                                                                         `json:"size,omitempty"`
}

type FindAdvancedComputerSearchesByName200ApplicationJSONSite struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the site
	Name string `json:"name"`
}

// FindAdvancedComputerSearchesByName200ApplicationJSON - OK
type FindAdvancedComputerSearchesByName200ApplicationJSON struct {
	Computers     []FindAdvancedComputerSearchesByName200ApplicationJSONComputers     `json:"computers,omitempty"`
	Criteria      []FindAdvancedComputerSearchesByName200ApplicationJSONCriteria      `json:"criteria,omitempty"`
	DisplayFields []FindAdvancedComputerSearchesByName200ApplicationJSONDisplayFields `json:"display_fields,omitempty"`
	ID            *int64                                                              `json:"id,omitempty"`
	// Name of the advanced computer search
	Name   string                                                    `json:"name"`
	Site   *FindAdvancedComputerSearchesByName200ApplicationJSONSite `json:"site,omitempty"`
	Sort1  *string                                                   `json:"sort_1,omitempty"`
	Sort2  *string                                                   `json:"sort_2,omitempty"`
	Sort3  *string                                                   `json:"sort_3,omitempty"`
	ViewAs *string                                                   `json:"view_as,omitempty"`
}

type FindAdvancedComputerSearchesByNameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindAdvancedComputerSearchesByName200ApplicationJSONObject *FindAdvancedComputerSearchesByName200ApplicationJSON
}
