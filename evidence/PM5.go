package evidence

import (
	"strconv"
	"strings"
)

// PM5
func CheckPM5(item map[string]string) (string, string) {
	if !ismissense.MatchString(item["Function"]) {
		return "0", "-"
	}
	var trans_phgvs = item["Transcript"] + ":" + item["pHGVS1"]
	var trans_aapos = item["Transcript"] + ":" + item["Protein_position"]
	var countPHGVS = phgvsCount[trans_phgvs]
	var countAAPos = aaPostCount[trans_aapos]
	//fmt.Printf("%s\t%s\t%d\t%d\n", trans_aapos, trans_phgvs, countAAPos, countPHGVS)

	if countAAPos > countPHGVS {
		var dataAAPos = aaPostdb[trans_aapos]

		itemRevel := item["REVEL"]
		itemGrantham := item["Grantham"]

		var itemRevelVal, itemGranthamVal float64
		if itemRevel != "NA" {
			if val, err := strconv.ParseFloat(itemRevel, 64); err == nil {
				itemRevelVal = val
			}
		}

		if itemGrantham != "NA" {
			if val, err := strconv.ParseFloat(itemGrantham, 64); err == nil {
				itemGranthamVal = val
			}
		}

		currentItemID := item["Transcript"] + "|" + item["Protein_position"] + "|" + item["cHGVS"] + "|" + item["pHGVS1"]

		records := strings.Split(dataAAPos, ";")
		for _, record := range records {
			if strings.HasPrefix(record, currentItemID) {
				continue
			}
			fields := strings.Split(record, "|")
			if len(fields) >= 7 {
				revelField := fields[4]
				revelParts := strings.Split(revelField, "=")
				if len(revelParts) == 2 && revelParts[0] == "Revel" && revelParts[1] != "NA" {
					revelValue := revelParts[1]
					if revelValue == "-" {
						revelValue = "0"
					}
					if dbRevel, err := strconv.ParseFloat(revelValue, 64); err == nil {
						if dbRevel <= itemRevelVal {
							return "1", dataAAPos
						}
					}
				}

				granthamField := fields[5]
				granthamParts := strings.Split(granthamField, "=")
				if len(granthamParts) == 2 && granthamParts[0] == "Grantham" && granthamParts[1] != "NA" {
					granthamValue := granthamParts[1]
					if granthamValue == "-" {
						granthamValue = "0"
					}
					if dbGrantham, err := strconv.ParseFloat(granthamValue, 64); err == nil {
						if dbGrantham <= itemGranthamVal {
							return "1", dataAAPos
						}
					}
				}
			}
		}


		return "0", dataAAPos
	} else {
		return "0", "-"
	}
}

func ComparePM5(item map[string]string, ClinVarPHGVSlist, ClinVarAAPosList, HGMDPHGVSlist, HGMDAAPosList map[string]int) {
	rule := "PM5"
	val, _ := CheckPM5(item)
	if val != item[rule] {
		PrintConflict(item, rule, val, "Function", "Transcript", "pHGVS")
	}
}
