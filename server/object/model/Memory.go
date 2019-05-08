package model

// MemoryLocation Memory connection information to sockets and memory controllers.
type MemoryLocation struct {
	Socket     *int // Socket number in which Memory is connected.
	Controller *int // Memory controller number in which Memory is connected.
	Channel    *int // Channel number in which Memory is connected.
	Slot       *int // Slot number in which Memory is connected.
}

// Memory o is the schema definition for definition of a Memory and its configuration.
type Memory struct {
	Resource
	ProductInfo
	MemoryType        *string         // The type of Memory.
	MemoryDeviceType  *string         // Type details of the Memory.
	CapacityMiB       *int            // Memory Capacity in MiB.
	DataWidthBits     *int            // Data Width in bits.
	BusWidthBits      *int            // Bus Width in bits.
	VendorID          *string         // Vendor ID.
	DeviceID          *string         // Device ID.
	SubsystemVendorID *string         // SubSystem Vendor ID.
	SubsystemDeviceID *string         // Subsystem Device ID.
	RankCount         *int            // Number of ranks available in the Memory.
	DeviceLocator     *string         // Location of the Memory in the platform.
	MemoryLocation    *MemoryLocation // Memory connection information to sockets and memory controllers.
	ErrorCorrection   *string         // Error correction scheme supported for o memory.
	OperatingSpeedMhz *int            // Operating speed of Memory in MHz.
}
