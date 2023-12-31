// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FindUsersByNameRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeType string

const (
	FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeTypeString  FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeType = "String"
	FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeTypeInteger FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeType = "Integer"
	FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeTypeDate    FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeType = "Date"
)

func (e FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeType) ToPointer() *FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeType {
	return &e
}

func (e *FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "String":
		fallthrough
	case "Integer":
		fallthrough
	case "Date":
		*e = FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeType: %v", v)
	}
}

type FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttribute struct {
	ID    *int64
	Name  *string
	Type  *FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttributeType
	Value *string
}

type FindUsersByName200ApplicationXMLExtensionAttributes struct {
	ExtensionAttribute *FindUsersByName200ApplicationXMLExtensionAttributesExtensionAttribute
}

type FindUsersByName200ApplicationXMLLdapServer struct {
	ID   *int64
	Name *string
}

type FindUsersByName200ApplicationXMLLinksComputersComputer struct {
	ID   *int64
	Name *string
}

type FindUsersByName200ApplicationXMLLinksComputers struct {
	Computer *FindUsersByName200ApplicationXMLLinksComputersComputer
}

type FindUsersByName200ApplicationXMLLinksMobileDevicesMobileDevice struct {
	ID   *int64
	Name *string
}

type FindUsersByName200ApplicationXMLLinksMobileDevices struct {
	MobileDevice *FindUsersByName200ApplicationXMLLinksMobileDevicesMobileDevice
}

type FindUsersByName200ApplicationXMLLinksPeripheralsPeripheral struct {
	ID   *int64
	Name *string
}

type FindUsersByName200ApplicationXMLLinksPeripherals struct {
	Peripheral *FindUsersByName200ApplicationXMLLinksPeripheralsPeripheral
}

type FindUsersByName200ApplicationXMLLinksVppAssignmentsVppAssignment struct {
	ID   *int64
	Name *string
}

type FindUsersByName200ApplicationXMLLinksVppAssignments struct {
	VppAssignment *FindUsersByName200ApplicationXMLLinksVppAssignmentsVppAssignment
}

type FindUsersByName200ApplicationXMLLinks struct {
	Computers         *FindUsersByName200ApplicationXMLLinksComputers
	MobileDevices     *FindUsersByName200ApplicationXMLLinksMobileDevices
	Peripherals       *FindUsersByName200ApplicationXMLLinksPeripherals
	TotalVppCodeCount *int64
	VppAssignments    *FindUsersByName200ApplicationXMLLinksVppAssignments
}

type FindUsersByName200ApplicationXMLSitesSite struct {
	ID *int64
	// Name of the site
	Name string
}

type FindUsersByName200ApplicationXMLSites struct {
	Site *FindUsersByName200ApplicationXMLSitesSite
}

// FindUsersByName200ApplicationXML - OK
type FindUsersByName200ApplicationXML struct {
	CustomPhotoURL       *string
	Email                *string
	EmailAddress         *string
	EnableCustomPhotoURL *bool
	ExtensionAttributes  []FindUsersByName200ApplicationXMLExtensionAttributes
	FullName             *string
	ID                   *int64
	LdapServer           *FindUsersByName200ApplicationXMLLdapServer
	Links                *FindUsersByName200ApplicationXMLLinks
	// Name of the user
	Name        string
	PhoneNumber *string
	Position    *string
	Sites       []FindUsersByName200ApplicationXMLSites
}

type FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeType string

const (
	FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeTypeString  FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeType = "String"
	FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeTypeInteger FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeType = "Integer"
	FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeTypeDate    FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeType = "Date"
)

func (e FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeType) ToPointer() *FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeType {
	return &e
}

func (e *FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "String":
		fallthrough
	case "Integer":
		fallthrough
	case "Date":
		*e = FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeType: %v", v)
	}
}

type FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttribute struct {
	ID    *int64                                                                      `json:"id,omitempty"`
	Name  *string                                                                     `json:"name,omitempty"`
	Type  *FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttributeType `json:"type,omitempty"`
	Value *string                                                                     `json:"value,omitempty"`
}

type FindUsersByName200ApplicationJSONExtensionAttributes struct {
	ExtensionAttribute *FindUsersByName200ApplicationJSONExtensionAttributesExtensionAttribute `json:"extension_attribute,omitempty"`
}

type FindUsersByName200ApplicationJSONLdapServer struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindUsersByName200ApplicationJSONLinksComputersComputer struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindUsersByName200ApplicationJSONLinksComputers struct {
	Computer *FindUsersByName200ApplicationJSONLinksComputersComputer `json:"computer,omitempty"`
}

type FindUsersByName200ApplicationJSONLinksMobileDevicesMobileDevice struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindUsersByName200ApplicationJSONLinksMobileDevices struct {
	MobileDevice *FindUsersByName200ApplicationJSONLinksMobileDevicesMobileDevice `json:"mobile_device,omitempty"`
}

type FindUsersByName200ApplicationJSONLinksPeripheralsPeripheral struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindUsersByName200ApplicationJSONLinksPeripherals struct {
	Peripheral *FindUsersByName200ApplicationJSONLinksPeripheralsPeripheral `json:"peripheral,omitempty"`
}

type FindUsersByName200ApplicationJSONLinksVppAssignmentsVppAssignment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindUsersByName200ApplicationJSONLinksVppAssignments struct {
	VppAssignment *FindUsersByName200ApplicationJSONLinksVppAssignmentsVppAssignment `json:"vpp_assignment,omitempty"`
}

type FindUsersByName200ApplicationJSONLinks struct {
	Computers         *FindUsersByName200ApplicationJSONLinksComputers      `json:"computers,omitempty"`
	MobileDevices     *FindUsersByName200ApplicationJSONLinksMobileDevices  `json:"mobile_devices,omitempty"`
	Peripherals       *FindUsersByName200ApplicationJSONLinksPeripherals    `json:"peripherals,omitempty"`
	TotalVppCodeCount *int64                                                `json:"total_vpp_code_count,omitempty"`
	VppAssignments    *FindUsersByName200ApplicationJSONLinksVppAssignments `json:"vpp_assignments,omitempty"`
}

type FindUsersByName200ApplicationJSONSitesSite struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the site
	Name string `json:"name"`
}

type FindUsersByName200ApplicationJSONSites struct {
	Site *FindUsersByName200ApplicationJSONSitesSite `json:"site,omitempty"`
}

// FindUsersByName200ApplicationJSON - OK
type FindUsersByName200ApplicationJSON struct {
	CustomPhotoURL       *string                                                `json:"custom_photo_url,omitempty"`
	Email                *string                                                `json:"email,omitempty"`
	EmailAddress         *string                                                `json:"email_address,omitempty"`
	EnableCustomPhotoURL *bool                                                  `json:"enable_custom_photo_url,omitempty"`
	ExtensionAttributes  []FindUsersByName200ApplicationJSONExtensionAttributes `json:"extension_attributes,omitempty"`
	FullName             *string                                                `json:"full_name,omitempty"`
	ID                   *int64                                                 `json:"id,omitempty"`
	LdapServer           *FindUsersByName200ApplicationJSONLdapServer           `json:"ldap_server,omitempty"`
	Links                *FindUsersByName200ApplicationJSONLinks                `json:"links,omitempty"`
	// Name of the user
	Name        string                                   `json:"name"`
	PhoneNumber *string                                  `json:"phone_number,omitempty"`
	Position    *string                                  `json:"position,omitempty"`
	Sites       []FindUsersByName200ApplicationJSONSites `json:"sites,omitempty"`
}

type FindUsersByNameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindUsersByName200ApplicationJSONObject *FindUsersByName200ApplicationJSON
}
