// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindComputerManagementByNameUsernameSubsetSubset - Subset to filter by
type FindComputerManagementByNameUsernameSubsetSubset string

const (
	FindComputerManagementByNameUsernameSubsetSubsetGeneral                      FindComputerManagementByNameUsernameSubsetSubset = "General"
	FindComputerManagementByNameUsernameSubsetSubsetPolicies                     FindComputerManagementByNameUsernameSubsetSubset = "Policies"
	FindComputerManagementByNameUsernameSubsetSubsetEbooks                       FindComputerManagementByNameUsernameSubsetSubset = "Ebooks"
	FindComputerManagementByNameUsernameSubsetSubsetMacAppStoreApps              FindComputerManagementByNameUsernameSubsetSubset = "MacAppStoreApps"
	FindComputerManagementByNameUsernameSubsetSubsetOsxConfigurationProfiles     FindComputerManagementByNameUsernameSubsetSubset = "OSXConfigurationProfiles"
	FindComputerManagementByNameUsernameSubsetSubsetManagedPreferenceProfiles    FindComputerManagementByNameUsernameSubsetSubset = "ManagedPreferenceProfiles"
	FindComputerManagementByNameUsernameSubsetSubsetRestrictedSoftware           FindComputerManagementByNameUsernameSubsetSubset = "RestrictedSoftware"
	FindComputerManagementByNameUsernameSubsetSubsetSmartGroups                  FindComputerManagementByNameUsernameSubsetSubset = "SmartGroups"
	FindComputerManagementByNameUsernameSubsetSubsetStaticGroups                 FindComputerManagementByNameUsernameSubsetSubset = "StaticGroups"
	FindComputerManagementByNameUsernameSubsetSubsetPatchReportingSoftwareTitles FindComputerManagementByNameUsernameSubsetSubset = "PatchReportingSoftwareTitles"
)

func (e FindComputerManagementByNameUsernameSubsetSubset) ToPointer() *FindComputerManagementByNameUsernameSubsetSubset {
	return &e
}

func (e *FindComputerManagementByNameUsernameSubsetSubset) UnmarshalJSON(data []byte) error {
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
		*e = FindComputerManagementByNameUsernameSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindComputerManagementByNameUsernameSubsetSubset: %v", v)
	}
}

