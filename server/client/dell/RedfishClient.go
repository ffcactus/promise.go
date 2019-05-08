package dell

import (
	"errors"
	// "promise/base"
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/server/client/dell/dto"
	"promise/server/client/redfish"
	"promise/server/object/constvalue"
	"promise/server/object/model"
	"strings"
)

// RedfishClient is the redfish client for HP servers.
type RedfishClient struct {
	base.Client
}

// GetInstance Get a new instance of Redfish client.
func GetInstance(address string, username string, password string) *RedfishClient {
	client := RedfishClient{}
	client.Protocol = "https"
	client.CurrentAddress = address
	client.Username = username
	client.Password = password
	return &client
}

// Support check if support.
func (c *RedfishClient) Support() bool {
	if err := c.Get("/redfish/v1", nil); err != nil {
		return false
	}
	return true
}

// String returns the client info.
func (c RedfishClient) String() string {
	return "Dell Redfish " + c.CurrentAddress
}

// GetProtocol Get the protocal used by this client.
func (c *RedfishClient) GetProtocol() string {
	return constvalue.RedfishV1
}

// GetBasicInfo Get server basic info.
// Just set parts of the properties.
func (c *RedfishClient) GetBasicInfo() (*model.ServerBasicInfo, error) {
	// First set the server type.
	var chassisCollection = redfish.Collection{}
	var systemURI, chassisURI string

	if err := c.Get("/redfish/v1/Chassis", &chassisCollection); err != nil {
		return nil, err
	}

	var systemCollection = redfish.Collection{}
	if err := c.Get("/redfish/v1/Systems", &systemCollection); err != nil {
		return nil, err
	}

	if systemCollection.Count != 1 {
		log.WithFields(log.Fields{"client": c}).Warn("GetBasicInfo failed, systems count not equal 1.")
		return nil, errors.New("can not find system URI")
	}
	systemURI = systemCollection.Members[0].Id

	for _, uri := range chassisCollection.Members {
		if strings.Contains(uri.Id, "System.Embedded") {
			chassisURI = uri.Id
		}
	}
	if chassisURI == "" {
		log.WithFields(log.Fields{"client": c}).Warn("GetBasicInfo failed, can not find chassis URI.")
		return nil, errors.New("can not find chassis URI")
	}

	// Get info from Computer system.
	var system = redfish.GetSystemResponse{}
	if err := c.Get(systemURI, &system); err != nil {
		return nil, err
	}
	ret := model.ServerBasicInfo{}
	ret.OriginURIs.System = systemURI
	ret.OriginURIs.Chassis = chassisURI
	ret.PhysicalUUID = system.UUID
	ret.Protocol = constvalue.RedfishV1
	ret.Vender = "Dell"
	ret.PowerState = system.PowerState

	// Get info from chassis.
	var chassis = redfish.GetChassisResponse{}
	if err := c.Get(chassisURI, &chassis); err != nil {
		return nil, err
	}
	ret.PhysicalState = chassis.Status.State
	ret.PhysicalHealth = chassis.Status.Health
	ret.Type = chassis.ChassisType
	ret.Model = chassis.Model
	ret.Manufacturer = chassis.Manufacturer
	ret.SKU = chassis.SKU
	ret.SerialNumber = chassis.SerialNumber
	ret.PartNumber = chassis.PartNumber
	ret.SparePartNumber = chassis.SparePartNumber
	ret.AssetTag = chassis.AssetTag
	ret.IndicatorLED = chassis.IndicatorLED

	return &ret, nil

}

// CreateManagementAccount Create Management account.
func (c *RedfishClient) CreateManagementAccount(username string, password string) error {
	return nil
}

// GetProcessors Get server's process info.
func (c *RedfishClient) GetProcessors(systemID string) ([]model.Processor, error) {
	collection := redfish.Collection{}
	if err := c.Get(systemID+"/Processors", &collection); err != nil {
		return nil, err
	}
	var ret []model.Processor
	for i := range collection.Members {
		each := new(dto.GetProcessorResponse)
		if err := c.Get(collection.Members[i].Id, each); err != nil {
			return nil, err
		}
		ret = append(ret, *createProcessorModel(each))
	}
	return ret, nil
}

