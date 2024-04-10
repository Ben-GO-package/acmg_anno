package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/liserjrqlxue/goUtil/textUtil"
)

// flag
var (
	acmgDb = flag.String(
		"acmgDb",
		"",
		"config file for ACMG database (default : cfg\\acmg.db.cfg)",
	)
	snv = flag.String(
		"snv",
		"",
		"input snv anno txt, comma as sep",
	)
	prefix = flag.String(
		"prefix",
		"",
		"output prefix, default is same to first file of -snv",
	)
	logfile = flag.String(
		"log",
		"",
		"output log to log.log ( prefix.acmg.log)",
	)
	temp_title = flag.String(
		"temp_title",
		"",
		"config file for final title of result (default : cfg\\final_result_title.cfg)",
	)
)

// bool flag
var (
	acmg = flag.Bool(
		"acmg",
		false,
		"if use new ACMG, fix PVS1, PS1,PS4, PM1,PM2,PM4,PM5 PP2,PP3, BA1, BS1,BS2, BP1,BP3,BP4,BP7",
	)
	autoPVS1 = flag.Bool(
		"autoPVS1",
		false,
		"if use autoPVS1 for acmg",
	)

	outTsv = flag.Bool(
		"tsv",
		true,
		"if output tsv format result ( prefix.acmg.tsv)",
	)
	outJson = flag.Bool(
		"json",
		false,
		"if output json format result  ( prefix.acmg.json)",
	)
)

func checkFlag() {
	if *snv == "" {
		flag.Usage()
		fmt.Println("\n -snv is required ")
		os.Exit(0)
	} else {
		snvs = strings.Split(*snv, ",")
		if *prefix == "" {
			*prefix = snvs[0]
		}
	}
	if *logfile == "" {
		*logfile = *prefix + ".acmg.log"
	}

	// 设置ACMG数据库路径
	if *acmgDb == "" {
		*acmgDb = filepath.Join(cfgPath, "acmg.db.cfg")
	}
	if *temp_title != "" {
		// 解析最终输出字段
		TempOutputTitle = textUtil.File2Array(*temp_title)
	}

}
