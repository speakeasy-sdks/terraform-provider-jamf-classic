// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindMobileDeviceEnrollmentProfilesByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type FindMobileDeviceEnrollmentProfilesByID200ApplicationXMLAttachmentsAttachment struct {
	Filename *string
	ID       *int64
	URI      *string
}

type FindMobileDeviceEnrollmentProfilesByID200ApplicationXMLAttachments struct {
	Attachment *FindMobileDeviceEnrollmentProfilesByID200ApplicationXMLAttachmentsAttachment
}

type FindMobileDeviceEnrollmentProfilesByID200ApplicationXMLGeneral struct {
	Description *string
	ID          *int64
	Invitation  *string
	// Name of the enrollment profile
	Name string
	Udid *string
}

type FindMobileDeviceEnrollmentProfilesByID200ApplicationXMLLocation struct {
	Building     *string
	Department   *string
	EmailAddress *string
	Phone        *string
	PhoneNumber  *string
	Position     *string
	RealName     *string
	Realname     *string
	Room         *string
	Username     *string
}

type FindMobileDeviceEnrollmentProfilesByID200ApplicationXMLPurchasing struct {
	ApplecareID          *string
	IsLeased             *bool
	IsPurchased          *bool
	LeaseExpires         *string
	LeaseExpiresEpoch    *int64
	LeaseExpiresUtc      *string
	LifeExpectancy       *int64
	PoDate               *string
	PoDateEpoch          *int64
	PoDateUtc            *string
	PoNumber             *string
	PurchasePrice        *string
	PurchasingAccount    *string
	PurchasingContact    *string
	Vendor               *string
	WarrantyExpires      *string
	WarrantyExpiresEpoch *int64
	WarrantyExpiresUtc   *string
}

// FindMobileDeviceEnrollmentProfilesByID200ApplicationXML - OK
type FindMobileDeviceEnrollmentProfilesByID200ApplicationXML struct {
	Attachments []FindMobileDeviceEnrollmentProfilesByID200ApplicationXMLAttachments
	General     *FindMobileDeviceEnrollmentProfilesByID200ApplicationXMLGeneral
	Location    *FindMobileDeviceEnrollmentProfilesByID200ApplicationXMLLocation
	Purchasing  *FindMobileDeviceEnrollmentProfilesByID200ApplicationXMLPurchasing
}

type FindMobileDeviceEnrollmentProfilesByID200ApplicationJSONAttachmentsAttachment struct {
	Filename *string `json:"filename,omitempty"`
	ID       *int64  `json:"id,omitempty"`
	URI      *string `json:"uri,omitempty"`
}

type FindMobileDeviceEnrollmentProfilesByID200ApplicationJSONAttachments struct {
	Attachment *FindMobileDeviceEnrollmentProfilesByID200ApplicationJSONAttachmentsAttachment `json:"attachment,omitempty"`
}

type FindMobileDeviceEnrollmentProfilesByID200ApplicationJSONGeneral struct {
	Description *string `json:"description,omitempty"`
	ID          *int64  `json:"id,omitempty"`
	Invitation  *string `json:"invitation,omitempty"`
	// Name of the enrollment profile
	Name string  `json:"name"`
	Udid *string `json:"udid,omitempty"`
}

type FindMobileDeviceEnrollmentProfilesByID200ApplicationJSONLocation struct {
	Building     *string `json:"building,omitempty"`
	Department   *string `json:"department,omitempty"`
	EmailAddress *string `json:"email_address,omitempty"`
	Phone        *string `json:"phone,omitempty"`
	PhoneNumber  *string `json:"phone_number,omitempty"`
	Position     *string `json:"position,omitempty"`
	RealName     *string `json:"real_name,omitempty"`
	Realname     *string `json:"realname,omitempty"`
	Room         *string `json:"room,omitempty"`
	Username     *string `json:"username,omitempty"`
}

type FindMobileDeviceEnrollmentProfilesByID200ApplicationJSONPurchasing struct {
	ApplecareID          *string `json:"applecare_id,omitempty"`
	IsLeased             *bool   `json:"is_leased,omitempty"`
	IsPurchased          *bool   `json:"is_purchased,omitempty"`
	LeaseExpires         *string `json:"lease_expires,omitempty"`
	LeaseExpiresEpoch    *int64  `json:"lease_expires_epoch,omitempty"`
	LeaseExpiresUtc      *string `json:"lease_expires_utc,omitempty"`
	LifeExpectancy       *int64  `json:"life_expectancy,omitempty"`
	PoDate               *string `json:"po_date,omitempty"`
	PoDateEpoch          *int64  `json:"po_date_epoch,omitempty"`
	PoDateUtc            *string `json:"po_date_utc,omitempty"`
	PoNumber             *string `json:"po_number,omitempty"`
	PurchasePrice        *string `json:"purchase_price,omitempty"`
	PurchasingAccount    *string `json:"purchasing_account,omitempty"`
	PurchasingContact    *string `json:"purchasing_contact,omitempty"`
	Vendor               *string `json:"vendor,omitempty"`
	WarrantyExpires      *string `json:"warranty_expires,omitempty"`
	WarrantyExpiresEpoch *int64  `json:"warranty_expires_epoch,omitempty"`
	WarrantyExpiresUtc   *string `json:"warranty_expires_utc,omitempty"`
}

// FindMobileDeviceEnrollmentProfilesByID200ApplicationJSON - OK
type FindMobileDeviceEnrollmentProfilesByID200ApplicationJSON struct {
	Attachments []FindMobileDeviceEnrollmentProfilesByID200ApplicationJSONAttachments `json:"attachments,omitempty"`
	General     *FindMobileDeviceEnrollmentProfilesByID200ApplicationJSONGeneral      `json:"general,omitempty"`
	Location    *FindMobileDeviceEnrollmentProfilesByID200ApplicationJSONLocation     `json:"location,omitempty"`
	Purchasing  *FindMobileDeviceEnrollmentProfilesByID200ApplicationJSONPurchasing   `json:"purchasing,omitempty"`
}

type FindMobileDeviceEnrollmentProfilesByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDeviceEnrollmentProfilesByID200ApplicationJSONObject *FindMobileDeviceEnrollmentProfilesByID200ApplicationJSON
}
