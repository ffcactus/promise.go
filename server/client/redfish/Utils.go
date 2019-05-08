package redfish

import (
	"promise/server/object/model"
	"strconv"
)

// CreateResourceModel creates resource from Redfish.
func CreateResourceModel(d *Resource, m *model.Resource) {
	m.URI = d.OdataID
	m.Name = d.Name
	m.Description = d.Description
	m.OriginID = d.Id
	m.PhysicalState = d.Status.State
	m.PhysicalHealth = d.Status.Health
}

// CreateMemberModel creates member from Redfish
func CreateMemberModel(d *Member, m *model.Member) {
	m.URI = d.OdataID
	m.Name = d.Name
	m.Description = d.Description
	m.OriginMemberID = d.MemberId
	m.PhysicalState = d.Status.State
	m.PhysicalHealth = d.Status.Health
}

// CreateThresholdModel creates threshold from Redfish.
func CreateThresholdModel(d *Threshold, m *model.Threshold) {
	m.UpperThresholdNonCritical = d.UpperThresholdNonCritical
	m.UpperThresholdCritical = d.UpperThresholdCritical
	m.UpperThresholdFatal = d.UpperThresholdFatal
	m.LowerThresholdNonCritical = d.LowerThresholdNonCritical
	m.LowerThresholdCritical = d.LowerThresholdCritical
	m.LowerThresholdFatal = d.LowerThresholdFatal
}

// CreateProductInfoModel do convert.
func CreateProductInfoModel(d *ProductInfo, m *model.ProductInfo) {
	m.Model = d.Model
	m.Manufacturer = d.Manufacturer
	m.SKU = d.SKU
	m.SerialNumber = d.SerialNumber
	m.PartNumber = d.PartNumber
	m.SparePartNumber = d.SparePartNumber
	m.AssetTag = d.AssetTag
}

// CreateLocationModel do convert.
func CreateLocationModel(d *Location, m *model.Location) {
	m.Info = d.Info
	m.InfoFormat = d.InfoFormat
}

// CreateProcessorModel do convert.
func CreateProcessorModel(d *GetProcessorResponse) *model.Processor {
	ret := model.Processor{}
	CreateResourceModel(&d.Resource, &ret.Resource)
	CreateProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	if d.Socket != nil {
		s := strconv.Itoa(*d.Socket)
		ret.Socket = &s
	}
	ret.ProcessorType = d.ProcessorType
	ret.ProcessorArchitecture = d.ProcessorArchitecture
	ret.InstructionSet = d.InstructionSet
	ret.MaxSpeedMHz = d.MaxSpeedMHz
	ret.TotalCores = d.TotalCores
	return &ret
}

// CreateMemoryModel do convert.
func CreateMemoryModel(d *GetMemoryResponse) *model.Memory {
	ret := model.Memory{}
	CreateResourceModel(&d.Resource, &ret.Resource)
	CreateProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	ret.CapacityMiB = d.CapacityMiB
	ret.OperatingSpeedMhz = d.OperatingSpeedMhz
	ret.MemoryDeviceType = d.MemoryDeviceType
	ret.DataWidthBits = d.DataWidthBits
	ret.RankCount = d.RankCount
	ret.DeviceLocator = d.DeviceLocator

	if d.MemoryLocation != nil {
		ret.MemoryLocation = new(model.MemoryLocation)
		ret.MemoryLocation.Socket = d.MemoryLocation.Socket
		ret.MemoryLocation.Controller = d.MemoryLocation.Controller
		ret.MemoryLocation.Channel = d.MemoryLocation.Channel
		ret.MemoryLocation.Slot = d.MemoryLocation.Slot
	}
	return &ret
}