// GetMemory Get server's memory info.
func (c *RedfishClient) GetMemory(systemID string) ([]model.Memory, error) {
	collection := redfish.Collection{}
	if err := c.Get(systemID+"/Memory", &collection); err != nil {
		return nil, err
	}

	var ret []model.Memory
	for i := range collection.Members {
		each := new(redfish.GetMemoryResponse)
		if err := c.Get(collection.Members[i].Id, each); err != nil {
			return nil, err
		}
		ret = append(ret, *redfish.CreateMemoryModel(each))
	}
	return ret, nil
}

// GetEthernetInterfaces Get server's ethernet interface info.
func (c *RedfishClient) GetEthernetInterfaces(systemID string) ([]model.EthernetInterface, error) {
	collection := redfish.Collection{}
	if err := c.Get(systemID+"/EthernetInterfaces", &collection); err != nil {
		return nil, err
	}
	var ret []model.EthernetInterface
	for i := range collection.Members {
		eachEthernet := new(redfish.GetEthernetInterfaceResponse)
		if err := c.Get(collection.Members[i].Id, eachEthernet); err != nil {
			return nil, err
		}
		// Get the VLANs
		vlanCollection := redfish.Collection{}
		vlanCollectionPageURI := systemID + "/EthernetInterfaces/" + eachEthernet.Id + "/VLANs"
		if err := c.Get(vlanCollectionPageURI, &vlanCollection); err != nil {
			return nil, err
		}
		var vlans []model.VLanNetworkInterface
		for j := range vlanCollection.Members {
			eachVlan := new(redfish.GetVLANResponse)
			if err := c.Get(collection.Members[j].Id, eachVlan); err != nil {
				return nil, err
			}
			vlans = append(vlans, *redfish.CreateVLanModel(eachVlan))
		}
		ethernetMode := *redfish.CreateEthernetInterfaceModel(eachEthernet)
		ethernetMode.VLANs = vlans
		ret = append(ret, ethernetMode)
	}
	return ret, nil
}

// GetNetworkInterfaces get network interfaces.
func (c *RedfishClient) GetNetworkInterfaces(systemID string) ([]model.NetworkInterface, error) {
	collection := redfish.Collection{}
	if err := c.Get(systemID+"/NetworkInterfaces", &collection); err != nil {
		return nil, err
	}
	var ret []model.NetworkInterface
	for i := range collection.Members {
		networkInterface := new(redfish.GetNetworkInterfaceResponse)
		if err := c.Get(collection.Members[i].Id, networkInterface); err != nil {
			return nil, err
		}

		ret = append(ret, *redfish.CreateNetworkInterfaceModel(networkInterface))
	}
	return ret, nil
}

// GetStorages get storages.
func (c *RedfishClient) GetStorages(systemID string) ([]model.Storage, error) {
	collection := redfish.Collection{}
	if err := c.Get(systemID+"/Storages", &collection); err != nil {
		return nil, err
	}
	ret := []model.Storage{}
	for i := range collection.Members {
		storage := new(redfish.GetStorageResponse)
		if err := c.Get(collection.Members[i].Id, storage); err != nil {
			return nil, err
		}

		ret = append(ret, *redfish.CreateStorageModel(storage))
	}
	return ret, nil
}

// GetPower get power.
func (c *RedfishClient) GetPower(chassisID string) (*model.Power, error) {
	power := new(redfish.GetPowerResponse)
	if err := c.Get(chassisID+"/Power", power); err != nil {
		return nil, err
	}
	model := redfish.CreatePowerModel(power)
	return model, nil
}

// GetThermal get thermal.
func (c *RedfishClient) GetThermal(chassisID string) (*model.Thermal, error) {
	thermal := new(redfish.GetThermalResponse)
	if err := c.Get(chassisID+"/Thermal", thermal); err != nil {
	}
	model := redfish.CreateThermalModel(thermal)
	return model, nil
}

// GetBoards get oem huawei boards.
func (c *RedfishClient) GetBoards(chassisID string) ([]model.Board, error) {
	ret := []model.Board{}
	return ret, nil
}

