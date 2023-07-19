// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FindOsxConfigurationProfilesByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type FindOsxConfigurationProfilesByID200ApplicationXMLGeneralCategory struct {
	ID *int64
	// Name of the category
	Name string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLGeneralDistributionMethod string

const (
	FindOsxConfigurationProfilesByID200ApplicationXMLGeneralDistributionMethodInstallAutomatically       FindOsxConfigurationProfilesByID200ApplicationXMLGeneralDistributionMethod = "Install Automatically"
	FindOsxConfigurationProfilesByID200ApplicationXMLGeneralDistributionMethodMakeAvailableInSelfService FindOsxConfigurationProfilesByID200ApplicationXMLGeneralDistributionMethod = "Make Available in Self Service"
)

func (e FindOsxConfigurationProfilesByID200ApplicationXMLGeneralDistributionMethod) ToPointer() *FindOsxConfigurationProfilesByID200ApplicationXMLGeneralDistributionMethod {
	return &e
}

func (e *FindOsxConfigurationProfilesByID200ApplicationXMLGeneralDistributionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Install Automatically":
		fallthrough
	case "Make Available in Self Service":
		*e = FindOsxConfigurationProfilesByID200ApplicationXMLGeneralDistributionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindOsxConfigurationProfilesByID200ApplicationXMLGeneralDistributionMethod: %v", v)
	}
}

type FindOsxConfigurationProfilesByID200ApplicationXMLGeneralLevel string

const (
	FindOsxConfigurationProfilesByID200ApplicationXMLGeneralLevelComputer FindOsxConfigurationProfilesByID200ApplicationXMLGeneralLevel = "computer"
	FindOsxConfigurationProfilesByID200ApplicationXMLGeneralLevelUser     FindOsxConfigurationProfilesByID200ApplicationXMLGeneralLevel = "user"
)

func (e FindOsxConfigurationProfilesByID200ApplicationXMLGeneralLevel) ToPointer() *FindOsxConfigurationProfilesByID200ApplicationXMLGeneralLevel {
	return &e
}

func (e *FindOsxConfigurationProfilesByID200ApplicationXMLGeneralLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "computer":
		fallthrough
	case "user":
		*e = FindOsxConfigurationProfilesByID200ApplicationXMLGeneralLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindOsxConfigurationProfilesByID200ApplicationXMLGeneralLevel: %v", v)
	}
}