// CreateEthernetInterfaceModel do convert.
func CreateEthernetInterfaceModel(d *GetEthernetInterfaceResponse) *model.EthernetInterface {
	ret := model.EthernetInterface{}
	CreateResourceModel(&d.Resource, &ret.Resource)
	ret.UefiDevicePath = d.UefiDevicePath
	ret.InterfaceEnabled = d.InterfaceEnabled
	ret.PermanentMACAddress = d.PermanentMACAddress
	ret.MACAddress = d.MACAddress
	ret.SpeedMbps = d.SpeedMbps
	ret.AutoNeg = d.AutoNeg
	ret.FullDuplex = d.FullDuplex
	ret.MTUSize = d.MTUSize
	ret.HostName = d.HostName
	ret.FQDN = d.FQDN
	ret.MaxIPv6StaticAddresses = d.MaxIPv6StaticAddresses
	ret.LinkStatus = d.LinkStatus
	if d.IPv4Addresses != nil {
		ipv4 := []model.IPv4Address{}
		for i := range *d.IPv4Addresses {
			each := model.IPv4Address{}
			each.Address = (*d.IPv4Addresses)[i].Address
			each.SubnetMask = (*d.IPv4Addresses)[i].SubnetMask
			each.AddressOrigin = (*d.IPv4Addresses)[i].AddressOrigin
			each.Gateway = (*d.IPv4Addresses)[i].Gateway
			ipv4 = append(ipv4, each)
		}
		ret.IPv4Addresses = ipv4
	}
	if d.IPv6Addresses != nil {
		ipv6 := []model.IPv6Address{}
		for i := range *d.IPv6Addresses {
			each := model.IPv6Address{}
			each.Address = (*d.IPv6Addresses)[i].Address
			each.PrefixLength = (*d.IPv6Addresses)[i].PrefixLength
			each.AddressOrigin = (*d.IPv6Addresses)[i].AddressOrigin
			each.AddressState = (*d.IPv6Addresses)[i].AddressState
			ipv6 = append(ipv6, each)
		}
		ret.IPv6Addresses = ipv6
	}
	return &ret
}

// CreateVLanModel do convert.
func CreateVLanModel(d *GetVLANResponse) *model.VLanNetworkInterface {
	ret := model.VLanNetworkInterface{}
	CreateResourceModel(&d.Resource, &ret.Resource)
	ret.VLANEnable = d.VLANEnable
	ret.VLANID = d.VLANID
	return &ret
}

// CreateNetworkInterfaceModel do convert.
func CreateNetworkInterfaceModel(d *GetNetworkInterfaceResponse) *model.NetworkInterface {
	ret := model.NetworkInterface{}
	CreateResourceModel(&d.Resource, &ret.Resource)
	ret.NetworkAdapterURI = d.Links.NetworkAdapter.OdataId
	return &ret
}

// CreateStorageControllerModel do convert.
func CreateStorageControllerModel(d *StorageController) *model.StorageController {
	ret := model.StorageController{}
	CreateMemberModel(&d.Member, &ret.Member)
	CreateProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	ret.SpeedGbps = d.SpeedGbps
	ret.FirmwareVersion = d.FirmwareVersion
	ret.SupportedDeviceProtocols = d.SupportedDeviceProtocols
	return &ret
}

// CreateStorageModel do convert.
func CreateStorageModel(d *GetStorageResponse) *model.Storage {
	ret := model.Storage{}
	CreateResourceModel(&d.Resource, &ret.Resource)
	for i := range d.Drives {
		ret.DriveURIs = append(ret.DriveURIs, d.Drives[i].OdataId)
	}
	for i := range d.StorageControllers {
		ret.StorageControllers = append(ret.StorageControllers, *CreateStorageControllerModel(&d.StorageControllers[i]))
	}
	return &ret
}

