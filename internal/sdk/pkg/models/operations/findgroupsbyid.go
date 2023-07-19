// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FindGroupsByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type FindGroupsByID200ApplicationXMLAccessLevel string

const (
	FindGroupsByID200ApplicationXMLAccessLevelFullAccess  FindGroupsByID200ApplicationXMLAccessLevel = "Full Access"
	FindGroupsByID200ApplicationXMLAccessLevelSiteAccess  FindGroupsByID200ApplicationXMLAccessLevel = "Site Access"
	FindGroupsByID200ApplicationXMLAccessLevelGroupAccess FindGroupsByID200ApplicationXMLAccessLevel = "Group Access"
)

func (e FindGroupsByID200ApplicationXMLAccessLevel) ToPointer() *FindGroupsByID200ApplicationXMLAccessLevel {
	return &e
}

func (e *FindGroupsByID200ApplicationXMLAccessLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Full Access":
		fallthrough
	case "Site Access":
		fallthrough
	case "Group Access":
		*e = FindGroupsByID200ApplicationXMLAccessLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindGroupsByID200ApplicationXMLAccessLevel: %v", v)
	}
}

type FindGroupsByID200ApplicationXMLMembersUser struct {
	ID   *int64
	Name *string
}

type FindGroupsByID200ApplicationXMLMembers struct {
	User *FindGroupsByID200ApplicationXMLMembersUser
}

type FindGroupsByID200ApplicationXMLPrivilegeSet string

const (
	FindGroupsByID200ApplicationXMLPrivilegeSetAdministrator  FindGroupsByID200ApplicationXMLPrivilegeSet = "Administrator"
	FindGroupsByID200ApplicationXMLPrivilegeSetAuditor        FindGroupsByID200ApplicationXMLPrivilegeSet = "Auditor"
	FindGroupsByID200ApplicationXMLPrivilegeSetEnrollmentOnly FindGroupsByID200ApplicationXMLPrivilegeSet = "Enrollment Only"
	FindGroupsByID200ApplicationXMLPrivilegeSetCustom         FindGroupsByID200ApplicationXMLPrivilegeSet = "Custom"
)

func (e FindGroupsByID200ApplicationXMLPrivilegeSet) ToPointer() *FindGroupsByID200ApplicationXMLPrivilegeSet {
	return &e
}

func (e *FindGroupsByID200ApplicationXMLPrivilegeSet) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Administrator":
		fallthrough
	case "Auditor":
		fallthrough
	case "Enrollment Only":
		fallthrough
	case "Custom":
		*e = FindGroupsByID200ApplicationXMLPrivilegeSet(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindGroupsByID200ApplicationXMLPrivilegeSet: %v", v)
	}
}

type FindGroupsByID200ApplicationXMLPrivilegesCasperAdmin struct {
	Privilege *string
}

type FindGroupsByID200ApplicationXMLPrivilegesCasperImaging struct {
	Privilege *string
}

type FindGroupsByID200ApplicationXMLPrivilegesCasperRemote struct {
	Privilege *string
}

type FindGroupsByID200ApplicationXMLPrivilegesJssActions struct {
	Privilege *string
}

type FindGroupsByID200ApplicationXMLPrivilegesJssObjects struct {
	Privilege *string
}

type FindGroupsByID200ApplicationXMLPrivilegesJssSettings struct {
	Privilege *string
}

type FindGroupsByID200ApplicationXMLPrivilegesRecon struct {
	Privilege *string
}

type FindGroupsByID200ApplicationXMLPrivileges struct {
	CasperAdmin   []FindGroupsByID200ApplicationXMLPrivilegesCasperAdmin
	CasperImaging []FindGroupsByID200ApplicationXMLPrivilegesCasperImaging
	CasperRemote  []FindGroupsByID200ApplicationXMLPrivilegesCasperRemote
	JssActions    []FindGroupsByID200ApplicationXMLPrivilegesJssActions
	JssObjects    []FindGroupsByID200ApplicationXMLPrivilegesJssObjects
	JssSettings   []FindGroupsByID200ApplicationXMLPrivilegesJssSettings
	Recon         []FindGroupsByID200ApplicationXMLPrivilegesRecon
}

type FindGroupsByID200ApplicationXMLSite struct {
	ID *int64
	// Name of the site
	Name string
}

// FindGroupsByID200ApplicationXML - OK
type FindGroupsByID200ApplicationXML struct {
	AccessLevel *FindGroupsByID200ApplicationXMLAccessLevel
	ID          *int64
	Members     []FindGroupsByID200ApplicationXMLMembers
	// Group name
	Name         string
	PrivilegeSet *FindGroupsByID200ApplicationXMLPrivilegeSet
	Privileges   *FindGroupsByID200ApplicationXMLPrivileges
	Site         *FindGroupsByID200ApplicationXMLSite
}

type FindGroupsByID200ApplicationJSONAccessLevel string

const (
	FindGroupsByID200ApplicationJSONAccessLevelFullAccess  FindGroupsByID200ApplicationJSONAccessLevel = "Full Access"
	FindGroupsByID200ApplicationJSONAccessLevelSiteAccess  FindGroupsByID200ApplicationJSONAccessLevel = "Site Access"
	FindGroupsByID200ApplicationJSONAccessLevelGroupAccess FindGroupsByID200ApplicationJSONAccessLevel = "Group Access"
)

