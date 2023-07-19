// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindComputerManagementByIDSubsetSubset - Subset to filter by
type FindComputerManagementByIDSubsetSubset string

const (
	FindComputerManagementByIDSubsetSubsetGeneral                      FindComputerManagementByIDSubsetSubset = "General"
	FindComputerManagementByIDSubsetSubsetPolicies                     FindComputerManagementByIDSubsetSubset = "Policies"
	FindComputerManagementByIDSubsetSubsetEbooks                       FindComputerManagementByIDSubsetSubset = "Ebooks"
	FindComputerManagementByIDSubsetSubsetMacAppStoreApps              FindComputerManagementByIDSubsetSubset = "MacAppStoreApps"
	FindComputerManagementByIDSubsetSubsetOsxConfigurationProfiles     FindComputerManagementByIDSubsetSubset = "OSXConfigurationProfiles"
	FindComputerManagementByIDSubsetSubsetManagedPreferenceProfiles    FindComputerManagementByIDSubsetSubset = "ManagedPreferenceProfiles"
	FindComputerManagementByIDSubsetSubsetRestrictedSoftware           FindComputerManagementByIDSubsetSubset = "RestrictedSoftware"
	FindComputerManagementByIDSubsetSubsetSmartGroups                  FindComputerManagementByIDSubsetSubset = "SmartGroups"
	FindComputerManagementByIDSubsetSubsetStaticGroups                 FindComputerManagementByIDSubsetSubset = "StaticGroups"
	FindComputerManagementByIDSubsetSubsetPatchReportingSoftwareTitles FindComputerManagementByIDSubsetSubset = "PatchReportingSoftwareTitles"
)

func (e FindComputerManagementByIDSubsetSubset) ToPointer() *FindComputerManagementByIDSubsetSubset {
	return &e
}

func (e *FindComputerManagementByIDSubsetSubset) UnmarshalJSON(data []byte) error {
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
		*e = FindComputerManagementByIDSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerManagementByIDSubsetSubset: %v", v)
	}
}

type FindComputerManagementByIDSubsetRequest struct {
	// Computer ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Subset to filter by
	Subset FindComputerManagementByIDSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindComputerManagementByIDSubset200ApplicationXMLEbooksEbook struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDSubset200ApplicationXMLEbooks struct {
	Ebook *FindComputerManagementByIDSubset200ApplicationXMLEbooksEbook
}

type FindComputerManagementByIDSubset200ApplicationXMLGeneral struct {
	ID         *int64
	MacAddress *string
	// Name of the computer
	Name         *string
	SerialNumber *string
	Udid         *string
}

type FindComputerManagementByIDSubset200ApplicationXMLMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDSubset200ApplicationXMLMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementByIDSubset200ApplicationXMLMacAppStoreAppsMacAppStoreApp
}

type FindComputerManagementByIDSubset200ApplicationXMLManagedPreferenceProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDSubset200ApplicationXMLManagedPreferenceProfiles struct {
	Profile *FindComputerManagementByIDSubset200ApplicationXMLManagedPreferenceProfilesProfile
}

type FindComputerManagementByIDSubset200ApplicationXMLOsXConfigurationProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDSubset200ApplicationXMLOsXConfigurationProfiles struct {
	Profile *FindComputerManagementByIDSubset200ApplicationXMLOsXConfigurationProfilesProfile
}

type FindComputerManagementByIDSubset200ApplicationXMLPatchPoliciesPatchPolicy struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDSubset200ApplicationXMLPatchPolicies struct {
	PatchPolicy *FindComputerManagementByIDSubset200ApplicationXMLPatchPoliciesPatchPolicy
}

type FindComputerManagementByIDSubset200ApplicationXMLPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string
	LatestVersion    *string
	Name             *string
}

type FindComputerManagementByIDSubset200ApplicationXMLPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementByIDSubset200ApplicationXMLPatchReportingSoftwareTitlesTitle
}

type FindComputerManagementByIDSubset200ApplicationXMLPoliciesPolicy struct {
	ID      *int64
	Name    *string
	Trigger *string
}

type FindComputerManagementByIDSubset200ApplicationXMLPolicies struct {
	Policy *FindComputerManagementByIDSubset200ApplicationXMLPoliciesPolicy
}

type FindComputerManagementByIDSubset200ApplicationXMLRestrictedSoftwareSoftware struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDSubset200ApplicationXMLRestrictedSoftware struct {
	Software *FindComputerManagementByIDSubset200ApplicationXMLRestrictedSoftwareSoftware
}

type FindComputerManagementByIDSubset200ApplicationXMLSmartGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDSubset200ApplicationXMLSmartGroups struct {
	Group *FindComputerManagementByIDSubset200ApplicationXMLSmartGroupsGroup
}

type FindComputerManagementByIDSubset200ApplicationXMLStaticGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDSubset200ApplicationXMLStaticGroups struct {
	Group *FindComputerManagementByIDSubset200ApplicationXMLStaticGroupsGroup
}