// CreatePowerModel do convert.
func CreatePowerModel(d *GetPowerResponse) *model.Power {
	r := *d
	ret := model.Power{}
	CreateResourceModel(&d.Resource, &ret.Resource)
	// PowerControl
	powerControl := []model.PowerControl{}
	for i := range *r.PowerControl {
		eachModel := model.PowerControl{}
		eachDto := (*r.PowerControl)[i]
		CreateResourceModel(&eachDto.Resource, &eachModel.Resource)
		CreateProductInfoModel(&eachDto.ProductInfo, &eachModel.ProductInfo)
		eachModel.PowerConsumedWatts = eachDto.PowerConsumedWatts
		eachModel.PowerRequestedWatts = eachDto.PowerRequestedWatts
		eachModel.PowerAvailableWatts = eachDto.PowerAvailableWatts
		eachModel.PowerCapacityWatts = eachDto.PowerCapacityWatts
		eachModel.PowerAllocatedWatts = eachDto.PowerAllocatedWatts
		if eachDto.PowerMetrics != nil {
			powerMetrics := model.PowerMetrics{}
			powerMetrics.MinConsumedWatts = eachDto.PowerMetrics.MinConsumedWatts
			powerMetrics.MaxConsumedWatts = eachDto.PowerMetrics.MaxConsumedWatts
			powerMetrics.AverageConsumedWatts = eachDto.PowerMetrics.AverageConsumedWatts
			eachModel.PowerMetrics = &powerMetrics
		}
		if eachDto.PowerLimit != nil {
			powerLimit := model.PowerLimit{}
			powerLimit.LimitInWatts = eachDto.PowerLimit.LimitInWatts
			powerLimit.LimitException = eachDto.PowerLimit.LimitException
			powerLimit.CorrectionInMs = eachDto.PowerLimit.CorrectionInMs
			eachModel.PowerLimit = &powerLimit
		}
		powerControl = append(powerControl, eachModel)
	}
	ret.PowerControl = powerControl

	// PowerSupplies
	powerSupplies := []model.PowerSupply{}
	for i := range *r.PowerSupplies {
		eachModel := model.PowerSupply{}
		eachDto := (*r.PowerSupplies)[i]
		CreateResourceModel(&eachDto.Resource, &eachModel.Resource)
		CreateProductInfoModel(&eachDto.ProductInfo, &eachModel.ProductInfo)
		eachModel.PowerSupplyType = eachDto.PowerSupplyType
		eachModel.LineInputVoltageType = eachDto.LineInputVoltageType
		eachModel.LineInputVoltage = eachDto.LineInputVoltage
		eachModel.PowerCapacityWatts = eachDto.PowerCapacityWatts
		eachModel.LastPowerOutputWatts = eachDto.LastPowerOutputWatts
		eachModel.FirmwareVersion = eachDto.FirmwareVersion
		eachModel.IndicatorLed = eachDto.IndicatorLed
		powerSupplies = append(powerSupplies, eachModel)
	}
	ret.PowerSupplies = powerSupplies

	// Redundancy
	redundancy := []model.Redundancy{}
	for i := range *r.Redundancy {
		eachModel := model.Redundancy{}
		eachDto := (*r.Redundancy)[i]
		CreateResourceModel(&eachDto.Resource, &eachModel.Resource)
		eachModel.Mode = eachDto.Mode
		eachModel.MaxNumSupported = eachDto.MaxNumSupported
		eachModel.MinNumNeeded = eachDto.MinNumNeeded
		eachModel.RedundancyEnabled = eachDto.RedundancyEnabled
		// only name is needed in the name of redundancy set.
		redundancySet := []string{}
		for j := range *eachDto.RedundancySet {
			for k := range *r.PowerSupplies {
				redundancyOdataID := (*eachDto.RedundancySet)[j].OdataId
				powerSupply := (*r.PowerSupplies)[k]
				if powerSupply.OdataID == redundancyOdataID {
					redundancySet = append(redundancySet, powerSupply.Name)
				}
			}
		}
		eachModel.RedundancySet = &redundancySet
		redundancy = append(redundancy, eachModel)
	}
	ret.Redundancy = redundancy
	return &ret
}

