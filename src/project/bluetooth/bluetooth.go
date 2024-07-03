package bluetooth

import (
	"shylinux.com/x/ice"
	"shylinux.com/x/icebergs/base/web"
	kit "shylinux.com/x/toolkits"
	"tinygo.org/x/bluetooth"
)

type Bluetooth struct {
	list string `name:"list path auto"`
}

func (s Bluetooth) Scan(m *ice.Message, arg ...string) {
	adapter := bluetooth.DefaultAdapter
	err := adapter.Enable()
	m.Warn(err)
	err = adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		web.PushNoticeRich(m.Message,
			"address", device.Address.String(),
			"id", device.RSSI,
			"name", device.LocalName(),
			"payload", kit.Format("%#v", device.AdvertisementPayload))
		m.Debug("found device: %v %v %v %v", device.Address.String(), device.RSSI, device.LocalName(), device.AdvertisementPayload)
	})
	m.Sleep("10s")
}
func (s Bluetooth) List(m *ice.Message, arg ...string) {
	m.Action(s.Scan)
}

func init() { ice.Cmd("web.chat.bluetooth.bluetooth", Bluetooth{}) }