type FindComputerManagementByNameUsernameSubsetRequest struct {
	// Computer name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Subset to filter by
	Subset FindComputerManagementByNameUsernameSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
	// Username to filter by
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLEbooksEbook struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLEbooks struct {
	Ebook *FindComputerManagementByNameUsernameSubset200ApplicationXMLEbooksEbook
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLGeneral struct {
	ID         *int64
	MacAddress *string
	// Name of the computer
	Name         *string
	SerialNumber *string
	Udid         *string
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementByNameUsernameSubset200ApplicationXMLMacAppStoreAppsMacAppStoreApp
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLManagedPreferenceProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLManagedPreferenceProfiles struct {
	Profile *FindComputerManagementByNameUsernameSubset200ApplicationXMLManagedPreferenceProfilesProfile
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLOsXConfigurationProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLOsXConfigurationProfiles struct {
	Profile *FindComputerManagementByNameUsernameSubset200ApplicationXMLOsXConfigurationProfilesProfile
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLPatchPoliciesPatchPolicy struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLPatchPolicies struct {
	PatchPolicy *FindComputerManagementByNameUsernameSubset200ApplicationXMLPatchPoliciesPatchPolicy
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string
	LatestVersion    *string
	Name             *string
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementByNameUsernameSubset200ApplicationXMLPatchReportingSoftwareTitlesTitle
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLPoliciesPolicy struct {
	ID      *int64
	Name    *string
	Trigger *string
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLPolicies struct {
	Policy *FindComputerManagementByNameUsernameSubset200ApplicationXMLPoliciesPolicy
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLRestrictedSoftwareSoftware struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLRestrictedSoftware struct {
	Software *FindComputerManagementByNameUsernameSubset200ApplicationXMLRestrictedSoftwareSoftware
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLSmartGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLSmartGroups struct {
	Group *FindComputerManagementByNameUsernameSubset200ApplicationXMLSmartGroupsGroup
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLStaticGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByNameUsernameSubset200ApplicationXMLStaticGroups struct {
	Group *FindComputerManagementByNameUsernameSubset200ApplicationXMLStaticGroupsGroup
}

// FindComputerManagementByNameUsernameSubset200ApplicationXML - OK
type FindComputerManagementByNameUsernameSubset200ApplicationXML struct {
	Ebooks                       []FindComputerManagementByNameUsernameSubset200ApplicationXMLEbooks
	General                      *FindComputerManagementByNameUsernameSubset200ApplicationXMLGeneral
	MacAppStoreApps              []FindComputerManagementByNameUsernameSubset200ApplicationXMLMacAppStoreApps
	ManagedPreferenceProfiles    []FindComputerManagementByNameUsernameSubset200ApplicationXMLManagedPreferenceProfiles
	OsXConfigurationProfiles     []FindComputerManagementByNameUsernameSubset200ApplicationXMLOsXConfigurationProfiles
	PatchPolicies                []FindComputerManagementByNameUsernameSubset200ApplicationXMLPatchPolicies
	PatchReportingSoftwareTitles []FindComputerManagementByNameUsernameSubset200ApplicationXMLPatchReportingSoftwareTitles
	Policies                     []FindComputerManagementByNameUsernameSubset200ApplicationXMLPolicies
	RestrictedSoftware           []FindComputerManagementByNameUsernameSubset200ApplicationXMLRestrictedSoftware
	SmartGroups                  []FindComputerManagementByNameUsernameSubset200ApplicationXMLSmartGroups
	StaticGroups                 []FindComputerManagementByNameUsernameSubset200ApplicationXMLStaticGroups
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONEbooksEbook struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONEbooks struct {
	Ebook *FindComputerManagementByNameUsernameSubset200ApplicationJSONEbooksEbook `json:"ebook,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONGeneral struct {
	ID         *int64  `json:"id,omitempty"`
	MacAddress *string `json:"mac_address,omitempty"`
	// Name of the computer
	Name         *string `json:"name,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Udid         *string `json:"udid,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementByNameUsernameSubset200ApplicationJSONMacAppStoreAppsMacAppStoreApp `json:"mac_app_store_app,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONManagedPreferenceProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONManagedPreferenceProfiles struct {
	Profile *FindComputerManagementByNameUsernameSubset200ApplicationJSONManagedPreferenceProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONOsXConfigurationProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONOsXConfigurationProfiles struct {
	Profile *FindComputerManagementByNameUsernameSubset200ApplicationJSONOsXConfigurationProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONPatchPoliciesPatchPolicy struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONPatchPolicies struct {
	PatchPolicy *FindComputerManagementByNameUsernameSubset200ApplicationJSONPatchPoliciesPatchPolicy `json:"patch_policy,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string `json:"installed_version,omitempty"`
	LatestVersion    *string `json:"latest_version,omitempty"`
	Name             *string `json:"name,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementByNameUsernameSubset200ApplicationJSONPatchReportingSoftwareTitlesTitle `json:"title,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONPoliciesPolicy struct {
	ID      *int64  `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Trigger *string `json:"trigger,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONPolicies struct {
	Policy *FindComputerManagementByNameUsernameSubset200ApplicationJSONPoliciesPolicy `json:"policy,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONRestrictedSoftwareSoftware struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONRestrictedSoftware struct {
	Software *FindComputerManagementByNameUsernameSubset200ApplicationJSONRestrictedSoftwareSoftware `json:"software,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONSmartGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONSmartGroups struct {
	Group *FindComputerManagementByNameUsernameSubset200ApplicationJSONSmartGroupsGroup `json:"group,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONStaticGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByNameUsernameSubset200ApplicationJSONStaticGroups struct {
	Group *FindComputerManagementByNameUsernameSubset200ApplicationJSONStaticGroupsGroup `json:"group,omitempty"`
}

// FindComputerManagementByNameUsernameSubset200ApplicationJSON - OK
type FindComputerManagementByNameUsernameSubset200ApplicationJSON struct {
	Ebooks                       []FindComputerManagementByNameUsernameSubset200ApplicationJSONEbooks                       `json:"ebooks,omitempty"`
	General                      *FindComputerManagementByNameUsernameSubset200ApplicationJSONGeneral                       `json:"general,omitempty"`
	MacAppStoreApps              []FindComputerManagementByNameUsernameSubset200ApplicationJSONMacAppStoreApps              `json:"mac_app_store_apps,omitempty"`
	ManagedPreferenceProfiles    []FindComputerManagementByNameUsernameSubset200ApplicationJSONManagedPreferenceProfiles    `json:"managed_preference_profiles,omitempty"`
	OsXConfigurationProfiles     []FindComputerManagementByNameUsernameSubset200ApplicationJSONOsXConfigurationProfiles     `json:"os_x_configuration_profiles,omitempty"`
	PatchPolicies                []FindComputerManagementByNameUsernameSubset200ApplicationJSONPatchPolicies                `json:"patch_policies,omitempty"`
	PatchReportingSoftwareTitles []FindComputerManagementByNameUsernameSubset200ApplicationJSONPatchReportingSoftwareTitles `json:"patch_reporting_software_titles,omitempty"`
	Policies                     []FindComputerManagementByNameUsernameSubset200ApplicationJSONPolicies                     `json:"policies,omitempty"`
	RestrictedSoftware           []FindComputerManagementByNameUsernameSubset200ApplicationJSONRestrictedSoftware           `json:"restricted_software,omitempty"`
	SmartGroups                  []FindComputerManagementByNameUsernameSubset200ApplicationJSONSmartGroups                  `json:"smart_groups,omitempty"`
	StaticGroups                 []FindComputerManagementByNameUsernameSubset200ApplicationJSONStaticGroups                 `json:"static_groups,omitempty"`
}

type FindComputerManagementByNameUsernameSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindComputerManagementByNameUsernameSubset200ApplicationJSONObject *FindComputerManagementByNameUsernameSubset200ApplicationJSON
}
