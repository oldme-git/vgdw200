package service

import (
	"github.com/oldme-git/vgdw200/internal/unit"
)

var SRV = NewService()

func init() {
	SRV.Register(0x01, func() SrvInterface {
		return &unit.FindDeviceStatus{}
	})
	SRV.Register(0x04, func() SrvInterface {
		return &unit.BuzzerAndLED{}
	})
	SRV.Register(0x29, func() SrvInterface {
		return &unit.AudioCtrl{}
	})
	SRV.Register(0x61, func() SrvInterface {
		return &unit.ShowMsg{}
	})
	SRV.Register(0x63, func() SrvInterface {
		return &unit.ShowImg{}
	})
	SRV.Register(0x64, func() SrvInterface {
		return &unit.InWindow{}
	})
	SRV.Register(0x81, func() SrvInterface {
		return &unit.UptResourceReady{}
	})
	SRV.Register(0x82, func() SrvInterface {
		return &unit.UptResourceTransmit{}
	})
	SRV.Register(0x83, func() SrvInterface {
		return &unit.UptResourceOver{}
	})
}
