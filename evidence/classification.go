package evidence

// output B LB VUS LP P
func PredACMG2015(item map[string]string, autoPVS1 bool, runPM1 bool) string {
	var sumPVS int
	var sumPS int
	var sumPM int
	var sumPP int
	var sumBA int
	var sumBS int
	var sumBM int
	var sumBP int
	PVS1 := item["PVS1"]
	// autoPVS1
	// 若已使用PVS1，则不可同时使用PP3
	// 任一强度的PVS1与PM4 不可共用的证据

	if autoPVS1 {
		switch item["AutoPVS1 Adjusted Strength"] {
		case "VeryStrong":
			sumPVS++
			item["PP3"] = "-1"
			item["PM4"] = "-1"
			item["BP4"] = "-1"
		case "Strong":
			sumPS++
			item["PP3"] = "-1"
			item["PM4"] = "-1"
			item["BP4"] = "-1"
		case "Moderate":
			sumPM++
			item["PP3"] = "-1"
			item["PM4"] = "-1"
			item["BP4"] = "-1"
		case "Supporting":
			sumPP++
			item["PP3"] = "-1"
			item["PM4"] = "-1"
			item["BP4"] = "-1"
		}
	} else {
		if PVS1 == "1" {
			sumPVS++
			item["PP3"] = "-1"
			item["PM4"] = "-1"
			item["BP4"] = "-1"
		}
	}
	// PP3 不与 PM4 共同得分
	if item["PM4"] == "1" {
		item["PP3"] = "-1"
	}
	// BP3 不与 BP4 共同得分
	if item["BP3"] == "1" {
		item["BP4"] = "-1"
	}
	PS1 := item["PS1"]
	PS2 := item["PS2"]
	PS3 := item["PS3"]
	PS4 := item["PS4"]
	PM1 := "0"
	if runPM1 {
		PM1 = item["PM1"]
	}
	PM2 := item["PM2"]
	PM3 := item["PM3"]
	PM4 := item["PM4"]
	PM5 := item["PM5"]
	PM6 := item["PM6"]

	PP1 := item["PP1"]
	PP2 := item["PP2"]
	PP3 := item["PP3"]
	PP4 := item["PP4"]
	PP5 := item["PP5"]

	BA1 := item["BA1"]

	BS1 := item["BS1"]
	BS2 := item["BS2"]
	BS3 := item["BS3"]
	BS4 := item["BS4"]

	BP1 := item["BP1"]
	BP2 := item["BP2"]
	BP3 := item["BP3"]
	BP4 := item["BP4"]
	BP5 := item["BP5"]
	BP6 := item["BP6"]
	BP7 := item["BP7"]

	// PVS
	//  PVS1 5 得分
	//  PVS1 6 不得分
	if PVS1 == "5" {
		PVS1 = "1"
	}
	if PVS1 == "6" {
		PVS1 = "0"
	}

	// PS
	//  PS1 1,2 暂时不得分
	//  PS1 3 得分
	//  PS1 4 不得分
	if PS1 == "1" || PS1 == "2" || PS1 == "4" {
		PS1 = "0"
	}
	if PS1 == "3" {
		PS1 = "1"
	}
	//  PS4 5 得分
	if PS4 == "5" {
		PS4 = "1"
	}
	if PS1 == "1" {
		sumPS++
	}
	if PS2 == "1" {
		sumPS++
	}
	if PS3 == "1" {
		sumPS++
	}
	if PS4 == "1" {
		sumPS++
	}

	//  PM3 2 升级到PS得分
	if PM3 == "2" {
		sumPS++
		PM3 = "0"
	}
	//  PM5 1,2 不得分
	//  PM5 3 得分
	//  PM5 4 不得分
	//  PM5 5 得分
	if PM5 == "1" || PM5 == "2" || PM5 == "4" {
		PM5 = "0"
	}
	if PM5 == "3" || PM5 == "5" {
		PM5 = "1"
	}

	if PM1 == "1" {
		sumPM++
	}
	if PM1 == "Supporting" {
		sumPP++
	}
	if PM2 == "Supporting" {
		sumPP++
	}
	if PM4 == "Supporting" {
		sumPP++
	}
	if PM3 == "1" {
		sumPM++
	}
	if PM4 == "1" {
		sumPM++
	}
	if PM5 == "1" {
		sumPM++
	}
	if PM6 == "1" {
		sumPM++
	}

	// PP
	//  PP1 2 升级到PM
	if PP1 == "2" {
		sumPM++
		PP1 = "0"
	}
	// PP3 不与 PS3 共同得分
	if PS3 == "1" {
		PP3 = "0"
	}

	//  ACMG 已取消证据 PP5
	PP5 = "0"
	if PP1 == "1" {
		sumPP++
	}
	if PP2 == "1" {
		sumPP++
	}
	switch PP3 {
	case "1":
		sumPP++
	case "Moderate":
		sumPM++
	case "Strong":
		sumPS++
	}

	if PP4 == "1" {
		sumPP++
	}
	if PP5 == "1" {
		sumPP++
	}

	// BA
	if BA1 == "1" {
		sumBA++
	}
	// BS
	if BS1 == "1" {
		sumBS++
	}
	if BS2 == "1" {
		sumBS++
	}
	if BS3 == "1" {
		sumBS++
	}
	if BS4 == "1" {
		sumBS++
	}
	// BP
	// ACMG 已取消证据 BP6
	BP6 = "0"
	if BP1 == "1" {
		sumBP++
	}
	if BP2 == "1" {
		sumBP++
	}
	if BP3 == "1" {
		sumBP++
	}
	switch BP4 {
	case "1":
		sumBP++
	case "Moderate":
		sumBM++
	case "Strong":
		sumBS++
	}

	if BP5 == "1" {
		sumBP++
	}
	if BP6 == "1" {
		sumBP++
	}
	if BP7 == "1" {
		sumBP++
	}

	var ACMG = make(map[string]bool)
	if sumPVS > 0 {
		if sumPS == 1 || (sumPM+sumPP > 1) {
			ACMG["P"] = true
		}
		if sumPM == 1 {
			ACMG["LP"] = true
		}
		if PM2 == "Supporting" {
			ACMG["LP"] = true
		}
	}
	if sumBM >= 1 && sumPVS == 0 && sumPS == 0 && sumPM == 0 && sumPP == 0 && sumBA == 0 && sumBS == 0 {
		ACMG["LB"] = true
	}
	if sumPS > 1 {
		ACMG["P"] = true
	}
	if sumPS == 1 {
		if sumPM > 2 || (sumPM == 2 && sumPP >= 2) || (sumPM == 1 && sumPP >= 4) {
			ACMG["P"] = true
		}
		if sumPM == 1 || sumPM == 2 || sumPP > 1 {
			ACMG["LP"] = true
		}
	}
	if sumPM > 2 || (sumPM == 2 && sumPP > 1) || (sumPM == 1 && sumPP > 3) {
		ACMG["LP"] = true
	}
	if sumBA > 0 || sumBS > 1 {
		ACMG["B"] = true
	}
	if sumBP > 1 || (sumBP == 1 && sumBS == 1) {
		ACMG["LB"] = true
	}
	var PLP, BLB bool
	if ACMG["P"] || ACMG["LP"] {
		PLP = true
	}
	if ACMG["B"] || ACMG["LB"] {
		BLB = true
	}
	if PLP && BLB {
		if !ACMG["B"] && BP4 == "1" {
			sumBP--
			if sumBP > 1 || (sumBP == 1 && sumBS == 1) {
				return "VUS"
			} else if ACMG["P"] {
				return "P"
			} else if ACMG["LP"] {
				return "LP"
			}
		}
		return "VUS"
	} else if ACMG["P"] {
		return "P"
	} else if ACMG["LP"] {
		return "LP"
	} else if ACMG["B"] {
		return "B"
	} else if ACMG["LB"] {
		return "LB"
	} else {
		return "VUS"
	}
}
