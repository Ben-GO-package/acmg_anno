package evidence

import (
	"regexp"

	"github.com/liserjrqlxue/goUtil/textUtil"
)

func FindPathogenicMissense(fileName, key string, pathogenicRegexp *regexp.Regexp) (map[string]int, map[string]int, map[string]int) {
	var varList = make(map[string]int)
	var pHGVSList = make(map[string]int)
	var pPosList = make(map[string]int)
	itemArray, _ := textUtil.File2MapArray(fileName, "\t", nil)
	for _, item := range itemArray {
		if !pathogenicRegexp.MatchString(item[key]) {
			continue
		}
		if !ismissense.MatchString(item["Function"]) {
			continue
		}
		var key = item["Transcript"] + ":" + item["pHGVS"]
		pHGVSList[key]++
		varList[item["MutationName"]]++
		AAPos := getAAPos.FindString(item["pHGVS"])
		if AAPos != "" {
			pPosList[item["Transcript"]+":"+AAPos]++
		}
	}
	return varList, pHGVSList, pPosList
}

// PS1
func CheckPS1(item map[string]string) (string, string) {
	if !PS1Function.MatchString(item["Function"]) {
		return "0", "-"
	}
	var trans_chgvs = item["Transcript"] + ":" + item["cHGVS"]
	var trans_phgvs = item["Transcript"] + ":" + item["pHGVS1"]
	var countHGVS = hgvsCount[trans_chgvs]
	var countPHGVS = phgvsCount[trans_phgvs]
	//fmt.Printf("%s\t%s\t%d\t%d\n", trans_chgvs, trans_phgvs, countHGVS,countPHGVS)
	if countPHGVS > countHGVS {
		var dataPHGVS = phgvsdb[trans_phgvs]
		return "1", dataPHGVS
	} else {
		return "0", "-"
	}
}

func ComparePS1(item map[string]string, ClinVarMissense, ClinVarPHGVSlist, HGMDMissense, HGMDPHGVSlist map[string]int) {
	rule := "PS1"
	val, _ := CheckPS1(item)
	if val != item[rule] {
		PrintConflict(item, rule, val, "Function", "Transcript", "pHGVS")
	}
}
