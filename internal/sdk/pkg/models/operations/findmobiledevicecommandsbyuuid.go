// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindMobileDeviceCommandsByUUIDRequest struct {
	// UUID value to filter by
	UUID string `pathParam:"style=simple,explode=false,name=uuid"`
}

type FindMobileDeviceCommandsByUUID200ApplicationXMLGeneral struct {
	ApnsResultStatus *string
	Command          *string
	DateSent         *string
	DateSentEpoch    *string
	DateSentUtc      *string
	ID               *int64
	ProfileID        *int64
	ProfileUdid      *string
	Udid             *string
	UUID             *string
}

type FindMobileDeviceCommandsByUUID200ApplicationXMLMobileDevicesMobileDevice struct {
	ID             *int64
	PhoneNumber    *string
	SerialNumber   *string
	Udid           *string
	WifiMacAddress *string
}

type FindMobileDeviceCommandsByUUID200ApplicationXMLMobileDevices struct {
	MobileDevice *FindMobileDeviceCommandsByUUID200ApplicationXMLMobileDevicesMobileDevice
	Size         *int64
}

// FindMobileDeviceCommandsByUUID200ApplicationXML - OK
type FindMobileDeviceCommandsByUUID200ApplicationXML struct {
	General       *FindMobileDeviceCommandsByUUID200ApplicationXMLGeneral
	MobileDevices *FindMobileDeviceCommandsByUUID200ApplicationXMLMobileDevices
}

type FindMobileDeviceCommandsByUUID200ApplicationJSONGeneral struct {
	ApnsResultStatus *string `json:"apns_result_status,omitempty"`
	Command          *string `json:"command,omitempty"`
	DateSent         *string `json:"date_sent,omitempty"`
	DateSentEpoch    *string `json:"date_sent_epoch,omitempty"`
	DateSentUtc      *string `json:"date_sent_utc,omitempty"`
	ID               *int64  `json:"id,omitempty"`
	ProfileID        *int64  `json:"profile_id,omitempty"`
	ProfileUdid      *string `json:"profile_udid,omitempty"`
	Udid             *string `json:"udid,omitempty"`
	UUID             *string `json:"uuid,omitempty"`
}

type FindMobileDeviceCommandsByUUID200ApplicationJSONMobileDevicesMobileDevice struct {
	ID             *int64  `json:"id,omitempty"`
	PhoneNumber    *string `json:"phone_number,omitempty"`
	SerialNumber   *string `json:"serial_number,omitempty"`
	Udid           *string `json:"udid,omitempty"`
	WifiMacAddress *string `json:"wifi_mac_address,omitempty"`
}

type FindMobileDeviceCommandsByUUID200ApplicationJSONMobileDevices struct {
	MobileDevice *FindMobileDeviceCommandsByUUID200ApplicationJSONMobileDevicesMobileDevice `json:"mobile_device,omitempty"`
	Size         *int64                                                                     `json:"size,omitempty"`
}

// FindMobileDeviceCommandsByUUID200ApplicationJSON - OK
type FindMobileDeviceCommandsByUUID200ApplicationJSON struct {
	General       *FindMobileDeviceCommandsByUUID200ApplicationJSONGeneral       `json:"general,omitempty"`
	MobileDevices *FindMobileDeviceCommandsByUUID200ApplicationJSONMobileDevices `json:"mobile_devices,omitempty"`
}

type FindMobileDeviceCommandsByUUIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDeviceCommandsByUUID200ApplicationJSONObject *FindMobileDeviceCommandsByUUID200ApplicationJSON
}
