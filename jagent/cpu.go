package main

import (
	"encoding/json"

	"github.com/shirou/gopsutil/cpu"
	"github.com/tidwall/gjson"
)

type cInfo struct {
	VendorID  string
	Family    string
	Model     string
	ModelName string
	Cores     int32
}

type cCount struct {
	CPUCount int
	CPUInfo  []cInfo
}

type cLoad struct {
	Agent       string
	Token       string
	Information string
	Data        cCount
}

func (item *cCount) AddItem(nitem cInfo) []cInfo {
	item.CPUInfo = append(item.CPUInfo, nitem)
	return item.CPUInfo
}

func cpuinfo() (string, error) {
	cpud, cpuerr := cpu.Info()
	if cpuerr != nil {
		logger.Println("Issue getting cpu information")
		return string(""), cpuerr
	}
	token := gjson.Get(configfile, "token").String()
	agent := gjson.Get(configfile, "agent").String()
	var CLoad cLoad
	CLoad.Agent = agent
	CLoad.Token = token
	CLoad.Information = "host_cpuinfo"
	var cpuC cCount
	for _, cpuI := range cpud {
		newIcpu := cInfo{}
		newIcpu.Cores = cpuI.Cores
		newIcpu.Family = cpuI.Family
		newIcpu.Model = cpuI.Model
		newIcpu.ModelName = cpuI.ModelName
		newIcpu.VendorID = cpuI.VendorID
		cpuC.AddItem(newIcpu)
	}
	cpuC.CPUCount = len(cpud)
	CLoad.Data = cpuC
	cinf, cerror := json.Marshal(CLoad)
	if cerror != nil {
		logger.Println("Error Marshalling disk partion data")
		return string(cinf), cerror
	}
	return string(cinf), cerror

}
