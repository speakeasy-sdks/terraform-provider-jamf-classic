// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"jamf/internal/sdk/pkg/types"
	"net/http"
)

// FindComputerHardwareSoftwareReportsBySerialSubsetSubset - Subset to filter by
type FindComputerHardwareSoftwareReportsBySerialSubsetSubset string

const (
	FindComputerHardwareSoftwareReportsBySerialSubsetSubsetSoftware FindComputerHardwareSoftwareReportsBySerialSubsetSubset = "Software"
	FindComputerHardwareSoftwareReportsBySerialSubsetSubsetHardwre  FindComputerHardwareSoftwareReportsBySerialSubsetSubset = "Hardwre"
	FindComputerHardwareSoftwareReportsBySerialSubsetSubsetFonts    FindComputerHardwareSoftwareReportsBySerialSubsetSubset = "Fonts"
	FindComputerHardwareSoftwareReportsBySerialSubsetSubsetPlugins  FindComputerHardwareSoftwareReportsBySerialSubsetSubset = "Plugins"
)

func (e FindComputerHardwareSoftwareReportsBySerialSubsetSubset) ToPointer() *FindComputerHardwareSoftwareReportsBySerialSubsetSubset {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsBySerialSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Software":
		fallthrough
	case "Hardwre":
		fallthrough
	case "Fonts":
		fallthrough
	case "Plugins":
		*e = FindComputerHardwareSoftwareReportsBySerialSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsBySerialSubsetSubset: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsBySerialSubsetRequest struct {
	// End date (e.g. yyyy-mm-dd)
	EndDate types.Date `pathParam:"style=simple,explode=false,name=end_date"`
	// Serial number to filter by
	Serialnumber string `pathParam:"style=simple,explode=false,name=serialnumber"`
	// Start date (e.g. yyyy-mm-dd)
	StartDate types.Date `pathParam:"style=simple,explode=false,name=start_date"`
	// Subset to filter by
	Subset FindComputerHardwareSoftwareReportsBySerialSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReportType string

const (
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReportTypeAdded   FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReportType = "Added"
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReportTypeDeleted FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReportType = "Deleted"
)

func (e FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReportType) ToPointer() *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Deleted":
		*e = FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReport struct {
	DateTime      *string
	DateTimeEpoch *string
	DateTimeUtc   *string
	Name          *string
	Path          *string
	Type          *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReportType
	Version       *string
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLHardwareReport struct {
	NICSpeed                 *string
	BootPartitionUsedPercent *int64
	CoreCount                *int64
	DateTime                 *string
	DateTimeEpoch            *string
	DateTimeUtc              *string
	Make                     *string
	ModelIdentifier          *string
	OpenRAMSlots             *int64
	OperatingSystem          *string
	OpticalDrive             *string
	ProcessorCount           *int64
	ProcessorSpeedMhz        *int64
	SerialNumber             *string
	ServicePack              *string
	TotalHarddriveSize       *string
	TotalRAMMb               *int64
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReportType string

const (
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReportTypeAdded   FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReportType = "Added"
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReportTypeRemoved FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReportType = "Removed"
)

func (e FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReportType) ToPointer() *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Removed":
		*e = FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReport struct {
	DateTime      *string
	DateTimeEpoch *string
	DateTimeUtc   *string
	Name          *string
	Path          *string
	Type          *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReportType
	Version       *string
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReportType string

const (
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReportTypeAdded   FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReportType = "Added"
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReportTypeDeleted FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReportType = "Deleted"
)

func (e FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReportType) ToPointer() *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Deleted":
		*e = FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReport struct {
	DateTime      *string
	DateTimeEpoch *string
	DateTimeUtc   *string
	Name          *string
	Path          *string
	Type          *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReportType
	Version       *string
}

// FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXML - OK
type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXML struct {
	FontReport     *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLFontReport
	HardwareReport *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLHardwareReport
	PluginReport   *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLPluginReport
	SoftwareReport *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationXMLSoftwareReport
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReportType string

const (
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReportTypeAdded   FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReportType = "Added"
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReportTypeDeleted FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReportType = "Deleted"
)

func (e FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReportType) ToPointer() *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Deleted":
		*e = FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReport struct {
	DateTime      *string                                                                            `json:"date_time,omitempty"`
	DateTimeEpoch *string                                                                            `json:"date_time_epoch,omitempty"`
	DateTimeUtc   *string                                                                            `json:"date_time_utc,omitempty"`
	Name          *string                                                                            `json:"name,omitempty"`
	Path          *string                                                                            `json:"path,omitempty"`
	Type          *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReportType `json:"type,omitempty"`
	Version       *string                                                                            `json:"version,omitempty"`
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONHardwareReport struct {
	NICSpeed                 *string `json:"NIC_speed,omitempty"`
	BootPartitionUsedPercent *int64  `json:"boot_partition_used_percent,omitempty"`
	CoreCount                *int64  `json:"core_count,omitempty"`
	DateTime                 *string `json:"date_time,omitempty"`
	DateTimeEpoch            *string `json:"date_time_epoch,omitempty"`
	DateTimeUtc              *string `json:"date_time_utc,omitempty"`
	Make                     *string `json:"make,omitempty"`
	ModelIdentifier          *string `json:"model_identifier,omitempty"`
	OpenRAMSlots             *int64  `json:"open_ram_slots,omitempty"`
	OperatingSystem          *string `json:"operating_system,omitempty"`
	OpticalDrive             *string `json:"optical_drive,omitempty"`
	ProcessorCount           *int64  `json:"processor_count,omitempty"`
	ProcessorSpeedMhz        *int64  `json:"processor_speed_mhz,omitempty"`
	SerialNumber             *string `json:"serial_number,omitempty"`
	ServicePack              *string `json:"service_pack,omitempty"`
	TotalHarddriveSize       *string `json:"total_harddrive_size,omitempty"`
	TotalRAMMb               *int64  `json:"total_ram_mb,omitempty"`
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReportType string

const (
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReportTypeAdded   FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReportType = "Added"
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReportTypeRemoved FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReportType = "Removed"
)

func (e FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReportType) ToPointer() *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Removed":
		*e = FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReport struct {
	DateTime      *string                                                                              `json:"date_time,omitempty"`
	DateTimeEpoch *string                                                                              `json:"date_time_epoch,omitempty"`
	DateTimeUtc   *string                                                                              `json:"date_time_utc,omitempty"`
	Name          *string                                                                              `json:"name,omitempty"`
	Path          *string                                                                              `json:"path,omitempty"`
	Type          *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReportType `json:"type,omitempty"`
	Version       *string                                                                              `json:"version,omitempty"`
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReportType string

const (
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReportTypeAdded   FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReportType = "Added"
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReportTypeDeleted FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReportType = "Deleted"
)

func (e FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReportType) ToPointer() *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Deleted":
		*e = FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReport struct {
	DateTime      *string                                                                                `json:"date_time,omitempty"`
	DateTimeEpoch *string                                                                                `json:"date_time_epoch,omitempty"`
	DateTimeUtc   *string                                                                                `json:"date_time_utc,omitempty"`
	Name          *string                                                                                `json:"name,omitempty"`
	Path          *string                                                                                `json:"path,omitempty"`
	Type          *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReportType `json:"type,omitempty"`
	Version       *string                                                                                `json:"version,omitempty"`
}

// FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSON - OK
type FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSON struct {
	FontReport     *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONFontReport     `json:"font_report,omitempty"`
	HardwareReport *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONHardwareReport `json:"hardware_report,omitempty"`
	PluginReport   *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONPluginReport   `json:"plugin_report,omitempty"`
	SoftwareReport *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONSoftwareReport `json:"software_report,omitempty"`
}

type FindComputerHardwareSoftwareReportsBySerialSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSONObject *FindComputerHardwareSoftwareReportsBySerialSubset200ApplicationJSON
}
