package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/load"
	"github.com/tidwall/gjson"
)

type pload struct {
	Token  string
	Metric string
	Load1  float64
	Load5  float64
	Load15 float64
	Rules  []string
}

type rinfo struct {
	Load1  float64
	Load5  float64
	Load15 float64
}
type rLoad struct {
	Agent  string
	Token  string
	Metric string
	Data   rinfo
}

func loadAvgM() (string, error) {
	l, loaderr := load.Avg()
	if loaderr != nil {
		logger.Println(formatLog("Error getting config"))
		return string(""), loaderr
	}
	token := gjson.Get(configfile, "token").String()
	nrule := gjson.Get(configfile, "rule")
	var load pload
	load.Token = token
	load.Metric = "load_metric"
	load.Load1 = l.Load1
	load.Load5 = l.Load5
	load.Load15 = l.Load15
	mload, _ := json.Marshal(load)
	for _, b := range nrule.Array() {
		got, err := ruleprocessor(b.String(), mload)
		if err != nil {
			str := fmt.Sprintf("%v", err)
			logger.Println(formatLog(str))
		} else {
			str := fmt.Sprintf("%v", got)
			load.Rules = append(load.Rules, str)
		}
	}
	inf, error := json.Marshal(load)
	if error != nil {
		logger.Println(formatLog("Error Marshalling mem metric data"))
		return string(""), error
	}
	return string(inf), error
}

func loadAvg() (string, error) {
	l, loaderr := load.Avg()
	if loaderr != nil {
		logger.Println(formatLog("Error getting config"))
		return string(""), loaderr
	}
	token := gjson.Get(configfile, "token").String()
	agent := gjson.Get(configfile, "agent").String()
	var lAvg rLoad
	var linfo rinfo
	lAvg.Agent = agent
	lAvg.Token = token
	lAvg.Metric = "loadavg"
	linfo.Load1 = l.Load1
	linfo.Load5 = l.Load5
	linfo.Load15 = l.Load15
	lAvg.Data = linfo
	f, _ := json.Marshal(lAvg)
	return string(f), loaderr

}
