// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindMobileDeviceHistoryByUDIDSubsetSubset - Subset to filter by
type FindMobileDeviceHistoryByUDIDSubsetSubset string

const (
	FindMobileDeviceHistoryByUDIDSubsetSubsetGeneral            FindMobileDeviceHistoryByUDIDSubsetSubset = "General"
	FindMobileDeviceHistoryByUDIDSubsetSubsetManagementCommands FindMobileDeviceHistoryByUDIDSubsetSubset = "ManagementCommands"
	FindMobileDeviceHistoryByUDIDSubsetSubsetUserLocation       FindMobileDeviceHistoryByUDIDSubsetSubset = "UserLocation"
	FindMobileDeviceHistoryByUDIDSubsetSubsetAudits             FindMobileDeviceHistoryByUDIDSubsetSubset = "Audits"
	FindMobileDeviceHistoryByUDIDSubsetSubsetApplications       FindMobileDeviceHistoryByUDIDSubsetSubset = "Applications"
	FindMobileDeviceHistoryByUDIDSubsetSubsetEbooks             FindMobileDeviceHistoryByUDIDSubsetSubset = "Ebooks"
)

func (e FindMobileDeviceHistoryByUDIDSubsetSubset) ToPointer() *FindMobileDeviceHistoryByUDIDSubsetSubset {
	return &e
}

func (e *FindMobileDeviceHistoryByUDIDSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "ManagementCommands":
		fallthrough
	case "UserLocation":
		fallthrough
	case "Audits":
		fallthrough
	case "Applications":
		fallthrough
	case "Ebooks":
		*e = FindMobileDeviceHistoryByUDIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByUDIDSubsetSubset: %v", v)
	}
}

