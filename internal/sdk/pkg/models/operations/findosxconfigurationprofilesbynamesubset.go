// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindOsxConfigurationProfilesByNameSubsetSubset - Subset to filter by
type FindOsxConfigurationProfilesByNameSubsetSubset string

const (
	FindOsxConfigurationProfilesByNameSubsetSubsetGeneral     FindOsxConfigurationProfilesByNameSubsetSubset = "General"
	FindOsxConfigurationProfilesByNameSubsetSubsetScope       FindOsxConfigurationProfilesByNameSubsetSubset = "Scope"
	FindOsxConfigurationProfilesByNameSubsetSubsetSelfService FindOsxConfigurationProfilesByNameSubsetSubset = "SelfService"
)

func (e FindOsxConfigurationProfilesByNameSubsetSubset) ToPointer() *FindOsxConfigurationProfilesByNameSubsetSubset {
	return &e
}

func (e *FindOsxConfigurationProfilesByNameSubsetSubset) UnmarshalJSON(data []byte) error {
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
		*e = FindOsxConfigurationProfilesByNameSubsetSubset(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindOsxConfigurationProfilesByNameSubsetSubset: %v", v)
	}
}

type FindOsxConfigurationProfilesByNameSubsetRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Subset to filter by
	Subset FindOsxConfigurationProfilesByNameSubsetSubset `pathParam:"style=simple,explode=false,name=subset"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralCategory struct {
	ID *int64
	// Name of the category
	Name string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralDistributionMethod string

const (
	FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralDistributionMethodInstallAutomatically       FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralDistributionMethod = "Install Automatically"
	FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralDistributionMethodMakeAvailableInSelfService FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralDistributionMethod = "Make Available in Self Service"
)

func (e FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralDistributionMethod) ToPointer() *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralDistributionMethod {
	return &e
}

func (e *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralDistributionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Install Automatically":
		fallthrough
	case "Make Available in Self Service":
		*e = FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralDistributionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralDistributionMethod: %v", v)
	}
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel string

const (
	FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevelComputer FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel = "computer"
	FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevelUser     FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel = "user"
)

func (e FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel) ToPointer() *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel {
	return &e
}

func (e *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "computer":
		fallthrough
	case "user":
		*e = FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel: %v", v)
	}
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralSite struct {
	ID *int64
	// Name of the site
	Name string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneral struct {
	Category           *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralCategory
	Description        *string
	DistributionMethod *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralDistributionMethod
	ID                 *int64
	Level              *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralLevel
	// Name of the configuration profile
	Name             string
	Payloads         *string
	RedeployOnUpdate *string
	Site             *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneralSite
	UserRemovable    *bool
	UUID             *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeBuildingsBuilding struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeBuildings struct {
	Building *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeBuildingsBuilding
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeComputerGroupsComputerGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeComputerGroups struct {
	ComputerGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeComputerGroupsComputerGroup
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeComputersComputer struct {
	ID *int64
	// Name of the computer
	Name *string
	Udid *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeComputers struct {
	Computer *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeComputersComputer
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeDepartmentsDepartment struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeDepartments struct {
	Department *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeDepartmentsDepartment
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsBuildingsBuilding struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsBuildings struct {
	Building *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsBuildingsBuilding
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsComputerGroupsComputerGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsComputerGroups struct {
	ComputerGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsComputerGroupsComputerGroup
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsComputersComputer struct {
	ID *int64
	// Name of the computer
	Name *string
	Udid *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsComputers struct {
	Computer *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsComputersComputer
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsDepartmentsDepartment struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsDepartments struct {
	Department *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsDepartmentsDepartment
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsIbeaconsIbeacon struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsIbeacons struct {
	Ibeacon *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsIbeaconsIbeacon
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUserGroupsJssUserGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUserGroups struct {
	JssUserGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUserGroupsJssUserGroup
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUsersJssUser struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUsers struct {
	JssUser *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUsersJssUser
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsNetworkSegmentsNetworkSegment struct {
	ID *int64
	// Name of the network segment
	Name *string
	UID  *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsNetworkSegments struct {
	NetworkSegment *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsNetworkSegmentsNetworkSegment
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUserGroups struct {
	UserGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUserGroupsUserGroup
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUsersUser struct {
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUsers struct {
	User *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUsersUser
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusions struct {
	Buildings       []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsBuildings
	ComputerGroups  []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsComputerGroups
	Computers       []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsComputers
	Departments     []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsDepartments
	Ibeacons        []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsIbeacons
	JssUserGroups   []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUserGroups
	JssUsers        []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsJssUsers
	NetworkSegments []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsNetworkSegments
	UserGroups      []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUserGroups
	Users           []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusionsUsers
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUserGroupsJssUserGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUserGroups struct {
	JssUserGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUserGroupsJssUserGroup
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUsersUsers struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUsers struct {
	Users *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUsersUsers
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsIbeaconsIbeacon struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsIbeacons struct {
	Ibeacon *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsIbeaconsIbeacon
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsNetworkSegmentsNetworkSegment struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsNetworkSegments struct {
	NetworkSegment *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsNetworkSegmentsNetworkSegment
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUserGroups struct {
	UserGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUserGroupsUserGroup
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUsersUser struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUsers struct {
	User *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUsersUser
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitations struct {
	Ibeacons        []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsIbeacons
	NetworkSegments []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsNetworkSegments
	UserGroups      []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUserGroups
	Users           []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitationsUsers
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScope struct {
	AllComputers   *bool
	AllJssUsers    *bool
	Buildings      []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeBuildings
	ComputerGroups []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeComputerGroups
	Computers      []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeComputers
	Departments    []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeDepartments
	Exclusions     *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeExclusions
	JssUserGroups  []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUserGroups
	JssUsers       []FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeJssUsers
	Limitations    *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScopeLimitations
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceCategoriesCategory struct {
	DisplayIn *bool
	FeatureIn *bool
	ID        *int64
	Name      *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceCategories struct {
	Category *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceCategoriesCategory
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceIcon struct {
	Data *string
	ID   *int64
	URI  *string
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationXMLSelfService struct {
	FeatureOnMainPage           *bool
	ForceUsersToViewDescription *bool
	InstallButtonText           *string
	Notification                *string
	NotificationMessage         *string
	NotificationSubject         *string
	SelfServiceCategories       *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceCategories
	SelfServiceDescription      *string
	SelfServiceIcon             *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLSelfServiceSelfServiceIcon
}

// FindOsxConfigurationProfilesByNameSubset200ApplicationXML - OK
type FindOsxConfigurationProfilesByNameSubset200ApplicationXML struct {
	General     *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLGeneral
	Scope       *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLScope
	SelfService *FindOsxConfigurationProfilesByNameSubset200ApplicationXMLSelfService
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralCategory struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the category
	Name string `json:"name"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralDistributionMethod string

