// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindMobileDeviceConfigurationProfilesByNameSubsetSubset - Subset to filter by
type FindMobileDeviceConfigurationProfilesByNameSubsetSubset string

const (
	FindMobileDeviceConfigurationProfilesByNameSubsetSubsetGeneral     FindMobileDeviceConfigurationProfilesByNameSubsetSubset = "General"
	FindMobileDeviceConfigurationProfilesByNameSubsetSubsetScope       FindMobileDeviceConfigurationProfilesByNameSubsetSubset = "Scope"
	FindMobileDeviceConfigurationProfilesByNameSubsetSubsetSelfService FindMobileDeviceConfigurationProfilesByNameSubsetSubset = "SelfService"
)

func (e FindMobileDeviceConfigurationProfilesByNameSubsetSubset) ToPointer() *FindMobileDeviceConfigurationProfilesByNameSubsetSubset {
	return &e
}

func (e *FindMobileDeviceConfigurationProfilesByNameSubsetSubset) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "General":
		fallthrough
	case "Scope":
		fallthrough
	case "SelfService":
		*e = FindMobileDeviceConfigurationProfilesByNameSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceConfigurationProfilesByNameSubsetSubset: %v", v)
	}
}

type FindMobileDeviceConfigurationProfilesByNameSubsetRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Subset to filter by
	Subset FindMobileDeviceConfigurationProfilesByNameSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralCategory struct {
	ID *int64
	// Name of the category
	Name string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralDeploymentMethod string

const (
	FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralDeploymentMethodInstallAutomatically       FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralDeploymentMethod = "Install Automatically"
	FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralDeploymentMethodMakeAvailableInSelfService FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralDeploymentMethod = "Make Available in Self Service"
)

func (e FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralDeploymentMethod) ToPointer() *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralDeploymentMethod {
	return &e
}

func (e *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralDeploymentMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Install Automatically":
		fallthrough
	case "Make Available in Self Service":
		*e = FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralDeploymentMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralDeploymentMethod: %v", v)
	}
}

// FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel - Level of the configuration profile
type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel string

const (
	FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevelSystem FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel = "System"
	FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevelUser   FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel = "User"
)

func (e FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel) ToPointer() *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel {
	return &e
}

