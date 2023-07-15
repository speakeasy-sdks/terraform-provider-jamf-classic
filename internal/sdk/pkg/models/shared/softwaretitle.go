// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SoftwareTitleNotifications struct {
	EmailNotification *bool `json:"email_notification,omitempty"`
	JssNotification   *bool `json:"jss_notification,omitempty"`
}

type SoftwareTitleVersionsVersionComputersComputer struct {
	AltMacAddress *string `json:"alt_mac_address,omitempty"`
	ID            *int64  `json:"id,omitempty"`
	MacAddress    *string `json:"mac_address,omitempty"`
	Name          *string `json:"name,omitempty"`
	SerialNumber  *string `json:"serial_number,omitempty"`
}

type SoftwareTitleVersionsVersionComputers struct {
	Computer *SoftwareTitleVersionsVersionComputersComputer `json:"computer,omitempty"`
}

type SoftwareTitleVersionsVersionPackage struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SoftwareTitleVersionsVersion struct {
	Computers       []SoftwareTitleVersionsVersionComputers `json:"computers,omitempty"`
	Package         *SoftwareTitleVersionsVersionPackage    `json:"package,omitempty"`
	Size            *int64                                  `json:"size,omitempty"`
	SoftwareVersion *string                                 `json:"software_version,omitempty"`
}

type SoftwareTitleVersions struct {
	Version *SoftwareTitleVersionsVersion `json:"version,omitempty"`
}

// SoftwareTitle - OK
type SoftwareTitle struct {
	Category *Category `json:"category,omitempty"`
	ID       *int64    `json:"id,omitempty"`
	// Name of the patch software title
	Name           *string                     `json:"name,omitempty"`
	Notifications  *SoftwareTitleNotifications `json:"notifications,omitempty"`
	TotalComputers *int64                      `json:"total_computers,omitempty"`
	TotalVersions  *int64                      `json:"total_versions,omitempty"`
	Versions       []SoftwareTitleVersions     `json:"versions,omitempty"`
}