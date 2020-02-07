package device

import (
	"github.com/am-x/homekit/app/hub"
	"github.com/am-x/homekit/app/shared/messages"
	"gobot.io/x/gobot"
)

type HubDevice interface {
	GetDevices() []gobot.Device
}

type Initializer func(h *hub.Hub, cfg *messages.DeviceConfig) HubDevice
