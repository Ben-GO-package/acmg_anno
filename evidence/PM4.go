package evidence

import "regexp"

// ture	:	"1"
// flase:	"0"
// nil	:	""
func CheckPM4(item map[string]string, autoPVS1 bool) string {
	if autoPVS1 {
		switch item["AutoPVS1 Adjusted Strength"] {
		case "VeryStrong":
			return "0"
		case "Strong":
			return "0"
		case "Moderate":
			return "0"
		case "Supporting":
			return "0"
		}
	} else if item["PVS1"] == "1" {
		return "0"
	}
	var isPM4Func = regexp.MustCompile(`cds-del|cds-ins|cds-indel|stop-loss|inframe_deletion|inframe_insertion|stop_lost|protein_altering_variant`)
	if isPM4Func.MatchString(item["Function"]) {
		if item["RepeatTag"] == "" || item["RepeatTag"] == "." || item["RepeatTag"] == "-" {
			return "1"
		} else {
			return "0"
		}
	} else {
		return "0"
	}
}

// func ComparePM4(item map[string]string) {
// 	rule := "PM4"
// 	val := CheckPM4(item, autoPVS1)
// 	if val != item[rule] {
// 		PrintConflict(item, rule, val, "Function", "RepeatTag")
// 	}
// }
