package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/mem"
	"github.com/tidwall/gjson"
)

type memMetricS struct {
	Token       string
	Metric      string
	Total       uint64
	Avail       uint64
	Used        uint64
	UsedPercent float64
	Free        uint64
	Rules       []string
}

type memInfo struct {
	Total       uint64
	Avail       uint64
	Used        uint64
	UsedPercent float64
	Free        uint64
}

type mLoad struct {
	Agent       string
	Token       string
	Information string
	Data        memInfo
}

func memMetric() (string, error) {
	v, merr := mem.VirtualMemory()
	if merr != nil {
		logger.Println(formatLog("Memory gathering error failed"))
		return string(""), merr
	}
	token := gjson.Get(configfile, "token").String()
	nrule := gjson.Get(configfile, "rule")
	var NmemInfo memMetricS
	NmemInfo.Token = token
	NmemInfo.Metric = "mem_metric"
	NmemInfo.Total = v.Total
	NmemInfo.Avail = v.Available
	NmemInfo.Used = v.Used
	NmemInfo.Free = v.Free
	NmemInfo.UsedPercent = v.UsedPercent
	MNinf, _ := json.Marshal(NmemInfo)
	for _, b := range nrule.Array() {
		got, err := ruleprocessor(b.String(), MNinf)
		if err != nil {
			str := fmt.Sprintf("%v", err)
			logger.Println(formatLog(str))
		} else {
			str := fmt.Sprintf("%v", got)
			NmemInfo.Rules = append(NmemInfo.Rules, str)
		}
	}
	inf, error := json.Marshal(NmemInfo)
	if error != nil {
		logger.Println(formatLog("Error Marshalling mem metric data"))
		return string(""), error
	}
	return string(inf), error
}

func memData() (string, error) {
	v, merr := mem.VirtualMemory()
	if merr != nil {
		logger.Println(formatLog("Memory gathering error failed"))
		return string(""), merr
	}
	token := gjson.Get(configfile, "token").String()
	agent := gjson.Get(configfile, "agent").String()
	var mNLoad mLoad
	mNLoad.Agent = agent
	mNLoad.Token = token
	mNLoad.Information = "host_memusage"
	var mInfo memInfo
	mInfo.Total = v.Total
	mInfo.Avail = v.Available
	mInfo.Used = v.Used
	mInfo.Free = v.Free
	mInfo.UsedPercent = v.UsedPercent
	mNLoad.Data = mInfo
	MNinf, IMerror := json.Marshal(mNLoad)
	if IMerror != nil {
		logger.Println(formatLog("Error Marshalling memory data"))
		return string(""), IMerror
	}
	return string(MNinf), IMerror

}
