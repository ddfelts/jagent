package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func sendPost(mdata string) (string, error) {
	URL := gjson.Get(configfile, "url").String()
	njson := []byte(mdata)
	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(njson))
	if err != nil {
		logger.Println(formatLog("Error connecting to web server"))
		return string(""), err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Println(formatLog("Error in response from web server"))
		return string(""), err
	}
	return string(body), err
}
