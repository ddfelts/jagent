package main

import (
	"encoding/json"
	"strings"

	"github.com/shirou/gopsutil/net"
	"github.com/tidwall/gjson"
)

type infInfo struct {
	Name   string
	Status string
	HAddr  string
	Addrv4 string
	Addrv6 string
}

type iLoad struct {
	Agent       string
	Token       string
	Information string
	Data        []infInfo
}

func (item *iLoad) AddItem(nitem infInfo) []infInfo {
	item.Data = append(item.Data, nitem)
	return item.Data
}

//IsIPv4 : checks to see if string is an ipv4
func IsIPv4(address string) bool {
	return strings.Count(address, ":") < 2
}

//IsIPv6 : checks to see if string is an ipv6
func IsIPv6(address string) bool {
	return strings.Count(address, ":") >= 2
}

func netWork() (string, error) {
	c, inerror := net.Interfaces()
	if inerror != nil {
		logger.Println(formatLog("Error getting network inteface config"))
		return string(""), inerror
	}
	token := gjson.Get(configfile, "token").String()
	agent := gjson.Get(configfile, "agent").String()
	var iInfo iLoad
	iInfo.Agent = agent
	iInfo.Token = token
	iInfo.Information = "host_ifaces"
	for _, infoI := range c {
		newinf := infInfo{}
		newinf.Name = infoI.Name
		if infoI.HardwareAddr == "" {
			newinf.HAddr = "None"
		} else {
			newinf.HAddr = infoI.HardwareAddr
		}
		if len(infoI.Flags) >= 1 {
			if infoI.Flags[0] == "up" {
				newinf.Status = "up"
			} else {
				newinf.Status = "down"
			}
		} else {
			newinf.Status = "None"
		}
		if len(infoI.Addrs) >= 1 {
			for _, ipaddrs := range infoI.Addrs {
				if IsIPv4(ipaddrs.Addr) {
					newinf.Addrv4 = ipaddrs.Addr
					continue
				}
				if IsIPv6(ipaddrs.Addr) {
					newinf.Addrv6 = ipaddrs.Addr
					continue
				}
			}

		}
		iInfo.AddItem(newinf)
	}
	INinf, INerror := json.Marshal(iInfo)
	if INerror != nil {
		logger.Println(formatLog("Error Marshalling network data"))
		return string(INinf), INerror
	}
	return string(INinf), inerror

}
