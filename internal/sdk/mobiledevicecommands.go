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

type mobiledevicecommands struct {
	sdkConfiguration sdkConfiguration
}

func newMobiledevicecommands(sdkConfig sdkConfiguration) *mobiledevicecommands {
	return &mobiledevicecommands{
		sdkConfiguration: sdkConfig,
	}
}

// CreateMobileDeviceCommand - Creates a new mobile device command
// The chart below includes additional requirements for usage of specific commands
//
// | command | Parameters | Requirements | Values |
// | ---------- | ------- | ------------ | ------ |
// | DeviceName | device_name | Supervised Device | string |
// | DeviceLocation | N/A | Supervised and in lost mode | N/A |
// | DeviceLock | lock_message | optional | string |
// | DisableLostMode | N/A | Supervised and in lost mode | N/A |
// | EnableLostMode | lost_mode_message | Supervised device (required if lost_mode_phone is not set) | string |
// | EnableLostMode | lost_mode_phone | Supervised device (required if lost_mode_message is not set) | string |
// | EnableLostMode | lost_mode_footnote | optional | string |
// | EnableLostMode | always_enforce_lost_mode | optional (defaults to 'true') | boolean |
// | EnableLostMode | lost_mode_with_sound | optional (defaults to 'false') | boolean |
// | EraseDevice | preserve_data_plan | optional (defaults to 'false') | boolean |
// | EraseDevice | disallow_proximity_setup | optional (defaults to 'false') | boolean |
// | EraseDevice | clear_activation_lock | optional (defaults to 'false') | boolean |
// | PasscodeLockGracePeriod | passcode_lock_grace_period | Shared iPad | seconds (integer) |
// | PlayLostModeSound | N/A | Supervised and in lost mode | N/A |
// | RestartDevice | N/A | Supervised device | N/A |
// | ScheduleOSUpdate (deprecated on 2022-10-17) | install_action | iOS 9 - 10.2, enrolled via a Prestage enrollment; and/or iOS 10.3 or later; tvOS 12 or later | 1 = Download the update for users to install, 2 = Download and install the update, and restart devices after installation |
// | ScheduleOSUpdate (deprecated on 2022-10-17) | product_version | iOS 11.3 or later, tvOS 12.2 or later | string |
// | SettingsDisableBluetooth | N/A | iOS 11.3+ and Supervised | N/A |
// | SettingsEnableBluetooth | N/A | iOS 11.3+ and Supervised | N/A |
// | ShutDownDevice | N/A | Supervised device | N/A |
// | Wallpaper | wallpaper_settings | Supervised device | 1 = Lock screen, 2 = Home screen, 3 = both |
// | Wallpaper | wallpaper_id | Supervised device (required if wallpaper_content is not set) | Jamf Pro icon ID (integer) |
// | Wallpaper | wallpaper_content | Supervised device (required if wallpaper_id is not set) | base64 encoded PNG or JPEG |
// | RefreshCellularPlans | e_sim_server_url | N/A | This URL is obtained from each carrier separately |
func (s *mobiledevicecommands) CreateMobileDeviceCommand(ctx context.Context, request []byte) (*operations.CreateMobileDeviceCommandResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/mobiledevicecommands/command"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "raw")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

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

	res := &operations.CreateMobileDeviceCommandResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
	}

	return res, nil
}

// CreateMobileDeviceCommandURL - Creates a new mobile device command
// The chart below includes additional requirements for usage of specific commands
//
// | command | Parameters | Requirements | Values |
// | ---------- | ------- | ------------ | ------ |
// | DeviceLocation | N/A | Supervised and in lost mode | N/A |
// | DisableLostMode | N/A | Supervised and in lost mode | N/A |
// | EnableLostMode | lost_mode_message | Supervised device (required if lost_mode_phone is not set) | string |
// | EnableLostMode | lost_mode_phone | Supervised device (required if lost_mode_message is not set) | string |
// | EnableLostMode | lost_mode_footnote | optional | string |
// | EnableLostMode | always_enforce_lost_mode | optional (defaults to 'true') | boolean |
// | EnableLostMode | lost_mode_with_sound | optional (defaults to 'false') | boolean |
// | EraseDevice | preserve_data_plan | optional (defaults to 'false') | boolean |
// | EraseDevice | disallow_proximity_setup | optional (defaults to 'false') | boolean |
// | PasscodeLockGracePeriod | passcode_lock_grace_period | Shared iPad | seconds (integer) |
// | PlayLostModeSound | N/A | Supervised and in lost mode | N/A |
// | RestartDevice | N/A | Supervised device | N/A |
// | SettingsDisableBluetooth | N/A | iOS 11.3+ and Supervised | N/A |
// | SettingsEnableBluetooth | N/A | iOS 11.3+ and Supervised | N/A |
// | ShutDownDevice | N/A | Supervised device | N/A |
func (s *mobiledevicecommands) CreateMobileDeviceCommandURL(ctx context.Context, request operations.CreateMobileDeviceCommandURLRequest) (*operations.CreateMobileDeviceCommandURLResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/mobiledevicecommands/command/{command}/id/{id_list}", request, nil)
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

	res := &operations.CreateMobileDeviceCommandURLResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
	}

	return res, nil
}

