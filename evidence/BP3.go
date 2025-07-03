package evidence

import (
	"strconv"
)

const BayesDelThreshold = -0.18 // 定义BayesDel_noAF_score阈值

// ture	:	"1"
// flase:	"0"
func CheckBP3(item map[string]string) string {
	if isBP3PM4Func.MatchString(item["Function"]) && item["VarType"] != "snv" {
		if item["RepeatTag"] == "" || item["RepeatTag"] == "." || item["RepeatTag"] == "-" {
			return "0"
		} else {
			BayesDel_noAF_score_float, _ := strconv.ParseFloat(item["BayesDel_noAF_score"], 64)
			if BayesDel_noAF_score_float <= BayesDelThreshold {
				return "1"
			}
		}
	}
	return "0"
}

func CompareBP3(item map[string]string) {
	rule := "BP3"
	val := CheckBP3(item)
	if val != item[rule] {
		if item[rule] == "0" && val == "" {
		} else {
			PrintConflict(item, rule, val, "Function", "RepeatTag", "VarType")
		}
	}
}
