// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindPeripheralsByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type FindPeripheralsByID200ApplicationXMLAttachmentsAttachment struct {
	Filename *string
	ID       *int64
	URI      *string
}

type FindPeripheralsByID200ApplicationXMLAttachments struct {
	Attachment *FindPeripheralsByID200ApplicationXMLAttachmentsAttachment
}

type FindPeripheralsByID200ApplicationXMLGeneralFieldsField struct {
	Name  *string
	Value *string
}

type FindPeripheralsByID200ApplicationXMLGeneralFields struct {
	Field *FindPeripheralsByID200ApplicationXMLGeneralFieldsField
}

type FindPeripheralsByID200ApplicationXMLGeneral struct {
	BarCode1 *string
	BarCode2 *string
	Fields   []FindPeripheralsByID200ApplicationXMLGeneralFields
	ID       *int64
	// Name of the peripheral type
	Type string
}

type FindPeripheralsByID200ApplicationXMLLocation struct {
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

type FindPeripheralsByID200ApplicationXMLPurchasing struct {
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

// FindPeripheralsByID200ApplicationXML - OK
type FindPeripheralsByID200ApplicationXML struct {
	Attachments []FindPeripheralsByID200ApplicationXMLAttachments
	General     *FindPeripheralsByID200ApplicationXMLGeneral
	Location    *FindPeripheralsByID200ApplicationXMLLocation
	Purchasing  *FindPeripheralsByID200ApplicationXMLPurchasing
}

type FindPeripheralsByID200ApplicationJSONAttachmentsAttachment struct {
	Filename *string `json:"filename,omitempty"`
	ID       *int64  `json:"id,omitempty"`
	URI      *string `json:"uri,omitempty"`
}

type FindPeripheralsByID200ApplicationJSONAttachments struct {
	Attachment *FindPeripheralsByID200ApplicationJSONAttachmentsAttachment `json:"attachment,omitempty"`
}

type FindPeripheralsByID200ApplicationJSONGeneralFieldsField struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

type FindPeripheralsByID200ApplicationJSONGeneralFields struct {
	Field *FindPeripheralsByID200ApplicationJSONGeneralFieldsField `json:"field,omitempty"`
}

type FindPeripheralsByID200ApplicationJSONGeneral struct {
	BarCode1 *string                                              `json:"bar_code_1,omitempty"`
	BarCode2 *string                                              `json:"bar_code_2,omitempty"`
	Fields   []FindPeripheralsByID200ApplicationJSONGeneralFields `json:"fields,omitempty"`
	ID       *int64                                               `json:"id,omitempty"`
	// Name of the peripheral type
	Type string `json:"type"`
}

type FindPeripheralsByID200ApplicationJSONLocation struct {
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

type FindPeripheralsByID200ApplicationJSONPurchasing struct {
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

// FindPeripheralsByID200ApplicationJSON - OK
type FindPeripheralsByID200ApplicationJSON struct {
	Attachments []FindPeripheralsByID200ApplicationJSONAttachments `json:"attachments,omitempty"`
	General     *FindPeripheralsByID200ApplicationJSONGeneral      `json:"general,omitempty"`
	Location    *FindPeripheralsByID200ApplicationJSONLocation     `json:"location,omitempty"`
	Purchasing  *FindPeripheralsByID200ApplicationJSONPurchasing   `json:"purchasing,omitempty"`
}

type FindPeripheralsByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindPeripheralsByID200ApplicationJSONObject *FindPeripheralsByID200ApplicationJSON
}
