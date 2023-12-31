// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindMobileDeviceEnrollmentProfilesByNameSubsetSubset - Subset to filter by
type FindMobileDeviceEnrollmentProfilesByNameSubsetSubset string

const (
	FindMobileDeviceEnrollmentProfilesByNameSubsetSubsetGeneral     FindMobileDeviceEnrollmentProfilesByNameSubsetSubset = "General"
	FindMobileDeviceEnrollmentProfilesByNameSubsetSubsetLocation    FindMobileDeviceEnrollmentProfilesByNameSubsetSubset = "Location"
	FindMobileDeviceEnrollmentProfilesByNameSubsetSubsetPurchasing  FindMobileDeviceEnrollmentProfilesByNameSubsetSubset = "Purchasing"
	FindMobileDeviceEnrollmentProfilesByNameSubsetSubsetAttachments FindMobileDeviceEnrollmentProfilesByNameSubsetSubset = "Attachments"
)

func (e FindMobileDeviceEnrollmentProfilesByNameSubsetSubset) ToPointer() *FindMobileDeviceEnrollmentProfilesByNameSubsetSubset {
	return &e
}

func (e *FindMobileDeviceEnrollmentProfilesByNameSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "Location":
		fallthrough
	case "Purchasing":
		fallthrough
	case "Attachments":
		*e = FindMobileDeviceEnrollmentProfilesByNameSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceEnrollmentProfilesByNameSubsetSubset: %v", v)
	}
}

type FindMobileDeviceEnrollmentProfilesByNameSubsetRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Subset to filter by
	Subset FindMobileDeviceEnrollmentProfilesByNameSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXMLAttachmentsAttachment struct {
	Filename *string
	ID       *int64
	URI      *string
}

type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXMLAttachments struct {
	Attachment *FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXMLAttachmentsAttachment
}

type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXMLGeneral struct {
	Description *string
	ID          *int64
	Invitation  *string
	// Name of the enrollment profile
	Name string
	Udid *string
}

type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXMLLocation struct {
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

type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXMLPurchasing struct {
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

// FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXML - OK
type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXML struct {
	Attachments []FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXMLAttachments
	General     *FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXMLGeneral
	Location    *FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXMLLocation
	Purchasing  *FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationXMLPurchasing
}

type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSONAttachmentsAttachment struct {
	Filename *string `json:"filename,omitempty"`
	ID       *int64  `json:"id,omitempty"`
	URI      *string `json:"uri,omitempty"`
}

type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSONAttachments struct {
	Attachment *FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSONAttachmentsAttachment `json:"attachment,omitempty"`
}

type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSONGeneral struct {
	Description *string `json:"description,omitempty"`
	ID          *int64  `json:"id,omitempty"`
	Invitation  *string `json:"invitation,omitempty"`
	// Name of the enrollment profile
	Name string  `json:"name"`
	Udid *string `json:"udid,omitempty"`
}

type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSONLocation struct {
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

type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSONPurchasing struct {
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

// FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSON - OK
type FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSON struct {
	Attachments []FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSONAttachments `json:"attachments,omitempty"`
	General     *FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSONGeneral      `json:"general,omitempty"`
	Location    *FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSONLocation     `json:"location,omitempty"`
	Purchasing  *FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSONPurchasing   `json:"purchasing,omitempty"`
}

type FindMobileDeviceEnrollmentProfilesByNameSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSONObject *FindMobileDeviceEnrollmentProfilesByNameSubset200ApplicationJSON
}