// CreateThermalModel do convert.
func CreateThermalModel(d *GetThermalResponse) *model.Thermal {
	ret := new(model.Thermal)
	CreateResourceModel(&d.Resource, &ret.Resource)

	fans := []model.Fan{}
	for i := range d.Fans {
		each := model.Fan{}
		CreateMemberModel(&d.Fans[i].Member, &each.Member)
		CreateProductInfoModel(&d.Fans[i].ProductInfo, &each.ProductInfo)
		CreateThresholdModel(&d.Fans[i].Threshold, &each.Threshold)
		each.Reading = d.Fans[i].Reading
		each.MinReadingRange = d.Fans[i].MinReadingRange
		each.MaxReadingRange = d.Fans[i].MaxReadingRange
		each.ReadingUnits = d.Fans[i].ReadingUnits
		// Redundancy is needed for Enclosure.
		fans = append(fans, each)
	}
	ret.Fans = fans

	return ret
}

// CreateNetworkAdapterModel do convert.
func CreateNetworkAdapterModel(d *GetNetworkAdapterResponse) *model.NetworkAdapter {
	ret := new(model.NetworkAdapter)
	CreateResourceModel(&d.Resource, &ret.Resource)
	CreateProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	return ret
}

// CreateControllerModel do convert.
func CreateControllerModel(d *Controller) *model.Controller {
	ret := new(model.Controller)
	ret.FirmwarePackageVersion = d.FirmwarePackageVersion
	ret.ControllerCapabilities.NetworkPortCount = d.ControllerCapabilities.NetworkPortCount
	return ret
}

// CreateNetworkPortModel do convert.
func CreateNetworkPortModel(d *NetworkPort) *model.NetworkPort {
	ret := new(model.NetworkPort)
	CreateResourceModel(&d.Resource, &ret.Resource)
	ret.PhysicalPortNumber = d.PhysicalPortNumber
	ret.LinkStatus = d.LinkStatus
	ret.AssociatedNetworkAddresses = d.AssociatedNetworkAddresses
	return ret
}

// CreateDriveModel do convert.
func CreateDriveModel(d *GetDriveResponse) *model.Drive {
	ret := new(model.Drive)
	CreateResourceModel(&d.Resource, &ret.Resource)
	CreateProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	ret.StatusIndicator = d.StatusIndicator
	ret.IndicatorLED = d.IndicatorLED
	ret.Revision = d.Revision
	ret.CapacityBytes = d.CapacityBytes
	ret.FailurePredicted = d.FailurePredicted
	ret.Protocol = d.Protocol
	ret.MediaType = d.MediaType
	ret.HotspareType = d.HotspareType
	ret.CapableSpeedGbs = d.CapableSpeedGbs
	ret.NegotiatedSpeedGbs = d.NegotiatedSpeedGbs
	ret.PredictedMediaLifeLeftPercent = d.PredictedMediaLifeLeftPercent
	for i := range d.Location {
		m := new(model.Location)
		CreateLocationModel(&d.Location[i], m)
		ret.Location = append(ret.Location, *m)
	}
	return ret
}

// CreatePCIeDeviceModel do convert.
func CreatePCIeDeviceModel(device *GetPCIeDeviceResponse, functions *[]GetPCIeFunctionResponse) *model.PCIeDevice {
	ret := new(model.PCIeDevice)
	CreateResourceModel(&device.Resource, &ret.Resource)
	CreateProductInfoModel(&device.ProductInfo, &ret.ProductInfo)
	ret.DeviceType = device.DeviceType
	ret.FirmwareVersion = device.FirmwareVersion
	for i := range *functions {
		ret.PCIeFunctions = append(ret.PCIeFunctions, *CreatePCIeFunctionModel(&(*functions)[i]))
	}
	return ret
}

// // CreateLocationModel do convert. do convert.
func CreatePCIeFunctionModel(d *GetPCIeFunctionResponse) *model.PCIeFunction {
	ret := new(model.PCIeFunction)
	CreateResourceModel(&d.Resource, &ret.Resource)
	ret.DeviceClass = d.DeviceClass
	ret.DeviceID = d.DeviceID
	ret.VendorID = d.VendorID
	ret.SubsystemID = d.SubsystemID
	ret.SubsystemVendorID = d.SubsystemVendorID
	for i := range d.Links.EthernetInterfaces {
		ret.EthernetInterfaces = append(ret.EthernetInterfaces, d.Links.EthernetInterfaces[i].OdataId)
	}
	return ret
}
