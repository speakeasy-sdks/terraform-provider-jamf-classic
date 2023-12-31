// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindMobileDeviceHistoryByNameSubsetSubset - Subset to filter by
type FindMobileDeviceHistoryByNameSubsetSubset string

const (
	FindMobileDeviceHistoryByNameSubsetSubsetGeneral            FindMobileDeviceHistoryByNameSubsetSubset = "General"
	FindMobileDeviceHistoryByNameSubsetSubsetManagementCommands FindMobileDeviceHistoryByNameSubsetSubset = "ManagementCommands"
	FindMobileDeviceHistoryByNameSubsetSubsetUserLocation       FindMobileDeviceHistoryByNameSubsetSubset = "UserLocation"
	FindMobileDeviceHistoryByNameSubsetSubsetAudits             FindMobileDeviceHistoryByNameSubsetSubset = "Audits"
	FindMobileDeviceHistoryByNameSubsetSubsetApplications       FindMobileDeviceHistoryByNameSubsetSubset = "Applications"
	FindMobileDeviceHistoryByNameSubsetSubsetEbooks             FindMobileDeviceHistoryByNameSubsetSubset = "Ebooks"
)

func (e FindMobileDeviceHistoryByNameSubsetSubset) ToPointer() *FindMobileDeviceHistoryByNameSubsetSubset {
	return &e
}

func (e *FindMobileDeviceHistoryByNameSubsetSubset) UnmarshalJSON(data []byte) error {
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
		*e = FindMobileDeviceHistoryByNameSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByNameSubsetSubset: %v", v)
	}
}

type FindMobileDeviceHistoryByNameSubsetRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Subset to filter by
	Subset FindMobileDeviceHistoryByNameSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsFailedApp struct {
	ManagementStatus *string
	Name             *string
	Version          *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsFailed struct {
	App *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsFailedApp
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus string

const (
	FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatusUnmanaged FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatusManaged   FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus) ToPointer() *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalog struct {
	BundleSize       *string
	DynamicSize      *string
	ManagementStatus *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus
	Name             *string
	Version          *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus string

const (
	FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatusUnmanaged FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatusManaged   FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus) ToPointer() *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalog struct {
	BundleSize       *string
	DynamicSize      *string
	ManagementStatus *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus
	Name             *string
	Version          *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus string

const (
	FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOtherManagementStatusUnmanaged FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOtherManagementStatusManaged   FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus) ToPointer() *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOther struct {
	BundleSize       *string
	DynamicSize      *string
	ManagementStatus *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOtherManagementStatus
	Name             *string
	Version          *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalled struct {
	AppStoreFromMobileDeviceAppCatalog *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledAppStoreFromMobileDeviceAppCatalog
	InHouseFromMobileDeviceAppCatalog  *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledInHouseFromMobileDeviceAppCatalog
	Other                              *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalledOther
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsPendingApp struct {
	ManagementStatus *string
	Name             *string
	Version          *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsPending struct {
	App *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsPendingApp
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplications struct {
	Failed    []FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsFailed
	Installed []FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsInstalled
	Pending   []FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplicationsPending
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLAuditsAudit struct {
	DateTime      *string
	DateTimeEpoch *int64
	Event         *string
	Username      *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLAudits struct {
	Audit *FindMobileDeviceHistoryByNameSubset200ApplicationXMLAuditsAudit
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooksFailed struct {
	Author           *string
	Kind             *string
	ManagementStatus *string
	Title            *string
	Version          *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooksInstalledIbookstore struct {
	Author           *string
	Kind             *string
	ManagementStatus *string
	Title            *string
	Version          *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooksInstalledInhouse struct {
	Author           *string
	Kind             *string
	ManagementStatus *string
	Title            *string
	Version          *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooksInstalled struct {
	Ibookstore []FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooksInstalledIbookstore
	Inhouse    []FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooksInstalledInhouse
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooksPending struct {
	Author           *string
	Kind             *string
	ManagementStatus *string
	Title            *string
	Version          *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooks struct {
	Failed    []FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooksFailed
	Installed *FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooksInstalled
	Pending   []FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooksPending
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLGeneral struct {
	ID         *int64
	MacAddress *string
	// Name of the device
	Name         *string
	SerialNumber *string
	Udid         *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsCompletedCommand struct {
	DateTimeCompleted      *string
	DateTimeCompletedEpoch *int64
	Name                   *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsCompleted struct {
	Command *FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsCompletedCommand
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsFailedCommand struct {
	DateTimeFailed      *string
	DateTimeFailedEpoch *int64
	DateTimeIssued      *string
	DateTimeIssuedEpoch *string
	Error               *string
	Name                *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsFailed struct {
	Command *FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsFailedCommand
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsPendingCommand struct {
	DateTimeFailed      *string
	DateTimeFailedEpoch *int64
	DateTimeIssued      *string
	DateTimeIssuedEpoch *string
	Name                *string
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsPending struct {
	Command *FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsPendingCommand
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommands struct {
	Completed []FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsCompleted
	Failed    []FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsFailed
	Pending   []FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommandsPending
}

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLUserLocationLocation struct {
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

type FindMobileDeviceHistoryByNameSubset200ApplicationXMLUserLocation struct {
	Location *FindMobileDeviceHistoryByNameSubset200ApplicationXMLUserLocationLocation
}

// FindMobileDeviceHistoryByNameSubset200ApplicationXML - OK
type FindMobileDeviceHistoryByNameSubset200ApplicationXML struct {
	Applications       *FindMobileDeviceHistoryByNameSubset200ApplicationXMLApplications
	Audits             []FindMobileDeviceHistoryByNameSubset200ApplicationXMLAudits
	Ebooks             *FindMobileDeviceHistoryByNameSubset200ApplicationXMLEbooks
	General            *FindMobileDeviceHistoryByNameSubset200ApplicationXMLGeneral
	ManagementCommands *FindMobileDeviceHistoryByNameSubset200ApplicationXMLManagementCommands
	UserLocation       []FindMobileDeviceHistoryByNameSubset200ApplicationXMLUserLocation
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsFailedApp struct {
	ManagementStatus *string `json:"management_status,omitempty"`
	Name             *string `json:"name,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsFailed struct {
	App *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsFailedApp `json:"app,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus string

const (
	FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatusUnmanaged FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatusManaged   FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus) ToPointer() *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalog struct {
	BundleSize       *string                                                                                                                       `json:"bundle_size,omitempty"`
	DynamicSize      *string                                                                                                                       `json:"dynamic_size,omitempty"`
	ManagementStatus *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalogManagementStatus `json:"management_status,omitempty"`
	Name             *string                                                                                                                       `json:"name,omitempty"`
	Version          *string                                                                                                                       `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus string

const (
	FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatusUnmanaged FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatusManaged   FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus) ToPointer() *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalog struct {
	BundleSize       *string                                                                                                                      `json:"bundle_size,omitempty"`
	DynamicSize      *string                                                                                                                      `json:"dynamic_size,omitempty"`
	ManagementStatus *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalogManagementStatus `json:"management_status,omitempty"`
	Name             *string                                                                                                                      `json:"name,omitempty"`
	Version          *string                                                                                                                      `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus string

const (
	FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOtherManagementStatusUnmanaged FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus = "Unmanaged"
	FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOtherManagementStatusManaged   FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus = "Managed"
)

func (e FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus) ToPointer() *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus {
	return &e
}

func (e *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Unmanaged":
		fallthrough
	case "Managed":
		*e = FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus: %v", v)
	}
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOther struct {
	BundleSize       *string                                                                                          `json:"bundle_size,omitempty"`
	DynamicSize      *string                                                                                          `json:"dynamic_size,omitempty"`
	ManagementStatus *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOtherManagementStatus `json:"management_status,omitempty"`
	Name             *string                                                                                          `json:"name,omitempty"`
	Version          *string                                                                                          `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalled struct {
	AppStoreFromMobileDeviceAppCatalog *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledAppStoreFromMobileDeviceAppCatalog `json:"app_store_from_mobile_device_app_catalog,omitempty"`
	InHouseFromMobileDeviceAppCatalog  *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledInHouseFromMobileDeviceAppCatalog  `json:"in_house_from_mobile_device_app_catalog,omitempty"`
	Other                              *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalledOther                              `json:"other,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsPendingApp struct {
	ManagementStatus *string `json:"management_status,omitempty"`
	Name             *string `json:"name,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsPending struct {
	App *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsPendingApp `json:"app,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplications struct {
	Failed    []FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsFailed    `json:"failed,omitempty"`
	Installed []FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsInstalled `json:"installed,omitempty"`
	Pending   []FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplicationsPending   `json:"pending,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONAuditsAudit struct {
	DateTime      *string `json:"date_time,omitempty"`
	DateTimeEpoch *int64  `json:"date_time_epoch,omitempty"`
	Event         *string `json:"event,omitempty"`
	Username      *string `json:"username,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONAudits struct {
	Audit *FindMobileDeviceHistoryByNameSubset200ApplicationJSONAuditsAudit `json:"audit,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooksFailed struct {
	Author           *string `json:"author,omitempty"`
	Kind             *string `json:"kind,omitempty"`
	ManagementStatus *string `json:"management_status,omitempty"`
	Title            *string `json:"title,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooksInstalledIbookstore struct {
	Author           *string `json:"author,omitempty"`
	Kind             *string `json:"kind,omitempty"`
	ManagementStatus *string `json:"management_status,omitempty"`
	Title            *string `json:"title,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooksInstalledInhouse struct {
	Author           *string `json:"author,omitempty"`
	Kind             *string `json:"kind,omitempty"`
	ManagementStatus *string `json:"management_status,omitempty"`
	Title            *string `json:"title,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooksInstalled struct {
	Ibookstore []FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooksInstalledIbookstore `json:"ibookstore,omitempty"`
	Inhouse    []FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooksInstalledInhouse    `json:"inhouse,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooksPending struct {
	Author           *string `json:"author,omitempty"`
	Kind             *string `json:"kind,omitempty"`
	ManagementStatus *string `json:"management_status,omitempty"`
	Title            *string `json:"title,omitempty"`
	Version          *string `json:"version,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooks struct {
	Failed    []FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooksFailed   `json:"failed,omitempty"`
	Installed *FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooksInstalled `json:"installed,omitempty"`
	Pending   []FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooksPending  `json:"pending,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONGeneral struct {
	ID         *int64  `json:"id,omitempty"`
	MacAddress *string `json:"mac_address,omitempty"`
	// Name of the device
	Name         *string `json:"name,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Udid         *string `json:"udid,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsCompletedCommand struct {
	DateTimeCompleted      *string `json:"date_time_completed,omitempty"`
	DateTimeCompletedEpoch *int64  `json:"date_time_completed_epoch,omitempty"`
	Name                   *string `json:"name,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsCompleted struct {
	Command *FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsCompletedCommand `json:"command,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsFailedCommand struct {
	DateTimeFailed      *string `json:"date_time_failed,omitempty"`
	DateTimeFailedEpoch *int64  `json:"date_time_failed_epoch,omitempty"`
	DateTimeIssued      *string `json:"date_time_issued,omitempty"`
	DateTimeIssuedEpoch *string `json:"date_time_issued_epoch,omitempty"`
	Error               *string `json:"error,omitempty"`
	Name                *string `json:"name,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsFailed struct {
	Command *FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsFailedCommand `json:"command,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsPendingCommand struct {
	DateTimeFailed      *string `json:"date_time_failed,omitempty"`
	DateTimeFailedEpoch *int64  `json:"date_time_failed_epoch,omitempty"`
	DateTimeIssued      *string `json:"date_time_issued,omitempty"`
	DateTimeIssuedEpoch *string `json:"date_time_issued_epoch,omitempty"`
	Name                *string `json:"name,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsPending struct {
	Command *FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsPendingCommand `json:"command,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommands struct {
	Completed []FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsCompleted `json:"completed,omitempty"`
	Failed    []FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsFailed    `json:"failed,omitempty"`
	Pending   []FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommandsPending   `json:"pending,omitempty"`
}

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONUserLocationLocation struct {
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

type FindMobileDeviceHistoryByNameSubset200ApplicationJSONUserLocation struct {
	Location *FindMobileDeviceHistoryByNameSubset200ApplicationJSONUserLocationLocation `json:"location,omitempty"`
}

// FindMobileDeviceHistoryByNameSubset200ApplicationJSON - OK
type FindMobileDeviceHistoryByNameSubset200ApplicationJSON struct {
	Applications       *FindMobileDeviceHistoryByNameSubset200ApplicationJSONApplications       `json:"applications,omitempty"`
	Audits             []FindMobileDeviceHistoryByNameSubset200ApplicationJSONAudits            `json:"audits,omitempty"`
	Ebooks             *FindMobileDeviceHistoryByNameSubset200ApplicationJSONEbooks             `json:"ebooks,omitempty"`
	General            *FindMobileDeviceHistoryByNameSubset200ApplicationJSONGeneral            `json:"general,omitempty"`
	ManagementCommands *FindMobileDeviceHistoryByNameSubset200ApplicationJSONManagementCommands `json:"management_commands,omitempty"`
	UserLocation       []FindMobileDeviceHistoryByNameSubset200ApplicationJSONUserLocation      `json:"user_location,omitempty"`
}

type FindMobileDeviceHistoryByNameSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDeviceHistoryByNameSubset200ApplicationJSONObject *FindMobileDeviceHistoryByNameSubset200ApplicationJSON
}
