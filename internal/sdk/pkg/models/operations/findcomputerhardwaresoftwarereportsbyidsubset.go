// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"jamf/internal/sdk/pkg/types"
	"net/http"
)

// FindComputerHardwareSoftwareReportsByIDSubsetSubset - Subset to filter by
type FindComputerHardwareSoftwareReportsByIDSubsetSubset string

const (
	FindComputerHardwareSoftwareReportsByIDSubsetSubsetSoftware FindComputerHardwareSoftwareReportsByIDSubsetSubset = "Software"
	FindComputerHardwareSoftwareReportsByIDSubsetSubsetHardwre  FindComputerHardwareSoftwareReportsByIDSubsetSubset = "Hardwre"
	FindComputerHardwareSoftwareReportsByIDSubsetSubsetFonts    FindComputerHardwareSoftwareReportsByIDSubsetSubset = "Fonts"
	FindComputerHardwareSoftwareReportsByIDSubsetSubsetPlugins  FindComputerHardwareSoftwareReportsByIDSubsetSubset = "Plugins"
)

func (e FindComputerHardwareSoftwareReportsByIDSubsetSubset) ToPointer() *FindComputerHardwareSoftwareReportsByIDSubsetSubset {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsByIDSubsetSubset) UnmarshalJSON(data []byte) error {
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
		*e = FindComputerHardwareSoftwareReportsByIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsByIDSubsetSubset: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsByIDSubsetRequest struct {
	// End date (e.g. yyyy-mm-dd)
	EndDate types.Date `pathParam:"style=simple,explode=false,name=end_date"`
	// Computer ID to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Start date (e.g. yyyy-mm-dd)
	StartDate types.Date `pathParam:"style=simple,explode=false,name=start_date"`
	// Subset to filter by
	Subset FindComputerHardwareSoftwareReportsByIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReportType string

const (
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReportTypeAdded   FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReportType = "Added"
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReportTypeDeleted FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReportType = "Deleted"
)

func (e FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReportType) ToPointer() *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Deleted":
		*e = FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReport struct {
	DateTime      *string
	DateTimeEpoch *string
	DateTimeUtc   *string
	Name          *string
	Path          *string
	Type          *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReportType
	Version       *string
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLHardwareReport struct {
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

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReportType string

const (
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReportTypeAdded   FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReportType = "Added"
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReportTypeRemoved FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReportType = "Removed"
)

func (e FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReportType) ToPointer() *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Removed":
		*e = FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReport struct {
	DateTime      *string
	DateTimeEpoch *string
	DateTimeUtc   *string
	Name          *string
	Path          *string
	Type          *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReportType
	Version       *string
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReportType string

const (
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReportTypeAdded   FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReportType = "Added"
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReportTypeDeleted FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReportType = "Deleted"
)

func (e FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReportType) ToPointer() *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Deleted":
		*e = FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReport struct {
	DateTime      *string
	DateTimeEpoch *string
	DateTimeUtc   *string
	Name          *string
	Path          *string
	Type          *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReportType
	Version       *string
}

// FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXML - OK
type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXML struct {
	FontReport     *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLFontReport
	HardwareReport *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLHardwareReport
	PluginReport   *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLPluginReport
	SoftwareReport *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationXMLSoftwareReport
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReportType string

const (
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReportTypeAdded   FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReportType = "Added"
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReportTypeDeleted FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReportType = "Deleted"
)

func (e FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReportType) ToPointer() *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Deleted":
		*e = FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReport struct {
	DateTime      *string                                                                        `json:"date_time,omitempty"`
	DateTimeEpoch *string                                                                        `json:"date_time_epoch,omitempty"`
	DateTimeUtc   *string                                                                        `json:"date_time_utc,omitempty"`
	Name          *string                                                                        `json:"name,omitempty"`
	Path          *string                                                                        `json:"path,omitempty"`
	Type          *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReportType `json:"type,omitempty"`
	Version       *string                                                                        `json:"version,omitempty"`
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONHardwareReport struct {
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

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReportType string

const (
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReportTypeAdded   FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReportType = "Added"
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReportTypeRemoved FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReportType = "Removed"
)

func (e FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReportType) ToPointer() *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Removed":
		*e = FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReport struct {
	DateTime      *string                                                                          `json:"date_time,omitempty"`
	DateTimeEpoch *string                                                                          `json:"date_time_epoch,omitempty"`
	DateTimeUtc   *string                                                                          `json:"date_time_utc,omitempty"`
	Name          *string                                                                          `json:"name,omitempty"`
	Path          *string                                                                          `json:"path,omitempty"`
	Type          *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReportType `json:"type,omitempty"`
	Version       *string                                                                          `json:"version,omitempty"`
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReportType string

const (
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReportTypeAdded   FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReportType = "Added"
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReportTypeDeleted FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReportType = "Deleted"
)

func (e FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReportType) ToPointer() *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReportType {
	return &e
}

func (e *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Added":
		fallthrough
	case "Deleted":
		*e = FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReportType: %v", v)
	}
}

type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReport struct {
	DateTime      *string                                                                            `json:"date_time,omitempty"`
	DateTimeEpoch *string                                                                            `json:"date_time_epoch,omitempty"`
	DateTimeUtc   *string                                                                            `json:"date_time_utc,omitempty"`
	Name          *string                                                                            `json:"name,omitempty"`
	Path          *string                                                                            `json:"path,omitempty"`
	Type          *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReportType `json:"type,omitempty"`
	Version       *string                                                                            `json:"version,omitempty"`
}

// FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSON - OK
type FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSON struct {
	FontReport     *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONFontReport     `json:"font_report,omitempty"`
	HardwareReport *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONHardwareReport `json:"hardware_report,omitempty"`
	PluginReport   *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONPluginReport   `json:"plugin_report,omitempty"`
	SoftwareReport *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONSoftwareReport `json:"software_report,omitempty"`
}

type FindComputerHardwareSoftwareReportsByIDSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSONObject *FindComputerHardwareSoftwareReportsByIDSubset200ApplicationJSON
}