func (e *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "System":
		fallthrough
	case "User":
		*e = FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel: %v", v)
	}
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralSite struct {
	ID *int64
	// Name of the site
	Name string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneral struct {
	Category         *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralCategory
	DeploymentMethod *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralDeploymentMethod
	Description      *string
	ID               *int64
	// Level of the configuration profile
	Level *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel
	// Name of the configuration profile
	Name                                  string
	Payloads                              *string
	RedeployDayssBeforeCertificateExpires *int64
	RedeployOnUpdate                      *string
	Site                                  *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneralSite
	UUID                                  *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeBuildingsBuilding struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeBuildings struct {
	Building *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeBuildingsBuilding
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeDepartmentsDepartment struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeDepartments struct {
	Department *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeDepartmentsDepartment
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsBuildingsBuilding struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsBuildings struct {
	Building *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsBuildingsBuilding
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsDepartmentsDepartment struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsDepartments struct {
	Department *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsDepartmentsDepartment
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsIbeaconsIbeacon struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsIbeacons struct {
	Ibeacon *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsIbeaconsIbeacon
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUserGroups struct {
	UserGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUserGroupsUserGroup
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUsersUser struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUsers struct {
	User *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUsersUser
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsMobileDeviceGroupsMobileDeviceGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsMobileDeviceGroups struct {
	MobileDeviceGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsMobileDeviceGroupsMobileDeviceGroup
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsMobileDevicesMobileDevice struct {
	ID *int64
	// Name of the device
	Name           *string
	Udid           *string
	WifiMacAddress *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsMobileDevices struct {
	MobileDevice *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsMobileDevicesMobileDevice
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsNetworkSegmentsNetworkSegment struct {
	ID *int64
	// Name of the network segment
	Name *string
	UID  *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsNetworkSegments struct {
	NetworkSegment *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsNetworkSegmentsNetworkSegment
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUserGroups struct {
	UserGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUserGroupsUserGroup
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUsersUser struct {
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUsers struct {
	User *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUsersUser
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusions struct {
	Buildings          []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsBuildings
	Departments        []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsDepartments
	Ibeacons           []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsIbeacons
	JssUserGroups      []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUserGroups
	JssUsers           []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUsers
	MobileDeviceGroups []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsMobileDeviceGroups
	MobileDevices      []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsMobileDevices
	NetworkSegments    []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsNetworkSegments
	UserGroups         []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUserGroups
	Users              []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUsers
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUserGroups struct {
	UserGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUserGroupsUserGroup
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUsersUser struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUsers struct {
	User *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUsersUser
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsIbeaconsIbeacon struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsIbeacons struct {
	Ibeacon *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsIbeaconsIbeacon
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsNetworkSegmentsNetworkSegment struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsNetworkSegments struct {
	NetworkSegment *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsNetworkSegmentsNetworkSegment
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUserGroups struct {
	UserGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUserGroupsUserGroup
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUsersUser struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUsers struct {
	User *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUsersUser
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitations struct {
	Ibeacons        []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsIbeacons
	NetworkSegments []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsNetworkSegments
	UserGroups      []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUserGroups
	Users           []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUsers
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeMobileDeviceGroupsMobileDeviceGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeMobileDeviceGroups struct {
	MobileDeviceGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeMobileDeviceGroupsMobileDeviceGroup
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeMobileDevicesMobileDevice struct {
	ID *int64
	// Name of the device
	Name           *string
	Udid           *string
	WifiMacAddress *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeMobileDevices struct {
	MobileDevice *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeMobileDevicesMobileDevice
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScope struct {
	AllJssUsers        *bool
	AllMobileDevices   *bool
	Buildings          []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeBuildings
	Departments        []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeDepartments
	Exclusions         *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusions
	JssUserGroups      []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUserGroups
	JssUsers           []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUsers
	Limitations        *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitations
	MobileDeviceGroups []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeMobileDeviceGroups
	MobileDevices      []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScopeMobileDevices
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSecurityName struct {
	RemovalDisallowed *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceCategoriesCategory struct {
	ID *int64
	// Name of the category
	Name     string
	Priority *int64
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceCategories struct {
	Category *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceCategoriesCategory
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceIcon struct {
	Data     *string
	Filename *string
	URI      *string
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLSelfService struct {
	FeatureOnMainPage      *bool
	SecurityName           *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSecurityName
	SelfServiceCategories  []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceCategories
	SelfServiceDescription *string
	SelfServiceIcon        *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceIcon
}

// FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXML - OK
type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXML struct {
	General     *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLGeneral
	Scope       *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLScope
	SelfService *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationXMLSelfService
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralCategory struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the category
	Name string `json:"name"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralDeploymentMethod string

const (
	FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralDeploymentMethodInstallAutomatically       FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralDeploymentMethod = "Install Automatically"
	FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralDeploymentMethodMakeAvailableInSelfService FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralDeploymentMethod = "Make Available in Self Service"
)

func (e FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralDeploymentMethod) ToPointer() *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralDeploymentMethod {
	return &e
}

func (e *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralDeploymentMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Install Automatically":
		fallthrough
	case "Make Available in Self Service":
		*e = FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralDeploymentMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralDeploymentMethod: %v", v)
	}
}

// FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel - Level of the configuration profile
type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel string

const (
	FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevelSystem FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel = "System"
	FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevelUser   FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel = "User"
)

func (e FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel) ToPointer() *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel {
	return &e
}

func (e *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "System":
		fallthrough
	case "User":
		*e = FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel: %v", v)
	}
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralSite struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the site
	Name string `json:"name"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneral struct {
	Category         *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralCategory         `json:"category,omitempty"`
	DeploymentMethod *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralDeploymentMethod `json:"deployment_method,omitempty"`
	Description      *string                                                                                     `json:"description,omitempty"`
	ID               *int64                                                                                      `json:"id,omitempty"`
	// Level of the configuration profile
	Level *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel `json:"level,omitempty"`
	// Name of the configuration profile
	Name                                  string                                                                          `json:"name"`
	Payloads                              *string                                                                         `json:"payloads,omitempty"`
	RedeployDayssBeforeCertificateExpires *int64                                                                          `json:"redeploy_Dayss_before_certificate_expires,omitempty"`
	RedeployOnUpdate                      *string                                                                         `json:"redeploy_on_update,omitempty"`
	Site                                  *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneralSite `json:"site,omitempty"`
	UUID                                  *string                                                                         `json:"uuid,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeBuildingsBuilding struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeBuildings struct {
	Building *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeBuildingsBuilding `json:"building,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeDepartmentsDepartment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeDepartments struct {
	Department *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeDepartmentsDepartment `json:"department,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsBuildingsBuilding struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsBuildings struct {
	Building *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsBuildingsBuilding `json:"building,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsDepartmentsDepartment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsDepartments struct {
	Department *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsDepartmentsDepartment `json:"department,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsIbeaconsIbeacon struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsIbeacons struct {
	Ibeacon *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsIbeaconsIbeacon `json:"ibeacon,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUserGroups struct {
	UserGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUsersUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUsers struct {
	User *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUsersUser `json:"user,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsMobileDeviceGroupsMobileDeviceGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsMobileDeviceGroups struct {
	MobileDeviceGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsMobileDeviceGroupsMobileDeviceGroup `json:"mobile_device_group,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsMobileDevicesMobileDevice struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the device
	Name           *string `json:"name,omitempty"`
	Udid           *string `json:"udid,omitempty"`
	WifiMacAddress *string `json:"wifi_mac_address,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsMobileDevices struct {
	MobileDevice *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsMobileDevicesMobileDevice `json:"mobile_device,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsNetworkSegmentsNetworkSegment struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the network segment
	Name *string `json:"name,omitempty"`
	UID  *string `json:"uid,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsNetworkSegments struct {
	NetworkSegment *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsNetworkSegmentsNetworkSegment `json:"network_segment,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUserGroups struct {
	UserGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUsersUser struct {
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUsers struct {
	User *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUsersUser `json:"user,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusions struct {
	Buildings          []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsBuildings          `json:"buildings,omitempty"`
	Departments        []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsDepartments        `json:"departments,omitempty"`
	Ibeacons           []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsIbeacons           `json:"ibeacons,omitempty"`
	JssUserGroups      []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUserGroups      `json:"jss_user_groups,omitempty"`
	JssUsers           []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUsers           `json:"jss_users,omitempty"`
	MobileDeviceGroups []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsMobileDeviceGroups `json:"mobile_device_groups,omitempty"`
	MobileDevices      []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsMobileDevices      `json:"mobile_devices,omitempty"`
	NetworkSegments    []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsNetworkSegments    `json:"network_segments,omitempty"`
	UserGroups         []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUserGroups         `json:"user_groups,omitempty"`
	Users              []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUsers              `json:"users,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUserGroups struct {
	UserGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUsersUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUsers struct {
	User *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUsersUser `json:"user,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsIbeaconsIbeacon struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsIbeacons struct {
	Ibeacon *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsIbeaconsIbeacon `json:"ibeacon,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsNetworkSegmentsNetworkSegment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsNetworkSegments struct {
	NetworkSegment *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsNetworkSegmentsNetworkSegment `json:"network_segment,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUserGroups struct {
	UserGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUsersUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUsers struct {
	User *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUsersUser `json:"user,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitations struct {
	Ibeacons        []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsIbeacons        `json:"ibeacons,omitempty"`
	NetworkSegments []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsNetworkSegments `json:"network_segments,omitempty"`
	UserGroups      []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUserGroups      `json:"user_groups,omitempty"`
	Users           []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUsers           `json:"users,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeMobileDeviceGroupsMobileDeviceGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeMobileDeviceGroups struct {
	MobileDeviceGroup *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeMobileDeviceGroupsMobileDeviceGroup `json:"mobile_device_group,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeMobileDevicesMobileDevice struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the device
	Name           *string `json:"name,omitempty"`
	Udid           *string `json:"udid,omitempty"`
	WifiMacAddress *string `json:"wifi_mac_address,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeMobileDevices struct {
	MobileDevice *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeMobileDevicesMobileDevice `json:"mobile_device,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScope struct {
	AllJssUsers        *bool                                                                                        `json:"all_jss_users,omitempty"`
	AllMobileDevices   *bool                                                                                        `json:"all_mobile_devices,omitempty"`
	Buildings          []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeBuildings          `json:"buildings,omitempty"`
	Departments        []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeDepartments        `json:"departments,omitempty"`
	Exclusions         *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusions          `json:"exclusions,omitempty"`
	JssUserGroups      []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUserGroups      `json:"jss_user_groups,omitempty"`
	JssUsers           []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUsers           `json:"jss_users,omitempty"`
	Limitations        *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitations         `json:"limitations,omitempty"`
	MobileDeviceGroups []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeMobileDeviceGroups `json:"mobile_device_groups,omitempty"`
	MobileDevices      []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScopeMobileDevices      `json:"mobile_devices,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSecurityName struct {
	RemovalDisallowed *string `json:"removal_disallowed,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceCategoriesCategory struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the category
	Name     string `json:"name"`
	Priority *int64 `json:"priority,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceCategories struct {
	Category *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceCategoriesCategory `json:"category,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceIcon struct {
	Data     *string `json:"data,omitempty"`
	Filename *string `json:"filename,omitempty"`
	URI      *string `json:"uri,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONSelfService struct {
	FeatureOnMainPage      *bool                                                                                                 `json:"feature_on_main_page,omitempty"`
	SecurityName           *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSecurityName           `json:"security_name,omitempty"`
	SelfServiceCategories  []FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceCategories `json:"self_service_categories,omitempty"`
	SelfServiceDescription *string                                                                                               `json:"self_service_description,omitempty"`
	SelfServiceIcon        *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceIcon        `json:"self_service_icon,omitempty"`
}

// FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSON - OK
type FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSON struct {
	General     *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONGeneral     `json:"general,omitempty"`
	Scope       *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONScope       `json:"scope,omitempty"`
	SelfService *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONSelfService `json:"self_service,omitempty"`
}

type FindMobileDeviceConfigurationProfilesByNameSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSONObject *FindMobileDeviceConfigurationProfilesByNameSubset200ApplicationJSON
}
