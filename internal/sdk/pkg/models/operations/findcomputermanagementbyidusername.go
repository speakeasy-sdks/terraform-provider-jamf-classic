// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindComputerManagementByIDUsernameRequest struct {
	// Computer ID to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
	// Username to filter by
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

type FindComputerManagementByIDUsername200ApplicationXMLEbooksEbook struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDUsername200ApplicationXMLEbooks struct {
	Ebook *FindComputerManagementByIDUsername200ApplicationXMLEbooksEbook
}

type FindComputerManagementByIDUsername200ApplicationXMLGeneral struct {
	ID         *int64
	MacAddress *string
	// Name of the computer
	Name         *string
	SerialNumber *string
	Udid         *string
}

type FindComputerManagementByIDUsername200ApplicationXMLMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDUsername200ApplicationXMLMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementByIDUsername200ApplicationXMLMacAppStoreAppsMacAppStoreApp
}

type FindComputerManagementByIDUsername200ApplicationXMLManagedPreferenceProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDUsername200ApplicationXMLManagedPreferenceProfiles struct {
	Profile *FindComputerManagementByIDUsername200ApplicationXMLManagedPreferenceProfilesProfile
}

type FindComputerManagementByIDUsername200ApplicationXMLOsXConfigurationProfilesProfile struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDUsername200ApplicationXMLOsXConfigurationProfiles struct {
	Profile *FindComputerManagementByIDUsername200ApplicationXMLOsXConfigurationProfilesProfile
}

type FindComputerManagementByIDUsername200ApplicationXMLPatchPoliciesPatchPolicy struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDUsername200ApplicationXMLPatchPolicies struct {
	PatchPolicy *FindComputerManagementByIDUsername200ApplicationXMLPatchPoliciesPatchPolicy
}

type FindComputerManagementByIDUsername200ApplicationXMLPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string
	LatestVersion    *string
	Name             *string
}

type FindComputerManagementByIDUsername200ApplicationXMLPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementByIDUsername200ApplicationXMLPatchReportingSoftwareTitlesTitle
}

type FindComputerManagementByIDUsername200ApplicationXMLPoliciesPolicy struct {
	ID      *int64
	Name    *string
	Trigger *string
}

type FindComputerManagementByIDUsername200ApplicationXMLPolicies struct {
	Policy *FindComputerManagementByIDUsername200ApplicationXMLPoliciesPolicy
}

type FindComputerManagementByIDUsername200ApplicationXMLRestrictedSoftwareSoftware struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDUsername200ApplicationXMLRestrictedSoftware struct {
	Software *FindComputerManagementByIDUsername200ApplicationXMLRestrictedSoftwareSoftware
}

type FindComputerManagementByIDUsername200ApplicationXMLSmartGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDUsername200ApplicationXMLSmartGroups struct {
	Group *FindComputerManagementByIDUsername200ApplicationXMLSmartGroupsGroup
}

type FindComputerManagementByIDUsername200ApplicationXMLStaticGroupsGroup struct {
	ID   *int64
	Name *string
}

type FindComputerManagementByIDUsername200ApplicationXMLStaticGroups struct {
	Group *FindComputerManagementByIDUsername200ApplicationXMLStaticGroupsGroup
}

// FindComputerManagementByIDUsername200ApplicationXML - OK
type FindComputerManagementByIDUsername200ApplicationXML struct {
	Ebooks                       []FindComputerManagementByIDUsername200ApplicationXMLEbooks
	General                      *FindComputerManagementByIDUsername200ApplicationXMLGeneral
	MacAppStoreApps              []FindComputerManagementByIDUsername200ApplicationXMLMacAppStoreApps
	ManagedPreferenceProfiles    []FindComputerManagementByIDUsername200ApplicationXMLManagedPreferenceProfiles
	OsXConfigurationProfiles     []FindComputerManagementByIDUsername200ApplicationXMLOsXConfigurationProfiles
	PatchPolicies                []FindComputerManagementByIDUsername200ApplicationXMLPatchPolicies
	PatchReportingSoftwareTitles []FindComputerManagementByIDUsername200ApplicationXMLPatchReportingSoftwareTitles
	Policies                     []FindComputerManagementByIDUsername200ApplicationXMLPolicies
	RestrictedSoftware           []FindComputerManagementByIDUsername200ApplicationXMLRestrictedSoftware
	SmartGroups                  []FindComputerManagementByIDUsername200ApplicationXMLSmartGroups
	StaticGroups                 []FindComputerManagementByIDUsername200ApplicationXMLStaticGroups
}

