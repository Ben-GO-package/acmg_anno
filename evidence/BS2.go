package evidence

import (
	"regexp"
	"strconv"
)

var BS2HomoList = []string{
	"ExAC HomoAlt Count",
	"GnomAD_nhomalt",
}

var BS2AF1List = []string{
	"GnomAD EAS AF",
	"GnomAD AF",
	// "1000G AF",
	// "ESP6500 AF",
	// "ExAC EAS AF",
	// "ExAC AF",
}

var (
	BS2LateOnsetHomoThreshold = 5
	BS2NoLateOnsetThreshold   = 0
	BS2HitCountThreshold      = 2
)

func CheckBS2(item map[string]string) string {
	if bs2GeneList[item["entrez_id"]] {
		return "0"
	}
	var inherit = item["ModeInheritance"] // AR;UNK;AD;AD,DD

	var (
		//Inheritance Model
		isADAR_AR_PR_DR = regexp.MustCompile(`AD,AR|AR|PR|DR`)
		isAD_DD         = regexp.MustCompile(`AD|DD`)
		isXL_YL         = regexp.MustCompile(`XL|YL`)
	)

	if isADAR_AR_PR_DR.MatchString(inherit) {
		if item["GnomAD_nhomalt"] != "" && item["GnomAD_nhomalt"] != "-" {
			value, err := strconv.Atoi(item["GnomAD_nhomalt"])
			if err == nil && value >= 5 {
				return "1"
			}
		}
	} else if isAD_DD.MatchString(inherit) {
		if item["GnomAD_AC"] != "" && item["GnomAD_AC"] != "-" {
			value, err := strconv.Atoi(item["GnomAD_AC"])
			if err == nil && value >= 5 {
				return "1"

			}
		}
	} else if isXL_YL.MatchString(inherit) {
		if item["GnomAD_AC_male"] != "" && item["GnomAD_AC_male"] != "-" {
			value, err := strconv.Atoi(item["GnomAD_AC_male"])
			if err == nil && value >= 5 {
				return "1"
			}
		}
	}
	return "0"
}

func CompareBS2(item map[string]string) {
	rule := "BS2"
	val := CheckBS2(item)
	if val != item[rule] {
		PrintConflict(item, rule, val, append([]string{"entrez_id", "ModeInheritance"}, append(BS2HomoList, BS2AF1List...)...)...)
	}
}
