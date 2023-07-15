// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// VppAccount - OK
type VppAccount struct {
	AccountName              *string `json:"account_name,omitempty"`
	AppleID                  *string `json:"apple_id,omitempty"`
	AutoRegisterManagedUsers *bool   `json:"auto_register_managed_users,omitempty"`
	Contact                  *string `json:"contact,omitempty"`
	Country                  *string `json:"country,omitempty"`
	ExpirationDate           *string `json:"expiration_date,omitempty"`
	ID                       *int64  `json:"id,omitempty"`
	// Name of the VPP account
	Name                          string      `json:"name"`
	NotifyDisassociation          *bool       `json:"notify_disassociation,omitempty"`
	PopulateCatalogFromVppContent *bool       `json:"populate_catalog_from_vpp_content,omitempty"`
	ServiceToken                  string      `json:"service_token"`
	Site                          *SiteObject `json:"site,omitempty"`
}