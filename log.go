package main

import (
	"fmt"
	"log"
	"os"
	"time"

	simple_util "github.com/liserjrqlxue/simple-util"
)

func logTierStats(stats map[string]int) {
	log.Printf("Total              Count : %7d\n", stats["Total"])
	if stats["Total"] == 0 {
		return
	}
	log.Printf("VarType:snv        Hit   : %7d\n", stats["snv"])
	log.Printf("VarType:ins        Hit   : %7d\n", stats["ins"])
	log.Printf("VarType:del        Hit   : %7d\n", stats["del"])

}

func logTime(message string) {
	ts = append(ts, time.Now())
	step++
	var trim = 3*8 - 1
	var str = simple_util.FormatWidth(trim, message, ' ')
	log.Printf("%s\ttook %7.3f/%7.3fs to run.\n", str, ts[step].Sub(ts[step-1]).Seconds(), ts[step].Sub(ts[0]).Seconds())
}

func LogVersion(gitDescribe string, buildStamp string, golangVersion string) {
	log.Printf("Build Time     : %s\n", buildStamp)
	log.Printf("Golang Version : %s\n", golangVersion)
	log.Printf("Git Describe   : %s\n", gitDescribe)
	var hostName, err = os.Hostname()
	log.Printf("Current Host   : %s%v\n", hostName, err)
}

// Version print out version info
func Version(gitDescribe string, buildStamp string, golangVersion string) {
	fmt.Printf("Build Time     : %s\n", buildStamp)
	fmt.Printf("Golang Version : %s\n", golangVersion)
	fmt.Printf("Git Describe   : %s \n", gitDescribe)
	var hostName, err = os.Hostname()
	fmt.Printf("Current Host   : %s%v\n", hostName, err)
}
