// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindComputerManagementByMacAddressPatchFilterRequest struct {
	// filter to apply
	Filter string `pathParam:"style=simple,explode=false,name=filter"`
	// Computer Mac address to filter by
	Macaddress string `pathParam:"style=simple,explode=false,name=macaddress"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLEbooksEbook struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLEbooks struct {
	Ebook *FindComputerManagementByMacAddressPatchFilter200ApplicationXMLEbooksEbook
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLGeneral struct {
	ID         *int64
	MacAddress *string
	// Name of the computer
	Name         *string
	SerialNumber *string
	Udid         *string
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementByMacAddressPatchFilter200ApplicationXMLMacAppStoreAppsMacAppStoreApp
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLManagedPreferenceProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLManagedPreferenceProfiles struct {
	Profile *FindComputerManagementByMacAddressPatchFilter200ApplicationXMLManagedPreferenceProfilesProfile
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLOsXConfigurationProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLOsXConfigurationProfiles struct {
	Profile *FindComputerManagementByMacAddressPatchFilter200ApplicationXMLOsXConfigurationProfilesProfile
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPatchPoliciesPatchPolicy struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPatchPolicies struct {
	PatchPolicy *FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPatchPoliciesPatchPolicy
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string
	LatestVersion    *string
	Name             *string
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPatchReportingSoftwareTitlesTitle
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPoliciesPolicy struct {
	ID      *int64
	Name    *string
	Trigger *string
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPolicies struct {
	Policy *FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPoliciesPolicy
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLRestrictedSoftwareSoftware struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLRestrictedSoftware struct {
	Software *FindComputerManagementByMacAddressPatchFilter200ApplicationXMLRestrictedSoftwareSoftware
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLSmartGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLSmartGroups struct {
	Group *FindComputerManagementByMacAddressPatchFilter200ApplicationXMLSmartGroupsGroup
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLStaticGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationXMLStaticGroups struct {
	Group *FindComputerManagementByMacAddressPatchFilter200ApplicationXMLStaticGroupsGroup
}

// FindComputerManagementByMacAddressPatchFilter200ApplicationXML - OK
type FindComputerManagementByMacAddressPatchFilter200ApplicationXML struct {
	Ebooks                       []FindComputerManagementByMacAddressPatchFilter200ApplicationXMLEbooks
	General                      *FindComputerManagementByMacAddressPatchFilter200ApplicationXMLGeneral
	MacAppStoreApps              []FindComputerManagementByMacAddressPatchFilter200ApplicationXMLMacAppStoreApps
	ManagedPreferenceProfiles    []FindComputerManagementByMacAddressPatchFilter200ApplicationXMLManagedPreferenceProfiles
	OsXConfigurationProfiles     []FindComputerManagementByMacAddressPatchFilter200ApplicationXMLOsXConfigurationProfiles
	PatchPolicies                []FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPatchPolicies
	PatchReportingSoftwareTitles []FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPatchReportingSoftwareTitles
	Policies                     []FindComputerManagementByMacAddressPatchFilter200ApplicationXMLPolicies
	RestrictedSoftware           []FindComputerManagementByMacAddressPatchFilter200ApplicationXMLRestrictedSoftware
	SmartGroups                  []FindComputerManagementByMacAddressPatchFilter200ApplicationXMLSmartGroups
	StaticGroups                 []FindComputerManagementByMacAddressPatchFilter200ApplicationXMLStaticGroups
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONEbooksEbook struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONEbooks struct {
	Ebook *FindComputerManagementByMacAddressPatchFilter200ApplicationJSONEbooksEbook `json:"ebook,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONGeneral struct {
	ID         *int64  `json:"id,omitempty"`
	MacAddress *string `json:"mac_address,omitempty"`
	// Name of the computer
	Name         *string `json:"name,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Udid         *string `json:"udid,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementByMacAddressPatchFilter200ApplicationJSONMacAppStoreAppsMacAppStoreApp `json:"mac_app_store_app,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONManagedPreferenceProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONManagedPreferenceProfiles struct {
	Profile *FindComputerManagementByMacAddressPatchFilter200ApplicationJSONManagedPreferenceProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONOsXConfigurationProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONOsXConfigurationProfiles struct {
	Profile *FindComputerManagementByMacAddressPatchFilter200ApplicationJSONOsXConfigurationProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPatchPoliciesPatchPolicy struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPatchPolicies struct {
	PatchPolicy *FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPatchPoliciesPatchPolicy `json:"patch_policy,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string `json:"installed_version,omitempty"`
	LatestVersion    *string `json:"latest_version,omitempty"`
	Name             *string `json:"name,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPatchReportingSoftwareTitlesTitle `json:"title,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPoliciesPolicy struct {
	ID      *int64  `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Trigger *string `json:"trigger,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPolicies struct {
	Policy *FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPoliciesPolicy `json:"policy,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONRestrictedSoftwareSoftware struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONRestrictedSoftware struct {
	Software *FindComputerManagementByMacAddressPatchFilter200ApplicationJSONRestrictedSoftwareSoftware `json:"software,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONSmartGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONSmartGroups struct {
	Group *FindComputerManagementByMacAddressPatchFilter200ApplicationJSONSmartGroupsGroup `json:"group,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONStaticGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilter200ApplicationJSONStaticGroups struct {
	Group *FindComputerManagementByMacAddressPatchFilter200ApplicationJSONStaticGroupsGroup `json:"group,omitempty"`
}

// FindComputerManagementByMacAddressPatchFilter200ApplicationJSON - OK
type FindComputerManagementByMacAddressPatchFilter200ApplicationJSON struct {
	Ebooks                       []FindComputerManagementByMacAddressPatchFilter200ApplicationJSONEbooks                       `json:"ebooks,omitempty"`
	General                      *FindComputerManagementByMacAddressPatchFilter200ApplicationJSONGeneral                       `json:"general,omitempty"`
	MacAppStoreApps              []FindComputerManagementByMacAddressPatchFilter200ApplicationJSONMacAppStoreApps              `json:"mac_app_store_apps,omitempty"`
	ManagedPreferenceProfiles    []FindComputerManagementByMacAddressPatchFilter200ApplicationJSONManagedPreferenceProfiles    `json:"managed_preference_profiles,omitempty"`
	OsXConfigurationProfiles     []FindComputerManagementByMacAddressPatchFilter200ApplicationJSONOsXConfigurationProfiles     `json:"os_x_configuration_profiles,omitempty"`
	PatchPolicies                []FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPatchPolicies                `json:"patch_policies,omitempty"`
	PatchReportingSoftwareTitles []FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPatchReportingSoftwareTitles `json:"patch_reporting_software_titles,omitempty"`
	Policies                     []FindComputerManagementByMacAddressPatchFilter200ApplicationJSONPolicies                     `json:"policies,omitempty"`
	RestrictedSoftware           []FindComputerManagementByMacAddressPatchFilter200ApplicationJSONRestrictedSoftware           `json:"restricted_software,omitempty"`
	SmartGroups                  []FindComputerManagementByMacAddressPatchFilter200ApplicationJSONSmartGroups                  `json:"smart_groups,omitempty"`
	StaticGroups                 []FindComputerManagementByMacAddressPatchFilter200ApplicationJSONStaticGroups                 `json:"static_groups,omitempty"`
}

type FindComputerManagementByMacAddressPatchFilterResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindComputerManagementByMacAddressPatchFilter200ApplicationJSONObject *FindComputerManagementByMacAddressPatchFilter200ApplicationJSON
}
