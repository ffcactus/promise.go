package entity

import (
	"promise/server/object/model"
)

// IPv6AddressPolicyEntry A entry in the RFC 6724 Address Selection Policy Table.
type IPv6AddressPolicyEntry struct {
	Prefix     *string // The IPv6 Address Prefix (as defined in RFC 6724 section 2.1).
	Precedence *int    // The IPv6 Precedence (as defined in RFC 6724 section 2.1.
	Label      *int    // The IPv6 Label (as defined in RFC 6724 section 2.1).
}

// EthernetInterfaceLinks Contains references to other resources that are related to this resource.
type EthernetInterfaceLinks struct {
	Endpoints     *[]string // An array of references to the endpoints that connect to this ethernet interface.
	HostInterface *string   // This is a reference to a Host Interface that is associated with this Ethernet Interface.
	Chassis       *string   // A reference to the Chassis which contains this Ethernet Interface.
}

// EthernetInterface This schema defines a simple ethernet NIC resource.
type EthernetInterface struct {
	ServerRef string
	EmbeddedResource
	UefiDevicePath         *string       // The UEFI device path for this interface.
	InterfaceEnabled       *bool         // This indicates whether this interface is enabled.
	PermanentMACAddress    *string       // This is the permanent MAC address assigned to this interface (port).
	MACAddress             *string       // This is the currently configured MAC address of the (logical port) interface.
	SpeedMbps              *int          // This is the current speed in Mbps of this interface.
	AutoNeg                *bool         // This indicates if the speed and duplex are automatically negotiated and configured on this interface.
	FullDuplex             *bool         // This indicates if the interface is in Full Duplex mode or not.
	MTUSize                *int          // This is the currently configured Maximum Transmission Unit (MTU) in bytes on this interface.
	HostName               *string       // The DNS Host Name, without any domain information.
	FQDN                   *string       // This is the complete, fully qualified domain name obtained by DNS for this interface.
	MaxIPv6StaticAddresses *string       // This indicates the maximum number of Static IPv6 addresses that can be configured on this interface.
	IPv4Addresses          []IPv4Address `gorm:"ForeignKey:EthernetInterfaceRef"` // The IPv4 addresses assigned to this interface.
	IPv6Addresses          []IPv6Address `gorm:"ForeignKey:EthernetInterfaceRef"` // This array of objects enumerates all of the currently assigned IPv6 addresses on this interface.
	// IPv6StaticAddresses    []IPv6Address          `gorm:"ForeignKey:EthernetInterfaceRef"` // This array of objects represents all of the IPv6 static addresses to be assigned on this interface.
	IPv6DefaultGateway *string                // This is the IPv6 default gateway address that is currently in use on this interface.
	VLANs              []VLanNetworkInterface `gorm:"ForeignKey:EthernetInterfaceRef"` // This is a reference to a collection of VLANs and is only used if the interface supports more than one VLANs.
	LinkStatus         *string                // The link status of this interface (port).
}

// IPv4Address This type describes an IPv4 Address.
type IPv4Address struct {
	EthernetInterfaceRef uint
	EmbeddedObject
	Address       *string // This is the IPv4 Address.
	SubnetMask    *string // This is the IPv4 Subnet mask.
	AddressOrigin *string // This indicates how the address was determined.
	Gateway       *string // This is the IPv4 gateway for this address.
}

// IPv6Address This type describes an IPv6 Address.
type IPv6Address struct {
	EthernetInterfaceRef uint
	EmbeddedObject
	Address       *string // This is the IPv6 Address.
	PrefixLength  *int    // This is the IPv6 Address Prefix Length.
	AddressOrigin *string // This indicates how the address was determined.
	AddressState  *string // The current state of this address as defined in RFC 4862.
}

// VLanNetworkInterface VLan network interface object.
type VLanNetworkInterface struct {
	EthernetInterfaceRef uint
	EmbeddedResource
	VLANEnable *bool // This indicates if this VLAN is enabled.
	VLANID     *int  // This indicates the VLAN identifier for this VLAN.
}

