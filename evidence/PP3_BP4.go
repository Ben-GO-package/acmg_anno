package evidence

import (
	"strconv"
)

// ture	:	"1"
// flase:	"0"

func CheckBayesDel(BayesDel_noAF_score_float float64) (string, string) { //PP3,BP4
	if BayesDel_noAF_score_float > -0.36 && BayesDel_noAF_score_float <= -0.18 {
		return "0", "1"
	} else if BayesDel_noAF_score_float <= -0.36 {
		return "0", "Moderate"
	} else if BayesDel_noAF_score_float >= 0.13 && BayesDel_noAF_score_float < 0.27 {
		return "1", "0"
	} else if BayesDel_noAF_score_float >= 0.27 && BayesDel_noAF_score_float < 0.5 {
		return "Moderate", "0"
	} else if BayesDel_noAF_score_float >= 0.5 {
		return "Strong", "0"
	}
	return "0", "0"
}

func CheckREVEL(REVEL_score_float float64) (string, string) { //PP3,BP4
	if REVEL_score_float >= 0.932 {
		return "Strong", "0"
	}
	if REVEL_score_float >= 0.773 && REVEL_score_float < 0.932 {
		return "Moderate", "0"
	}
	if REVEL_score_float >= 0.644 && REVEL_score_float < 0.773 {
		return "1", "0"
	}
	if REVEL_score_float >= 0.183 && REVEL_score_float <= 0.290 {
		return "0", "1"
	}
	if REVEL_score_float > 0.016 && REVEL_score_float <= 0.183 {
		return "0", "Moderate"
	}
	if REVEL_score_float <= 0.016 {
		return "0", "Strong"
	}
	return "0", "0"
}

func CheckPP3_BP4(item map[string]string) (string, string) {
	PP3 := "0"
	BP4 := "0"
	if isNoPP3BP4.MatchString(item["Function"]) {
		return PP3, BP4
	}
	if item["SpliceAI_Max_Score"] == "-" {
		//fmt.Printf("SpliceAI_Max_Score_float >0.100000: %f", SpliceAI_Max_Score_float)
		if ismissense.MatchString(item["Function"]) {
			if item["REVEL"] == "-" { // REVEL 缺失
				if item["BayesDel_noAF_score"] == "-" {
					return "0", "0"
				} else {
					BayesDel_noAF_score_float, _ := strconv.ParseFloat(item["BayesDel_noAF_score"], 64)
					PP3, BP4 = CheckBayesDel(BayesDel_noAF_score_float)
					return PP3, BP4
				}
			} else {
				REVEL_score_float, _ := strconv.ParseFloat(item["REVEL"], 64)
				PP3, BP4 = CheckREVEL(REVEL_score_float)
				return PP3, BP4
			}
		} else {
			return "0", "0"
		}
	} else {
		SpliceAI_Max_Score_float, _ := strconv.ParseFloat(item["SpliceAI_Max_Score"], 64)
		SpliceAI_Max_Score_float = SpliceAI_Max_Score_float + 0.0
		//fmt.Printf("SpliceAI_Max_Score_float: %f", SpliceAI_Max_Score_float)
		if SpliceAI_Max_Score_float >= 0.2 {
			if ismissense.MatchString(item["Function"]) {
				if item["REVEL"] == "-" { // REVEL 缺失
					if item["BayesDel_noAF_score"] == "-" {
						return PP3, BP4
					} else {
						BayesDel_noAF_score_float, _ := strconv.ParseFloat(item["BayesDel_noAF_score"], 64)
						if BayesDel_noAF_score_float < 0.27 {
							return "1", "0"
						} else if BayesDel_noAF_score_float >= 0.27 && BayesDel_noAF_score_float < 0.5 {
							return "Moderate", "0"
						} else if BayesDel_noAF_score_float >= 0.5 {
							return "Strong", "0"
						}
					}
				} else {
					REVEL_score_float, _ := strconv.ParseFloat(item["REVEL"], 64)
					if REVEL_score_float < 0.773 {
						return "1", "0"
					} else if REVEL_score_float < 0.932 {
						PP3 = "Moderate"
						return "Moderate", "0"
					} else if REVEL_score_float >= 0.932 {
						PP3 = "Strong"
						return "Strong", "0"
					}
				}
			} else {
				PP3 = "1"
				return PP3, BP4
			}

		} else if SpliceAI_Max_Score_float > 0.1 {
			//fmt.Printf("SpliceAI_Max_Score_float >0.100000: %f", SpliceAI_Max_Score_float)
			if ismissense.MatchString(item["Function"]) {
				if item["REVEL"] == "-" { // REVEL 缺失
					if item["BayesDel_noAF_score"] == "-" {
						return "0", "0"
					} else {
						BayesDel_noAF_score_float, _ := strconv.ParseFloat(item["BayesDel_noAF_score"], 64)
						PP3, BP4 = CheckBayesDel(BayesDel_noAF_score_float)
						return PP3, BP4
					}
				} else {
					REVEL_score_float, _ := strconv.ParseFloat(item["REVEL"], 64)
					PP3, BP4 = CheckREVEL(REVEL_score_float)
					return PP3, BP4
				}
			} else {
				return "0", "0" // ToDo 提供的逻辑缺失
			}
		} else if SpliceAI_Max_Score_float <= 0.1 {
			//fmt.Printf("SpliceAI_Max_Score_float <=0.100000: %f", SpliceAI_Max_Score_float)
			if ismissense.MatchString(item["Function"]) {
				if item["REVEL"] == "-" { // REVEL 缺失
					if item["BayesDel_noAF_score"] == "-" {
						return "0", "0"
					} else {
						BayesDel_noAF_score_float, _ := strconv.ParseFloat(item["BayesDel_noAF_score"], 64)
						PP3, BP4 = CheckBayesDel(BayesDel_noAF_score_float)
						return PP3, BP4
					}
				} else {
					REVEL_score_float, _ := strconv.ParseFloat(item["REVEL"], 64)
					PP3, BP4 = CheckREVEL(REVEL_score_float)
					return PP3, BP4
				}
			} else {
				return "0", "1"
			}
		}

	}
	return "0", "0"
}
