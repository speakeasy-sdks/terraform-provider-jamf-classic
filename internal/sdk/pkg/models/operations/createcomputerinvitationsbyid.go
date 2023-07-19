// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CreateComputerInvitationsByIDRequest struct {
	// ID value to filter by
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type CreateComputerInvitationsByID201ApplicationXMLEnrollIntoSite struct {
	ID   *int64
	Name *string
}

type CreateComputerInvitationsByID201ApplicationXMLSite struct {
	ID *int64
	// Name of the site
	Name string
}

// CreateComputerInvitationsByID201ApplicationXML - Created
type CreateComputerInvitationsByID201ApplicationXML struct {
	CreateAccountIfDoesNotExist *bool
	EnrollIntoSite              *CreateComputerInvitationsByID201ApplicationXMLEnrollIntoSite
	// Use 'Unlimited' to specify no expiration date
	ExpirationDate             *string
	ExpirationDateEpoch        *int64
	ExpirationDateUtc          *string
	HideAccount                *bool
	ID                         *int64
	Invitation                 *int64
	InvitationStatus           *string
	InvitationType             *string
	InvitedUserUUID            *string
	KeepExistingSiteMembership *bool
	LockDownSSH                *bool
	MultipleUsersAllowed       *bool
	Site                       *CreateComputerInvitationsByID201ApplicationXMLSite
	SSHPassword                *string
	SSHUsername                *string
	TimesUsed                  *int64
}

type CreateComputerInvitationsByIDResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
