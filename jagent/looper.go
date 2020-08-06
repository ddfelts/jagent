package main

import (
	"fmt"
	"strconv"
	"time"
)

func formatLog(d string) string {

	t := time.Now()
	finalstr := t.Format("2006-01-02 15:04:05") + " : [jagent] : " + d
	return finalstr
}

func doSendPost(c chan string) {
	for {
		msg := <-c
		_, senderr := sendPost(msg)
		if senderr != nil {
			logger.Println(formatLog(senderr.Error()))
		} else {
			sD := len(msg)
			sizeofdata := strconv.FormatInt(int64(sD), 10)
			logger.Println(formatLog("Data sent to server: size(" + sizeofdata + ")"))
		}
	}
}

func doCPU(c chan string) {
	cpuload, cpuloaderror := loadAvg()
	if cpuloaderror == nil {
		c <- cpuload
	} else {
		logger.Println(cpuloaderror)
	}
}

func doNetwork(c chan string) {
	netI, netIerror := netWork()
	if netIerror == nil {
		c <- netI
	} else {
		logger.Println(netIerror)
	}
}

func doHost(c chan string) {
	hosti, hostierror := hostInfo()
	if hostierror == nil {
		c <- hosti
	} else {
		logger.Println(hostierror)
	}
}

func doMem(c chan string) {
	memI, memIerror := memData()
	if memIerror == nil {
		c <- memI
	} else {
		logger.Println(memIerror)
	}
}

func doDisk(c chan string) {
	diskI, diskIerror := diskU()
	if diskIerror == nil {
		c <- diskI
	} else {
		logger.Println(diskIerror)
	}
}

func doCPUInfo(c chan string) {
	cpuID, cpuIderror := cpuinfo()
	if cpuIderror == nil {
		c <- cpuID
	} else {
		logger.Println(cpuIderror)
	}
}

func doOnce(c chan string) {
	doHost(c)
	doNetwork(c)
	doCPUInfo(c)
	doCPU(c)
	doMem(c)
	doDisk(c)
}

func doMetricMem(c chan string) {
	mMet, mError := memMetric()
	if mError == nil {
		c <- mMet
	} else {
		str := fmt.Sprintf("%v", mError)
		logger.Println(formatLog(str))
	}
}

func doLoadMem(c chan string) {
	ml, me := loadAvgM()
	if me == nil {
		c <- ml
	} else {
		str := fmt.Sprintf("%v", me)
		logger.Println(formatLog(str))
	}
}

func doDiskM(c chan string) {
	ml, me := mdiskU()
	if me == nil {
		c <- ml
	} else {
		str := fmt.Sprintf("%v", me)
		logger.Println(formatLog(str))
	}
}

func doEvery(c chan string, t time.Duration) {
	for {
		time.Sleep(t * time.Second)
		doMetricMem(c)
		doLoadMem(c)
		doDiskM(c)
	}
}
