package main

import (
	"log"

	"github.com/Ben-GO-package/acmg_anno/anno"
	"github.com/Ben-GO-package/acmg_anno/evidence"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
	"github.com/liserjrqlxue/goUtil/textUtil"
)

// initAcmg2015 初始化ACMG相关的配置。
// 这个函数不接受参数，也不返回任何值。
// 它主要根据全局标志*acmg的值来决定是否进行初始化。
// 如果*acmg标志为真，则会读取ACMG数据库文件，进行路径解析，并初始化ACMG模块。
func initAcmg2015() {
	if *acmg {
		// 从文件中读取ACMG配置，并处理每一项的路径，确保它们相对于数据库路径是正确的
		var acmgCfg = simpleUtil.HandleError(textUtil.File2Map(*acmgDb, "\t", false)).(map[string]string)
		for k, v := range acmgCfg {
			log.Printf("cfg[%s]:%s\n", k, v)
			acmgCfg[k] = anno.GuessPath(v, cfgPath) // 猜测并修正配置项的路径
		}
		evidence.Database_Init(acmgCfg, *autoPVS1, *runPM1) // 使用处理后的配置初始化ACMG模块
	}
}

func annotate1(item map[string]string) {

	// score to prediction ,get item["dbscSNV_ADA_pred"]
	anno.Score2Pred(item)

	// ues acmg of go
	if *acmg {
		evidence.AddEvidences(item, *autoPVS1, *runPM1)
		item["automated_judgment"] = evidence.PredACMG2015(item, *autoPVS1, *runPM1)
	}

	anno.UpdateSnv(item)
	anno.UpdateAutoRule(item)

	stats[item["#Chr"]]++
	stats[item["VarType"]]++
}
