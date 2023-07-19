// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type FindMobileDevicesByNameRequest struct {
	// Name to filter by
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

type FindMobileDevicesByName200ApplicationXMLApplicationsApplication struct {
	ApplicationName    *string
	ApplicationVersion *string
	Identifier         *string
}

type FindMobileDevicesByName200ApplicationXMLApplications struct {
	Application *FindMobileDevicesByName200ApplicationXMLApplicationsApplication
	Size        *int64
}

type FindMobileDevicesByName200ApplicationXMLCertificatesCertificate struct {
	CommonName *string
	Identity   *bool
}

type FindMobileDevicesByName200ApplicationXMLCertificates struct {
	Certificate *FindMobileDevicesByName200ApplicationXMLCertificatesCertificate
	Size        *int64
}

type FindMobileDevicesByName200ApplicationXMLConfigurationProfilesConfigurationProfile struct {
	DisplayName *string
	Identifier  *string
	UUID        *string
	Version     *int64
}

type FindMobileDevicesByName200ApplicationXMLConfigurationProfiles struct {
	ConfigurationProfile *FindMobileDevicesByName200ApplicationXMLConfigurationProfilesConfigurationProfile
	Size                 *int64
}

type FindMobileDevicesByName200ApplicationXMLExtensionAttributes struct {
	ID    *int64
	Name  *string
	Type  *string
	Value *string
}

type FindMobileDevicesByName200ApplicationXMLGeneral struct {
	AssetTag                           *string
	Available                          *int64
	AvailableMb                        *int64
	BatteryLevel                       *int64
	BleCapable                         *bool
	BluetoothMacAddress                *string
	Capacity                           *int64
	CapacityMb                         *int64
	CloudBackupEnabled                 *bool
	DeviceLocatorServiceEnabled        *bool
	DeviceName                         *string
	DeviceOwnershipLevel               *string
	DisplayName                        *string
	DoNotDisturbEnabled                *bool
	ExchangeActivesyncDeviceIdentifier *string
	ID                                 *int64
	InitialEntryDateEpoch              *int64
	InitialEntryDateUtc                *string
	IPAddress                          *string
	ItunesStoreAccountIsActive         *bool
	LastBackupTimeEpoch                *int64
	LastBackupTimeUtc                  *string
	LastCloudBackupDateEpoch           *int64
	LastCloudBackupDateUtc             *string
	LastEnrollmentEpoch                *int64
	LastEnrollmentUtc                  *string
	LastInventoryUpdate                *string
	LastInventoryUpdateEpoch           *int64
	LastInventoryUpdateUtc             *string
	LocationServicesEnabled            *bool
	Managed                            *bool
	Model                              *string
	ModelDisplay                       *string
	ModelIdentifier                    *string
	ModelNumber                        *string
	ModemFirmware                      *string
	Name                               string
	OsBuild                            *string
	OsType                             *string
	OsVersion                          *string
	PercentageUsed                     *int64
	PhoneNumber                        *string
	SerialNumber                       string
	Shared                             *string
	Supervised                         *bool
	Tethered                           *string
	Udid                               string
	WifiMacAddress                     *string
}

type FindMobileDevicesByName200ApplicationXMLLocation struct {
	Building     *string
	Department   *string
	EmailAddress *string
	Phone        *string
	PhoneNumber  *string
	Position     *string
	RealName     *string
	Realname     *string
	Room         *string
	Username     *string
}

type FindMobileDevicesByName200ApplicationXMLMobileDeviceGroupsMobileDeviceGroup struct {
	ID   *int64
	Name *string
}

type FindMobileDevicesByName200ApplicationXMLMobileDeviceGroups struct {
	MobileDeviceGroup *FindMobileDevicesByName200ApplicationXMLMobileDeviceGroupsMobileDeviceGroup
	Size              *int64
}

type FindMobileDevicesByName200ApplicationXMLNetwork struct {
	CarrierSettingsVersion   *string
	CellularTechnology       *string
	CurrentCarrierNetwork    *string
	CurrentMobileCountryCode *string
	CurrentMobileNetworkCode *string
	DataRoamingEnabled       *bool
	HomeCarrierNetwork       *string
	HomeMobileCountryCode    *string
	HomeMobileNetworkCode    *string
	Iccid                    *string
	Imei                     *string
	PhoneNumber              *string
	VoiceRoamingEnabled      *string
}

