package main

import (
	"encoding/json"
	"strconv"

	"github.com/shirou/gopsutil/disk"
	"github.com/tidwall/gjson"
)

type pDload struct {
	Token  string
	Metric string
	Data   []dMPoint
}

type dmPoint struct {
	Path    string
	Total   string
	Free    string
	Used    string
	Percent string
}

type dLoad struct {
	Agent       string
	Token       string
	Information string
	Data        []dmPoint
}

func (item *dLoad) AddItem(nitem dmPoint) []dmPoint {
	item.Data = append(item.Data, nitem)
	return item.Data
}

func (item *pDload) AddItem(nitem dMPoint) []dMPoint {
	item.Data = append(item.Data, nitem)
	return item.Data
}

func mdiskU() (string, error) {
	dparts, derr := disk.Partitions(false)
	if derr != nil {
		logger.Println("Error getting disk information")
		return string(""), derr
	}
	token := gjson.Get(configfile, "token").String()
	var cload pDload
	cload.Token = token
	cload.Metric = "Disk Metric"
	for _, part := range dparts {
		dus, err := disk.Usage(part.Mountpoint)
		if err != nil {
			logger.Println("Error getting partion information")
		} else {
			ndload := dMPoint{}
			ndload.Path = dus.Path
			ndload.Total = dus.Total
			ndload.Free = dus.Free
			ndload.Used = dus.Used
			ndload.Percent = dus.UsedPercent
			cload.AddItem(ndload)
		}
	}
	DNinf, DNerror := json.Marshal(cload)
	if DNerror != nil {
		logger.Println("Error Marshalling disk partion data")
		return string(DNinf), DNerror
	}
	return string(DNinf), DNerror
}

func diskU() (string, error) {
	dparts, derr := disk.Partitions(false)
	if derr != nil {
		logger.Println("Error getting disk information")
		return string(""), derr
	}
	var DLoad dLoad
	token := gjson.Get(configfile, "token").String()
	agent := gjson.Get(configfile, "agent").String()
	DLoad.Agent = agent
	DLoad.Token = token
	DLoad.Information = "host_diskusage"
	for _, part := range dparts {
		dus, err := disk.Usage(part.Mountpoint)
		if err != nil {
			logger.Println("Error getting partion information")
		} else {
			ndload := dmPoint{}
			ndload.Path = dus.Path
			ndload.Total = strconv.FormatUint(dus.Total/1024/1024/1024, 10)
			ndload.Free = strconv.FormatUint(dus.Free/1024/1024/1024, 10)
			ndload.Used = strconv.FormatUint(dus.Used/1024/1024/1024, 10)
			ndload.Percent = strconv.FormatFloat(dus.UsedPercent, 'f', 2, 64)
			DLoad.AddItem(ndload)
		}
	}
	DNinf, DNerror := json.Marshal(DLoad)
	if DNerror != nil {
		logger.Println("Error Marshalling disk partion data")
		return string(DNinf), DNerror
	}
	return string(DNinf), DNerror

}
