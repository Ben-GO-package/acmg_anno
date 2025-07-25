package main

import (
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"

	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

// init函数是程序启动时调用的初始化函数。
// 该函数主要完成以下功能：
// 1. 打印版本信息。
// 2. 解析命令行参数。
// 3. 初始化日志系统，包括创建日志文件并设置日志格式。
// 4. 解析配置文件。
// 5. 读取并处理函数级别配置。
// 6. 处理产品英文列表，并标记当前产品是否为英文产品。
func init() {
	// 打印版本信息
	gitDescribe := "https://gitlab.genomics.cn/bi-procreate/acmg"
	buildStamp := "2025.7.8"
	golangVersion := "1.21.0"
	fmt.Printf("acmg_anno      :'v1.0.6 @ 2025.7.8'\n")
	Version(gitDescribe, buildStamp, golangVersion)

	// 解析命令行参数
	flag.Parse()
	// 检查命令行参数是否合法
	checkFlag()

	// 初始化日志系统
	logFile, err = os.Create(*logfile)
	simpleUtil.CheckErr(err)
	log.Printf("Log file         : %v\n", *logfile)
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)
	log.Printf("Log file         : %v\n", *logfile)
	LogVersion(gitDescribe, buildStamp, golangVersion)

	initAcmg2015()
	//fmt.Print("Finish ACMG Init : \n")
}

func main() {
	defer simpleUtil.DeferClose(logFile)
	// anno
	if *snv != "" {
		var data, title = loadData()
		finalOutputTitle = title
		check_transverTitle_relation_map(finalOutputTitle)
		finalOutputTitle = append(finalOutputTitle, "autoRuleName", "automated_judgment")
		if *outpred {
			finalOutputTitle = append(finalOutputTitle, "dbscSNV_ADA_pred", "dbscSNV_RF_pred", "GERP_RS_pred", "PhyloP_Vertebrates_Pred", "PhyloP_Placental_Mammals_Pred")
		}
		fmt.Print("Finish Mutation Loading : ", len(data), "\n")

		stats["Total"] = len(data)
		for _, raw_item := range data {
			item := transverTitle(raw_item)
			annotate1(item)
			raw_item["autoRuleName"] = item["autoRuleName"]
			raw_item["automated_judgment"] = item["automated_judgment"]
			raw_item["GERP_RS_pred"] = item["GERP++_RS_pred"]
			raw_item["dbscSNV_ADA_pred"] = item["dbscSNV_ADA_pred"]
			raw_item["dbscSNV_RF_pred"] = item["dbscSNV_RF_pred"]
			raw_item["PhyloP_Vertebrates_Pred"] = item["PhyloP Vertebrates Pred"]
			raw_item["PhyloP_Placental_Mammals_Pred"] = item["PhyloP Placental Mammals Pred"]

			for _, col := range TempOutputTitle {
				_, exists := raw_item[col]
				temp_value, temp_exists := item[col]
				if !exists && temp_exists {
					raw_item[col] = temp_value
				}
			}
			WholeResultData = append(WholeResultData, raw_item)

			cycle1Count++
			if cycle1Count%1000 == 0 {
				log.Printf("cycle1 progress %d/%d", cycle1Count, len(data))
			}
		}
		logTierStats(stats)
		logTime("update info")

	}
	// Update By Liu.Bo @  2024/03/15 15:22:30 增加tsv格式输出，为便于观察增加*import.tsv仅输出两个最终需求字段(autoRuleName	自动化判断)确保后续精简
	if *outTsv {
		// 输出特定字段格式的tier1.tsv
		mapArray2tsv(WholeResultData, finalOutputTitle, *prefix+".acmg.tsv")
		mapArray2tsv(WholeResultData, TempOutputTitle, *prefix+".acmg.temp.tsv")
	}
	// 输出json
	if *outJson {
		if *snv != "" {
			// hash array 输出 json list
			mapArray2jsonList(WholeResultData, *prefix+".acmg.json")
		}
	}
}
