package main

import (
	"encoding/json"

	"github.com/shirou/gopsutil/host"
	"github.com/tidwall/gjson"
)

type hinfo struct {
	Host            string
	Uptime          uint64
	Boottime        uint64
	Os              string
	PlatForm        string
	PlatFormFamily  string
	PlatFormVersion string
	KernelVersion   string
	KernelArch      string
}

type hLoad struct {
	Agent       string
	Token       string
	Information string
	Data        hinfo
}

func hostInfo() (string, error) {
	h, herror := host.Info()
	if herror != nil {
		return string(""), herror
	}

	var Hload hLoad
	var Hinfo hinfo
	token := gjson.Get(configfile, "token").String()
	agent := gjson.Get(configfile, "agent").String()
	Hload.Agent = agent
	Hload.Token = token
	Hload.Information = "host_info"
	Hinfo.Host = h.Hostname
	Hinfo.Uptime = h.Uptime
	Hinfo.Boottime = h.BootTime
	Hinfo.Os = h.OS
	Hinfo.PlatForm = h.Platform
	Hinfo.PlatFormFamily = h.PlatformFamily
	Hinfo.PlatFormVersion = h.PlatformVersion
	Hinfo.KernelVersion = h.KernelVersion
	Hinfo.KernelArch = h.KernelArch
	Hload.Data = Hinfo
	H, _ := json.Marshal(Hload)
	return string(H), herror

}
