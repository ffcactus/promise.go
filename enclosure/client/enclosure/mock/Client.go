package mock

import (
	log "github.com/sirupsen/logrus"
	"promise/base"
	"promise/enclosure/object/model"
	"time"
)

// Client implements EnclosureClient interface.
type Client struct {
	Address  string
	Username string
	Password string
}

// NewClient creates a client for enclosure.
func NewClient(enclosure *model.Enclosure) *Client {
	client := Client{}
	if len(enclosure.Addresses) > 0 {
		client.Address = enclosure.Addresses[0]
	}
	return &client
}

// String returns the debug info of the client.
func (c Client) String() string {
	return "Mock"
}

// Ready returns if the enclosure is ready.
func (c Client) Ready() bool {
	return true
}

// Claim should make make a flag on the enclosure that indicate it is exclusively managed by this system.
func (c Client) Claim() base.ClientError {
	return nil
}

// Unclaim should remove the flag that indicate the enclosure is managed by this system.
func (c Client) Unclaim() base.ClientError {
	return nil
}

// DeviceIdentity returns the device identity.
func (c Client) DeviceIdentity() (*base.DeviceIdentity, base.ClientError) {
	time.Sleep(time.Duration(1) * time.Second)
	identity := base.DeviceIdentity{}
	identity.UUID = base.RandUUID()
	identity.SerialNumber = base.RandString(10)
	identity.PartNumber = base.RandString(10)
	log.WithFields(log.Fields{"client": c, "identity": identity}).Info("Client get device identity.")
	return &identity, nil
}

// ServerSlot returns the server slot info.
func (c Client) ServerSlot() ([]model.ServerSlot, base.ClientError) {
	time.Sleep(time.Duration(10) * time.Second)
	slots := make([]model.ServerSlot, 0)
	for i := 1; i <= 8; i++ {
		slot := model.ServerSlot{}
		slot.Index = i
		slot.Inserted = true
		slot.ProductName = "CH121 V5"
		slot.SerialNumber = base.RandString(10)
		slot.Height = 1
		slot.Width = 1
		slots = append(slots, slot)
	}
	log.WithFields(log.Fields{"client": c}).Info("Client get server slot.")
	return slots, nil
}

// SwitchSlot returns the switch ade slot info.
func (c Client) SwitchSlot() ([]model.SwitchSlot, base.ClientError) {
	time.Sleep(time.Duration(3) * time.Second)
	slots := make([]model.SwitchSlot, 0)
	for i := 1; i <= 4; i++ {
		slot := model.SwitchSlot{}
		slot.Index = i
		slot.Inserted = true
		slot.ProductName = "CX920"
		slot.SerialNumber = base.RandString(10)
		slots = append(slots, slot)
	}
	log.WithFields(log.Fields{"client": c}).Info("Client get switch slot.")
	return slots, nil
}

// FanSlot returns the fan slot info.
func (c Client) FanSlot() ([]model.FanSlot, base.ClientError) {
	time.Sleep(time.Duration(3) * time.Second)
	slots := make([]model.FanSlot, 0)
	for i := 1; i <= 14; i++ {
		slot := model.FanSlot{}
		slot.Index = i
		slot.Inserted = true
		slot.Health = "OK"
		slot.PCBVersion = base.RandString(10)
		slot.SoftwareVersion = base.RandString(10)
		slots = append(slots, slot)
	}
	log.WithFields(log.Fields{"client": c}).Info("Client get fan slot.")
	return slots, nil
}

// PowerSlot returns the power slot info.
func (c Client) PowerSlot() ([]model.PowerSlot, base.ClientError) {
	time.Sleep(time.Duration(3) * time.Second)
	slots := make([]model.PowerSlot, 0)
	for i := 1; i <= 6; i++ {
		slot := model.PowerSlot{}
		slot.Index = i
		slot.Inserted = true
		slot.Health = "OK"
		slot.PowerSupplyType = "AC"
		slot.SerialNumber = base.RandString(10)
		slot.FirmwareVersion = base.RandString(10)
		slot.SleepStatus = "Enable"
		slots = append(slots, slot)
	}
	log.WithFields(log.Fields{"client": c}).Info("Client get power slot.")
	return slots, nil
}

// ManagerSlot returns the manager slot info.
func (c Client) ManagerSlot() ([]model.ManagerSlot, base.ClientError) {
	time.Sleep(time.Duration(5) * time.Second)
	slots := make([]model.ManagerSlot, 0)
	for i := 1; i <= 2; i++ {
		slot := model.ManagerSlot{}
		slot.Index = i
		slot.Inserted = true
		slot.ProductName = "MM920"
		slot.SerialNumber = base.RandString(10)
		slot.FirmwareVersion = base.RandString(10)
		slot.CPLDVersion = base.RandString(10)
		slots = append(slots, slot)
	}
	log.WithFields(log.Fields{"client": c}).Info("Client get manager slot.")
	return slots, nil
}

// ApplianceSlot returns the manager slot info.
func (c Client) ApplianceSlot() ([]model.ApplianceSlot, base.ClientError) {
	time.Sleep(time.Duration(3) * time.Second)
	slots := make([]model.ApplianceSlot, 0)
	for i := 1; i <= 2; i++ {
		slot := model.ApplianceSlot{}
		slot.Index = i
		slot.Inserted = true
		slot.SerialNumber = base.RandString(10)
		slot.FirmwareVersion = base.RandString(10)
		slot.BIOSVersion = base.RandString(10)
		slots = append(slots, slot)
	}
	log.WithFields(log.Fields{"client": c}).Info("Client get appliance slot.")
	return slots, nil
}
