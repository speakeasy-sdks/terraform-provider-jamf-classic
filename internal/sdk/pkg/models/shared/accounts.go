// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AccountsGroupsGroup struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the group
	Name string `json:"name"`
}

type AccountsGroups struct {
	Group []AccountsGroupsGroup `json:"group,omitempty"`
}

type AccountsUsersUser struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the account
	Name string `json:"name"`
}

type AccountsUsers struct {
	User []AccountsUsersUser `json:"user,omitempty"`
}

// Accounts - OK
type Accounts struct {
	Groups *AccountsGroups `json:"groups,omitempty"`
	Users  *AccountsUsers  `json:"users,omitempty"`
}