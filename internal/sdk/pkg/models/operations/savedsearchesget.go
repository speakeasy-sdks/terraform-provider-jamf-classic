// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SavedsearchesGet200ApplicationXMLSavedSearchType string

const (
	SavedsearchesGet200ApplicationXMLSavedSearchTypeComputers     SavedsearchesGet200ApplicationXMLSavedSearchType = "Computers"
	SavedsearchesGet200ApplicationXMLSavedSearchTypeMobileDevices SavedsearchesGet200ApplicationXMLSavedSearchType = "Mobile Devices"
	SavedsearchesGet200ApplicationXMLSavedSearchTypeUsers         SavedsearchesGet200ApplicationXMLSavedSearchType = "Users"
)

func (e SavedsearchesGet200ApplicationXMLSavedSearchType) ToPointer() *SavedsearchesGet200ApplicationXMLSavedSearchType {
	return &e
}

func (e *SavedsearchesGet200ApplicationXMLSavedSearchType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Computers":
		fallthrough
	case "Mobile Devices":
		fallthrough
	case "Users":
		*e = SavedsearchesGet200ApplicationXMLSavedSearchType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedsearchesGet200ApplicationXMLSavedSearchType: %v", v)
	}
}

type SavedsearchesGet200ApplicationXMLSavedSearch struct {
	ID   *int64
	Name *string
	Type *SavedsearchesGet200ApplicationXMLSavedSearchType
}

type SavedsearchesGet200ApplicationXML struct {
	SavedSearch *SavedsearchesGet200ApplicationXMLSavedSearch
	Size        *int64
}

type SavedsearchesGet200ApplicationJSONSavedSearchType string

const (
	SavedsearchesGet200ApplicationJSONSavedSearchTypeComputers     SavedsearchesGet200ApplicationJSONSavedSearchType = "Computers"
	SavedsearchesGet200ApplicationJSONSavedSearchTypeMobileDevices SavedsearchesGet200ApplicationJSONSavedSearchType = "Mobile Devices"
	SavedsearchesGet200ApplicationJSONSavedSearchTypeUsers         SavedsearchesGet200ApplicationJSONSavedSearchType = "Users"
)

func (e SavedsearchesGet200ApplicationJSONSavedSearchType) ToPointer() *SavedsearchesGet200ApplicationJSONSavedSearchType {
	return &e
}

func (e *SavedsearchesGet200ApplicationJSONSavedSearchType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Computers":
		fallthrough
	case "Mobile Devices":
		fallthrough
	case "Users":
		*e = SavedsearchesGet200ApplicationJSONSavedSearchType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SavedsearchesGet200ApplicationJSONSavedSearchType: %v", v)
	}
}

type SavedsearchesGet200ApplicationJSONSavedSearch struct {
	ID   *int64                                             `json:"id,omitempty"`
	Name *string                                            `json:"name,omitempty"`
	Type *SavedsearchesGet200ApplicationJSONSavedSearchType `json:"type,omitempty"`
}

type SavedsearchesGet200ApplicationJSON struct {
	SavedSearch *SavedsearchesGet200ApplicationJSONSavedSearch `json:"saved_search,omitempty"`
	Size        *int64                                         `json:"size,omitempty"`
}

type SavedsearchesGetResponse struct {
	Body        []byte
	ContentType string
	// OK
	SavedsearchesGet200ApplicationJSONObjects []SavedsearchesGet200ApplicationJSON
	StatusCode                                int
	RawResponse                               *http.Response
}
