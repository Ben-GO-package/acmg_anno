package evidence

import (
	"github.com/liserjrqlxue/goUtil/jsonUtil"
)

var (
	transcriptInfo map[string][]Region
)

func Database_Init(cfg map[string]string, AutoPVS1 bool, runPM1 bool) {
	LoadPS1PM5Database(cfg["PS1PM5.database"])
	if !AutoPVS1 {
		LOFGeneList = LoadLOF(cfg["LOFList"])
		jsonUtil.JsonFile2Data(cfg["transcriptInfo"], &transcriptInfo)
	}
	LoadPP2PM1_special(cfg["PP2PM1_special"])
	LoadPP2(cfg["PP2GeneList"])
	LoadBS2(cfg["LateOnset"])
	LoadBP1(cfg["BP1GeneList"])
	LoadBA1(cfg["BA1ExceptionList"])
}

func AddEvidences(item map[string]string, AutoPVS1 bool, runPM1 bool) {
	if !AutoPVS1 {
		item["PVS1"] = "0" //CheckPVS1(item, LOFGeneList, transcriptInfo, tbx)
	}
	item["PS1"], item["PS1_evidence"] = CheckPS1(item)
	item["PM5"], item["PM5_evidence"] = CheckPM5(item)
	item["PS4"] = CheckPS4(item)
	if runPM1 {
		item["PM1"] = CheckPM1(item)
	}
	if item["PM1"] != "1" && item["PM1"] != "Supporting" { // PM1 和PP2不共用
		item["PP2"] = CheckPP2(item)
	} else {
		item["PP2"] = "-1"
	}
	item["PM2"] = CheckPM2(item)
	item["PM4"] = CheckPM4(item, AutoPVS1)
	//item["PP3"] = CheckPP3(item, AutoPVS1)
	item["BA1"] = CheckBA1(item)
	item["BS1"] = CheckBS1(item)
	item["BS2"] = CheckBS2(item)
	item["BP1"] = CheckBP1(item)
	item["BP3"] = CheckBP3(item)
	item["PP3"], item["BP4"] = CheckPP3_BP4(item)
	item["BP7"] = CheckBP7(item)
}
