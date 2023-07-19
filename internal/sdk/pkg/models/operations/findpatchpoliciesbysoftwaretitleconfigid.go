// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FindPatchPoliciesBySoftwareTitleConfigIDRequest struct {
	// Patch software title config ID value to filter by
	Softwaretitleconfigid int64 `pathParam:"style=simple,explode=false,name=softwaretitleconfigid"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralDistributionMethod string

const (
	FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralDistributionMethodSelfservice FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralDistributionMethod = "selfservice"
	FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralDistributionMethodPrompt      FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralDistributionMethod = "prompt"
)

func (e FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralDistributionMethod) ToPointer() *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralDistributionMethod {
	return &e
}

func (e *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralDistributionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "selfservice":
		fallthrough
	case "prompt":
		*e = FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralDistributionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralDistributionMethod: %v", v)
	}
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralKillAppsKillApp struct {
	KillAppBundleID *string
	KillAppName     *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralKillApps struct {
	KillApp *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralKillAppsKillApp
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneral struct {
	AllowDowngrade     *bool
	DistributionMethod *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralDistributionMethod
	Enabled            *bool
	ID                 *int64
	IncrementalUpdates *bool
	KillApps           []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneralKillApps
	MinimumOs          *string
	Name               string
	// Set to true to patch versions unidentified by Jamf Pro patch reporting
	PatchUnknown  *bool
	Reboot        *bool
	ReleaseDate   *int64
	TargetVersion string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeBuildingsBuilding struct {
	ID   *int64
	Name *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeBuildings struct {
	Building *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeBuildingsBuilding
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeComputerGroupsComputerGroup struct {
	ID   *int64
	Name *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeComputerGroups struct {
	ComputerGroup *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeComputerGroupsComputerGroup
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeComputersComputer struct {
	ID   *int64
	Name *string
	Udid *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeComputers struct {
	Computer *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeComputersComputer
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeDepartmentsDepartment struct {
	ID   *int64
	Name *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeDepartments struct {
	Department *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeDepartmentsDepartment
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsBuildingsBuilding struct {
	ID   *int64
	Name *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsBuildings struct {
	Building *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsBuildingsBuilding
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsComputerGroupsComputerGroup struct {
	ID   *int64
	Name *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsComputerGroups struct {
	ComputerGroup *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsComputerGroupsComputerGroup
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsComputersComputer struct {
	ID   *int64
	Name *string
	Udid *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsComputers struct {
	Computer *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsComputersComputer
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsDepartmentsDepartment struct {
	ID   *int64
	Name *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsDepartments struct {
	Department *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsDepartmentsDepartment
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsIbeaconsIbeacon struct {
	ID   *int64
	Name *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsIbeacons struct {
	Ibeacon *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsIbeaconsIbeacon
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsNetworkSegmentsNetworkSegment struct {
	ID   *int64
	Name *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsNetworkSegments struct {
	NetworkSegment *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsNetworkSegmentsNetworkSegment
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusions struct {
	Buildings       []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsBuildings
	ComputerGroups  []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsComputerGroups
	Computers       []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsComputers
	Departments     []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsDepartments
	Ibeacons        []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsIbeacons
	NetworkSegments []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusionsNetworkSegments
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeLimitationsIbeaconsIbeacon struct {
	ID   *int64
	Name *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeLimitationsIbeacons struct {
	Ibeacon *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeLimitationsIbeaconsIbeacon
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeLimitationsNetworkSegmentsNetworkSegment struct {
	ID   *int64
	Name *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeLimitationsNetworkSegments struct {
	NetworkSegment *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeLimitationsNetworkSegmentsNetworkSegment
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeLimitations struct {
	Ibeacons        []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeLimitationsIbeacons
	NetworkSegments []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeLimitationsNetworkSegments
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScope struct {
	AllComputers   *bool
	Buildings      []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeBuildings
	ComputerGroups []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeComputerGroups
	Computers      []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeComputers
	Departments    []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeDepartments
	Exclusions     *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeExclusions
	Limitations    *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScopeLimitations
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionDeadlines struct {
	DeadlineEnabled *bool
	DeadlinePeriod  *int64
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionGracePeriod struct {
	// Number of minutes to wait before automatically closing all apps required to be closed for an update
	GracePeriodDuration       *int64
	Message                   *string
	NotificationCenterSubject *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsNotificationType string

const (
	FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsNotificationTypeSelfService                      FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsNotificationType = "Self Service"
	FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsNotificationTypeSelfServiceAndNotificationCenter FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsNotificationType = "Self Service and Notification Center"
)

func (e FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsNotificationType) ToPointer() *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsNotificationType {
	return &e
}

func (e *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsNotificationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Self Service":
		fallthrough
	case "Self Service and Notification Center":
		*e = FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsNotificationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsNotificationType: %v", v)
	}
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsReminders struct {
	NotificationReminderFrequency *int64
	NotificationRemindersEnabled  *bool
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotifications struct {
	NotificationEnabled *bool
	NotificationMessage *string
	NotificationSubject *string
	NotificationType    *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsNotificationType
	Reminders           *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotificationsReminders
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionSelfServiceIcon struct {
	Filename *string
	ID       *int64
	URI      *string
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteraction struct {
	Deadlines              *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionDeadlines
	GracePeriod            *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionGracePeriod
	InstallButtonText      *string
	Notifications          *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionNotifications
	SelfServiceDescription *string
	SelfServiceIcon        *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteractionSelfServiceIcon
}

// FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXML - OK
type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXML struct {
	General                      *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLGeneral
	Scope                        *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLScope
	SoftwareTitleConfigurationID *int64
	UserInteraction              *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationXMLUserInteraction
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralDistributionMethod string

const (
	FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralDistributionMethodSelfservice FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralDistributionMethod = "selfservice"
	FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralDistributionMethodPrompt      FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralDistributionMethod = "prompt"
)

func (e FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralDistributionMethod) ToPointer() *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralDistributionMethod {
	return &e
}

func (e *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralDistributionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "selfservice":
		fallthrough
	case "prompt":
		*e = FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralDistributionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralDistributionMethod: %v", v)
	}
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralKillAppsKillApp struct {
	KillAppBundleID *string `json:"kill_app_bundle_id,omitempty"`
	KillAppName     *string `json:"kill_app_name,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralKillApps struct {
	KillApp *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralKillAppsKillApp `json:"kill_app,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneral struct {
	AllowDowngrade     *bool                                                                                `json:"allow_downgrade,omitempty"`
	DistributionMethod *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralDistributionMethod `json:"distribution_method,omitempty"`
	Enabled            *bool                                                                                `json:"enabled,omitempty"`
	ID                 *int64                                                                               `json:"id,omitempty"`
	IncrementalUpdates *bool                                                                                `json:"incremental_updates,omitempty"`
	KillApps           []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneralKillApps          `json:"kill_apps,omitempty"`
	MinimumOs          *string                                                                              `json:"minimum_os,omitempty"`
	Name               string                                                                               `json:"name"`
	// Set to true to patch versions unidentified by Jamf Pro patch reporting
	PatchUnknown  *bool  `json:"patch_unknown,omitempty"`
	Reboot        *bool  `json:"reboot,omitempty"`
	ReleaseDate   *int64 `json:"release_date,omitempty"`
	TargetVersion string `json:"target_version"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeBuildingsBuilding struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeBuildings struct {
	Building *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeBuildingsBuilding `json:"building,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeComputerGroupsComputerGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeComputerGroups struct {
	ComputerGroup *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeComputerGroupsComputerGroup `json:"computer_group,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeComputersComputer struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Udid *string `json:"udid,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeComputers struct {
	Computer *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeComputersComputer `json:"computer,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeDepartmentsDepartment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeDepartments struct {
	Department *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeDepartmentsDepartment `json:"department,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsBuildingsBuilding struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsBuildings struct {
	Building *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsBuildingsBuilding `json:"building,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsComputerGroupsComputerGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsComputerGroups struct {
	ComputerGroup *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsComputerGroupsComputerGroup `json:"computer_group,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsComputersComputer struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Udid *string `json:"udid,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsComputers struct {
	Computer *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsComputersComputer `json:"computer,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsDepartmentsDepartment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsDepartments struct {
	Department *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsDepartmentsDepartment `json:"department,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsIbeaconsIbeacon struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsIbeacons struct {
	Ibeacon *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsIbeaconsIbeacon `json:"ibeacon,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsNetworkSegmentsNetworkSegment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsNetworkSegments struct {
	NetworkSegment *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsNetworkSegmentsNetworkSegment `json:"network_segment,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusions struct {
	Buildings       []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsBuildings       `json:"buildings,omitempty"`
	ComputerGroups  []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsComputerGroups  `json:"computer_groups,omitempty"`
	Computers       []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsComputers       `json:"computers,omitempty"`
	Departments     []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsDepartments     `json:"departments,omitempty"`
	Ibeacons        []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsIbeacons        `json:"ibeacons,omitempty"`
	NetworkSegments []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusionsNetworkSegments `json:"network_segments,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeLimitationsIbeaconsIbeacon struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeLimitationsIbeacons struct {
	Ibeacon *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeLimitationsIbeaconsIbeacon `json:"ibeacon,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeLimitationsNetworkSegmentsNetworkSegment struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeLimitationsNetworkSegments struct {
	NetworkSegment *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeLimitationsNetworkSegmentsNetworkSegment `json:"network_segment,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeLimitations struct {
	Ibeacons        []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeLimitationsIbeacons        `json:"ibeacons,omitempty"`
	NetworkSegments []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeLimitationsNetworkSegments `json:"network_segments,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScope struct {
	AllComputers   *bool                                                                           `json:"all_computers,omitempty"`
	Buildings      []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeBuildings      `json:"buildings,omitempty"`
	ComputerGroups []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeComputerGroups `json:"computer_groups,omitempty"`
	Computers      []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeComputers      `json:"computers,omitempty"`
	Departments    []FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeDepartments    `json:"departments,omitempty"`
	Exclusions     *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeExclusions      `json:"exclusions,omitempty"`
	Limitations    *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScopeLimitations     `json:"limitations,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionDeadlines struct {
	DeadlineEnabled *bool  `json:"deadline_enabled,omitempty"`
	DeadlinePeriod  *int64 `json:"deadline_period,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionGracePeriod struct {
	// Number of minutes to wait before automatically closing all apps required to be closed for an update
	GracePeriodDuration       *int64  `json:"grace_period_duration,omitempty"`
	Message                   *string `json:"message,omitempty"`
	NotificationCenterSubject *string `json:"notification_center_subject,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsNotificationType string

const (
	FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsNotificationTypeSelfService                      FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsNotificationType = "Self Service"
	FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsNotificationTypeSelfServiceAndNotificationCenter FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsNotificationType = "Self Service and Notification Center"
)

func (e FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsNotificationType) ToPointer() *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsNotificationType {
	return &e
}

func (e *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsNotificationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Self Service":
		fallthrough
	case "Self Service and Notification Center":
		*e = FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsNotificationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsNotificationType: %v", v)
	}
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsReminders struct {
	NotificationReminderFrequency *int64 `json:"notification_reminder_frequency,omitempty"`
	NotificationRemindersEnabled  *bool  `json:"notification_reminders_enabled,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotifications struct {
	NotificationEnabled *bool                                                                                                   `json:"notification_enabled,omitempty"`
	NotificationMessage *string                                                                                                 `json:"notification_message,omitempty"`
	NotificationSubject *string                                                                                                 `json:"notification_subject,omitempty"`
	NotificationType    *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsNotificationType `json:"notification_type,omitempty"`
	Reminders           *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotificationsReminders        `json:"reminders,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionSelfServiceIcon struct {
	Filename *string `json:"filename,omitempty"`
	ID       *int64  `json:"id,omitempty"`
	URI      *string `json:"uri,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteraction struct {
	Deadlines              *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionDeadlines       `json:"deadlines,omitempty"`
	GracePeriod            *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionGracePeriod     `json:"grace_period,omitempty"`
	InstallButtonText      *string                                                                                   `json:"install_button_text,omitempty"`
	Notifications          *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionNotifications   `json:"notifications,omitempty"`
	SelfServiceDescription *string                                                                                   `json:"self_service_description,omitempty"`
	SelfServiceIcon        *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteractionSelfServiceIcon `json:"self_service_icon,omitempty"`
}

// FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSON - OK
type FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSON struct {
	General                      *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONGeneral         `json:"general,omitempty"`
	Scope                        *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONScope           `json:"scope,omitempty"`
	SoftwareTitleConfigurationID *int64                                                                     `json:"software_title_configuration_id,omitempty"`
	UserInteraction              *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONUserInteraction `json:"user_interaction,omitempty"`
}

type FindPatchPoliciesBySoftwareTitleConfigIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSONObject *FindPatchPoliciesBySoftwareTitleConfigID200ApplicationJSON
}
