// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindComputerManagementBySerialNumberUsernameSubsetSubset - Subset to filter by
type FindComputerManagementBySerialNumberUsernameSubsetSubset string

const (
	FindComputerManagementBySerialNumberUsernameSubsetSubsetGeneral                      FindComputerManagementBySerialNumberUsernameSubsetSubset = "General"
	FindComputerManagementBySerialNumberUsernameSubsetSubsetPolicies                     FindComputerManagementBySerialNumberUsernameSubsetSubset = "Policies"
	FindComputerManagementBySerialNumberUsernameSubsetSubsetEbooks                       FindComputerManagementBySerialNumberUsernameSubsetSubset = "Ebooks"
	FindComputerManagementBySerialNumberUsernameSubsetSubsetMacAppStoreApps              FindComputerManagementBySerialNumberUsernameSubsetSubset = "MacAppStoreApps"
	FindComputerManagementBySerialNumberUsernameSubsetSubsetOsxConfigurationProfiles     FindComputerManagementBySerialNumberUsernameSubsetSubset = "OSXConfigurationProfiles"
	FindComputerManagementBySerialNumberUsernameSubsetSubsetManagedPreferenceProfiles    FindComputerManagementBySerialNumberUsernameSubsetSubset = "ManagedPreferenceProfiles"
	FindComputerManagementBySerialNumberUsernameSubsetSubsetRestrictedSoftware           FindComputerManagementBySerialNumberUsernameSubsetSubset = "RestrictedSoftware"
	FindComputerManagementBySerialNumberUsernameSubsetSubsetSmartGroups                  FindComputerManagementBySerialNumberUsernameSubsetSubset = "SmartGroups"
	FindComputerManagementBySerialNumberUsernameSubsetSubsetStaticGroups                 FindComputerManagementBySerialNumberUsernameSubsetSubset = "StaticGroups"
	FindComputerManagementBySerialNumberUsernameSubsetSubsetPatchReportingSoftwareTitles FindComputerManagementBySerialNumberUsernameSubsetSubset = "PatchReportingSoftwareTitles"
)

func (e FindComputerManagementBySerialNumberUsernameSubsetSubset) ToPointer() *FindComputerManagementBySerialNumberUsernameSubsetSubset {
	return &e
}

func (e *FindComputerManagementBySerialNumberUsernameSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "Policies":
		fallthrough
	case "Ebooks":
		fallthrough
	case "MacAppStoreApps":
		fallthrough
	case "OSXConfigurationProfiles":
		fallthrough
	case "ManagedPreferenceProfiles":
		fallthrough
	case "RestrictedSoftware":
		fallthrough
	case "SmartGroups":
		fallthrough
	case "StaticGroups":
		fallthrough
	case "PatchReportingSoftwareTitles":
		*e = FindComputerManagementBySerialNumberUsernameSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerManagementBySerialNumberUsernameSubsetSubset: %v", v)
	}
}