type FindMobileDeviceHistoryByUDIDSubsetRequest struct {
	// Subset to filter by
	Subset FindMobileDeviceHistoryByUDIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
	// UDID to filter by
	Udid string `pathParam:"style=simple,explode=false,name=udid"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsFailedApp struct {
	ManagementStatus *string
	Name             *string
	Version          *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsFailed struct {
	App *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsFailedApp
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus string

const (
	FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatusUnmanaged FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatusManaged   FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus) ToPointer() *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalog struct {
	BundleSize       *string
	DynamicSize      *string
	ManagementStatus *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus
	Name             *string
	Version          *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus string

const (
	FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatusUnmanaged FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatusManaged   FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus) ToPointer() *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalog struct {
	BundleSize       *string
	DynamicSize      *string
	ManagementStatus *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus
	Name             *string
	Version          *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus string

const (
	FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOtherManagementStatusUnmanaged FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOtherManagementStatusManaged   FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus) ToPointer() *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOther struct {
	BundleSize       *string
	DynamicSize      *string
	ManagementStatus *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus
	Name             *string
	Version          *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalled struct {
	AppStoreFromMobileDeviceAppCatalog *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalog
	InHouseFromMobileDeviceAppCatalog  *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalog
	Other                              *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalledOther
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsPendingApp struct {
	ManagementStatus *string
	Name             *string
	Version          *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsPending struct {
	App *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsPendingApp
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplications struct {
	Failed    []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsFailed
	Installed []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsInstalled
	Pending   []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplicationsPending
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLAuditsAudit struct {
	DateTime      *string
	DateTimeEpoch *int64
	Event         *string
	Username      *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLAudits struct {
	Audit *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLAuditsAudit
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooksFailed struct {
	Author           *string
	Kind             *string
	ManagementStatus *string
	Title            *string
	Version          *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooksInstalledIbookstore struct {
	Author           *string
	Kind             *string
	ManagementStatus *string
	Title            *string
	Version          *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooksInstalledInhouse struct {
	Author           *string
	Kind             *string
	ManagementStatus *string
	Title            *string
	Version          *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooksInstalled struct {
	Ibookstore []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooksInstalledIbookstore
	Inhouse    []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooksInstalledInhouse
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooksPending struct {
	Author           *string
	Kind             *string
	ManagementStatus *string
	Title            *string
	Version          *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooks struct {
	Failed    []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooksFailed
	Installed *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooksInstalled
	Pending   []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooksPending
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLGeneral struct {
	ID         *int64
	MacAddress *string
	// Name of the device
	Name         *string
	SerialNumber *string
	Udid         *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsCompletedCommand struct {
	DateTimeCompleted      *string
	DateTimeCompletedEpoch *int64
	Name                   *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsCompleted struct {
	Command *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsCompletedCommand
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsFailedCommand struct {
	DateTimeFailed      *string
	DateTimeFailedEpoch *int64
	DateTimeIssued      *string
	DateTimeIssuedEpoch *string
	Error               *string
	Name                *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsFailed struct {
	Command *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsFailedCommand
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsPendingCommand struct {
	DateTimeFailed      *string
	DateTimeFailedEpoch *int64
	DateTimeIssued      *string
	DateTimeIssuedEpoch *string
	Name                *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsPending struct {
	Command *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsPendingCommand
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommands struct {
	Completed []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsCompleted
	Failed    []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsFailed
	Pending   []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommandsPending
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLUserLocationLocation struct {
	Building      *string
	DateTime      *string
	DateTimeEpoch *int64
	Department    *string
	EmailAddress  *string
	FullName      *string
	PhoneNumber   *string
	Position      *string
	Room          *string
	Username      *string
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLUserLocation struct {
	Location *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLUserLocationLocation
}

// FindMobileDeviceHistoryByUDIDSubset200ApplicationXML - OK
type FindMobileDeviceHistoryByUDIDSubset200ApplicationXML struct {
	Applications       *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLApplications
	Audits             []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLAudits
	Ebooks             *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLEbooks
	General            *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLGeneral
	ManagementCommands *FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLManagementCommands
	UserLocation       []FindMobileDeviceHistoryByUDIDSubset200ApplicationXMLUserLocation
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsFailedApp struct {
	ManagementStatus *string `json:"management_status,omitempty"`
	Name             *string `json:"name,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsFailed struct {
	App *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsFailedApp `json:"app,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus string

const (
	FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatusUnmanaged FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatusManaged   FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus) ToPointer() *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalog struct {
	BundleSize       *string                                                                                                                       `json:"bundle_size,omitempty"`
	DynamicSize      *string                                                                                                                       `json:"dynamic_size,omitempty"`
	ManagementStatus *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus `json:"management_status,omitempty"`
	Name             *string                                                                                                                       `json:"name,omitempty"`
	Version          *string                                                                                                                       `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus string

const (
	FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatusUnmanaged FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatusManaged   FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus) ToPointer() *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalog struct {
	BundleSize       *string                                                                                                                      `json:"bundle_size,omitempty"`
	DynamicSize      *string                                                                                                                      `json:"dynamic_size,omitempty"`
	ManagementStatus *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus `json:"management_status,omitempty"`
	Name             *string                                                                                                                      `json:"name,omitempty"`
	Version          *string                                                                                                                      `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus string

const (
	FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOtherManagementStatusUnmanaged FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOtherManagementStatusManaged   FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus) ToPointer() *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOther struct {
	BundleSize       *string                                                                                          `json:"bundle_size,omitempty"`
	DynamicSize      *string                                                                                          `json:"dynamic_size,omitempty"`
	ManagementStatus *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus `json:"management_status,omitempty"`
	Name             *string                                                                                          `json:"name,omitempty"`
	Version          *string                                                                                          `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalled struct {
	AppStoreFromMobileDeviceAppCatalog *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalog `json:"app_store_from_mobile_device_app_catalog,omitempty"`
	InHouseFromMobileDeviceAppCatalog  *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalog  `json:"in_house_from_mobile_device_app_catalog,omitempty"`
	Other                              *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalledOther                              `json:"other,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsPendingApp struct {
	ManagementStatus *string `json:"management_status,omitempty"`
	Name             *string `json:"name,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsPending struct {
	App *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsPendingApp `json:"app,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplications struct {
	Failed    []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsFailed    `json:"failed,omitempty"`
	Installed []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsInstalled `json:"installed,omitempty"`
	Pending   []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplicationsPending   `json:"pending,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONAuditsAudit struct {
	DateTime      *string `json:"date_time,omitempty"`
	DateTimeEpoch *int64  `json:"date_time_epoch,omitempty"`
	Event         *string `json:"event,omitempty"`
	Username      *string `json:"username,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONAudits struct {
	Audit *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONAuditsAudit `json:"audit,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooksFailed struct {
	Author           *string `json:"author,omitempty"`
	Kind             *string `json:"kind,omitempty"`
	ManagementStatus *string `json:"management_status,omitempty"`
	Title            *string `json:"title,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooksInstalledIbookstore struct {
	Author           *string `json:"author,omitempty"`
	Kind             *string `json:"kind,omitempty"`
	ManagementStatus *string `json:"management_status,omitempty"`
	Title            *string `json:"title,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooksInstalledInhouse struct {
	Author           *string `json:"author,omitempty"`
	Kind             *string `json:"kind,omitempty"`
	ManagementStatus *string `json:"management_status,omitempty"`
	Title            *string `json:"title,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooksInstalled struct {
	Ibookstore []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooksInstalledIbookstore `json:"ibookstore,omitempty"`
	Inhouse    []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooksInstalledInhouse    `json:"inhouse,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooksPending struct {
	Author           *string `json:"author,omitempty"`
	Kind             *string `json:"kind,omitempty"`
	ManagementStatus *string `json:"management_status,omitempty"`
	Title            *string `json:"title,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooks struct {
	Failed    []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooksFailed   `json:"failed,omitempty"`
	Installed *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooksInstalled `json:"installed,omitempty"`
	Pending   []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooksPending  `json:"pending,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONGeneral struct {
	ID         *int64  `json:"id,omitempty"`
	MacAddress *string `json:"mac_address,omitempty"`
	// Name of the device
	Name         *string `json:"name,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Udid         *string `json:"udid,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsCompletedCommand struct {
	DateTimeCompleted      *string `json:"date_time_completed,omitempty"`
	DateTimeCompletedEpoch *int64  `json:"date_time_completed_epoch,omitempty"`
	Name                   *string `json:"name,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsCompleted struct {
	Command *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsCompletedCommand `json:"command,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsFailedCommand struct {
	DateTimeFailed      *string `json:"date_time_failed,omitempty"`
	DateTimeFailedEpoch *int64  `json:"date_time_failed_epoch,omitempty"`
	DateTimeIssued      *string `json:"date_time_issued,omitempty"`
	DateTimeIssuedEpoch *string `json:"date_time_issued_epoch,omitempty"`
	Error               *string `json:"error,omitempty"`
	Name                *string `json:"name,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsFailed struct {
	Command *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsFailedCommand `json:"command,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsPendingCommand struct {
	DateTimeFailed      *string `json:"date_time_failed,omitempty"`
	DateTimeFailedEpoch *int64  `json:"date_time_failed_epoch,omitempty"`
	DateTimeIssued      *string `json:"date_time_issued,omitempty"`
	DateTimeIssuedEpoch *string `json:"date_time_issued_epoch,omitempty"`
	Name                *string `json:"name,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsPending struct {
	Command *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsPendingCommand `json:"command,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommands struct {
	Completed []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsCompleted `json:"completed,omitempty"`
	Failed    []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsFailed    `json:"failed,omitempty"`
	Pending   []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommandsPending   `json:"pending,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONUserLocationLocation struct {
	Building      *string `json:"building,omitempty"`
	DateTime      *string `json:"date_time,omitempty"`
	DateTimeEpoch *int64  `json:"date_time_epoch,omitempty"`
	Department    *string `json:"department,omitempty"`
	EmailAddress  *string `json:"email_address,omitempty"`
	FullName      *string `json:"full_name,omitempty"`
	PhoneNumber   *string `json:"phone_number,omitempty"`
	Position      *string `json:"position,omitempty"`
	Room          *string `json:"room,omitempty"`
	Username      *string `json:"username,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONUserLocation struct {
	Location *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONUserLocationLocation `json:"location,omitempty"`
}

// FindMobileDeviceHistoryByUDIDSubset200ApplicationJSON - OK
type FindMobileDeviceHistoryByUDIDSubset200ApplicationJSON struct {
	Applications       *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONApplications       `json:"applications,omitempty"`
	Audits             []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONAudits            `json:"audits,omitempty"`
	Ebooks             *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONEbooks             `json:"ebooks,omitempty"`
	General            *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONGeneral            `json:"general,omitempty"`
	ManagementCommands *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONManagementCommands `json:"management_commands,omitempty"`
	UserLocation       []FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONUserLocation      `json:"user_location,omitempty"`
}

type FindMobileDeviceHistoryByUDIDSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDeviceHistoryByUDIDSubset200ApplicationJSONObject *FindMobileDeviceHistoryByUDIDSubset200ApplicationJSON
}
