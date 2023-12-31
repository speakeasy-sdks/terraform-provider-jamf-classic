// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindComputerManagementByUdidUsernameRequest struct {
	// Computer UDID to filter by
	Udid string `pathParam:"style=simple,explode=false,name=udid"`
	// Username to filter by
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

type FindComputerManagementByUdidUsername200ApplicationXMLEbooksEbook struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByUdidUsername200ApplicationXMLEbooks struct {
	Ebook *FindComputerManagementByUdidUsername200ApplicationXMLEbooksEbook
}

type FindComputerManagementByUdidUsername200ApplicationXMLGeneral struct {
	ID         *int64
	MacAddress *string
	// Name of the computer
	Name         *string
	SerialNumber *string
	Udid         *string
}

type FindComputerManagementByUdidUsername200ApplicationXMLMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByUdidUsername200ApplicationXMLMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementByUdidUsername200ApplicationXMLMacAppStoreAppsMacAppStoreApp
}

type FindComputerManagementByUdidUsername200ApplicationXMLManagedPreferenceProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByUdidUsername200ApplicationXMLManagedPreferenceProfiles struct {
	Profile *FindComputerManagementByUdidUsername200ApplicationXMLManagedPreferenceProfilesProfile
}

type FindComputerManagementByUdidUsername200ApplicationXMLOsXConfigurationProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByUdidUsername200ApplicationXMLOsXConfigurationProfiles struct {
	Profile *FindComputerManagementByUdidUsername200ApplicationXMLOsXConfigurationProfilesProfile
}

type FindComputerManagementByUdidUsername200ApplicationXMLPatchPoliciesPatchPolicy struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByUdidUsername200ApplicationXMLPatchPolicies struct {
	PatchPolicy *FindComputerManagementByUdidUsername200ApplicationXMLPatchPoliciesPatchPolicy
}

type FindComputerManagementByUdidUsername200ApplicationXMLPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string
	LatestVersion    *string
	Name             *string
}

type FindComputerManagementByUdidUsername200ApplicationXMLPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementByUdidUsername200ApplicationXMLPatchReportingSoftwareTitlesTitle
}

type FindComputerManagementByUdidUsername200ApplicationXMLPoliciesPolicy struct {
	ID      *int64
	Name    *string
	Trigger *string
}

type FindComputerManagementByUdidUsername200ApplicationXMLPolicies struct {
	Policy *FindComputerManagementByUdidUsername200ApplicationXMLPoliciesPolicy
}

type FindComputerManagementByUdidUsername200ApplicationXMLRestrictedSoftwareSoftware struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByUdidUsername200ApplicationXMLRestrictedSoftware struct {
	Software *FindComputerManagementByUdidUsername200ApplicationXMLRestrictedSoftwareSoftware
}

type FindComputerManagementByUdidUsername200ApplicationXMLSmartGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByUdidUsername200ApplicationXMLSmartGroups struct {
	Group *FindComputerManagementByUdidUsername200ApplicationXMLSmartGroupsGroup
}

type FindComputerManagementByUdidUsername200ApplicationXMLStaticGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByUdidUsername200ApplicationXMLStaticGroups struct {
	Group *FindComputerManagementByUdidUsername200ApplicationXMLStaticGroupsGroup
}

// FindComputerManagementByUdidUsername200ApplicationXML - OK
type FindComputerManagementByUdidUsername200ApplicationXML struct {
	Ebooks                       []FindComputerManagementByUdidUsername200ApplicationXMLEbooks
	General                      *FindComputerManagementByUdidUsername200ApplicationXMLGeneral
	MacAppStoreApps              []FindComputerManagementByUdidUsername200ApplicationXMLMacAppStoreApps
	ManagedPreferenceProfiles    []FindComputerManagementByUdidUsername200ApplicationXMLManagedPreferenceProfiles
	OsXConfigurationProfiles     []FindComputerManagementByUdidUsername200ApplicationXMLOsXConfigurationProfiles
	PatchPolicies                []FindComputerManagementByUdidUsername200ApplicationXMLPatchPolicies
	PatchReportingSoftwareTitles []FindComputerManagementByUdidUsername200ApplicationXMLPatchReportingSoftwareTitles
	Policies                     []FindComputerManagementByUdidUsername200ApplicationXMLPolicies
	RestrictedSoftware           []FindComputerManagementByUdidUsername200ApplicationXMLRestrictedSoftware
	SmartGroups                  []FindComputerManagementByUdidUsername200ApplicationXMLSmartGroups
	StaticGroups                 []FindComputerManagementByUdidUsername200ApplicationXMLStaticGroups
}