type FindComputerManagementBySerialNumberUsernameSubsetRequest struct {
	// Computer serial number to filter by
	Serialnumber string `pathParam:"style=simple,explode=false,name=serialnumber"`
	// Subset to filter by
	Subset FindComputerManagementBySerialNumberUsernameSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
	// Username to filter by
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLEbooksEbook struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLEbooks struct {
	Ebook *FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLEbooksEbook
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLGeneral struct {
	ID         *int64
	MacAddress *string
	// Name of the computer
	Name         *string
	SerialNumber *string
	Udid         *string
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLMacAppStoreAppsMacAppStoreApp
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLManagedPreferenceProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLManagedPreferenceProfiles struct {
	Profile *FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLManagedPreferenceProfilesProfile
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLOsXConfigurationProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLOsXConfigurationProfiles struct {
	Profile *FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLOsXConfigurationProfilesProfile
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPatchPoliciesPatchPolicy struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPatchPolicies struct {
	PatchPolicy *FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPatchPoliciesPatchPolicy
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string
	LatestVersion    *string
	Name             *string
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPatchReportingSoftwareTitlesTitle
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPoliciesPolicy struct {
	ID      *int64
	Name    *string
	Trigger *string
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPolicies struct {
	Policy *FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPoliciesPolicy
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLRestrictedSoftwareSoftware struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLRestrictedSoftware struct {
	Software *FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLRestrictedSoftwareSoftware
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLSmartGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLSmartGroups struct {
	Group *FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLSmartGroupsGroup
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLStaticGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLStaticGroups struct {
	Group *FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLStaticGroupsGroup
}

// FindComputerManagementBySerialNumberUsernameSubset200ApplicationXML - OK
type FindComputerManagementBySerialNumberUsernameSubset200ApplicationXML struct {
	Ebooks                       []FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLEbooks
	General                      *FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLGeneral
	MacAppStoreApps              []FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLMacAppStoreApps
	ManagedPreferenceProfiles    []FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLManagedPreferenceProfiles
	OsXConfigurationProfiles     []FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLOsXConfigurationProfiles
	PatchPolicies                []FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPatchPolicies
	PatchReportingSoftwareTitles []FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPatchReportingSoftwareTitles
	Policies                     []FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLPolicies
	RestrictedSoftware           []FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLRestrictedSoftware
	SmartGroups                  []FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLSmartGroups
	StaticGroups                 []FindComputerManagementBySerialNumberUsernameSubset200ApplicationXMLStaticGroups
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONEbooksEbook struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONEbooks struct {
	Ebook *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONEbooksEbook `json:"ebook,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONGeneral struct {
	ID         *int64  `json:"id,omitempty"`
	MacAddress *string `json:"mac_address,omitempty"`
	// Name of the computer
	Name         *string `json:"name,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Udid         *string `json:"udid,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONMacAppStoreAppsMacAppStoreApp `json:"mac_app_store_app,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONManagedPreferenceProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONManagedPreferenceProfiles struct {
	Profile *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONManagedPreferenceProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONOsXConfigurationProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONOsXConfigurationProfiles struct {
	Profile *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONOsXConfigurationProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPatchPoliciesPatchPolicy struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPatchPolicies struct {
	PatchPolicy *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPatchPoliciesPatchPolicy `json:"patch_policy,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string `json:"installed_version,omitempty"`
	LatestVersion    *string `json:"latest_version,omitempty"`
	Name             *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPatchReportingSoftwareTitlesTitle `json:"title,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPoliciesPolicy struct {
	ID      *int64  `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Trigger *string `json:"trigger,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPolicies struct {
	Policy *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPoliciesPolicy `json:"policy,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONRestrictedSoftwareSoftware struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONRestrictedSoftware struct {
	Software *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONRestrictedSoftwareSoftware `json:"software,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONSmartGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONSmartGroups struct {
	Group *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONSmartGroupsGroup `json:"group,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONStaticGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONStaticGroups struct {
	Group *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONStaticGroupsGroup `json:"group,omitempty"`
}

// FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSON - OK
type FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSON struct {
	Ebooks                       []FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONEbooks                       `json:"ebooks,omitempty"`
	General                      *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONGeneral                       `json:"general,omitempty"`
	MacAppStoreApps              []FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONMacAppStoreApps              `json:"mac_app_store_apps,omitempty"`
	ManagedPreferenceProfiles    []FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONManagedPreferenceProfiles    `json:"managed_preference_profiles,omitempty"`
	OsXConfigurationProfiles     []FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONOsXConfigurationProfiles     `json:"os_x_configuration_profiles,omitempty"`
	PatchPolicies                []FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPatchPolicies                `json:"patch_policies,omitempty"`
	PatchReportingSoftwareTitles []FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPatchReportingSoftwareTitles `json:"patch_reporting_software_titles,omitempty"`
	Policies                     []FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONPolicies                     `json:"policies,omitempty"`
	RestrictedSoftware           []FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONRestrictedSoftware           `json:"restricted_software,omitempty"`
	SmartGroups                  []FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONSmartGroups                  `json:"smart_groups,omitempty"`
	StaticGroups                 []FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONStaticGroups                 `json:"static_groups,omitempty"`
}

type FindComputerManagementBySerialNumberUsernameSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSONObject *FindComputerManagementBySerialNumberUsernameSubset200ApplicationJSON
}