func (e FindGroupsByID200ApplicationJSONAccessLevel) ToPointer() *FindGroupsByID200ApplicationJSONAccessLevel {
	return &e
}

func (e *FindGroupsByID200ApplicationJSONAccessLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Full Access":
		fallthrough
	case "Site Access":
		fallthrough
	case "Group Access":
		*e = FindGroupsByID200ApplicationJSONAccessLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindGroupsByID200ApplicationJSONAccessLevel: %v", v)
	}
}

type FindGroupsByID200ApplicationJSONMembersUser struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindGroupsByID200ApplicationJSONMembers struct {
	User *FindGroupsByID200ApplicationJSONMembersUser `json:"user,omitempty"`
}

type FindGroupsByID200ApplicationJSONPrivilegeSet string

const (
	FindGroupsByID200ApplicationJSONPrivilegeSetAdministrator  FindGroupsByID200ApplicationJSONPrivilegeSet = "Administrator"
	FindGroupsByID200ApplicationJSONPrivilegeSetAuditor        FindGroupsByID200ApplicationJSONPrivilegeSet = "Auditor"
	FindGroupsByID200ApplicationJSONPrivilegeSetEnrollmentOnly FindGroupsByID200ApplicationJSONPrivilegeSet = "Enrollment Only"
	FindGroupsByID200ApplicationJSONPrivilegeSetCustom         FindGroupsByID200ApplicationJSONPrivilegeSet = "Custom"
)

func (e FindGroupsByID200ApplicationJSONPrivilegeSet) ToPointer() *FindGroupsByID200ApplicationJSONPrivilegeSet {
	return &e
}

func (e *FindGroupsByID200ApplicationJSONPrivilegeSet) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Administrator":
		fallthrough
	case "Auditor":
		fallthrough
	case "Enrollment Only":
		fallthrough
	case "Custom":
		*e = FindGroupsByID200ApplicationJSONPrivilegeSet(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindGroupsByID200ApplicationJSONPrivilegeSet: %v", v)
	}
}

type FindGroupsByID200ApplicationJSONPrivilegesCasperAdmin struct {
	Privilege *string `json:"privilege,omitempty"`
}

type FindGroupsByID200ApplicationJSONPrivilegesCasperImaging struct {
	Privilege *string `json:"privilege,omitempty"`
}

type FindGroupsByID200ApplicationJSONPrivilegesCasperRemote struct {
	Privilege *string `json:"privilege,omitempty"`
}

type FindGroupsByID200ApplicationJSONPrivilegesJssActions struct {
	Privilege *string `json:"privilege,omitempty"`
}

type FindGroupsByID200ApplicationJSONPrivilegesJssObjects struct {
	Privilege *string `json:"privilege,omitempty"`
}

type FindGroupsByID200ApplicationJSONPrivilegesJssSettings struct {
	Privilege *string `json:"privilege,omitempty"`
}

type FindGroupsByID200ApplicationJSONPrivilegesRecon struct {
	Privilege *string `json:"privilege,omitempty"`
}

type FindGroupsByID200ApplicationJSONPrivileges struct {
	CasperAdmin   []FindGroupsByID200ApplicationJSONPrivilegesCasperAdmin   `json:"casper_admin,omitempty"`
	CasperImaging []FindGroupsByID200ApplicationJSONPrivilegesCasperImaging `json:"casper_imaging,omitempty"`
	CasperRemote  []FindGroupsByID200ApplicationJSONPrivilegesCasperRemote  `json:"casper_remote,omitempty"`
	JssActions    []FindGroupsByID200ApplicationJSONPrivilegesJssActions    `json:"jss_actions,omitempty"`
	JssObjects    []FindGroupsByID200ApplicationJSONPrivilegesJssObjects    `json:"jss_objects,omitempty"`
	JssSettings   []FindGroupsByID200ApplicationJSONPrivilegesJssSettings   `json:"jss_settings,omitempty"`
	Recon         []FindGroupsByID200ApplicationJSONPrivilegesRecon         `json:"recon,omitempty"`
}

type FindGroupsByID200ApplicationJSONSite struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the site
	Name string `json:"name"`
}

// FindGroupsByID200ApplicationJSON - OK
type FindGroupsByID200ApplicationJSON struct {
	AccessLevel *FindGroupsByID200ApplicationJSONAccessLevel `json:"access_level,omitempty"`
	ID          *int64                                       `json:"id,omitempty"`
	Members     []FindGroupsByID200ApplicationJSONMembers    `json:"members,omitempty"`
	// Group name
	Name         string                                        `json:"name"`
	PrivilegeSet *FindGroupsByID200ApplicationJSONPrivilegeSet `json:"privilege_set,omitempty"`
	Privileges   *FindGroupsByID200ApplicationJSONPrivileges   `json:"privileges,omitempty"`
	Site         *FindGroupsByID200ApplicationJSONSite         `json:"site,omitempty"`
}

type FindGroupsByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindGroupsByID200ApplicationJSONObject *FindGroupsByID200ApplicationJSON
}
