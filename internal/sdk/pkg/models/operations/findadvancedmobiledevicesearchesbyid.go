// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FindAdvancedMobileDeviceSearchesByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterionAndOr string

const (
	FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterionAndOrAnd FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterionAndOr = "and"
	FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterionAndOrOr  FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterionAndOr = "or"
)

func (e FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterionAndOr) ToPointer() *FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterionAndOr {
	return &e
}

func (e *FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterionAndOr) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "and":
		fallthrough
	case "or":
		*e = FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterionAndOr(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterionAndOr: %v", v)
	}
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterion struct {
	AndOr        *FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterionAndOr
	ClosingParen *bool
	// Name of the criteria
	Name         *string
	OpeningParen *bool
	Priority     *int64
	// Operator
	SearchType *string
	Value      *string
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteria struct {
	Criterion *FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteriaCriterion
	Size      *int64
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationXMLDisplayFieldsDisplayField struct {
	// Name of the display field
	Name *string
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationXMLDisplayFields struct {
	DisplayField *FindAdvancedMobileDeviceSearchesByID200ApplicationXMLDisplayFieldsDisplayField
	Size         *int64
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationXMLMobileDevicesMobileDevice struct {
	DisplayName *string
	ID          *int64
	// Name of the mobile device
	Name *string
	Udid *string
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationXMLMobileDevices struct {
	MobileDevice *FindAdvancedMobileDeviceSearchesByID200ApplicationXMLMobileDevicesMobileDevice
	Size         *int64
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationXMLSite struct {
	ID *int64
	// Name of the site
	Name string
}

// FindAdvancedMobileDeviceSearchesByID200ApplicationXML - OK
type FindAdvancedMobileDeviceSearchesByID200ApplicationXML struct {
	Criteria      []FindAdvancedMobileDeviceSearchesByID200ApplicationXMLCriteria
	DisplayFields []FindAdvancedMobileDeviceSearchesByID200ApplicationXMLDisplayFields
	ID            *int64
	MobileDevices []FindAdvancedMobileDeviceSearchesByID200ApplicationXMLMobileDevices
	// Name of the advanced mobile device search
	Name   string
	Site   *FindAdvancedMobileDeviceSearchesByID200ApplicationXMLSite
	Sort1  *string
	Sort2  *string
	Sort3  *string
	ViewAs *string
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterionAndOr string

const (
	FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterionAndOrAnd FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterionAndOr = "and"
	FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterionAndOrOr  FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterionAndOr = "or"
)

func (e FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterionAndOr) ToPointer() *FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterionAndOr {
	return &e
}

func (e *FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterionAndOr) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "and":
		fallthrough
	case "or":
		*e = FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterionAndOr(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterionAndOr: %v", v)
	}
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterion struct {
	AndOr        *FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterionAndOr `json:"and_or,omitempty"`
	ClosingParen *bool                                                                         `json:"closing_paren,omitempty"`
	// Name of the criteria
	Name         *string `json:"name,omitempty"`
	OpeningParen *bool   `json:"opening_paren,omitempty"`
	Priority     *int64  `json:"priority,omitempty"`
	// Operator
	SearchType *string `json:"search_type,omitempty"`
	Value      *string `json:"value,omitempty"`
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteria struct {
	Criterion *FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteriaCriterion `json:"criterion,omitempty"`
	Size      *int64                                                                   `json:"size,omitempty"`
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationJSONDisplayFieldsDisplayField struct {
	// Name of the display field
	Name *string `json:"name,omitempty"`
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationJSONDisplayFields struct {
	DisplayField *FindAdvancedMobileDeviceSearchesByID200ApplicationJSONDisplayFieldsDisplayField `json:"display_field,omitempty"`
	Size         *int64                                                                           `json:"size,omitempty"`
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationJSONMobileDevicesMobileDevice struct {
	DisplayName *string `json:"Display_Name,omitempty"`
	ID          *int64  `json:"id,omitempty"`
	// Name of the mobile device
	Name *string `json:"name,omitempty"`
	Udid *string `json:"udid,omitempty"`
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationJSONMobileDevices struct {
	MobileDevice *FindAdvancedMobileDeviceSearchesByID200ApplicationJSONMobileDevicesMobileDevice `json:"mobile_device,omitempty"`
	Size         *int64                                                                           `json:"size,omitempty"`
}

type FindAdvancedMobileDeviceSearchesByID200ApplicationJSONSite struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the site
	Name string `json:"name"`
}

// FindAdvancedMobileDeviceSearchesByID200ApplicationJSON - OK
type FindAdvancedMobileDeviceSearchesByID200ApplicationJSON struct {
	Criteria      []FindAdvancedMobileDeviceSearchesByID200ApplicationJSONCriteria      `json:"criteria,omitempty"`
	DisplayFields []FindAdvancedMobileDeviceSearchesByID200ApplicationJSONDisplayFields `json:"display_fields,omitempty"`
	ID            *int64                                                                `json:"id,omitempty"`
	MobileDevices []FindAdvancedMobileDeviceSearchesByID200ApplicationJSONMobileDevices `json:"mobile_devices,omitempty"`
	// Name of the advanced mobile device search
	Name   string                                                      `json:"name"`
	Site   *FindAdvancedMobileDeviceSearchesByID200ApplicationJSONSite `json:"site,omitempty"`
	Sort1  *string                                                     `json:"sort_1,omitempty"`
	Sort2  *string                                                     `json:"sort_2,omitempty"`
	Sort3  *string                                                     `json:"sort_3,omitempty"`
	ViewAs *string                                                     `json:"view_as,omitempty"`
}

type FindAdvancedMobileDeviceSearchesByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindAdvancedMobileDeviceSearchesByID200ApplicationJSONObject *FindAdvancedMobileDeviceSearchesByID200ApplicationJSON
}