type FindComputerManagementByIDUsername200ApplicationJSONEbooksEbook struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONEbooks struct {
	Ebook *FindComputerManagementByIDUsername200ApplicationJSONEbooksEbook `json:"ebook,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONGeneral struct {
	ID         *int64  `json:"id,omitempty"`
	MacAddress *string `json:"mac_address,omitempty"`
	// Name of the computer
	Name         *string `json:"name,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty"`
	Udid         *string `json:"udid,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONMacAppStoreAppsMacAppStoreApp struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONMacAppStoreApps struct {
	MacAppStoreApp *FindComputerManagementByIDUsername200ApplicationJSONMacAppStoreAppsMacAppStoreApp `json:"mac_app_store_app,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONManagedPreferenceProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONManagedPreferenceProfiles struct {
	Profile *FindComputerManagementByIDUsername200ApplicationJSONManagedPreferenceProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONOsXConfigurationProfilesProfile struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONOsXConfigurationProfiles struct {
	Profile *FindComputerManagementByIDUsername200ApplicationJSONOsXConfigurationProfilesProfile `json:"profile,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONPatchPoliciesPatchPolicy struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONPatchPolicies struct {
	PatchPolicy *FindComputerManagementByIDUsername200ApplicationJSONPatchPoliciesPatchPolicy `json:"patch_policy,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONPatchReportingSoftwareTitlesTitle struct {
	InstalledVersion *string `json:"installed_version,omitempty"`
	LatestVersion    *string `json:"latest_version,omitempty"`
	Name             *string `json:"name,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONPatchReportingSoftwareTitles struct {
	Title *FindComputerManagementByIDUsername200ApplicationJSONPatchReportingSoftwareTitlesTitle `json:"title,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONPoliciesPolicy struct {
	ID      *int64  `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Trigger *string `json:"trigger,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONPolicies struct {
	Policy *FindComputerManagementByIDUsername200ApplicationJSONPoliciesPolicy `json:"policy,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONRestrictedSoftwareSoftware struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONRestrictedSoftware struct {
	Software *FindComputerManagementByIDUsername200ApplicationJSONRestrictedSoftwareSoftware `json:"software,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONSmartGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONSmartGroups struct {
	Group *FindComputerManagementByIDUsername200ApplicationJSONSmartGroupsGroup `json:"group,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONStaticGroupsGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindComputerManagementByIDUsername200ApplicationJSONStaticGroups struct {
	Group *FindComputerManagementByIDUsername200ApplicationJSONStaticGroupsGroup `json:"group,omitempty"`
}

// FindComputerManagementByIDUsername200ApplicationJSON - OK
type FindComputerManagementByIDUsername200ApplicationJSON struct {
	Ebooks                       []FindComputerManagementByIDUsername200ApplicationJSONEbooks                       `json:"ebooks,omitempty"`
	General                      *FindComputerManagementByIDUsername200ApplicationJSONGeneral                       `json:"general,omitempty"`
	MacAppStoreApps              []FindComputerManagementByIDUsername200ApplicationJSONMacAppStoreApps              `json:"mac_app_store_apps,omitempty"`
	ManagedPreferenceProfiles    []FindComputerManagementByIDUsername200ApplicationJSONManagedPreferenceProfiles    `json:"managed_preference_profiles,omitempty"`
	OsXConfigurationProfiles     []FindComputerManagementByIDUsername200ApplicationJSONOsXConfigurationProfiles     `json:"os_x_configuration_profiles,omitempty"`
	PatchPolicies                []FindComputerManagementByIDUsername200ApplicationJSONPatchPolicies                `json:"patch_policies,omitempty"`
	PatchReportingSoftwareTitles []FindComputerManagementByIDUsername200ApplicationJSONPatchReportingSoftwareTitles `json:"patch_reporting_software_titles,omitempty"`
	Policies                     []FindComputerManagementByIDUsername200ApplicationJSONPolicies                     `json:"policies,omitempty"`
	RestrictedSoftware           []FindComputerManagementByIDUsername200ApplicationJSONRestrictedSoftware           `json:"restricted_software,omitempty"`
	SmartGroups                  []FindComputerManagementByIDUsername200ApplicationJSONSmartGroups                  `json:"smart_groups,omitempty"`
	StaticGroups                 []FindComputerManagementByIDUsername200ApplicationJSONStaticGroups                 `json:"static_groups,omitempty"`
}

type FindComputerManagementByIDUsernameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindComputerManagementByIDUsername200ApplicationJSONObject *FindComputerManagementByIDUsername200ApplicationJSON
}