// GetNetworkAdapters get networkadapters.
func (c *RedfishClient) GetNetworkAdapters(chassisID string) ([]model.NetworkAdapter, error) {
	collection := redfish.Collection{}
	if err := c.Get(chassisID+"/NetworkAdapters", &collection); err != nil {
		return nil, err
	}
	var ret []model.NetworkAdapter
	for i := range collection.Members {
		resp := new(redfish.GetNetworkAdapterResponse)
		if err := c.Get(collection.Members[i].Id, resp); err != nil {
			return nil, err
		}
		networkAdpter := redfish.CreateNetworkAdapterModel(resp)
		for j := range resp.Controllers {
			eachController := redfish.CreateControllerModel(&resp.Controllers[j])
			portsResp := resp.Controllers[j].Links.NetworkPorts
			for k := range portsResp {
				portPageURI := portsResp[k].OdataId
				portResp := new(redfish.NetworkPort)
				if err := c.Get(portPageURI, portResp); err != nil {
					return nil, err
				}
				eachController.NetworkPorts = append(eachController.NetworkPorts, *redfish.CreateNetworkPortModel(portResp))
			}
			networkAdpter.Controllers = append(networkAdpter.Controllers, *eachController)
		}
		ret = append(ret, *networkAdpter)
	}
	// util.PrintJson(ret)
	return ret, nil
}

// GetDrives get drives.
func (c *RedfishClient) GetDrives(chassisID string) ([]model.Drive, error) {
	// Get the Drive links from chassis.
	// chassis := new(redfish.GetChassisResponse)
	// if err := c.Get(chassisID, chassis); err != nil {
	// 	return nil, err
	// }
	ret := []model.Drive{}
	// for i := range chassis.Links.Drives {
	// 	uri := chassis.Links.Drives[i].OdataId
	// 	drive := new(redfish.GetDriveResponse)
	// 	if err := c.Get(uri, drive); err != nil {
	// 		return nil, err
	// 	}
	// 	ret = append(ret, *redfish.CreateDriveModel(drive))
	// }
	return ret, nil
}

// GetPCIeDevices get PCIeDevices.
func (c *RedfishClient) GetPCIeDevices(chassisID string) ([]model.PCIeDevice, error) {
	// Get the Drive links from chassis.
	chassis := new(redfish.GetChassisResponse)
	if err := c.Get(chassisID, chassis); err != nil {
		return nil, err
	}
	ret := []model.PCIeDevice{}
	// for i := range chassis.Links.PCIeDevices {
	// 	uri := chassis.Links.PCIeDevices[i].OdataId
	// 	pcieDevice := new(redfish.GetPCIeDeviceResponse)
	// 	if err := c.Get(uri, pcieDevice); err != nil {
	// 		return nil, err
	// 	}
	// 	pcieFunctions := new([]redfish.GetPCIeFunctionResponse)
	// 	for j := range pcieDevice.Links.PCIeFunctions {
	// 		uri := pcieDevice.Links.PCIeFunctions[j].OdataId
	// 		pcieFunction := new(redfish.GetPCIeFunctionResponse)
	// 		if err := c.Get(uri, pcieFunction); err != nil {
	// 			return nil, err
	// 		}
	// 		*pcieFunctions = append(*pcieFunctions, *pcieFunction)
	// 	}

	// 	ret = append(ret, *redfish.CreatePCIeDeviceModel(pcieDevice, pcieFunctions))
	// }
	return ret, nil
}

// GetNetworkPorts get network ports.
func (c *RedfishClient) GetNetworkPorts(uri string) ([]model.NetworkPort, error) {
	collection := redfish.Collection{}
	if err := c.Get(uri, &collection); err != nil {
		return nil, err
	}
	var ret []model.NetworkPort
	for i := range collection.Members {
		resp := new(redfish.NetworkPort)
		if err := c.Get(collection.Members[i].Id, resp); err != nil {
			return nil, err
		}
		ret = append(ret, *redfish.CreateNetworkPortModel(resp))
	}
	return ret, nil
}

// createProcessorModel is used for Dell redfish.
func createProcessorModel(d *dto.GetProcessorResponse) *model.Processor {
	ret := model.Processor{}
	redfish.CreateResourceModel(&d.Resource, &ret.Resource)
	redfish.CreateProductInfoModel(&d.ProductInfo, &ret.ProductInfo)
	ret.Socket = d.Socket
	ret.ProcessorType = d.ProcessorType
	ret.ProcessorArchitecture = d.ProcessorArchitecture
	ret.InstructionSet = d.InstructionSet
	ret.MaxSpeedMHz = d.MaxSpeedMHz
	ret.TotalCores = d.TotalCores
	return &ret
}