const (
	FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralDistributionMethodInstallAutomatically       FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralDistributionMethod = "Install Automatically"
	FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralDistributionMethodMakeAvailableInSelfService FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralDistributionMethod = "Make Available in Self Service"
)

func (e FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralDistributionMethod) ToPointer() *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralDistributionMethod {
	return &e
}

func (e *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralDistributionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Install Automatically":
		fallthrough
	case "Make Available in Self Service":
		*e = FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralDistributionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralDistributionMethod: %v", v)
	}
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel string

const (
	FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevelComputer FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel = "computer"
	FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevelUser     FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel = "user"
)

func (e FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel) ToPointer() *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel {
	return &e
}

func (e *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "computer":
		fallthrough
	case "user":
		*e = FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel: %v", v)
	}
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralSite struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the site
	Name string `json:"name"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneral struct {
	Category           *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralCategory           `json:"category,omitempty"`
	Description        *string                                                                              `json:"description,omitempty"`
	DistributionMethod *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralDistributionMethod `json:"distribution_method,omitempty"`
	ID                 *int64                                                                               `json:"id,omitempty"`
	Level              *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralLevel              `json:"level,omitempty"`
	// Name of the configuration profile
	Name             string                                                                 `json:"name"`
	Payloads         *string                                                                `json:"payloads,omitempty"`
	RedeployOnUpdate *string                                                                `json:"redeploy_on_update,omitempty"`
	Site             *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneralSite `json:"site,omitempty"`
	UserRemovable    *bool                                                                  `json:"user_removable,omitempty"`
	UUID             *string                                                                `json:"uuid,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeBuildingsBuilding struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeBuildings struct {
	Building *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeBuildingsBuilding `json:"building,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeComputerGroupsComputerGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeComputerGroups struct {
	ComputerGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeComputerGroupsComputerGroup `json:"computer_group,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeComputersComputer struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the computer
	Name *string `json:"name,omitempty"`
	Udid *string `json:"udid,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeComputers struct {
	Computer *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeComputersComputer `json:"computer,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeDepartmentsDepartment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeDepartments struct {
	Department *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeDepartmentsDepartment `json:"department,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsBuildingsBuilding struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsBuildings struct {
	Building *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsBuildingsBuilding `json:"building,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsComputerGroupsComputerGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsComputerGroups struct {
	ComputerGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsComputerGroupsComputerGroup `json:"computer_group,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsComputersComputer struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the computer
	Name *string `json:"name,omitempty"`
	Udid *string `json:"udid,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsComputers struct {
	Computer *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsComputersComputer `json:"computer,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsDepartmentsDepartment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsDepartments struct {
	Department *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsDepartmentsDepartment `json:"department,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsIbeaconsIbeacon struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsIbeacons struct {
	Ibeacon *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsIbeaconsIbeacon `json:"ibeacon,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUserGroupsJssUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUserGroups struct {
	JssUserGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUserGroupsJssUserGroup `json:"jss_user_group,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUsersJssUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUsers struct {
	JssUser *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUsersJssUser `json:"jss_user,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsNetworkSegmentsNetworkSegment struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the network segment
	Name *string `json:"name,omitempty"`
	UID  *string `json:"uid,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsNetworkSegments struct {
	NetworkSegment *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsNetworkSegmentsNetworkSegment `json:"network_segment,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUserGroups struct {
	UserGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUsersUser struct {
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUsers struct {
	User *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUsersUser `json:"user,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusions struct {
	Buildings       []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsBuildings       `json:"buildings,omitempty"`
	ComputerGroups  []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsComputerGroups  `json:"computer_groups,omitempty"`
	Computers       []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsComputers       `json:"computers,omitempty"`
	Departments     []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsDepartments     `json:"departments,omitempty"`
	Ibeacons        []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsIbeacons        `json:"ibeacons,omitempty"`
	JssUserGroups   []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUserGroups   `json:"jss_user_groups,omitempty"`
	JssUsers        []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsJssUsers        `json:"jss_users,omitempty"`
	NetworkSegments []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsNetworkSegments `json:"network_segments,omitempty"`
	UserGroups      []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUserGroups      `json:"user_groups,omitempty"`
	Users           []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusionsUsers           `json:"users,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUserGroupsJssUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUserGroups struct {
	JssUserGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUserGroupsJssUserGroup `json:"jss_user_group,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUsersUsers struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUsers struct {
	Users *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUsersUsers `json:"users,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsIbeaconsIbeacon struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsIbeacons struct {
	Ibeacon *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsIbeaconsIbeacon `json:"ibeacon,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsNetworkSegmentsNetworkSegment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsNetworkSegments struct {
	NetworkSegment *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsNetworkSegmentsNetworkSegment `json:"network_segment,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUserGroups struct {
	UserGroup *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUsersUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUsers struct {
	User *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUsersUser `json:"user,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitations struct {
	Ibeacons        []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsIbeacons        `json:"ibeacons,omitempty"`
	NetworkSegments []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsNetworkSegments `json:"network_segments,omitempty"`
	UserGroups      []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUserGroups      `json:"user_groups,omitempty"`
	Users           []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitationsUsers           `json:"users,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScope struct {
	AllComputers   *bool                                                                           `json:"all_computers,omitempty"`
	AllJssUsers    *bool                                                                           `json:"all_jss_users,omitempty"`
	Buildings      []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeBuildings      `json:"buildings,omitempty"`
	ComputerGroups []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeComputerGroups `json:"computer_groups,omitempty"`
	Computers      []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeComputers      `json:"computers,omitempty"`
	Departments    []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeDepartments    `json:"departments,omitempty"`
	Exclusions     *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeExclusions      `json:"exclusions,omitempty"`
	JssUserGroups  []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUserGroups  `json:"jss_user_groups,omitempty"`
	JssUsers       []FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeJssUsers       `json:"jss_users,omitempty"`
	Limitations    *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScopeLimitations     `json:"limitations,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceCategoriesCategory struct {
	DisplayIn *bool   `json:"display_in,omitempty"`
	FeatureIn *bool   `json:"feature_in,omitempty"`
	ID        *int64  `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceCategories struct {
	Category *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceCategoriesCategory `json:"category,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceIcon struct {
	Data *string `json:"data,omitempty"`
	ID   *int64  `json:"id,omitempty"`
	URI  *string `json:"uri,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubset200ApplicationJSONSelfService struct {
	FeatureOnMainPage           *bool                                                                                       `json:"feature_on_main_page,omitempty"`
	ForceUsersToViewDescription *bool                                                                                       `json:"force_users_to_view_description,omitempty"`
	InstallButtonText           *string                                                                                     `json:"install_button_text,omitempty"`
	Notification                *string                                                                                     `json:"notification,omitempty"`
	NotificationMessage         *string                                                                                     `json:"notification_message,omitempty"`
	NotificationSubject         *string                                                                                     `json:"notification_subject,omitempty"`
	SelfServiceCategories       *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceCategories `json:"self_service_categories,omitempty"`
	SelfServiceDescription      *string                                                                                     `json:"self_service_description,omitempty"`
	SelfServiceIcon             *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONSelfServiceSelfServiceIcon       `json:"self_service_icon,omitempty"`
}

// FindOsxConfigurationProfilesByNameSubset200ApplicationJSON - OK
type FindOsxConfigurationProfilesByNameSubset200ApplicationJSON struct {
	General     *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONGeneral     `json:"general,omitempty"`
	Scope       *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONScope       `json:"scope,omitempty"`
	SelfService *FindOsxConfigurationProfilesByNameSubset200ApplicationJSONSelfService `json:"self_service,omitempty"`
}

type FindOsxConfigurationProfilesByNameSubsetResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindOsxConfigurationProfilesByNameSubset200ApplicationJSONObject *FindOsxConfigurationProfilesByNameSubset200ApplicationJSON
}
