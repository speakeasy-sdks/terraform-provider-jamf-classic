// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FindDiskEncryptionConfigurationsByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type FindDiskEncryptionConfigurationsByID200ApplicationXMLFileVaultEnabledUsers string

const (
	FindDiskEncryptionConfigurationsByID200ApplicationXMLFileVaultEnabledUsersCurrentOrNextUser FindDiskEncryptionConfigurationsByID200ApplicationXMLFileVaultEnabledUsers = "Current or Next User"
	FindDiskEncryptionConfigurationsByID200ApplicationXMLFileVaultEnabledUsersManagementAccount FindDiskEncryptionConfigurationsByID200ApplicationXMLFileVaultEnabledUsers = "Management Account"
)

func (e FindDiskEncryptionConfigurationsByID200ApplicationXMLFileVaultEnabledUsers) ToPointer() *FindDiskEncryptionConfigurationsByID200ApplicationXMLFileVaultEnabledUsers {
	return &e
}

func (e *FindDiskEncryptionConfigurationsByID200ApplicationXMLFileVaultEnabledUsers) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Current or Next User":
		fallthrough
	case "Management Account":
		*e = FindDiskEncryptionConfigurationsByID200ApplicationXMLFileVaultEnabledUsers(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindDiskEncryptionConfigurationsByID200ApplicationXMLFileVaultEnabledUsers: %v", v)
	}
}

type FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyType string

const (
	FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyTypeIndividual                 FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyType = "Individual"
	FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyTypeInstitutional              FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyType = "Institutional"
	FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyTypeIndividualAndInstitutional FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyType = "Individual And Institutional"
)

func (e FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyType) ToPointer() *FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyType {
	return &e
}

func (e *FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Individual":
		fallthrough
	case "Institutional":
		fallthrough
	case "Individual And Institutional":
		*e = FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyType: %v", v)
	}
}

// FindDiskEncryptionConfigurationsByID200ApplicationXML - OK
type FindDiskEncryptionConfigurationsByID200ApplicationXML struct {
	FileVaultEnabledUsers *FindDiskEncryptionConfigurationsByID200ApplicationXMLFileVaultEnabledUsers
	ID                    *int64
	KeyType               *FindDiskEncryptionConfigurationsByID200ApplicationXMLKeyType
	// Name of the disk encryption configuration
	Name string
}

type FindDiskEncryptionConfigurationsByID200ApplicationJSONFileVaultEnabledUsers string

const (
	FindDiskEncryptionConfigurationsByID200ApplicationJSONFileVaultEnabledUsersCurrentOrNextUser FindDiskEncryptionConfigurationsByID200ApplicationJSONFileVaultEnabledUsers = "Current or Next User"
	FindDiskEncryptionConfigurationsByID200ApplicationJSONFileVaultEnabledUsersManagementAccount FindDiskEncryptionConfigurationsByID200ApplicationJSONFileVaultEnabledUsers = "Management Account"
)

func (e FindDiskEncryptionConfigurationsByID200ApplicationJSONFileVaultEnabledUsers) ToPointer() *FindDiskEncryptionConfigurationsByID200ApplicationJSONFileVaultEnabledUsers {
	return &e
}

func (e *FindDiskEncryptionConfigurationsByID200ApplicationJSONFileVaultEnabledUsers) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Current or Next User":
		fallthrough
	case "Management Account":
		*e = FindDiskEncryptionConfigurationsByID200ApplicationJSONFileVaultEnabledUsers(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindDiskEncryptionConfigurationsByID200ApplicationJSONFileVaultEnabledUsers: %v", v)
	}
}

type FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyType string

const (
	FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyTypeIndividual                 FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyType = "Individual"
	FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyTypeInstitutional              FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyType = "Institutional"
	FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyTypeIndividualAndInstitutional FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyType = "Individual And Institutional"
)

func (e FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyType) ToPointer() *FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyType {
	return &e
}

func (e *FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Individual":
		fallthrough
	case "Institutional":
		fallthrough
	case "Individual And Institutional":
		*e = FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyType: %v", v)
	}
}

// FindDiskEncryptionConfigurationsByID200ApplicationJSON - OK
type FindDiskEncryptionConfigurationsByID200ApplicationJSON struct {
	FileVaultEnabledUsers *FindDiskEncryptionConfigurationsByID200ApplicationJSONFileVaultEnabledUsers `json:"file_vault_enabled_users,omitempty"`
	ID                    *int64                                                                       `json:"id,omitempty"`
	KeyType               *FindDiskEncryptionConfigurationsByID200ApplicationJSONKeyType               `json:"key_type,omitempty"`
	// Name of the disk encryption configuration
	Name string `json:"name"`
}

type FindDiskEncryptionConfigurationsByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindDiskEncryptionConfigurationsByID200ApplicationJSONObject *FindDiskEncryptionConfigurationsByID200ApplicationJSON
}