type FindOsxConfigurationProfilesByID200ApplicationXMLGeneralSite struct {
	ID *int64
	// Name of the site
	Name string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLGeneral struct {
	Category           *FindOsxConfigurationProfilesByID200ApplicationXMLGeneralCategory
	Description        *string
	DistributionMethod *FindOsxConfigurationProfilesByID200ApplicationXMLGeneralDistributionMethod
	ID                 *int64
	Level              *FindOsxConfigurationProfilesByID200ApplicationXMLGeneralLevel
	// Name of the configuration profile
	Name             string
	Payloads         *string
	RedeployOnUpdate *string
	Site             *FindOsxConfigurationProfilesByID200ApplicationXMLGeneralSite
	UserRemovable    *bool
	UUID             *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeBuildingsBuilding struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeBuildings struct {
	Building *FindOsxConfigurationProfilesByID200ApplicationXMLScopeBuildingsBuilding
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeComputerGroupsComputerGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeComputerGroups struct {
	ComputerGroup *FindOsxConfigurationProfilesByID200ApplicationXMLScopeComputerGroupsComputerGroup
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeComputersComputer struct {
	ID *int64
	// Name of the computer
	Name *string
	Udid *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeComputers struct {
	Computer *FindOsxConfigurationProfilesByID200ApplicationXMLScopeComputersComputer
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeDepartmentsDepartment struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeDepartments struct {
	Department *FindOsxConfigurationProfilesByID200ApplicationXMLScopeDepartmentsDepartment
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsBuildingsBuilding struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsBuildings struct {
	Building *FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsBuildingsBuilding
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsComputerGroupsComputerGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsComputerGroups struct {
	ComputerGroup *FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsComputerGroupsComputerGroup
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsComputersComputer struct {
	ID *int64
	// Name of the computer
	Name *string
	Udid *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsComputers struct {
	Computer *FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsComputersComputer
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsDepartmentsDepartment struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsDepartments struct {
	Department *FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsDepartmentsDepartment
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsIbeaconsIbeacon struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsIbeacons struct {
	Ibeacon *FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsIbeaconsIbeacon
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsJssUserGroupsJssUserGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsJssUserGroups struct {
	JssUserGroup *FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsJssUserGroupsJssUserGroup
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsJssUsersJssUser struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsJssUsers struct {
	JssUser *FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsJssUsersJssUser
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsNetworkSegmentsNetworkSegment struct {
	ID *int64
	// Name of the network segment
	Name *string
	UID  *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsNetworkSegments struct {
	NetworkSegment *FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsNetworkSegmentsNetworkSegment
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsUserGroups struct {
	UserGroup *FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsUserGroupsUserGroup
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsUsersUser struct {
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsUsers struct {
	User *FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsUsersUser
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusions struct {
	Buildings       []FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsBuildings
	ComputerGroups  []FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsComputerGroups
	Computers       []FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsComputers
	Departments     []FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsDepartments
	Ibeacons        []FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsIbeacons
	JssUserGroups   []FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsJssUserGroups
	JssUsers        []FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsJssUsers
	NetworkSegments []FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsNetworkSegments
	UserGroups      []FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsUserGroups
	Users           []FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusionsUsers
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeJssUserGroupsJssUserGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeJssUserGroups struct {
	JssUserGroup *FindOsxConfigurationProfilesByID200ApplicationXMLScopeJssUserGroupsJssUserGroup
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeJssUsersUsers struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeJssUsers struct {
	Users *FindOsxConfigurationProfilesByID200ApplicationXMLScopeJssUsersUsers
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsIbeaconsIbeacon struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsIbeacons struct {
	Ibeacon *FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsIbeaconsIbeacon
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsNetworkSegmentsNetworkSegment struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsNetworkSegments struct {
	NetworkSegment *FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsNetworkSegmentsNetworkSegment
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsUserGroupsUserGroup struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsUserGroups struct {
	UserGroup *FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsUserGroupsUserGroup
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsUsersUser struct {
	ID   *int64
	Name *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsUsers struct {
	User *FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsUsersUser
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitations struct {
	Ibeacons        []FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsIbeacons
	NetworkSegments []FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsNetworkSegments
	UserGroups      []FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsUserGroups
	Users           []FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitationsUsers
}

type FindOsxConfigurationProfilesByID200ApplicationXMLScope struct {
	AllComputers   *bool
	AllJssUsers    *bool
	Buildings      []FindOsxConfigurationProfilesByID200ApplicationXMLScopeBuildings
	ComputerGroups []FindOsxConfigurationProfilesByID200ApplicationXMLScopeComputerGroups
	Computers      []FindOsxConfigurationProfilesByID200ApplicationXMLScopeComputers
	Departments    []FindOsxConfigurationProfilesByID200ApplicationXMLScopeDepartments
	Exclusions     *FindOsxConfigurationProfilesByID200ApplicationXMLScopeExclusions
	JssUserGroups  []FindOsxConfigurationProfilesByID200ApplicationXMLScopeJssUserGroups
	JssUsers       []FindOsxConfigurationProfilesByID200ApplicationXMLScopeJssUsers
	Limitations    *FindOsxConfigurationProfilesByID200ApplicationXMLScopeLimitations
}

type FindOsxConfigurationProfilesByID200ApplicationXMLSelfServiceSelfServiceCategoriesCategory struct {
	DisplayIn *bool
	FeatureIn *bool
	ID        *int64
	Name      *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLSelfServiceSelfServiceCategories struct {
	Category *FindOsxConfigurationProfilesByID200ApplicationXMLSelfServiceSelfServiceCategoriesCategory
}

type FindOsxConfigurationProfilesByID200ApplicationXMLSelfServiceSelfServiceIcon struct {
	Data *string
	ID   *int64
	URI  *string
}

type FindOsxConfigurationProfilesByID200ApplicationXMLSelfService struct {
	FeatureOnMainPage           *bool
	ForceUsersToViewDescription *bool
	InstallButtonText           *string
	Notification                *string
	NotificationMessage         *string
	NotificationSubject         *string
	SelfServiceCategories       *FindOsxConfigurationProfilesByID200ApplicationXMLSelfServiceSelfServiceCategories
	SelfServiceDescription      *string
	SelfServiceIcon             *FindOsxConfigurationProfilesByID200ApplicationXMLSelfServiceSelfServiceIcon
}

// FindOsxConfigurationProfilesByID200ApplicationXML - OK
type FindOsxConfigurationProfilesByID200ApplicationXML struct {
	General     *FindOsxConfigurationProfilesByID200ApplicationXMLGeneral
	Scope       *FindOsxConfigurationProfilesByID200ApplicationXMLScope
	SelfService *FindOsxConfigurationProfilesByID200ApplicationXMLSelfService
}

type FindOsxConfigurationProfilesByID200ApplicationJSONGeneralCategory struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the category
	Name string `json:"name"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONGeneralDistributionMethod string

const (
	FindOsxConfigurationProfilesByID200ApplicationJSONGeneralDistributionMethodInstallAutomatically       FindOsxConfigurationProfilesByID200ApplicationJSONGeneralDistributionMethod = "Install Automatically"
	FindOsxConfigurationProfilesByID200ApplicationJSONGeneralDistributionMethodMakeAvailableInSelfService FindOsxConfigurationProfilesByID200ApplicationJSONGeneralDistributionMethod = "Make Available in Self Service"
)

func (e FindOsxConfigurationProfilesByID200ApplicationJSONGeneralDistributionMethod) ToPointer() *FindOsxConfigurationProfilesByID200ApplicationJSONGeneralDistributionMethod {
	return &e
}

func (e *FindOsxConfigurationProfilesByID200ApplicationJSONGeneralDistributionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Install Automatically":
		fallthrough
	case "Make Available in Self Service":
		*e = FindOsxConfigurationProfilesByID200ApplicationJSONGeneralDistributionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindOsxConfigurationProfilesByID200ApplicationJSONGeneralDistributionMethod: %v", v)
	}
}

type FindOsxConfigurationProfilesByID200ApplicationJSONGeneralLevel string

const (
	FindOsxConfigurationProfilesByID200ApplicationJSONGeneralLevelComputer FindOsxConfigurationProfilesByID200ApplicationJSONGeneralLevel = "computer"
	FindOsxConfigurationProfilesByID200ApplicationJSONGeneralLevelUser     FindOsxConfigurationProfilesByID200ApplicationJSONGeneralLevel = "user"
)

func (e FindOsxConfigurationProfilesByID200ApplicationJSONGeneralLevel) ToPointer() *FindOsxConfigurationProfilesByID200ApplicationJSONGeneralLevel {
	return &e
}

func (e *FindOsxConfigurationProfilesByID200ApplicationJSONGeneralLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "computer":
		fallthrough
	case "user":
		*e = FindOsxConfigurationProfilesByID200ApplicationJSONGeneralLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindOsxConfigurationProfilesByID200ApplicationJSONGeneralLevel: %v", v)
	}
}

type FindOsxConfigurationProfilesByID200ApplicationJSONGeneralSite struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the site
	Name string `json:"name"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONGeneral struct {
	Category           *FindOsxConfigurationProfilesByID200ApplicationJSONGeneralCategory           `json:"category,omitempty"`
	Description        *string                                                                      `json:"description,omitempty"`
	DistributionMethod *FindOsxConfigurationProfilesByID200ApplicationJSONGeneralDistributionMethod `json:"distribution_method,omitempty"`
	ID                 *int64                                                                       `json:"id,omitempty"`
	Level              *FindOsxConfigurationProfilesByID200ApplicationJSONGeneralLevel              `json:"level,omitempty"`
	// Name of the configuration profile
	Name             string                                                         `json:"name"`
	Payloads         *string                                                        `json:"payloads,omitempty"`
	RedeployOnUpdate *string                                                        `json:"redeploy_on_update,omitempty"`
	Site             *FindOsxConfigurationProfilesByID200ApplicationJSONGeneralSite `json:"site,omitempty"`
	UserRemovable    *bool                                                          `json:"user_removable,omitempty"`
	UUID             *string                                                        `json:"uuid,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeBuildingsBuilding struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeBuildings struct {
	Building *FindOsxConfigurationProfilesByID200ApplicationJSONScopeBuildingsBuilding `json:"building,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeComputerGroupsComputerGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeComputerGroups struct {
	ComputerGroup *FindOsxConfigurationProfilesByID200ApplicationJSONScopeComputerGroupsComputerGroup `json:"computer_group,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeComputersComputer struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the computer
	Name *string `json:"name,omitempty"`
	Udid *string `json:"udid,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeComputers struct {
	Computer *FindOsxConfigurationProfilesByID200ApplicationJSONScopeComputersComputer `json:"computer,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeDepartmentsDepartment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeDepartments struct {
	Department *FindOsxConfigurationProfilesByID200ApplicationJSONScopeDepartmentsDepartment `json:"department,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsBuildingsBuilding struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsBuildings struct {
	Building *FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsBuildingsBuilding `json:"building,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsComputerGroupsComputerGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsComputerGroups struct {
	ComputerGroup *FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsComputerGroupsComputerGroup `json:"computer_group,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsComputersComputer struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the computer
	Name *string `json:"name,omitempty"`
	Udid *string `json:"udid,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsComputers struct {
	Computer *FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsComputersComputer `json:"computer,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsDepartmentsDepartment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsDepartments struct {
	Department *FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsDepartmentsDepartment `json:"department,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsIbeaconsIbeacon struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsIbeacons struct {
	Ibeacon *FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsIbeaconsIbeacon `json:"ibeacon,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsJssUserGroupsJssUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsJssUserGroups struct {
	JssUserGroup *FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsJssUserGroupsJssUserGroup `json:"jss_user_group,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsJssUsersJssUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsJssUsers struct {
	JssUser *FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsJssUsersJssUser `json:"jss_user,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsNetworkSegmentsNetworkSegment struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the network segment
	Name *string `json:"name,omitempty"`
	UID  *string `json:"uid,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsNetworkSegments struct {
	NetworkSegment *FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsNetworkSegmentsNetworkSegment `json:"network_segment,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsUserGroups struct {
	UserGroup *FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsUsersUser struct {
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsUsers struct {
	User *FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsUsersUser `json:"user,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusions struct {
	Buildings       []FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsBuildings       `json:"buildings,omitempty"`
	ComputerGroups  []FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsComputerGroups  `json:"computer_groups,omitempty"`
	Computers       []FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsComputers       `json:"computers,omitempty"`
	Departments     []FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsDepartments     `json:"departments,omitempty"`
	Ibeacons        []FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsIbeacons        `json:"ibeacons,omitempty"`
	JssUserGroups   []FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsJssUserGroups   `json:"jss_user_groups,omitempty"`
	JssUsers        []FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsJssUsers        `json:"jss_users,omitempty"`
	NetworkSegments []FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsNetworkSegments `json:"network_segments,omitempty"`
	UserGroups      []FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsUserGroups      `json:"user_groups,omitempty"`
	Users           []FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusionsUsers           `json:"users,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeJssUserGroupsJssUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeJssUserGroups struct {
	JssUserGroup *FindOsxConfigurationProfilesByID200ApplicationJSONScopeJssUserGroupsJssUserGroup `json:"jss_user_group,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeJssUsersUsers struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeJssUsers struct {
	Users *FindOsxConfigurationProfilesByID200ApplicationJSONScopeJssUsersUsers `json:"users,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsIbeaconsIbeacon struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsIbeacons struct {
	Ibeacon *FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsIbeaconsIbeacon `json:"ibeacon,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsNetworkSegmentsNetworkSegment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsNetworkSegments struct {
	NetworkSegment *FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsNetworkSegmentsNetworkSegment `json:"network_segment,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsUserGroupsUserGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsUserGroups struct {
	UserGroup *FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsUserGroupsUserGroup `json:"user_group,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsUsersUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsUsers struct {
	User *FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsUsersUser `json:"user,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitations struct {
	Ibeacons        []FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsIbeacons        `json:"ibeacons,omitempty"`
	NetworkSegments []FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsNetworkSegments `json:"network_segments,omitempty"`
	UserGroups      []FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsUserGroups      `json:"user_groups,omitempty"`
	Users           []FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitationsUsers           `json:"users,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONScope struct {
	AllComputers   *bool                                                                   `json:"all_computers,omitempty"`
	AllJssUsers    *bool                                                                   `json:"all_jss_users,omitempty"`
	Buildings      []FindOsxConfigurationProfilesByID200ApplicationJSONScopeBuildings      `json:"buildings,omitempty"`
	ComputerGroups []FindOsxConfigurationProfilesByID200ApplicationJSONScopeComputerGroups `json:"computer_groups,omitempty"`
	Computers      []FindOsxConfigurationProfilesByID200ApplicationJSONScopeComputers      `json:"computers,omitempty"`
	Departments    []FindOsxConfigurationProfilesByID200ApplicationJSONScopeDepartments    `json:"departments,omitempty"`
	Exclusions     *FindOsxConfigurationProfilesByID200ApplicationJSONScopeExclusions      `json:"exclusions,omitempty"`
	JssUserGroups  []FindOsxConfigurationProfilesByID200ApplicationJSONScopeJssUserGroups  `json:"jss_user_groups,omitempty"`
	JssUsers       []FindOsxConfigurationProfilesByID200ApplicationJSONScopeJssUsers       `json:"jss_users,omitempty"`
	Limitations    *FindOsxConfigurationProfilesByID200ApplicationJSONScopeLimitations     `json:"limitations,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONSelfServiceSelfServiceCategoriesCategory struct {
	DisplayIn *bool   `json:"display_in,omitempty"`
	FeatureIn *bool   `json:"feature_in,omitempty"`
	ID        *int64  `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONSelfServiceSelfServiceCategories struct {
	Category *FindOsxConfigurationProfilesByID200ApplicationJSONSelfServiceSelfServiceCategoriesCategory `json:"category,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONSelfServiceSelfServiceIcon struct {
	Data *string `json:"data,omitempty"`
	ID   *int64  `json:"id,omitempty"`
	URI  *string `json:"uri,omitempty"`
}

type FindOsxConfigurationProfilesByID200ApplicationJSONSelfService struct {
	FeatureOnMainPage           *bool                                                                               `json:"feature_on_main_page,omitempty"`
	ForceUsersToViewDescription *bool                                                                               `json:"force_users_to_view_description,omitempty"`
	InstallButtonText           *string                                                                             `json:"install_button_text,omitempty"`
	Notification                *string                                                                             `json:"notification,omitempty"`
	NotificationMessage         *string                                                                             `json:"notification_message,omitempty"`
	NotificationSubject         *string                                                                             `json:"notification_subject,omitempty"`
	SelfServiceCategories       *FindOsxConfigurationProfilesByID200ApplicationJSONSelfServiceSelfServiceCategories `json:"self_service_categories,omitempty"`
	SelfServiceDescription      *string                                                                             `json:"self_service_description,omitempty"`
	SelfServiceIcon             *FindOsxConfigurationProfilesByID200ApplicationJSONSelfServiceSelfServiceIcon       `json:"self_service_icon,omitempty"`
}

// FindOsxConfigurationProfilesByID200ApplicationJSON - OK
type FindOsxConfigurationProfilesByID200ApplicationJSON struct {
	General     *FindOsxConfigurationProfilesByID200ApplicationJSONGeneral     `json:"general,omitempty"`
	Scope       *FindOsxConfigurationProfilesByID200ApplicationJSONScope       `json:"scope,omitempty"`
	SelfService *FindOsxConfigurationProfilesByID200ApplicationJSONSelfService `json:"self_service,omitempty"`
}

type FindOsxConfigurationProfilesByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindOsxConfigurationProfilesByID200ApplicationJSONObject *FindOsxConfigurationProfilesByID200ApplicationJSON
}
