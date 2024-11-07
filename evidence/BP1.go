package evidence

// BP1
func CheckBP1(item map[string]string) string {
	if !ismissense.MatchString(item["Function"]) {
		return "0"
	}
	if bp1GeneList[item["entrez_id"]] {
		return "1"
	} else {
		return "0"
	}
}

func CompareBP1(item map[string]string, ClinVarBP1GeneList, HgmdBP1GeneList map[string]float64) {
	rule := "BP1"
	val := CheckBP1(item)
	if val != item[rule] {
		PrintConflict(item, rule, val, "Function", "entrez_id")
	}
}
