package main

import (
	"github.com/Ben-GO-package/acmg_anno/anno"
	"github.com/Ben-GO-package/acmg_anno/evidence"
	"github.com/brentp/bix"
	"github.com/liserjrqlxue/goUtil/jsonUtil"
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
			acmgCfg[k] = anno.GuessPath(v, cfgPath) // 猜测并修正配置项的路径
		}
		Database_Init(acmgCfg, *autoPVS1, *runPM1) // 使用处理后的配置初始化ACMG模块
	}
}

var (
	tbx            *bix.Bix
	LOFGeneList    map[string]int
	transcriptInfo map[string][]evidence.Region
)

func Database_Init(cfg map[string]string, AutoPVS1 bool, runPM1 bool) {
	evidence.LoadPS1PM5Database(cfg["PS1PM5.database"])
	if runPM1 {
		evidence.LoadPM1(cfg["PM1InterproDomain"], cfg["PM1PfamIdDomain"])
	}
	if !AutoPVS1 {
		LOFGeneList = evidence.LoadLOF(cfg["LOFList"])
		jsonUtil.JsonFile2Data(cfg["transcriptInfo"], &transcriptInfo)
	}
	if runPM1 || !AutoPVS1 {
		tbx = simpleUtil.HandleError(bix.New(cfg["PathogenicLite"])).(*bix.Bix)
	}
	evidence.LoadPP2(cfg["PP2GeneList"])
	evidence.LoadBS2(cfg["LateOnset"])
	evidence.LoadBP1(cfg["BP1GeneList"])
	evidence.LoadBA1(cfg["BA1ExceptionList"])
}

func AddEvidences(item map[string]string, AutoPVS1 bool, runPM1 bool) {
	if !AutoPVS1 {
		item["PVS1"] = evidence.CheckPVS1(item, LOFGeneList, transcriptInfo, tbx)
	}
	item["PS1"] = evidence.CheckPS1(item)
	item["PM5"] = evidence.CheckPM5(item)
	item["PS4"] = evidence.CheckPS4(item)
	if runPM1 {
		item["PM1"] = evidence.CheckPM1(item, tbx)
	}
	item["PM2"] = evidence.CheckPM2(item)
	item["PM4"] = evidence.CheckPM4(item, AutoPVS1)
	item["PP2"] = evidence.CheckPP2(item)
	item["PP3"] = evidence.CheckPP3(item, AutoPVS1)
	item["BA1"] = evidence.CheckBA1(item)
	item["BS1"] = evidence.CheckBS1(item)
	item["BS2"] = evidence.CheckBS2(item)
	item["BP1"] = evidence.CheckBP1(item)
	item["BP3"] = evidence.CheckBP3(item)
	item["BP4"] = evidence.CheckBP4(item)
	item["BP7"] = evidence.CheckBP7(item)
}

func annotate1(item map[string]string) {

	// score to prediction ,get item["dbscSNV_ADA_pred"]
	anno.Score2Pred(item)

	// ues acmg of go
	if *acmg {
		AddEvidences(item, *autoPVS1, *runPM1)
		item["automated_judgment"] = evidence.PredACMG2015(item, *autoPVS1, *runPM1)
	}

	anno.UpdateSnv(item)

	anno.UpdateAutoRule(item)

	stats[item["#Chr"]]++
	stats[item["VarType"]]++
}