// FindComputerManagementByIDSubset200ApplicationXML - OK
type FindComputerManagementByIDSubset200ApplicationXML struct {
	Ebooks                       []FindComputerManagementByIDSubset200ApplicationXMLEbooks
	General                      *FindComputerManagementByIDSubset200ApplicationXMLGeneral
	MacAppStoreApps              []FindComputerManagementByIDSubset200ApplicationXMLMacAppStoreApps
	ManagedPreferenceProfiles    []FindComputerManagementByIDSubset200ApplicationXMLManagedPreferenceProfiles
	OsXConfigurationProfiles     []FindComputerManagementByIDSubset200ApplicationXMLOsXConfigurationProfiles
	PatchPolicies                []FindComputerManagementByIDSubset200ApplicationXMLPatchPolicies
	PatchReportingSoftwareTitles []FindComputerManagementByIDSubset200ApplicationXMLPatchReportingSoftwareTitles
	Policies                     []FindComputerManagementByIDSubset200ApplicationXMLPolicies
	RestrictedSoftware           []FindComputerManagementByIDSubset200ApplicationXMLRestrictedSoftware
	SmartGroups                  []FindComputerManagementByIDSubset200ApplicationXMLSmartGroups
	StaticGroups                 []FindComputerManagementByIDSubset200ApplicationXMLStaticGroups
}

type FindComputerManagementByIDSubset200ApplicationJSONEbooksEbook struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONEbooks struct {
	Ebook *FindComputerManagementByIDSubset200ApplicationJSONEbooksEbook `json:"ebook,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONGeneral struct {
	ID         *int64  `json:"id,omitempty"`
	MacAddress *string `json:"mac_address,omitempty"`
	// Name of the computer
	Name         *string `json:"name,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Udid         *string `json:"udid,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementByIDSubset200ApplicationJSONMacAppStoreAppsMacAppStoreApp `json:"mac_app_store_app,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONManagedPreferenceProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONManagedPreferenceProfiles struct {
	Profile *FindComputerManagementByIDSubset200ApplicationJSONManagedPreferenceProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONOsXConfigurationProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONOsXConfigurationProfiles struct {
	Profile *FindComputerManagementByIDSubset200ApplicationJSONOsXConfigurationProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONPatchPoliciesPatchPolicy struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONPatchPolicies struct {
	PatchPolicy *FindComputerManagementByIDSubset200ApplicationJSONPatchPoliciesPatchPolicy `json:"patch_policy,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string `json:"installed_version,omitempty"`
	LatestVersion    *string `json:"latest_version,omitempty"`
	Name             *string `json:"name,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementByIDSubset200ApplicationJSONPatchReportingSoftwareTitlesTitle `json:"title,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONPoliciesPolicy struct {
	ID      *int64  `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Trigger *string `json:"trigger,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONPolicies struct {
	Policy *FindComputerManagementByIDSubset200ApplicationJSONPoliciesPolicy `json:"policy,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONRestrictedSoftwareSoftware struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONRestrictedSoftware struct {
	Software *FindComputerManagementByIDSubset200ApplicationJSONRestrictedSoftwareSoftware `json:"software,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONSmartGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONSmartGroups struct {
	Group *FindComputerManagementByIDSubset200ApplicationJSONSmartGroupsGroup `json:"group,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONStaticGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDSubset200ApplicationJSONStaticGroups struct {
	Group *FindComputerManagementByIDSubset200ApplicationJSONStaticGroupsGroup `json:"group,omitempty"`
}

// FindComputerManagementByIDSubset200ApplicationJSON - OK
type FindComputerManagementByIDSubset200ApplicationJSON struct {
	Ebooks                       []FindComputerManagementByIDSubset200ApplicationJSONEbooks                       `json:"ebooks,omitempty"`
	General                      *FindComputerManagementByIDSubset200ApplicationJSONGeneral                       `json:"general,omitempty"`
	MacAppStoreApps              []FindComputerManagementByIDSubset200ApplicationJSONMacAppStoreApps              `json:"mac_app_store_apps,omitempty"`
	ManagedPreferenceProfiles    []FindComputerManagementByIDSubset200ApplicationJSONManagedPreferenceProfiles    `json:"managed_preference_profiles,omitempty"`
	OsXConfigurationProfiles     []FindComputerManagementByIDSubset200ApplicationJSONOsXConfigurationProfiles     `json:"os_x_configuration_profiles,omitempty"`
	PatchPolicies                []FindComputerManagementByIDSubset200ApplicationJSONPatchPolicies                `json:"patch_policies,omitempty"`
	PatchReportingSoftwareTitles []FindComputerManagementByIDSubset200ApplicationJSONPatchReportingSoftwareTitles `json:"patch_reporting_software_titles,omitempty"`
	Policies                     []FindComputerManagementByIDSubset200ApplicationJSONPolicies                     `json:"policies,omitempty"`
	RestrictedSoftware           []FindComputerManagementByIDSubset200ApplicationJSONRestrictedSoftware           `json:"restricted_software,omitempty"`
	SmartGroups                  []FindComputerManagementByIDSubset200ApplicationJSONSmartGroups                  `json:"smart_groups,omitempty"`
	StaticGroups                 []FindComputerManagementByIDSubset200ApplicationJSONStaticGroups                 `json:"static_groups,omitempty"`
}

type FindComputerManagementByIDSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindComputerManagementByIDSubset200ApplicationJSONObject *FindComputerManagementByIDSubset200ApplicationJSON
}
