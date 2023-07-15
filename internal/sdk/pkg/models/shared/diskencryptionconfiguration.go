// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DiskEncryptionConfigurationFileVaultEnabledUsers string

const (
	DiskEncryptionConfigurationFileVaultEnabledUsersCurrentOrNextUser DiskEncryptionConfigurationFileVaultEnabledUsers = "Current or Next User"
	DiskEncryptionConfigurationFileVaultEnabledUsersManagementAccount DiskEncryptionConfigurationFileVaultEnabledUsers = "Management Account"
)

func (e DiskEncryptionConfigurationFileVaultEnabledUsers) ToPointer() *DiskEncryptionConfigurationFileVaultEnabledUsers {
	return &e
}

func (e *DiskEncryptionConfigurationFileVaultEnabledUsers) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Current or Next User":
		fallthrough
	case "Management Account":
		*e = DiskEncryptionConfigurationFileVaultEnabledUsers(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DiskEncryptionConfigurationFileVaultEnabledUsers: %v", v)
	}
}

type DiskEncryptionConfigurationKeyType string

const (
	DiskEncryptionConfigurationKeyTypeIndividual                 DiskEncryptionConfigurationKeyType = "Individual"
	DiskEncryptionConfigurationKeyTypeInstitutional              DiskEncryptionConfigurationKeyType = "Institutional"
	DiskEncryptionConfigurationKeyTypeIndividualAndInstitutional DiskEncryptionConfigurationKeyType = "Individual And Institutional"
)

func (e DiskEncryptionConfigurationKeyType) ToPointer() *DiskEncryptionConfigurationKeyType {
	return &e
}

func (e *DiskEncryptionConfigurationKeyType) UnmarshalJSON(data []byte) error {
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
		*e = DiskEncryptionConfigurationKeyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DiskEncryptionConfigurationKeyType: %v", v)
	}
}

// DiskEncryptionConfiguration - OK
type DiskEncryptionConfiguration struct {
	FileVaultEnabledUsers *DiskEncryptionConfigurationFileVaultEnabledUsers `json:"file_vault_enabled_users,omitempty"`
	ID                    *int64                                            `json:"id,omitempty"`
	KeyType               *DiskEncryptionConfigurationKeyType               `json:"key_type,omitempty"`
	// Name of the disk encryption configuration
	Name string `json:"name"`
}
