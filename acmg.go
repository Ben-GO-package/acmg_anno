package main

import (
	"github.com/Ben-GO-package/acmg2015"
	"github.com/Ben-GO-package/auto_acmg/anno"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
	"github.com/liserjrqlxue/goUtil/textUtil"
)

// initAcmg2015 初始化ACMG相关的配置。
// 这个函数不接受参数，也不返回任何值。
// 它主要根据全局标志*acmg的值来决定是否进行初始化。
// 如果*acmg标志为真，则会读取ACMG数据库文件，进行路径解析，并初始化ACMG模块。
func initAcmg2015() {
	if *acmg {
		acmg2015.AutoPVS1 = *autoPVS1 // 设置自动PVS1的值
		// 从文件中读取ACMG配置，并处理每一项的路径，确保它们相对于数据库路径是正确的
		var acmgCfg = simpleUtil.HandleError(textUtil.File2Map(*acmgDb, "\t", false)).(map[string]string)
		for k, v := range acmgCfg {
			acmgCfg[k] = anno.GuessPath(v, cfgPath) // 猜测并修正配置项的路径
		}
		acmg2015.Init(acmgCfg) // 使用处理后的配置初始化ACMG模块
	}
}

func annotate1(item map[string]string) {
	// inhouse_AF -> frequency
	item["frequency"] = item["inhouse_AF"]

	// score to prediction ,get item["dbscSNV_ADA_pred"]
	anno.Score2Pred(item)

	// update Function update item["Function"]
	anno.UpdateFunction(item)

	item["Gene"] = item["Omim Gene"]
	item["OMIM"] = item["OMIM_Phenotype_ID"]

	//anno.ParseSpliceAI(item)

	// ues acmg of go
	if *acmg {
		if item["cHGVS_org"] == "" {
			item["cHGVS_org"] = item["cHGVS"]
		}
		acmg2015.AddEvidences(item)
		item["自动化判断"] = acmg2015.PredACMG2015(item, *autoPVS1)
	}

	anno.UpdateSnv(item)

	anno.UpdateAutoRule(item)

	stats[item["#Chr"]]++
	stats[item["VarType"]]++
}
