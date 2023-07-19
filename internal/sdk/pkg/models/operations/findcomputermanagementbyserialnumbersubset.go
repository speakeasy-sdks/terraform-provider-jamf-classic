// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindComputerManagementBySerialNumberSubsetSubset - Subset to filter by
type FindComputerManagementBySerialNumberSubsetSubset string

const (
	FindComputerManagementBySerialNumberSubsetSubsetGeneral                      FindComputerManagementBySerialNumberSubsetSubset = "General"
	FindComputerManagementBySerialNumberSubsetSubsetPolicies                     FindComputerManagementBySerialNumberSubsetSubset = "Policies"
	FindComputerManagementBySerialNumberSubsetSubsetEbooks                       FindComputerManagementBySerialNumberSubsetSubset = "Ebooks"
	FindComputerManagementBySerialNumberSubsetSubsetMacAppStoreApps              FindComputerManagementBySerialNumberSubsetSubset = "MacAppStoreApps"
	FindComputerManagementBySerialNumberSubsetSubsetOsxConfigurationProfiles     FindComputerManagementBySerialNumberSubsetSubset = "OSXConfigurationProfiles"
	FindComputerManagementBySerialNumberSubsetSubsetManagedPreferenceProfiles    FindComputerManagementBySerialNumberSubsetSubset = "ManagedPreferenceProfiles"
	FindComputerManagementBySerialNumberSubsetSubsetRestrictedSoftware           FindComputerManagementBySerialNumberSubsetSubset = "RestrictedSoftware"
	FindComputerManagementBySerialNumberSubsetSubsetSmartGroups                  FindComputerManagementBySerialNumberSubsetSubset = "SmartGroups"
	FindComputerManagementBySerialNumberSubsetSubsetStaticGroups                 FindComputerManagementBySerialNumberSubsetSubset = "StaticGroups"
	FindComputerManagementBySerialNumberSubsetSubsetPatchReportingSoftwareTitles FindComputerManagementBySerialNumberSubsetSubset = "PatchReportingSoftwareTitles"
)

func (e FindComputerManagementBySerialNumberSubsetSubset) ToPointer() *FindComputerManagementBySerialNumberSubsetSubset {
	return &e
}

func (e *FindComputerManagementBySerialNumberSubsetSubset) UnmarshalJSON(data []byte) error {
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
		*e = FindComputerManagementBySerialNumberSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerManagementBySerialNumberSubsetSubset: %v", v)
	}
}

