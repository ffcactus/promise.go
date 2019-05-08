package redfish

// Boot information.
type Boot struct {
	BootSourceOverrideTarget     string `json:"BootSourceOverrideTarget"`     // The current boot source to be used at next boot instead of the normal boot device, if BootSourceOverrideEnabled is true.
	BootSourceOverrideEnabled    string `json:"BootSourceOverrideEnabled"`    // Describes the state of the Boot Source Override feature.
	UefiTargetBootSourceOverride string `json:"UefiTargetBootSourceOverride"` // This property is the UEFI Device Path of the device to boot from when BootSourceOverrideSupported is UefiTarget.
	BootSourceOverrideMode       string `json:"BootSourceOverrideMode"`       // The BIOS Boot Mode (either Legacy or UEFI) to be used when BootSourceOverrideTarget boot source is booted from.
}

// GetSystemResponse is System response in Redfish.
type GetSystemResponse struct {
	Resource
	UUID         string `json:"UUID"`
	SystemType   string `json:"SystemType"`
	HostName     string `json:"HostName"`
	IndicatorLED string `json:"IndicatorLED"`
	PowerState   string `json:"PowerState"`
	Boot         Boot   `json:"Boot"`
	BiosVersion  string `json:"BiosVersion"`
}
