package evidence

// PM5
func CheckPM5(item map[string]string) string {
	if !ismissense.MatchString(item["Function"]) {
		return "0"
	}
	var trans_phgvs = item["Transcript"] + ":" + item["pHGVS1"]
	var trans_aapos = item["Transcript"] + ":" + item["Protein_position"]
	var countPHGVS = phgvsCount[trans_phgvs]
	var countAAPos = aaPostCount[trans_aapos]
	//fmt.Printf("%s\t%s\t%d\t%d\n", trans_aapos, trans_phgvs, countAAPos, countPHGVS)

	if countAAPos > countPHGVS {
		return "1"
	} else {
		return "0"
	}
}

func ComparePM5(item map[string]string, ClinVarPHGVSlist, ClinVarAAPosList, HGMDPHGVSlist, HGMDAAPosList map[string]int) {
	rule := "PM5"
	val := CheckPM5(item)
	if val != item[rule] {
		PrintConflict(item, rule, val, "Function", "Transcript", "pHGVS")
	}
}