type FindComputerManagementByUdidUsername200ApplicationJSONEbooksEbook struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONEbooks struct {
	Ebook *FindComputerManagementByUdidUsername200ApplicationJSONEbooksEbook `json:"ebook,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONGeneral struct {
	ID         *int64  `json:"id,omitempty"`
	MacAddress *string `json:"mac_address,omitempty"`
	// Name of the computer
	Name         *string `json:"name,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Udid         *string `json:"udid,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementByUdidUsername200ApplicationJSONMacAppStoreAppsMacAppStoreApp `json:"mac_app_store_app,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONManagedPreferenceProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONManagedPreferenceProfiles struct {
	Profile *FindComputerManagementByUdidUsername200ApplicationJSONManagedPreferenceProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONOsXConfigurationProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONOsXConfigurationProfiles struct {
	Profile *FindComputerManagementByUdidUsername200ApplicationJSONOsXConfigurationProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONPatchPoliciesPatchPolicy struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONPatchPolicies struct {
	PatchPolicy *FindComputerManagementByUdidUsername200ApplicationJSONPatchPoliciesPatchPolicy `json:"patch_policy,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string `json:"installed_version,omitempty"`
	LatestVersion    *string `json:"latest_version,omitempty"`
	Name             *string `json:"name,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementByUdidUsername200ApplicationJSONPatchReportingSoftwareTitlesTitle `json:"title,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONPoliciesPolicy struct {
	ID      *int64  `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Trigger *string `json:"trigger,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONPolicies struct {
	Policy *FindComputerManagementByUdidUsername200ApplicationJSONPoliciesPolicy `json:"policy,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONRestrictedSoftwareSoftware struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONRestrictedSoftware struct {
	Software *FindComputerManagementByUdidUsername200ApplicationJSONRestrictedSoftwareSoftware `json:"software,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONSmartGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONSmartGroups struct {
	Group *FindComputerManagementByUdidUsername200ApplicationJSONSmartGroupsGroup `json:"group,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONStaticGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByUdidUsername200ApplicationJSONStaticGroups struct {
	Group *FindComputerManagementByUdidUsername200ApplicationJSONStaticGroupsGroup `json:"group,omitempty"`
}

// FindComputerManagementByUdidUsername200ApplicationJSON - OK
type FindComputerManagementByUdidUsername200ApplicationJSON struct {
	Ebooks                       []FindComputerManagementByUdidUsername200ApplicationJSONEbooks                       `json:"ebooks,omitempty"`
	General                      *FindComputerManagementByUdidUsername200ApplicationJSONGeneral                       `json:"general,omitempty"`
	MacAppStoreApps              []FindComputerManagementByUdidUsername200ApplicationJSONMacAppStoreApps              `json:"mac_app_store_apps,omitempty"`
	ManagedPreferenceProfiles    []FindComputerManagementByUdidUsername200ApplicationJSONManagedPreferenceProfiles    `json:"managed_preference_profiles,omitempty"`
	OsXConfigurationProfiles     []FindComputerManagementByUdidUsername200ApplicationJSONOsXConfigurationProfiles     `json:"os_x_configuration_profiles,omitempty"`
	PatchPolicies                []FindComputerManagementByUdidUsername200ApplicationJSONPatchPolicies                `json:"patch_policies,omitempty"`
	PatchReportingSoftwareTitles []FindComputerManagementByUdidUsername200ApplicationJSONPatchReportingSoftwareTitles `json:"patch_reporting_software_titles,omitempty"`
	Policies                     []FindComputerManagementByUdidUsername200ApplicationJSONPolicies                     `json:"policies,omitempty"`
	RestrictedSoftware           []FindComputerManagementByUdidUsername200ApplicationJSONRestrictedSoftware           `json:"restricted_software,omitempty"`
	SmartGroups                  []FindComputerManagementByUdidUsername200ApplicationJSONSmartGroups                  `json:"smart_groups,omitempty"`
	StaticGroups                 []FindComputerManagementByUdidUsername200ApplicationJSONStaticGroups                 `json:"static_groups,omitempty"`
}

type FindComputerManagementByUdidUsernameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindComputerManagementByUdidUsername200ApplicationJSONObject *FindComputerManagementByUdidUsername200ApplicationJSON
}
