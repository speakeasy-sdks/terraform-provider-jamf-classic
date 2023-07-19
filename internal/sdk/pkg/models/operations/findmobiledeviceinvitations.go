// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindMobileDeviceInvitations200ApplicationXMLMobileDeviceInvitation struct {
	ExpirationDate      *string
	ExpirationDateEpoch *int64
	ExpirationDateUtc   *string
	ID                  *int64
	Invitation          *int64
	InvitationType      *string
	LastAction          *string
}

type FindMobileDeviceInvitations200ApplicationXML struct {
	MobileDeviceInvitation *FindMobileDeviceInvitations200ApplicationXMLMobileDeviceInvitation
	Size                   *int64
}

type FindMobileDeviceInvitations200ApplicationJSONMobileDeviceInvitation struct {
	ExpirationDate      *string `json:"expiration_date,omitempty"`
	ExpirationDateEpoch *int64  `json:"expiration_date_epoch,omitempty"`
	ExpirationDateUtc   *string `json:"expiration_date_utc,omitempty"`
	ID                  *int64  `json:"id,omitempty"`
	Invitation          *int64  `json:"invitation,omitempty"`
	InvitationType      *string `json:"invitation_type,omitempty"`
	LastAction          *string `json:"last_action,omitempty"`
}

type FindMobileDeviceInvitations200ApplicationJSON struct {
	MobileDeviceInvitation *FindMobileDeviceInvitations200ApplicationJSONMobileDeviceInvitation `json:"mobile_device_invitation,omitempty"`
	Size                   *int64                                                               `json:"size,omitempty"`
}

type FindMobileDeviceInvitationsResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDeviceInvitations200ApplicationJSONObjects []FindMobileDeviceInvitations200ApplicationJSON
}
