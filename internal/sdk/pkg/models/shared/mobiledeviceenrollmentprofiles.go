// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MobileDeviceEnrollmentProfilesMobileDeviceEnrollmentProfile struct {
	ID         *int64  `json:"id,omitempty"`
	Invitation *string `json:"invitation,omitempty"`
	// Name of the enrollment profile
	Name *string `json:"name,omitempty"`
}

type MobileDeviceEnrollmentProfiles struct {
	MobileDeviceEnrollmentProfile *MobileDeviceEnrollmentProfilesMobileDeviceEnrollmentProfile `json:"mobile_device_enrollment_profile,omitempty"`
	Size                          *int64                                                       `json:"size,omitempty"`
}