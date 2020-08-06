package main

import (
	"encoding/json"

	"github.com/HuanTeng/go-jsonlogic"
)

//ruleprocessor : used to process rules for errors.
func ruleprocessor(prule string, MNinf []byte) (interface{}, error) {
	var rule interface{}
	var rdata interface{}
	json.Unmarshal([]byte(prule), &rule)
	json.Unmarshal(MNinf, &rdata)
	got, err := jsonlogic.Apply(rule, rdata)
	if err != nil {
		return got, err
	}
	return got, err
}
