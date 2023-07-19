// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindPatchPolicies200ApplicationXMLPatchPolicy struct {
	ID   *int64
	Name *string
}

type FindPatchPolicies200ApplicationXML struct {
	PatchPolicy *FindPatchPolicies200ApplicationXMLPatchPolicy
	Size        *int64
}

type FindPatchPolicies200ApplicationJSONPatchPolicy struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindPatchPolicies200ApplicationJSON struct {
	PatchPolicy *FindPatchPolicies200ApplicationJSONPatchPolicy `json:"patch_policy,omitempty"`
	Size        *int64                                          `json:"size,omitempty"`
}

type FindPatchPoliciesResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindPatchPolicies200ApplicationJSONObjects []FindPatchPolicies200ApplicationJSON
}
