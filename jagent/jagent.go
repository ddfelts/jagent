package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

// Agent log file
var (
	outfile, _ = os.Create("agent.log")
	logger     = log.New(outfile, "", 0)
	r          string
	c          string
	configfile string
)

// PostRet : Used to store the return value from the web server
type PostRet struct {
	Status string `json:"status"`
}

// Config : Used to store the structor of the config file
type Config struct {
	URL       string   `json:"url"`
	Token     string   `json:"token"`
	Agent     string   `json:"agent"`
	Delay     int64    `json:"delaySec"`
	Processes []string `json:"processes"`
	Logs      []string `json:"logs"`
}

// ReadConfig : Read the configuration file
func ReadConfg(cfile string) {
	content, err := ioutil.ReadFile(cfile)
	if err != nil {
		logger.Println(err)
	} else {
		logger.Println(formatLog("Config parsed for configuration"))
	}
	configfile = string(content)
}

// main : used to launch main program
func main() {
	flag.StringVar(&r, "r", "", "Register Agent with Server")
	flag.StringVar(&c, "c", "config.json", "Defined config file")
	flag.Parse()
	var sInfo chan string = make(chan string)
	ReadConfg(c)
	for i := 0; i < 3; i++ {
		tstring := "SendPost thread Started:" + strconv.FormatInt(int64(i), 10)
		logger.Println(formatLog(tstring))
		go doSendPost(sInfo)
	}
	delay := gjson.Get(configfile, "delaySec").Int()
	logger.Println(formatLog("Setting delay to: " + strconv.FormatInt(int64(delay), 10) + " Seconds"))
	doOnce(sInfo)
	doEvery(sInfo, time.Duration(delay))
}