type FindMobileDevicesByName200ApplicationXMLProvisioningProfilesMobileDeviceProvisioningProfile struct {
	DisplayName         *string
	ExpirationDate      *string
	ExpirationDateEpoch *int64
	ExpirationDateUtc   *string
	UUID                *string
}

type FindMobileDevicesByName200ApplicationXMLProvisioningProfiles struct {
	MobileDeviceProvisioningProfile *FindMobileDevicesByName200ApplicationXMLProvisioningProfilesMobileDeviceProvisioningProfile
	Size                            *int64
}

type FindMobileDevicesByName200ApplicationXMLPurchasing struct {
	ApplecareID          *string
	IsLeased             *bool
	IsPurchased          *bool
	LeaseExpires         *string
	LeaseExpiresEpoch    *int64
	LeaseExpiresUtc      *string
	LifeExpectancy       *int64
	PoDate               *string
	PoDateEpoch          *int64
	PoDateUtc            *string
	PoNumber             *string
	PurchasePrice        *string
	PurchasingAccount    *string
	PurchasingContact    *string
	Vendor               *string
	WarrantyExpires      *string
	WarrantyExpiresEpoch *int64
	WarrantyExpiresUtc   *string
}

type FindMobileDevicesByName200ApplicationXMLSecurityObject struct {
	ActivationLockEnabled           *bool
	BlockLevelEncryptionCapable     *bool
	DataProtection                  *bool
	FileLevelEncryptionCapable      *bool
	HardwareEncryption              *string
	JailbreakDetected               *string
	LostLocationAltitude            *string
	LostLocationCourse              *string
	LostLocationEpoch               *int64
	LostLocationHorizontalAccuracy  *string
	LostLocationLatitude            *string
	LostLocationLongitude           *string
	LostLocationSpeed               *string
	LostLocationUtc                 *string
	LostLocationVerticalAccuracy    *string
	LostModeEnableIssuedEpoch       *int64
	LostModeEnableIssuedUtc         *string
	LostModeEnabled                 *string
	LostModeEnforced                *bool
	LostModeFootnote                *string
	LostModeMessage                 *string
	LostModePhone                   *string
	PasscodeCompliant               *bool
	PasscodeCompliantWithProfile    *bool
	PasscodeLockGracePeriodEnforced *string
	PasscodePresent                 *bool
}

// FindMobileDevicesByName200ApplicationXML - OK
type FindMobileDevicesByName200ApplicationXML struct {
	Applications          []FindMobileDevicesByName200ApplicationXMLApplications
	Certificates          []FindMobileDevicesByName200ApplicationXMLCertificates
	ConfigurationProfiles []FindMobileDevicesByName200ApplicationXMLConfigurationProfiles
	ExtensionAttributes   []FindMobileDevicesByName200ApplicationXMLExtensionAttributes
	General               *FindMobileDevicesByName200ApplicationXMLGeneral
	Location              *FindMobileDevicesByName200ApplicationXMLLocation
	MobileDeviceGroups    []FindMobileDevicesByName200ApplicationXMLMobileDeviceGroups
	Network               *FindMobileDevicesByName200ApplicationXMLNetwork
	ProvisioningProfiles  []FindMobileDevicesByName200ApplicationXMLProvisioningProfiles
	Purchasing            *FindMobileDevicesByName200ApplicationXMLPurchasing
	SecurityObject        *FindMobileDevicesByName200ApplicationXMLSecurityObject
}

