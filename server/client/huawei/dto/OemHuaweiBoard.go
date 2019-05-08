package dto

import (
	"promise/server/client/redfish"
)

// GetBoardsResponse DTO.
type GetBoardsResponse struct {
	redfish.Resource
	redfish.ProductInfo
	CardNo          *int
	DeviceLocator   *string
	DeviceType      *string
	Location        *string
	CPLDVersion     *string
	PCBVersion      *string
	BoardName       *string
	BoardID         *string
	ManufactureDate *string
}
