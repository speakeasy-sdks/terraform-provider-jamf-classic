// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindPrintersByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

// FindPrintersByID200ApplicationXML - OK
type FindPrintersByID200ApplicationXML struct {
	CUPSName    *string
	Category    *string
	ID          *int64
	Info        *string
	Location    *string
	MakeDefault *bool
	Model       *string
	// Name of the printer
	Name        string
	Notes       *string
	Ppd         *string
	PpdContents *string
	PpdPath     *string
	URI         *string
	UseGeneric  *bool
}

// FindPrintersByID200ApplicationJSON - OK
type FindPrintersByID200ApplicationJSON struct {
	CUPSName    *string `json:"CUPS_name,omitempty"`
	Category    *string `json:"category,omitempty"`
	ID          *int64  `json:"id,omitempty"`
	Info        *string `json:"info,omitempty"`
	Location    *string `json:"location,omitempty"`
	MakeDefault *bool   `json:"make_default,omitempty"`
	Model       *string `json:"model,omitempty"`
	// Name of the printer
	Name        string  `json:"name"`
	Notes       *string `json:"notes,omitempty"`
	Ppd         *string `json:"ppd,omitempty"`
	PpdContents *string `json:"ppd_contents,omitempty"`
	PpdPath     *string `json:"ppd_path,omitempty"`
	URI         *string `json:"uri,omitempty"`
	UseGeneric  *bool   `json:"use_generic,omitempty"`
}

type FindPrintersByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindPrintersByID200ApplicationJSONObject *FindPrintersByID200ApplicationJSON
}