// CreateMobileDeviceLockCommandURL - Sends a new lock command to a list of mobile devices
func (s *mobiledevicecommands) CreateMobileDeviceLockCommandURL(ctx context.Context, request operations.CreateMobileDeviceLockCommandURLRequest) (*operations.CreateMobileDeviceLockCommandURLResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/mobiledevicecommands/command/DeviceLock/{lock_message}/id/{id_list}", request, nil)
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

	res := &operations.CreateMobileDeviceLockCommandURLResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
	}

	return res, nil
}

// CreateMobileDeviceNameCommandURL - Creates a new command to set the name of a mobile device.
func (s *mobiledevicecommands) CreateMobileDeviceNameCommandURL(ctx context.Context, request operations.CreateMobileDeviceNameCommandURLRequest) (*operations.CreateMobileDeviceNameCommandURLResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/mobiledevicecommands/command/DeviceName/{device_name}/id/{id_list}", request, nil)
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

	res := &operations.CreateMobileDeviceNameCommandURLResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
	}

	return res, nil
}

// CreateMobileScheduleOSUpdateCommandURL - Creates a new command to request that a mobile device update its OS. Command and mobile device list specified in URL. Device will be updated to the latest OS version based on device eligibility. (deprecated on 2022-10-17)
func (s *mobiledevicecommands) CreateMobileScheduleOSUpdateCommandURL(ctx context.Context, request operations.CreateMobileScheduleOSUpdateCommandURLRequest) (*operations.CreateMobileScheduleOSUpdateCommandURLResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/mobiledevicecommands/command/ScheduleOSUpdate/{install_action}/id/{id_list}", request, nil)
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

	res := &operations.CreateMobileScheduleOSUpdateCommandURLResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
	}

	return res, nil
}

// CreateMobileScheduleOSUpdateCommandWithProductVersionURL - Creates a new command to request that a mobile device update its OS. Command and mobile device list specified in URL. Mixing iOS and tvOS devices in ID list is not advised, as product version is specific to OS type. (deprecated on 2022-10-17)
func (s *mobiledevicecommands) CreateMobileScheduleOSUpdateCommandWithProductVersionURL(ctx context.Context, request operations.CreateMobileScheduleOSUpdateCommandWithProductVersionURLRequest) (*operations.CreateMobileScheduleOSUpdateCommandWithProductVersionURLResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/mobiledevicecommands/command/ScheduleOSUpdate/{install_action}/{product_version}/id/{id_list}", request, nil)
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

	res := &operations.CreateMobileScheduleOSUpdateCommandWithProductVersionURLResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
	}

	return res, nil
}

// FindMobileDeviceCommands - Finds all mobile device commands
func (s *mobiledevicecommands) FindMobileDeviceCommands(ctx context.Context) (*operations.FindMobileDeviceCommandsResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/mobiledevicecommands"

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

	res := &operations.FindMobileDeviceCommandsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []operations.FindMobileDeviceCommands200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.FindMobileDeviceCommands200ApplicationJSONObjects = out
		case utils.MatchContentType(contentType, `application/xml`):
			res.Body = rawBody
		}
	}

	return res, nil
}

// FindMobileDeviceCommandsByCommand - Finds all mobile device commands for specified command
func (s *mobiledevicecommands) FindMobileDeviceCommandsByCommand(ctx context.Context, request operations.FindMobileDeviceCommandsByCommandRequest) (*operations.FindMobileDeviceCommandsByCommandResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/mobiledevicecommands/command/{command}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

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

	res := &operations.FindMobileDeviceCommandsByCommandResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.FindMobileDeviceCommandsByCommand200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.FindMobileDeviceCommandsByCommand200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/xml`):
			res.Body = rawBody
		}
	}

	return res, nil
}

// FindMobileDeviceCommandsByName - Finds all mobile device commands by command name
func (s *mobiledevicecommands) FindMobileDeviceCommandsByName(ctx context.Context, request operations.FindMobileDeviceCommandsByNameRequest) (*operations.FindMobileDeviceCommandsByNameResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/mobiledevicecommands/name/{name}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

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

	res := &operations.FindMobileDeviceCommandsByNameResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.FindMobileDeviceCommandsByName200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.FindMobileDeviceCommandsByName200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/xml`):
			res.Body = rawBody
		}
	}

	return res, nil
}

// FindMobileDeviceCommandsByUUID - Finds a mobile device command by UUID
func (s *mobiledevicecommands) FindMobileDeviceCommandsByUUID(ctx context.Context, request operations.FindMobileDeviceCommandsByUUIDRequest) (*operations.FindMobileDeviceCommandsByUUIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/mobiledevicecommands/uuid/{uuid}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

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

	res := &operations.FindMobileDeviceCommandsByUUIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.FindMobileDeviceCommandsByUUID200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return nil, err
			}

			res.FindMobileDeviceCommandsByUUID200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/xml`):
			res.Body = rawBody
		}
	}

	return res, nil
}
