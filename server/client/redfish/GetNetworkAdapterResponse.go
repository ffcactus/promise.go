package redfish

// ControllersLinks the links used in Controllers.
type ControllersLinks struct {
	//	PCIeDevices            []string
	NetworkPorts []ResourceRef `json:"NetworkPorts"`
	//	NetworkDeviceFunctions []string
}

// DataCenterBridging data Center Bridging (DCB) for capabilities of a controller.
type DataCenterBridging struct {
	Capable bool `json:"Capable"` // Whether this controller is capable of Data Center Bridging (DCB).
}

// VirtualFunction a virtual function of a controller.
type VirtualFunction struct {
	DeviceMaxCount         int `json:"DeviceMaxCount"`         // The maximum number of Virtual Functions (VFs) supported by this controller.
	NetworkPortMaxCount    int `json:"NetworkPortMaxCount"`    // The maximum number of Virtual Functions (VFs) supported per network port for this controller.
	MinAssignmentGroupSize int `json:"MinAssignmentGroupSize"` // The minimum number of Virtual Functions (VFs) that can be allocated or moved between physical functions for this controller.
}

// SRIOV Single-Root Input/Output Virtualization (SR-IOV) capabilities.
type SRIOV struct {
	SRIOVVEPACapable bool `json:"SRIOVVEPACapable"` // Whether this controller supports Single Root Input/Output Virtualization (SR-IOV) in Virtual Ethernet Port Aggregator (VEPA) mode.
}

// NPIV N_Port ID Virtualization (NPIV) capabilties for a controller.
type NPIV struct {
	MaxDeviceLogins int `json:"MaxDeviceLogins"` // The maximum number of N_Port ID Virtualization (NPIV) logins allowed simultaneously from all ports on this controller.
	MaxPortLogins   int `json:"MaxPortLogins"`   // The maximum number of N_Port ID Virtualization (NPIV) logins allowed per physical port on this controller.
}

// VirtualizationOffload a Virtualization offload capability of a controller.
type VirtualizationOffload struct {
	VirtualFunction VirtualFunction `json:"VirtualFunction"` // A virtual function of a controller.
	SRIOV           SRIOV           `json:"SRIOV"`           // Single-Root Input/Output Virtualization (SR-IOV) capabilities.
}

// ControllerCapabilities the capabilities of a controller.
type ControllerCapabilities struct {
	NetworkPortCount int `json:"NetworkPortCount"` // The number of physical ports on this controller.
	// NetworkDeviceFunctionCount int                   // The maximum number of physical functions available on this controller.
	// DataCenterBridging         DataCenterBridging    // Data Center Bridging (DCB) for this controller.
	// VirtualizationOffload      VirtualizationOffload // Virtualization offload for this controller.
	// NPIV                       NPIV                  // N_Port ID Virtualization (NPIV) capabilties for this controller.
}

// Controller a network controller ASIC that makes up part of a NetworkAdapter.
type Controller struct {
	FirmwarePackageVersion string                 `json:"FirmwarePackageVersion"` // The version of the user-facing firmware package.
	ControllerCapabilities ControllerCapabilities `json:"ControllerCapabilities"` // The capabilities of this controller.
	Links                  ControllersLinks       `json:"Link"`                   // Links for this controller.
}

// GetNetworkAdapterResponse represents the physical network adapter capable of connecting to a computer network.  Examples include but are not limited to Ethernet, Fibre Channel, and converged network adapters.
type GetNetworkAdapterResponse struct {
	Resource
	ProductInfo
	Controllers  []Controller `json:"Controllers"`  // The set of network controllers ASICs that make up this NetworkAdapter.
	NetworkPorts ResourceRef  `json:"NetworkPorts"` // Contains the members of this collection.
}