type FindComputerManagementBySerialNumberSubsetRequest struct {
	// Computer Serial Number to filter by
	Serialnumber string `pathParam:"style=simple,explode=false,name=serialnumber"`
	// Subset to filter by
	Subset FindComputerManagementBySerialNumberSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLEbooksEbook struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLEbooks struct {
	Ebook *FindComputerManagementBySerialNumberSubset200ApplicationXMLEbooksEbook
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLGeneral struct {
	ID         *int64
	MacAddress *string
	// Name of the computer
	Name         *string
	SerialNumber *string
	Udid         *string
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementBySerialNumberSubset200ApplicationXMLMacAppStoreAppsMacAppStoreApp
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLManagedPreferenceProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLManagedPreferenceProfiles struct {
	Profile *FindComputerManagementBySerialNumberSubset200ApplicationXMLManagedPreferenceProfilesProfile
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLOsXConfigurationProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLOsXConfigurationProfiles struct {
	Profile *FindComputerManagementBySerialNumberSubset200ApplicationXMLOsXConfigurationProfilesProfile
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLPatchPoliciesPatchPolicy struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLPatchPolicies struct {
	PatchPolicy *FindComputerManagementBySerialNumberSubset200ApplicationXMLPatchPoliciesPatchPolicy
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string
	LatestVersion    *string
	Name             *string
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementBySerialNumberSubset200ApplicationXMLPatchReportingSoftwareTitlesTitle
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLPoliciesPolicy struct {
	ID      *int64
	Name    *string
	Trigger *string
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLPolicies struct {
	Policy *FindComputerManagementBySerialNumberSubset200ApplicationXMLPoliciesPolicy
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLRestrictedSoftwareSoftware struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLRestrictedSoftware struct {
	Software *FindComputerManagementBySerialNumberSubset200ApplicationXMLRestrictedSoftwareSoftware
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLSmartGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLSmartGroups struct {
	Group *FindComputerManagementBySerialNumberSubset200ApplicationXMLSmartGroupsGroup
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLStaticGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementBySerialNumberSubset200ApplicationXMLStaticGroups struct {
	Group *FindComputerManagementBySerialNumberSubset200ApplicationXMLStaticGroupsGroup
}

// FindComputerManagementBySerialNumberSubset200ApplicationXML - OK
type FindComputerManagementBySerialNumberSubset200ApplicationXML struct {
	Ebooks                       []FindComputerManagementBySerialNumberSubset200ApplicationXMLEbooks
	General                      *FindComputerManagementBySerialNumberSubset200ApplicationXMLGeneral
	MacAppStoreApps              []FindComputerManagementBySerialNumberSubset200ApplicationXMLMacAppStoreApps
	ManagedPreferenceProfiles    []FindComputerManagementBySerialNumberSubset200ApplicationXMLManagedPreferenceProfiles
	OsXConfigurationProfiles     []FindComputerManagementBySerialNumberSubset200ApplicationXMLOsXConfigurationProfiles
	PatchPolicies                []FindComputerManagementBySerialNumberSubset200ApplicationXMLPatchPolicies
	PatchReportingSoftwareTitles []FindComputerManagementBySerialNumberSubset200ApplicationXMLPatchReportingSoftwareTitles
	Policies                     []FindComputerManagementBySerialNumberSubset200ApplicationXMLPolicies
	RestrictedSoftware           []FindComputerManagementBySerialNumberSubset200ApplicationXMLRestrictedSoftware
	SmartGroups                  []FindComputerManagementBySerialNumberSubset200ApplicationXMLSmartGroups
	StaticGroups                 []FindComputerManagementBySerialNumberSubset200ApplicationXMLStaticGroups
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONEbooksEbook struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONEbooks struct {
	Ebook *FindComputerManagementBySerialNumberSubset200ApplicationJSONEbooksEbook `json:"ebook,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONGeneral struct {
	ID         *int64  `json:"id,omitempty"`
	MacAddress *string `json:"mac_address,omitempty"`
	// Name of the computer
	Name         *string `json:"name,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Udid         *string `json:"udid,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementBySerialNumberSubset200ApplicationJSONMacAppStoreAppsMacAppStoreApp `json:"mac_app_store_app,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONManagedPreferenceProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONManagedPreferenceProfiles struct {
	Profile *FindComputerManagementBySerialNumberSubset200ApplicationJSONManagedPreferenceProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONOsXConfigurationProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONOsXConfigurationProfiles struct {
	Profile *FindComputerManagementBySerialNumberSubset200ApplicationJSONOsXConfigurationProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONPatchPoliciesPatchPolicy struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONPatchPolicies struct {
	PatchPolicy *FindComputerManagementBySerialNumberSubset200ApplicationJSONPatchPoliciesPatchPolicy `json:"patch_policy,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string `json:"installed_version,omitempty"`
	LatestVersion    *string `json:"latest_version,omitempty"`
	Name             *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementBySerialNumberSubset200ApplicationJSONPatchReportingSoftwareTitlesTitle `json:"title,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONPoliciesPolicy struct {
	ID      *int64  `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Trigger *string `json:"trigger,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONPolicies struct {
	Policy *FindComputerManagementBySerialNumberSubset200ApplicationJSONPoliciesPolicy `json:"policy,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONRestrictedSoftwareSoftware struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONRestrictedSoftware struct {
	Software *FindComputerManagementBySerialNumberSubset200ApplicationJSONRestrictedSoftwareSoftware `json:"software,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONSmartGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONSmartGroups struct {
	Group *FindComputerManagementBySerialNumberSubset200ApplicationJSONSmartGroupsGroup `json:"group,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONStaticGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementBySerialNumberSubset200ApplicationJSONStaticGroups struct {
	Group *FindComputerManagementBySerialNumberSubset200ApplicationJSONStaticGroupsGroup `json:"group,omitempty"`
}

// FindComputerManagementBySerialNumberSubset200ApplicationJSON - OK
type FindComputerManagementBySerialNumberSubset200ApplicationJSON struct {
	Ebooks                       []FindComputerManagementBySerialNumberSubset200ApplicationJSONEbooks                       `json:"ebooks,omitempty"`
	General                      *FindComputerManagementBySerialNumberSubset200ApplicationJSONGeneral                       `json:"general,omitempty"`
	MacAppStoreApps              []FindComputerManagementBySerialNumberSubset200ApplicationJSONMacAppStoreApps              `json:"mac_app_store_apps,omitempty"`
	ManagedPreferenceProfiles    []FindComputerManagementBySerialNumberSubset200ApplicationJSONManagedPreferenceProfiles    `json:"managed_preference_profiles,omitempty"`
	OsXConfigurationProfiles     []FindComputerManagementBySerialNumberSubset200ApplicationJSONOsXConfigurationProfiles     `json:"os_x_configuration_profiles,omitempty"`
	PatchPolicies                []FindComputerManagementBySerialNumberSubset200ApplicationJSONPatchPolicies                `json:"patch_policies,omitempty"`
	PatchReportingSoftwareTitles []FindComputerManagementBySerialNumberSubset200ApplicationJSONPatchReportingSoftwareTitles `json:"patch_reporting_software_titles,omitempty"`
	Policies                     []FindComputerManagementBySerialNumberSubset200ApplicationJSONPolicies                     `json:"policies,omitempty"`
	RestrictedSoftware           []FindComputerManagementBySerialNumberSubset200ApplicationJSONRestrictedSoftware           `json:"restricted_software,omitempty"`
	SmartGroups                  []FindComputerManagementBySerialNumberSubset200ApplicationJSONSmartGroups                  `json:"smart_groups,omitempty"`
	StaticGroups                 []FindComputerManagementBySerialNumberSubset200ApplicationJSONStaticGroups                 `json:"static_groups,omitempty"`
}

type FindComputerManagementBySerialNumberSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindComputerManagementBySerialNumberSubset200ApplicationJSONObject *FindComputerManagementBySerialNumberSubset200ApplicationJSON
}
