package model

import (
	"promise/base"
)

// OriginURIs The
type OriginURIs struct {
	Chassis string
	System  string
}

// ComputerSystem Computer system object.
type ComputerSystem struct {
	Processors         []Processor
	Memory             []Memory
	EthernetInterfaces []EthernetInterface
	NetworkInterfaces  []NetworkInterface
	Storages           []Storage
}

// Chassis Chassis object.
type Chassis struct {
	Power           Power
	Thermal         Thermal
	Boards          []Board
	NetworkAdapters []NetworkAdapter
	Drives          []Drive
	PCIeDevices     []PCIeDevice
}

// ServerBasicInfo It represents the basic info about a server(Rack, Enclosure, Blade, Switch)
type ServerBasicInfo struct {
	Resource
	ProductInfo
	OriginURIs     OriginURIs // The URIs that we retrieve info from.
	PhysicalUUID   string
	Hostname       string
	Type           string
	Protocol       string
	Vender         string
	PowerState     string
	IndicatorLED   string
	OriginUsername *string
	OriginPassword *string
}

// Server is the model of server.
type Server struct {
	base.Model
	Resource
	ProductInfo
	OriginURIs     OriginURIs // The URIs that we retrieve info from.
	PhysicalUUID   string
	Hostname       string
	Type           string
	Protocol       string
	Vender         string
	PowerState     string
	IndicatorLED   string
	OriginUsername *string
	OriginPassword *string
	Credential     string
	ComputerSystem ComputerSystem
	Chassis        Chassis
}

// String return the debug name the model.
func (m Server) String() string {
	return m.Name
}

// ValueForDuplicationCheck return the value for duplication check.
func (m *Server) ValueForDuplicationCheck() string {
	return m.Name
}

// ServerCollectionMember is the member in collection.
type ServerCollectionMember struct {
	base.CollectionMemberModel
	Name   string
	State  string
	Health string
}

// ServerCollection is the model of collection.
type ServerCollection struct {
	base.CollectionModel
}

// NewModelMember return a new ModelMember
func (m *ServerCollection) NewModelMember() interface{} {
	return new(ServerCollectionMember)
}

// CreateServer Create servr object.
func (o *ServerBasicInfo) CreateServer() *Server {
	server := new(Server)
	server.Category = base.CategoryServer
	server.Name = o.Name
	server.Description = o.Description
	server.State = "State???"
	server.Health = "Health???"
	server.PhysicalState = o.PhysicalState
	server.PhysicalHealth = o.PhysicalHealth
	server.OriginURIs.Chassis = o.OriginURIs.Chassis
	server.OriginURIs.System = o.OriginURIs.System
	server.PhysicalUUID = o.PhysicalUUID
	server.Hostname = o.Hostname
	server.Type = o.Type
	server.Protocol = o.Protocol
	server.Vender = o.Vender
	server.PowerState = o.PowerState
	server.IndicatorLED = o.IndicatorLED
	server.OriginUsername = o.OriginUsername
	server.OriginPassword = o.OriginPassword

	server.ProductInfo.Model = o.Model
	server.Manufacturer = o.Manufacturer
	server.SKU = o.SKU
	server.SerialNumber = o.SerialNumber
	server.PartNumber = o.PartNumber
	server.SparePartNumber = o.SparePartNumber
	server.AssetTag = o.AssetTag
	return server
}