// ToModel will create a new model from
func (e *EthernetInterface) ToModel() *model.EthernetInterface {
	m := model.EthernetInterface{}
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	m.UefiDevicePath = e.UefiDevicePath
	m.InterfaceEnabled = e.InterfaceEnabled
	m.PermanentMACAddress = e.PermanentMACAddress
	m.MACAddress = e.MACAddress
	m.SpeedMbps = e.SpeedMbps
	m.AutoNeg = e.AutoNeg
	m.FullDuplex = e.FullDuplex
	m.MTUSize = e.MTUSize
	m.HostName = e.HostName
	m.FQDN = e.FQDN
	m.MaxIPv6StaticAddresses = e.MaxIPv6StaticAddresses
	m.LinkStatus = e.LinkStatus
	m.IPv4Addresses = []model.IPv4Address{}
	for i := range e.IPv4Addresses {
		each := model.IPv4Address{}
		each.Address = e.IPv4Addresses[i].Address
		each.SubnetMask = e.IPv4Addresses[i].SubnetMask
		each.AddressOrigin = e.IPv4Addresses[i].AddressOrigin
		each.Gateway = e.IPv4Addresses[i].Gateway
		m.IPv4Addresses = append(m.IPv4Addresses, each)

	}
	m.IPv6Addresses = []model.IPv6Address{}
	for i := range e.IPv6Addresses {
		each := model.IPv6Address{}
		each.Address = e.IPv6Addresses[i].Address
		each.PrefixLength = e.IPv6Addresses[i].PrefixLength
		each.AddressOrigin = e.IPv6Addresses[i].AddressOrigin
		each.AddressState = e.IPv6Addresses[i].AddressState
		m.IPv6Addresses = append(m.IPv6Addresses, each)

	}
	m.VLANs = []model.VLanNetworkInterface{}
	for i := range e.VLANs {
		each := model.VLanNetworkInterface{}
		each.VLANEnable = e.VLANs[i].VLANEnable
		each.VLANID = e.VLANs[i].VLANID
		m.VLANs = append(m.VLANs, each)
	}
	return &m
}

// Load will load data from model.
func (e *EthernetInterface) Load(m *model.EthernetInterface) {
	updateResourceEntity(&(*e).EmbeddedResource, &(*m).Resource)
	e.UefiDevicePath = m.UefiDevicePath
	e.InterfaceEnabled = m.InterfaceEnabled
	e.PermanentMACAddress = m.PermanentMACAddress
	e.MACAddress = m.MACAddress
	e.SpeedMbps = m.SpeedMbps
	e.AutoNeg = m.AutoNeg
	e.FullDuplex = m.FullDuplex
	e.MTUSize = m.MTUSize
	e.HostName = m.HostName
	e.FQDN = m.FQDN
	e.MaxIPv6StaticAddresses = m.MaxIPv6StaticAddresses
	e.LinkStatus = m.LinkStatus

	for i := range m.IPv4Addresses {
		each := IPv4Address{}
		each.Address = m.IPv4Addresses[i].Address
		each.SubnetMask = m.IPv4Addresses[i].SubnetMask
		each.AddressOrigin = m.IPv4Addresses[i].AddressOrigin
		each.Gateway = m.IPv4Addresses[i].Gateway
		e.IPv4Addresses = append(e.IPv4Addresses, each)
	}

	for i := range m.IPv6Addresses {
		each := IPv6Address{}
		each.Address = m.IPv6Addresses[i].Address
		each.PrefixLength = m.IPv6Addresses[i].PrefixLength
		each.AddressOrigin = m.IPv6Addresses[i].AddressOrigin
		each.AddressState = m.IPv6Addresses[i].AddressState
		e.IPv6Addresses = append(e.IPv6Addresses, each)
	}
	for i := range m.VLANs {
		each := VLanNetworkInterface{}
		updateResourceEntity(&each.EmbeddedResource, &m.VLANs[i].Resource)
		each.VLANEnable = m.VLANs[i].VLANEnable
		each.VLANID = m.VLANs[i].VLANID
		e.VLANs = append(e.VLANs, each)
	}
}
