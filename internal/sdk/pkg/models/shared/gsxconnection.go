// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type GsxConnectionRegion string

const (
	GsxConnectionRegionAmericas     GsxConnectionRegion = "Americas"
	GsxConnectionRegionApac         GsxConnectionRegion = "APAC"
	GsxConnectionRegionEmea         GsxConnectionRegion = "EMEA"
	GsxConnectionRegionLatinAmerica GsxConnectionRegion = "LatinAmerica"
)

func (e GsxConnectionRegion) ToPointer() *GsxConnectionRegion {
	return &e
}

func (e *GsxConnectionRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Americas":
		fallthrough
	case "APAC":
		fallthrough
	case "EMEA":
		fallthrough
	case "LatinAmerica":
		*e = GsxConnectionRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GsxConnectionRegion: %v", v)
	}
}

// GsxConnection - OK
type GsxConnection struct {
	AccountNumber *int64               `json:"account_number,omitempty"`
	Enabled       *bool                `json:"enabled,omitempty"`
	Region        *GsxConnectionRegion `json:"region,omitempty"`
	URI           *string              `json:"uri,omitempty"`
	Username      *string              `json:"username,omitempty"`
}
