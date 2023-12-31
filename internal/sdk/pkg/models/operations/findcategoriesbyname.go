// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindCategoriesByNameRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

// FindCategoriesByNameCategory - OK
type FindCategoriesByNameCategory struct {
	ID *int64 `json:"id,omitempty"`
	// Name of the category
	Name     string `json:"name"`
	Priority *int64 `json:"priority,omitempty"`
}

type FindCategoriesByNameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	Category *FindCategoriesByNameCategory
}
