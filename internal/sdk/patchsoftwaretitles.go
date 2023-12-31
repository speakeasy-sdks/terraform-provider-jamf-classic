// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"jamf/internal/sdk/pkg/models/operations"
	"jamf/internal/sdk/pkg/utils"
	"net/http"
	"strings"
)

type patchsoftwaretitles struct {
	sdkConfiguration sdkConfiguration
}

func newPatchsoftwaretitles(sdkConfig sdkConfiguration) *patchsoftwaretitles {
	return &patchsoftwaretitles{
		sdkConfiguration: sdkConfig,
	}
}

// PatchsoftwaretitlesGet - Finds all patch software titles. (Deprecated) Please transition use to Jamf Pro API endpoint "/v2/patch-software-title-configurations".
//
// Deprecated: this method will be removed in a future release, please migrate away from it as soon as possible.
func (s *patchsoftwaretitles) PatchsoftwaretitlesGet(ctx context.Context) (*operations.PatchsoftwaretitlesGetResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/patchsoftwaretitles"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/xml;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PatchsoftwaretitlesGetResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []operations.PatchsoftwaretitlesGet200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PatchsoftwaretitlesGet200ApplicationJSONObjects = out
		case utils.MatchContentType(contentType, `application/xml`):
			res.Body = rawBody
		}
	}

	return res, nil
}

// PatchsoftwaretitlesIDByIDDelete - Deletes a patch software title by ID. (Deprecated) Please transition use to Jamf Pro API endpoint "/v2/patch-software-title-configurations/{id}".
//
// Deprecated: this method will be removed in a future release, please migrate away from it as soon as possible.
func (s *patchsoftwaretitles) PatchsoftwaretitlesIDByIDDelete(ctx context.Context, request operations.PatchsoftwaretitlesIDByIDDeleteRequest) (*operations.PatchsoftwaretitlesIDByIDDeleteResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/patchsoftwaretitles/id/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PatchsoftwaretitlesIDByIDDeleteResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// PatchsoftwaretitlesIDByIDGet - Finds a patch software titles by ID. (Deprecated) Please transition use to Jamf Pro API endpoint "/v2/patch-software-title-configurations/{id}".
//
// Deprecated: this method will be removed in a future release, please migrate away from it as soon as possible.
func (s *patchsoftwaretitles) PatchsoftwaretitlesIDByIDGet(ctx context.Context, request operations.PatchsoftwaretitlesIDByIDGetRequest) (*operations.PatchsoftwaretitlesIDByIDGetResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/patchsoftwaretitles/id/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PatchsoftwaretitlesIDByIDGetResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.PatchsoftwaretitlesIDByIDGet200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.PatchsoftwaretitlesIDByIDGet200ApplicationJSONObject = out
		}
	}

	return res, nil
}

// PatchsoftwaretitlesIDByIDPost - Creates new patch software title by ID. (Deprecated) Please transition use to Jamf Pro API endpoint "/v2/patch-software-title-configurations/{id}".
//
// Deprecated: this method will be removed in a future release, please migrate away from it as soon as possible.
func (s *patchsoftwaretitles) PatchsoftwaretitlesIDByIDPost(ctx context.Context, request operations.PatchsoftwaretitlesIDByIDPostRequest) (*operations.PatchsoftwaretitlesIDByIDPostResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/patchsoftwaretitles/id/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PatchsoftwaretitlesIDByIDPostResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
	}

	return res, nil
}

// PatchsoftwaretitlesIDByIDPut - Updates a patch software title by ID. (Deprecated) Please transition use to Jamf Pro API endpoint "/v2/patch-software-title-configurations/{id}".
// Including an empty package element <package/> within the version element will disassociate any previously associated packages with that version of the software title.
//
// Deprecated: this method will be removed in a future release, please migrate away from it as soon as possible.
func (s *patchsoftwaretitles) PatchsoftwaretitlesIDByIDPut(ctx context.Context, request operations.PatchsoftwaretitlesIDByIDPutRequest) (*operations.PatchsoftwaretitlesIDByIDPutResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/patchsoftwaretitles/id/{id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PatchsoftwaretitlesIDByIDPutResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
	}

	return res, nil
}
