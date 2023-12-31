// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindDiskEncryptionConfigurations200ApplicationXMLDiskEncryptionConfiguration struct {
	ID   *int64
	Name *string
}

type FindDiskEncryptionConfigurations200ApplicationXML struct {
	DiskEncryptionConfiguration *FindDiskEncryptionConfigurations200ApplicationXMLDiskEncryptionConfiguration
	Size                        *int64
}

type FindDiskEncryptionConfigurations200ApplicationJSONDiskEncryptionConfiguration struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindDiskEncryptionConfigurations200ApplicationJSON struct {
	DiskEncryptionConfiguration *FindDiskEncryptionConfigurations200ApplicationJSONDiskEncryptionConfiguration `json:"disk_encryption_configuration,omitempty"`
	Size                        *int64                                                                         `json:"size,omitempty"`
}

type FindDiskEncryptionConfigurationsResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindDiskEncryptionConfigurations200ApplicationJSONObjects []FindDiskEncryptionConfigurations200ApplicationJSON
}
