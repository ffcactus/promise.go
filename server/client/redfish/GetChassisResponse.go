package redfish

// GetChassisResponse chassis response dto.
type GetChassisResponse struct {
	Resource
	ProductInfo
	IndicatorLED string
	ChassisType  string `json:"ChassisType"`
}
