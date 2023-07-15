// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AccountAccessLevel string

const (
	AccountAccessLevelFullAccess  AccountAccessLevel = "Full Access"
	AccountAccessLevelSiteAccess  AccountAccessLevel = "Site Access"
	AccountAccessLevelGroupAccess AccountAccessLevel = "Group Access"
)

func (e AccountAccessLevel) ToPointer() *AccountAccessLevel {
	return &e
}

func (e *AccountAccessLevel) UnmarshalJSON(data []byte) error {
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
		*e = AccountAccessLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountAccessLevel: %v", v)
	}
}

type AccountEnabled string

const (
	AccountEnabledEnabled  AccountEnabled = "Enabled"
	AccountEnabledDisabled AccountEnabled = "Disabled"
)

func (e AccountEnabled) ToPointer() *AccountEnabled {
	return &e
}

func (e *AccountEnabled) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Enabled":
		fallthrough
	case "Disabled":
		*e = AccountEnabled(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountEnabled: %v", v)
	}
}

type AccountLdapServer struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type AccountPrivilegeSet string

const (
	AccountPrivilegeSetAdministrator  AccountPrivilegeSet = "Administrator"
	AccountPrivilegeSetAuditor        AccountPrivilegeSet = "Auditor"
	AccountPrivilegeSetEnrollmentOnly AccountPrivilegeSet = "Enrollment Only"
	AccountPrivilegeSetCustom         AccountPrivilegeSet = "Custom"
)

func (e AccountPrivilegeSet) ToPointer() *AccountPrivilegeSet {
	return &e
}

func (e *AccountPrivilegeSet) UnmarshalJSON(data []byte) error {
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
		*e = AccountPrivilegeSet(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountPrivilegeSet: %v", v)
	}
}

type AccountPrivilegesCasperAdmin struct {
	Privilege *string `json:"privilege,omitempty"`
}

type AccountPrivilegesCasperImaging struct {
	Privilege *string `json:"privilege,omitempty"`
}

type AccountPrivilegesCasperRemote struct {
	Privilege *string `json:"privilege,omitempty"`
}

type AccountPrivilegesJssActions struct {
	Privilege *string `json:"privilege,omitempty"`
}

type AccountPrivilegesJssObjects struct {
	Privilege *string `json:"privilege,omitempty"`
}

type AccountPrivilegesJssSettings struct {
	Privilege *string `json:"privilege,omitempty"`
}

type AccountPrivilegesRecon struct {
	Privilege *string `json:"privilege,omitempty"`
}

type AccountPrivileges struct {
	CasperAdmin   []AccountPrivilegesCasperAdmin   `json:"casper_admin,omitempty"`
	CasperImaging []AccountPrivilegesCasperImaging `json:"casper_imaging,omitempty"`
	CasperRemote  []AccountPrivilegesCasperRemote  `json:"casper_remote,omitempty"`
	JssActions    []AccountPrivilegesJssActions    `json:"jss_actions,omitempty"`
	JssObjects    []AccountPrivilegesJssObjects    `json:"jss_objects,omitempty"`
	JssSettings   []AccountPrivilegesJssSettings   `json:"jss_settings,omitempty"`
	Recon         []AccountPrivilegesRecon         `json:"recon,omitempty"`
}

// Account - OK
type Account struct {
	AccessLevel         *AccountAccessLevel `json:"access_level,omitempty"`
	DirectoryUser       *bool               `json:"directory_user,omitempty"`
	Email               *string             `json:"email,omitempty"`
	EmailAddress        *string             `json:"email_address,omitempty"`
	Enabled             *AccountEnabled     `json:"enabled,omitempty"`
	ForcePasswordChange *bool               `json:"force_password_change,omitempty"`
	FullName            *string             `json:"full_name,omitempty"`
	ID                  *int64              `json:"id,omitempty"`
	LdapServer          *AccountLdapServer  `json:"ldap_server,omitempty"`
	// Name of the account
	Name         string               `json:"name"`
	PrivilegeSet *AccountPrivilegeSet `json:"privilege_set,omitempty"`
	Privileges   *AccountPrivileges   `json:"privileges,omitempty"`
	Site         *SiteObject          `json:"site,omitempty"`
}