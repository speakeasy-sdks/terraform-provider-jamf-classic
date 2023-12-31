// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FindMobileDeviceApplicationsByBundleIDRequest struct {
	// Bundle ID to filter by
	Bundleid string `pathParam:"style=simple,explode=false,name=bundleid"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLAppConfiguration struct {
	Preferences *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralCategory struct {
	ID *int64
	// Name of the category
	Name string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralDeploymentType string

const (
	FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralDeploymentTypeMakeAvailableInSelfService               FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralDeploymentType = "Make Available in Self Service"
	FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralDeploymentTypeInstallAutomaticallyPromptUsersToInstall FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralDeploymentType = "Install Automatically/Prompt Users to Install"
)

func (e FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralDeploymentType) ToPointer() *FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralDeploymentType {
	return &e
}

func (e *FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralDeploymentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Make Available in Self Service":
		fallthrough
	case "Install Automatically/Prompt Users to Install":
		*e = FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralDeploymentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralDeploymentType: %v", v)
	}
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralIcon struct {
	// base64 encoded
	Data *string
	ID   *int64
	Name *string
	URI  *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralIpa struct {
	Data *string
	Name *string
	URI  *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralSite struct {
	ID *int64
	// Name of the site
	Name string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneral struct {
	BundleID                         string
	Category                         *FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralCategory
	DeployAsManagedApp               *bool
	DeployAutomatically              *bool
	DeploymentType                   *FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralDeploymentType
	Description                      *string
	DisplayName                      *string
	ExternalURL                      *string
	Free                             *bool
	HostExternally                   *bool
	Icon                             *FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralIcon
	ID                               *int64
	InternalApp                      *bool
	Ipa                              *FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralIpa
	ItunesCountryRegion              *string
	ItunesStoreURL                   *string
	ItunesSyncTime                   *int64
	KeepDescriptionAndIconUpToDate   *bool
	MakeAvailableAfterInstall        *bool
	MobileDeviceProvisioningProfile  *int64
	Name                             string
	PreventBackupOfAppData           *bool
	RemoveAppWhenMdmProfileIsRemoved *bool
	Site                             *FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneralSite
	TakeOverManagement               *bool
	Version                          string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeBuildingsBuilding struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeBuildings struct {
	Building *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeBuildingsBuilding
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeDepartmentsDepartment struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeDepartments struct {
	Department *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeDepartmentsDepartment
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsBuildingsBuilding struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsBuildings struct {
	Building *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsBuildingsBuilding
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsDepartmentsDepartment struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsDepartments struct {
	Department *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsDepartmentsDepartment
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsJssUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsJssUserGroups struct {
	UserGroup *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsJssUserGroupsUserGroup
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsJssUsersUser struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsJssUsers struct {
	User *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsJssUsersUser
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsMobileDeviceGroupsMobileDeviceGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsMobileDeviceGroups struct {
	MobileDeviceGroup *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsMobileDeviceGroupsMobileDeviceGroup
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsMobileDevicesMobileDevice struct {
	ID *int64
	// Name of the device
	Name           *string
	Udid           *string
	WifiMacAddress *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsMobileDevices struct {
	MobileDevice *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsMobileDevicesMobileDevice
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsNetworkSegmentsNetworkSegment struct {
	ID *int64
	// Name of the network segment
	Name *string
	UID  *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsNetworkSegments struct {
	NetworkSegment *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsNetworkSegmentsNetworkSegment
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsUserGroups struct {
	UserGroup *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsUserGroupsUserGroup
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsUsersUser struct {
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsUsers struct {
	User *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsUsersUser
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusions struct {
	Buildings          []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsBuildings
	Departments        []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsDepartments
	JssUserGroups      []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsJssUserGroups
	JssUsers           []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsJssUsers
	MobileDeviceGroups []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsMobileDeviceGroups
	MobileDevices      []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsMobileDevices
	NetworkSegments    []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsNetworkSegments
	UserGroups         []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsUserGroups
	Users              []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusionsUsers
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeJssUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeJssUserGroups struct {
	UserGroup *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeJssUserGroupsUserGroup
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeJssUsersUser struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeJssUsers struct {
	User *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeJssUsersUser
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsNetworkSegmentsNetworkSegment struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsNetworkSegments struct {
	NetworkSegment *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsNetworkSegmentsNetworkSegment
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsUserGroups struct {
	UserGroup *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsUserGroupsUserGroup
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsUsersUser struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsUsers struct {
	User *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsUsersUser
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitations struct {
	NetworkSegments []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsNetworkSegments
	UserGroups      []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsUserGroups
	Users           []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitationsUsers
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeMobileDeviceGroupsMobileDeviceGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeMobileDeviceGroups struct {
	MobileDeviceGroup *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeMobileDeviceGroupsMobileDeviceGroup
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeMobileDevicesMobileDevice struct {
	ID *int64
	// Name of the device
	Name           *string
	Udid           *string
	WifiMacAddress *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeMobileDevices struct {
	MobileDevice *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeMobileDevicesMobileDevice
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLScope struct {
	AllJssUsers        *bool
	AllMobileDevices   *bool
	Buildings          []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeBuildings
	Departments        []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeDepartments
	Exclusions         *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeExclusions
	JssUserGroups      []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeJssUserGroups
	JssUsers           []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeJssUsers
	Limitations        *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeLimitations
	MobileDeviceGroups []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeMobileDeviceGroups
	MobileDevices      []FindMobileDeviceApplicationsByBundleID200ApplicationXMLScopeMobileDevices
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLSelfServiceSelfServiceCategoriesCategory struct {
	DisplayIn *bool
	ID        *int64
	Name      *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLSelfServiceSelfServiceCategories struct {
	Category *FindMobileDeviceApplicationsByBundleID200ApplicationXMLSelfServiceSelfServiceCategoriesCategory
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLSelfServiceSelfServiceIcon struct {
	// base64 encoded
	Data *string
	ID   *int64
	Name *string
	URI  *string
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLSelfService struct {
	FeatureOnMainPage      *bool
	Notification           *bool
	NotificationMessage    *string
	NotificationSubject    *string
	SelfServiceCategories  []FindMobileDeviceApplicationsByBundleID200ApplicationXMLSelfServiceSelfServiceCategories
	SelfServiceDescription *string
	SelfServiceIcon        *FindMobileDeviceApplicationsByBundleID200ApplicationXMLSelfServiceSelfServiceIcon
}

type FindMobileDeviceApplicationsByBundleID200ApplicationXMLVpp struct {
	AssignVppDeviceBasedLicenses *bool
	VppAdminAccountID            *int64
}

// FindMobileDeviceApplicationsByBundleID200ApplicationXML - OK
type FindMobileDeviceApplicationsByBundleID200ApplicationXML struct {
	AppConfiguration *FindMobileDeviceApplicationsByBundleID200ApplicationXMLAppConfiguration
	General          *FindMobileDeviceApplicationsByBundleID200ApplicationXMLGeneral
	Scope            *FindMobileDeviceApplicationsByBundleID200ApplicationXMLScope
	SelfService      *FindMobileDeviceApplicationsByBundleID200ApplicationXMLSelfService
	Vpp              *FindMobileDeviceApplicationsByBundleID200ApplicationXMLVpp
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONAppConfiguration struct {
	Preferences *string `json:"preferences,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralCategory struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the category
	Name string `json:"name"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralDeploymentType string

const (
	FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralDeploymentTypeMakeAvailableInSelfService               FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralDeploymentType = "Make Available in Self Service"
	FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralDeploymentTypeInstallAutomaticallyPromptUsersToInstall FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralDeploymentType = "Install Automatically/Prompt Users to Install"
)

func (e FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralDeploymentType) ToPointer() *FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralDeploymentType {
	return &e
}

func (e *FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralDeploymentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Make Available in Self Service":
		fallthrough
	case "Install Automatically/Prompt Users to Install":
		*e = FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralDeploymentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralDeploymentType: %v", v)
	}
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralIcon struct {
	// base64 encoded
	Data *string `json:"data,omitempty"`
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	URI  *string `json:"uri,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralIpa struct {
	Data *string `json:"data,omitempty"`
	Name *string `json:"name,omitempty"`
	URI  *string `json:"uri,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralSite struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the site
	Name string `json:"name"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneral struct {
	BundleID                         string                                                                         `json:"bundle_id"`
	Category                         *FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralCategory       `json:"category,omitempty"`
	DeployAsManagedApp               *bool                                                                          `json:"deploy_as_managed_app,omitempty"`
	DeployAutomatically              *bool                                                                          `json:"deploy_automatically,omitempty"`
	DeploymentType                   *FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralDeploymentType `json:"deployment_type,omitempty"`
	Description                      *string                                                                        `json:"description,omitempty"`
	DisplayName                      *string                                                                        `json:"display_name,omitempty"`
	ExternalURL                      *string                                                                        `json:"external_url,omitempty"`
	Free                             *bool                                                                          `json:"free,omitempty"`
	HostExternally                   *bool                                                                          `json:"host_externally,omitempty"`
	Icon                             *FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralIcon           `json:"icon,omitempty"`
	ID                               *int64                                                                         `json:"id,omitempty"`
	InternalApp                      *bool                                                                          `json:"internal_app,omitempty"`
	Ipa                              *FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralIpa            `json:"ipa,omitempty"`
	ItunesCountryRegion              *string                                                                        `json:"itunes_country_region,omitempty"`
	ItunesStoreURL                   *string                                                                        `json:"itunes_store_url,omitempty"`
	ItunesSyncTime                   *int64                                                                         `json:"itunes_sync_time,omitempty"`
	KeepDescriptionAndIconUpToDate   *bool                                                                          `json:"keep_description_and_icon_up_to_date,omitempty"`
	MakeAvailableAfterInstall        *bool                                                                          `json:"make_available_after_install,omitempty"`
	MobileDeviceProvisioningProfile  *int64                                                                         `json:"mobile_device_provisioning_profile,omitempty"`
	Name                             string                                                                         `json:"name"`
	PreventBackupOfAppData           *bool                                                                          `json:"prevent_backup_of_app_data,omitempty"`
	RemoveAppWhenMdmProfileIsRemoved *bool                                                                          `json:"remove_app_when_mdm_profile_is_removed,omitempty"`
	Site                             *FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneralSite           `json:"site,omitempty"`
	TakeOverManagement               *bool                                                                          `json:"take_over_management,omitempty"`
	Version                          string                                                                         `json:"version"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeBuildingsBuilding struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeBuildings struct {
	Building *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeBuildingsBuilding `json:"building,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeDepartmentsDepartment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeDepartments struct {
	Department *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeDepartmentsDepartment `json:"department,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsBuildingsBuilding struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsBuildings struct {
	Building *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsBuildingsBuilding `json:"building,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsDepartmentsDepartment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsDepartments struct {
	Department *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsDepartmentsDepartment `json:"department,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsJssUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsJssUserGroups struct {
	UserGroup *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsJssUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsJssUsersUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsJssUsers struct {
	User *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsJssUsersUser `json:"user,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsMobileDeviceGroupsMobileDeviceGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsMobileDeviceGroups struct {
	MobileDeviceGroup *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsMobileDeviceGroupsMobileDeviceGroup `json:"mobile_device_group,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsMobileDevicesMobileDevice struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the device
	Name           *string `json:"name,omitempty"`
	Udid           *string `json:"udid,omitempty"`
	WifiMacAddress *string `json:"wifi_mac_address,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsMobileDevices struct {
	MobileDevice *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsMobileDevicesMobileDevice `json:"mobile_device,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsNetworkSegmentsNetworkSegment struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the network segment
	Name *string `json:"name,omitempty"`
	UID  *string `json:"uid,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsNetworkSegments struct {
	NetworkSegment *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsNetworkSegmentsNetworkSegment `json:"network_segment,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsUserGroups struct {
	UserGroup *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsUsersUser struct {
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsUsers struct {
	User *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsUsersUser `json:"user,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusions struct {
	Buildings          []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsBuildings          `json:"buildings,omitempty"`
	Departments        []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsDepartments        `json:"departments,omitempty"`
	JssUserGroups      []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsJssUserGroups      `json:"jss_user_groups,omitempty"`
	JssUsers           []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsJssUsers           `json:"jss_users,omitempty"`
	MobileDeviceGroups []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsMobileDeviceGroups `json:"mobile_device_groups,omitempty"`
	MobileDevices      []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsMobileDevices      `json:"mobile_devices,omitempty"`
	NetworkSegments    []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsNetworkSegments    `json:"network_segments,omitempty"`
	UserGroups         []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsUserGroups         `json:"user_groups,omitempty"`
	Users              []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusionsUsers              `json:"users,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeJssUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeJssUserGroups struct {
	UserGroup *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeJssUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeJssUsersUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeJssUsers struct {
	User *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeJssUsersUser `json:"user,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsNetworkSegmentsNetworkSegment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsNetworkSegments struct {
	NetworkSegment *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsNetworkSegmentsNetworkSegment `json:"network_segment,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsUserGroups struct {
	UserGroup *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsUsersUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsUsers struct {
	User *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsUsersUser `json:"user,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitations struct {
	NetworkSegments []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsNetworkSegments `json:"network_segments,omitempty"`
	UserGroups      []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsUserGroups      `json:"user_groups,omitempty"`
	Users           []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitationsUsers           `json:"users,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeMobileDeviceGroupsMobileDeviceGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeMobileDeviceGroups struct {
	MobileDeviceGroup *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeMobileDeviceGroupsMobileDeviceGroup `json:"mobile_device_group,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeMobileDevicesMobileDevice struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the device
	Name           *string `json:"name,omitempty"`
	Udid           *string `json:"udid,omitempty"`
	WifiMacAddress *string `json:"wifi_mac_address,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeMobileDevices struct {
	MobileDevice *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeMobileDevicesMobileDevice `json:"mobile_device,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONScope struct {
	AllJssUsers        *bool                                                                             `json:"all_jss_users,omitempty"`
	AllMobileDevices   *bool                                                                             `json:"all_mobile_devices,omitempty"`
	Buildings          []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeBuildings          `json:"buildings,omitempty"`
	Departments        []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeDepartments        `json:"departments,omitempty"`
	Exclusions         *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeExclusions          `json:"exclusions,omitempty"`
	JssUserGroups      []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeJssUserGroups      `json:"jss_user_groups,omitempty"`
	JssUsers           []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeJssUsers           `json:"jss_users,omitempty"`
	Limitations        *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeLimitations         `json:"limitations,omitempty"`
	MobileDeviceGroups []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeMobileDeviceGroups `json:"mobile_device_groups,omitempty"`
	MobileDevices      []FindMobileDeviceApplicationsByBundleID200ApplicationJSONScopeMobileDevices      `json:"mobile_devices,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONSelfServiceSelfServiceCategoriesCategory struct {
	DisplayIn *bool   `json:"display_in,omitempty"`
	ID        *int64  `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONSelfServiceSelfServiceCategories struct {
	Category *FindMobileDeviceApplicationsByBundleID200ApplicationJSONSelfServiceSelfServiceCategoriesCategory `json:"category,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONSelfServiceSelfServiceIcon struct {
	// base64 encoded
	Data *string `json:"data,omitempty"`
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	URI  *string `json:"uri,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONSelfService struct {
	FeatureOnMainPage      *bool                                                                                      `json:"feature_on_main_page,omitempty"`
	Notification           *bool                                                                                      `json:"notification,omitempty"`
	NotificationMessage    *string                                                                                    `json:"notification_message,omitempty"`
	NotificationSubject    *string                                                                                    `json:"notification_subject,omitempty"`
	SelfServiceCategories  []FindMobileDeviceApplicationsByBundleID200ApplicationJSONSelfServiceSelfServiceCategories `json:"self_service_categories,omitempty"`
	SelfServiceDescription *string                                                                                    `json:"self_service_description,omitempty"`
	SelfServiceIcon        *FindMobileDeviceApplicationsByBundleID200ApplicationJSONSelfServiceSelfServiceIcon        `json:"self_service_icon,omitempty"`
}

type FindMobileDeviceApplicationsByBundleID200ApplicationJSONVpp struct {
	AssignVppDeviceBasedLicenses *bool  `json:"assign_vpp_device_based_licenses,omitempty"`
	VppAdminAccountID            *int64 `json:"vpp_admin_account_id,omitempty"`
}

// FindMobileDeviceApplicationsByBundleID200ApplicationJSON - OK
type FindMobileDeviceApplicationsByBundleID200ApplicationJSON struct {
	AppConfiguration *FindMobileDeviceApplicationsByBundleID200ApplicationJSONAppConfiguration `json:"app_configuration,omitempty"`
	General          *FindMobileDeviceApplicationsByBundleID200ApplicationJSONGeneral          `json:"general,omitempty"`
	Scope            *FindMobileDeviceApplicationsByBundleID200ApplicationJSONScope            `json:"scope,omitempty"`
	SelfService      *FindMobileDeviceApplicationsByBundleID200ApplicationJSONSelfService      `json:"self_service,omitempty"`
	Vpp              *FindMobileDeviceApplicationsByBundleID200ApplicationJSONVpp              `json:"vpp,omitempty"`
}

type FindMobileDeviceApplicationsByBundleIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDeviceApplicationsByBundleID200ApplicationJSONObject *FindMobileDeviceApplicationsByBundleID200ApplicationJSON
}
