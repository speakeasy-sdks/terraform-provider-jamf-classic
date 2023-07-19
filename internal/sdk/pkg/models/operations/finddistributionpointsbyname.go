// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FindDistributionPointsByNameRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type FindDistributionPointsByName200ApplicationXMLConnectionType string

const (
	FindDistributionPointsByName200ApplicationXMLConnectionTypeSmb FindDistributionPointsByName200ApplicationXMLConnectionType = "SMB"
	FindDistributionPointsByName200ApplicationXMLConnectionTypeAfp FindDistributionPointsByName200ApplicationXMLConnectionType = "AFP"
)

func (e FindDistributionPointsByName200ApplicationXMLConnectionType) ToPointer() *FindDistributionPointsByName200ApplicationXMLConnectionType {
	return &e
}

func (e *FindDistributionPointsByName200ApplicationXMLConnectionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SMB":
		fallthrough
	case "AFP":
		*e = FindDistributionPointsByName200ApplicationXMLConnectionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindDistributionPointsByName200ApplicationXMLConnectionType: %v", v)
	}
}

type FindDistributionPointsByName200ApplicationXMLProtocol string

const (
	FindDistributionPointsByName200ApplicationXMLProtocolHTTP  FindDistributionPointsByName200ApplicationXMLProtocol = "http"
	FindDistributionPointsByName200ApplicationXMLProtocolHTTPS FindDistributionPointsByName200ApplicationXMLProtocol = "https"
)

func (e FindDistributionPointsByName200ApplicationXMLProtocol) ToPointer() *FindDistributionPointsByName200ApplicationXMLProtocol {
	return &e
}

func (e *FindDistributionPointsByName200ApplicationXMLProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http":
		fallthrough
	case "https":
		*e = FindDistributionPointsByName200ApplicationXMLProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindDistributionPointsByName200ApplicationXMLProtocol: %v", v)
	}
}

// FindDistributionPointsByName200ApplicationXML - OK
type FindDistributionPointsByName200ApplicationXML struct {
	ConnectionType *FindDistributionPointsByName200ApplicationXMLConnectionType
	// Path to the share
	Context              *string
	EnableLoadBalancing  *bool
	FailoverPoint        *string
	FailoverPointURL     *string
	HTTPDownloadsEnabled *bool
	// Password for basic authentication
	HTTPPassword *string
	// URL to download packages from
	HTTPURL *string
	// Username to authenticate with for basic authentication
	HTTPUsername *string
	ID           *int64
	// IP address or hostname of distribution point
	IPAddress *string
	// Only one share can be set as master
	IsMaster  *bool
	LocalPath *string
	// Name of the distribution point
	Name                     string
	NoAuthenticationRequired *bool
	Password                 *string
	Port                     *int64
	Protocol                 *FindDistributionPointsByName200ApplicationXMLProtocol
	// Password for the read only account
	ReadOnlyPassword *string
	// Account with read only privileges to the share
	ReadOnlyUsername string
	// Password for the read/write account
	ReadWritePassword *string
	// Account with read/write privileges to the share
	ReadWriteUsername string
	// Name of the share
	ShareName                string
	SharePort                *int64
	SSHUsername              *string
	UsernamePasswordRequired *bool
	// Workgroup or domain of the accounts that have access to the share (SMB only)
	WorkgroupOrDomain *string
}

type FindDistributionPointsByName200ApplicationJSONConnectionType string

const (
	FindDistributionPointsByName200ApplicationJSONConnectionTypeSmb FindDistributionPointsByName200ApplicationJSONConnectionType = "SMB"
	FindDistributionPointsByName200ApplicationJSONConnectionTypeAfp FindDistributionPointsByName200ApplicationJSONConnectionType = "AFP"
)

func (e FindDistributionPointsByName200ApplicationJSONConnectionType) ToPointer() *FindDistributionPointsByName200ApplicationJSONConnectionType {
	return &e
}

func (e *FindDistributionPointsByName200ApplicationJSONConnectionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SMB":
		fallthrough
	case "AFP":
		*e = FindDistributionPointsByName200ApplicationJSONConnectionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindDistributionPointsByName200ApplicationJSONConnectionType: %v", v)
	}
}

type FindDistributionPointsByName200ApplicationJSONProtocol string

const (
	FindDistributionPointsByName200ApplicationJSONProtocolHTTP  FindDistributionPointsByName200ApplicationJSONProtocol = "http"
	FindDistributionPointsByName200ApplicationJSONProtocolHTTPS FindDistributionPointsByName200ApplicationJSONProtocol = "https"
)

func (e FindDistributionPointsByName200ApplicationJSONProtocol) ToPointer() *FindDistributionPointsByName200ApplicationJSONProtocol {
	return &e
}

func (e *FindDistributionPointsByName200ApplicationJSONProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http":
		fallthrough
	case "https":
		*e = FindDistributionPointsByName200ApplicationJSONProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindDistributionPointsByName200ApplicationJSONProtocol: %v", v)
	}
}

// FindDistributionPointsByName200ApplicationJSON - OK
type FindDistributionPointsByName200ApplicationJSON struct {
	ConnectionType *FindDistributionPointsByName200ApplicationJSONConnectionType `json:"connection_type,omitempty"`
	// Path to the share
	Context              *string `json:"context,omitempty"`
	EnableLoadBalancing  *bool   `json:"enable_load_balancing,omitempty"`
	FailoverPoint        *string `json:"failover_point,omitempty"`
	FailoverPointURL     *string `json:"failover_point_url,omitempty"`
	HTTPDownloadsEnabled *bool   `json:"http_downloads_enabled,omitempty"`
	// Password for basic authentication
	HTTPPassword *string `json:"http_password,omitempty"`
	// URL to download packages from
	HTTPURL *string `json:"http_url,omitempty"`
	// Username to authenticate with for basic authentication
	HTTPUsername *string `json:"http_username,omitempty"`
	ID           *int64  `json:"id,omitempty"`
	// IP address or hostname of distribution point
	IPAddress *string `json:"ip_address,omitempty"`
	// Only one share can be set as master
	IsMaster  *bool   `json:"is_master,omitempty"`
	LocalPath *string `json:"local_path,omitempty"`
	// Name of the distribution point
	Name                     string                                                  `json:"name"`
	NoAuthenticationRequired *bool                                                   `json:"no_authentication_required,omitempty"`
	Password                 *string                                                 `json:"password,omitempty"`
	Port                     *int64                                                  `json:"port,omitempty"`
	Protocol                 *FindDistributionPointsByName200ApplicationJSONProtocol `json:"protocol,omitempty"`
	// Password for the read only account
	ReadOnlyPassword *string `json:"read_only_password,omitempty"`
	// Account with read only privileges to the share
	ReadOnlyUsername string `json:"read_only_username"`
	// Password for the read/write account
	ReadWritePassword *string `json:"read_write_password,omitempty"`
	// Account with read/write privileges to the share
	ReadWriteUsername string `json:"read_write_username"`
	// Name of the share
	ShareName                string  `json:"share_name"`
	SharePort                *int64  `json:"share_port,omitempty"`
	SSHUsername              *string `json:"ssh_username,omitempty"`
	UsernamePasswordRequired *bool   `json:"username_password_required,omitempty"`
	// Workgroup or domain of the accounts that have access to the share (SMB only)
	WorkgroupOrDomain *string `json:"workgroup_or_domain,omitempty"`
}

type FindDistributionPointsByNameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindDistributionPointsByName200ApplicationJSONObject *FindDistributionPointsByName200ApplicationJSON
}