type FindMobileDevicesByName200ApplicationJSONApplicationsApplication struct {
	ApplicationName    *string `json:"application_name,omitempty"`
	ApplicationVersion *string `json:"application_version,omitempty"`
	Identifier         *string `json:"identifier,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONApplications struct {
	Application *FindMobileDevicesByName200ApplicationJSONApplicationsApplication `json:"application,omitempty"`
	Size        *int64                                                            `json:"size,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONCertificatesCertificate struct {
	CommonName *string `json:"common_name,omitempty"`
	Identity   *bool   `json:"identity,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONCertificates struct {
	Certificate *FindMobileDevicesByName200ApplicationJSONCertificatesCertificate `json:"certificate,omitempty"`
	Size        *int64                                                            `json:"size,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONConfigurationProfilesConfigurationProfile struct {
	DisplayName *string `json:"display_name,omitempty"`
	Identifier  *string `json:"identifier,omitempty"`
	UUID        *string `json:"uuid,omitempty"`
	Version     *int64  `json:"version,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONConfigurationProfiles struct {
	ConfigurationProfile *FindMobileDevicesByName200ApplicationJSONConfigurationProfilesConfigurationProfile `json:"configuration_profile,omitempty"`
	Size                 *int64                                                                              `json:"size,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONExtensionAttributes struct {
	ID    *int64  `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONGeneral struct {
	AssetTag                           *string `json:"asset_tag,omitempty"`
	Available                          *int64  `json:"available,omitempty"`
	AvailableMb                        *int64  `json:"available_mb,omitempty"`
	BatteryLevel                       *int64  `json:"battery_level,omitempty"`
	BleCapable                         *bool   `json:"ble_capable,omitempty"`
	BluetoothMacAddress                *string `json:"bluetooth_mac_address,omitempty"`
	Capacity                           *int64  `json:"capacity,omitempty"`
	CapacityMb                         *int64  `json:"capacity_mb,omitempty"`
	CloudBackupEnabled                 *bool   `json:"cloud_backup_enabled,omitempty"`
	DeviceLocatorServiceEnabled        *bool   `json:"device_locator_service_enabled,omitempty"`
	DeviceName                         *string `json:"device_name,omitempty"`
	DeviceOwnershipLevel               *string `json:"device_ownership_level,omitempty"`
	DisplayName                        *string `json:"display_name,omitempty"`
	DoNotDisturbEnabled                *bool   `json:"do_not_disturb_enabled,omitempty"`
	ExchangeActivesyncDeviceIdentifier *string `json:"exchange_activesync_device_identifier,omitempty"`
	ID                                 *int64  `json:"id,omitempty"`
	InitialEntryDateEpoch              *int64  `json:"initial_entry_date_epoch,omitempty"`
	InitialEntryDateUtc                *string `json:"initial_entry_date_utc,omitempty"`
	IPAddress                          *string `json:"ip_address,omitempty"`
	ItunesStoreAccountIsActive         *bool   `json:"itunes_store_account_is_active,omitempty"`
	LastBackupTimeEpoch                *int64  `json:"last_backup_time_epoch,omitempty"`
	LastBackupTimeUtc                  *string `json:"last_backup_time_utc,omitempty"`
	LastCloudBackupDateEpoch           *int64  `json:"last_cloud_backup_date_epoch,omitempty"`
	LastCloudBackupDateUtc             *string `json:"last_cloud_backup_date_utc,omitempty"`
	LastEnrollmentEpoch                *int64  `json:"last_enrollment_epoch,omitempty"`
	LastEnrollmentUtc                  *string `json:"last_enrollment_utc,omitempty"`
	LastInventoryUpdate                *string `json:"last_inventory_update,omitempty"`
	LastInventoryUpdateEpoch           *int64  `json:"last_inventory_update_epoch,omitempty"`
	LastInventoryUpdateUtc             *string `json:"last_inventory_update_utc,omitempty"`
	LocationServicesEnabled            *bool   `json:"location_services_enabled,omitempty"`
	Managed                            *bool   `json:"managed,omitempty"`
	Model                              *string `json:"model,omitempty"`
	ModelDisplay                       *string `json:"model_display,omitempty"`
	ModelIdentifier                    *string `json:"model_identifier,omitempty"`
	ModelNumber                        *string `json:"model_number,omitempty"`
	ModemFirmware                      *string `json:"modem_firmware,omitempty"`
	Name                               string  `json:"name"`
	OsBuild                            *string `json:"os_build,omitempty"`
	OsType                             *string `json:"os_type,omitempty"`
	OsVersion                          *string `json:"os_version,omitempty"`
	PercentageUsed                     *int64  `json:"percentage_used,omitempty"`
	PhoneNumber                        *string `json:"phone_number,omitempty"`
	SerialNumber                       string  `json:"serial_number"`
	Shared                             *string `json:"shared,omitempty"`
	Supervised                         *bool   `json:"supervised,omitempty"`
	Tethered                           *string `json:"tethered,omitempty"`
	Udid                               string  `json:"udid"`
	WifiMacAddress                     *string `json:"wifi_mac_address,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONLocation struct {
	Building     *string `json:"building,omitempty"`
	Department   *string `json:"department,omitempty"`
	EmailAddress *string `json:"email_address,omitempty"`
	Phone        *string `json:"phone,omitempty"`
	PhoneNumber  *string `json:"phone_number,omitempty"`
	Position     *string `json:"position,omitempty"`
	RealName     *string `json:"real_name,omitempty"`
	Realname     *string `json:"realname,omitempty"`
	Room         *string `json:"room,omitempty"`
	Username     *string `json:"username,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONMobileDeviceGroupsMobileDeviceGroup struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONMobileDeviceGroups struct {
	MobileDeviceGroup *FindMobileDevicesByName200ApplicationJSONMobileDeviceGroupsMobileDeviceGroup `json:"mobile_device_group,omitempty"`
	Size              *int64                                                                        `json:"size,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONNetwork struct {
	CarrierSettingsVersion   *string `json:"carrier_settings_version,omitempty"`
	CellularTechnology       *string `json:"cellular_technology,omitempty"`
	CurrentCarrierNetwork    *string `json:"current_carrier_network,omitempty"`
	CurrentMobileCountryCode *string `json:"current_mobile_country_code,omitempty"`
	CurrentMobileNetworkCode *string `json:"current_mobile_network_code,omitempty"`
	DataRoamingEnabled       *bool   `json:"data_roaming_enabled,omitempty"`
	HomeCarrierNetwork       *string `json:"home_carrier_network,omitempty"`
	HomeMobileCountryCode    *string `json:"home_mobile_country_code,omitempty"`
	HomeMobileNetworkCode    *string `json:"home_mobile_network_code,omitempty"`
	Iccid                    *string `json:"iccid,omitempty"`
	Imei                     *string `json:"imei,omitempty"`
	PhoneNumber              *string `json:"phone_number,omitempty"`
	VoiceRoamingEnabled      *string `json:"voice_roaming_enabled,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONProvisioningProfilesMobileDeviceProvisioningProfile struct {
	DisplayName         *string `json:"display_name,omitempty"`
	ExpirationDate      *string `json:"expiration_date,omitempty"`
	ExpirationDateEpoch *int64  `json:"expiration_date_epoch,omitempty"`
	ExpirationDateUtc   *string `json:"expiration_date_utc,omitempty"`
	UUID                *string `json:"uuid,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONProvisioningProfiles struct {
	MobileDeviceProvisioningProfile *FindMobileDevicesByName200ApplicationJSONProvisioningProfilesMobileDeviceProvisioningProfile `json:"mobile_device_provisioning_profile,omitempty"`
	Size                            *int64                                                                                        `json:"size,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONPurchasing struct {
	ApplecareID          *string `json:"applecare_id,omitempty"`
	IsLeased             *bool   `json:"is_leased,omitempty"`
	IsPurchased          *bool   `json:"is_purchased,omitempty"`
	LeaseExpires         *string `json:"lease_expires,omitempty"`
	LeaseExpiresEpoch    *int64  `json:"lease_expires_epoch,omitempty"`
	LeaseExpiresUtc      *string `json:"lease_expires_utc,omitempty"`
	LifeExpectancy       *int64  `json:"life_expectancy,omitempty"`
	PoDate               *string `json:"po_date,omitempty"`
	PoDateEpoch          *int64  `json:"po_date_epoch,omitempty"`
	PoDateUtc            *string `json:"po_date_utc,omitempty"`
	PoNumber             *string `json:"po_number,omitempty"`
	PurchasePrice        *string `json:"purchase_price,omitempty"`
	PurchasingAccount    *string `json:"purchasing_account,omitempty"`
	PurchasingContact    *string `json:"purchasing_contact,omitempty"`
	Vendor               *string `json:"vendor,omitempty"`
	WarrantyExpires      *string `json:"warranty_expires,omitempty"`
	WarrantyExpiresEpoch *int64  `json:"warranty_expires_epoch,omitempty"`
	WarrantyExpiresUtc   *string `json:"warranty_expires_utc,omitempty"`
}

type FindMobileDevicesByName200ApplicationJSONSecurityObject struct {
	ActivationLockEnabled           *bool   `json:"activation_lock_enabled,omitempty"`
	BlockLevelEncryptionCapable     *bool   `json:"block_level_encryption_capable,omitempty"`
	DataProtection                  *bool   `json:"data_protection,omitempty"`
	FileLevelEncryptionCapable      *bool   `json:"file_level_encryption_capable,omitempty"`
	HardwareEncryption              *string `json:"hardware_encryption,omitempty"`
	JailbreakDetected               *string `json:"jailbreak_detected,omitempty"`
	LostLocationAltitude            *string `json:"lost_location_altitude,omitempty"`
	LostLocationCourse              *string `json:"lost_location_course,omitempty"`
	LostLocationEpoch               *int64  `json:"lost_location_epoch,omitempty"`
	LostLocationHorizontalAccuracy  *string `json:"lost_location_horizontal_accuracy,omitempty"`
	LostLocationLatitude            *string `json:"lost_location_latitude,omitempty"`
	LostLocationLongitude           *string `json:"lost_location_longitude,omitempty"`
	LostLocationSpeed               *string `json:"lost_location_speed,omitempty"`
	LostLocationUtc                 *string `json:"lost_location_utc,omitempty"`
	LostLocationVerticalAccuracy    *string `json:"lost_location_vertical_accuracy,omitempty"`
	LostModeEnableIssuedEpoch       *int64  `json:"lost_mode_enable_issued_epoch,omitempty"`
	LostModeEnableIssuedUtc         *string `json:"lost_mode_enable_issued_utc,omitempty"`
	LostModeEnabled                 *string `json:"lost_mode_enabled,omitempty"`
	LostModeEnforced                *bool   `json:"lost_mode_enforced,omitempty"`
	LostModeFootnote                *string `json:"lost_mode_footnote,omitempty"`
	LostModeMessage                 *string `json:"lost_mode_message,omitempty"`
	LostModePhone                   *string `json:"lost_mode_phone,omitempty"`
	PasscodeCompliant               *bool   `json:"passcode_compliant,omitempty"`
	PasscodeCompliantWithProfile    *bool   `json:"passcode_compliant_with_profile,omitempty"`
	PasscodeLockGracePeriodEnforced *string `json:"passcode_lock_grace_period_enforced,omitempty"`
	PasscodePresent                 *bool   `json:"passcode_present,omitempty"`
}

// FindMobileDevicesByName200ApplicationJSON - OK
type FindMobileDevicesByName200ApplicationJSON struct {
	Applications          []FindMobileDevicesByName200ApplicationJSONApplications          `json:"applications,omitempty"`
	Certificates          []FindMobileDevicesByName200ApplicationJSONCertificates          `json:"certificates,omitempty"`
	ConfigurationProfiles []FindMobileDevicesByName200ApplicationJSONConfigurationProfiles `json:"configuration_profiles,omitempty"`
	ExtensionAttributes   []FindMobileDevicesByName200ApplicationJSONExtensionAttributes   `json:"extension_attributes,omitempty"`
	General               *FindMobileDevicesByName200ApplicationJSONGeneral                `json:"general,omitempty"`
	Location              *FindMobileDevicesByName200ApplicationJSONLocation               `json:"location,omitempty"`
	MobileDeviceGroups    []FindMobileDevicesByName200ApplicationJSONMobileDeviceGroups    `json:"mobile_device_groups,omitempty"`
	Network               *FindMobileDevicesByName200ApplicationJSONNetwork                `json:"network,omitempty"`
	ProvisioningProfiles  []FindMobileDevicesByName200ApplicationJSONProvisioningProfiles  `json:"provisioning_profiles,omitempty"`
	Purchasing            *FindMobileDevicesByName200ApplicationJSONPurchasing             `json:"purchasing,omitempty"`
	SecurityObject        *FindMobileDevicesByName200ApplicationJSONSecurityObject         `json:"security_object,omitempty"`
}

type FindMobileDevicesByNameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	FindMobileDevicesByName200ApplicationJSONObject *FindMobileDevicesByName200ApplicationJSON
}
