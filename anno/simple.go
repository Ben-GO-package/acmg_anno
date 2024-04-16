package anno

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/liserjrqlxue/goUtil/stringsUtil"
)

// regexp
var (
	inBrackets = regexp.MustCompile(`\(\S+\)`)

	rmChr = regexp.MustCompile(`^chr`)
)

// Score2Pred add _pred for scores
// Score2Pred 根据给定的项目条目中的分数信息，计算并更新项目的预测标签。
// 参数:
//
//	item - 一个map[string]string类型的变量，包含各种评分数据和对应的标签。
func Score2Pred(item map[string]string) {
	// 尝试将dbscSNV_ADA_SCORE转换为浮点数并根据分数设置预测标签
	score, e := strconv.ParseFloat(item["dbscSNV_ADA_SCORE"], 32)
	if e != nil {
		// 如果转换失败，则将原始分数用作标签
		item["dbscSNV_ADA_pred"] = item["dbscSNV_ADA_SCORE"]
	} else {
		// 根据分数值设置预测标签为"D"或"P"
		if score >= 0.6 {
			item["dbscSNV_ADA_pred"] = "D"
		} else {
			item["dbscSNV_ADA_pred"] = "P"
		}
	}

	// 同上，处理dbscSNV_RF_SCORE
	score, e = strconv.ParseFloat(item["dbscSNV_RF_SCORE"], 32)
	if e != nil {
		item["dbscSNV_RF_pred"] = item["dbscSNV_RF_SCORE"]
	} else {
		if score >= 0.6 {
			item["dbscSNV_RF_pred"] = "D"
		} else {
			item["dbscSNV_RF_pred"] = "P"
		}
	}

	// 处理GERP++_RS分数，判断是否为保守突变
	score, e = strconv.ParseFloat(item["GERP++_RS"], 32)
	if e != nil {
		item["GERP++_RS_pred"] = item["GERP++_RS"]
	} else {
		if score >= 2.0 {
			item["GERP++_RS_pred"] = "保守"
		} else {
			item["GERP++_RS_pred"] = "不保守"
		}
	}

	// 分别处理PhyloP Vertebrates和PhyloP Placental Mammals分数，判断是否为保守区域
	// 此部分代码与上述处理GERP++_RS分数的逻辑相似，故省略详细注释
	score, e = strconv.ParseFloat(item["PhyloP Vertebrates"], 32)
	if e != nil {
		item["PhyloP Vertebrates Pred"] = item["PhyloP Vertebrates"]
	} else {
		if score >= 2.0 {
			item["PhyloP Vertebrates Pred"] = "保守"
		} else {
			item["PhyloP Vertebrates Pred"] = "不保守"
		}
	}

	score, e = strconv.ParseFloat(item["PhyloP Placental Mammals"], 32)
	if e != nil {
		item["PhyloP Placental Mammals Pred"] = item["PhyloP Placental Mammals"]
	} else {
		if score >= 2.0 {
			item["PhyloP Placental Mammals Pred"] = "保守"
		} else {
			item["PhyloP Placental Mammals Pred"] = "不保守"
		}
	}
}

var (
	isDel = regexp.MustCompile(`del`)
)

func updatePos(item map[string]string) {
	item["chromosome"] = rmChr.ReplaceAllString(item["#Chr"], "")
	item["#Chr"] = "chr" + item["chromosome"]
	if item["VarType"] == "snv" || item["VarType"] == "ref" {
		item["#Chr+Stop"] = item["#Chr"] + ":" + item["Stop"]
		item["chr-show"] = item["#Chr"] + ":" + item["Stop"]
	} else {
		item["#Chr+Stop"] = item["#Chr"] + ":" + item["Start"] + "-" + item["Stop"]
		if isDel.MatchString(item["VarType"]) {
			item["chr-show"] = item["#Chr"] + ":" + stringsUtil.StringPlus(item["Start"], 1) + ".." + item["Stop"]
		} else {
			item["chr-show"] = item["#Chr"] + ":" + item["Start"] + ".." + stringsUtil.StringPlus(item["Stop"], 1)
		}
	}
}

// pHGVS= pHGVS1+"|"+pHGVS3
func getPhgvs(item map[string]string) string {
	if item["pHGVS1"] != "" && item["pHGVS3"] != "" && item["pHGVS1"] != "." && item["pHGVS3"] != "." {
		return item["pHGVS1"] + " | " + item["pHGVS3"]
	}
	return item["pHGVS"]
}

func getMNlite(item map[string]string) string {
	var MutationNameArray = strings.Split(item["MutationName"], ":")
	if len(MutationNameArray) > 1 {
		return inBrackets.ReplaceAllString(MutationNameArray[0], "") + ":" + MutationNameArray[1]
	}
	return item["MutationName"]
}

// UpdateSnv add info for all variant
func UpdateSnv(item map[string]string) {
	updatePos(item)
	item["pHGVS"] = getPhgvs(item)
	item["MutationNameLite"] = getMNlite(item)
}

// UpdateFunction convert intron to [splice+10,splice-10,splice+20,splice-20]

var chgvsReg = regexp.MustCompile(`c\.\d+([+-])(\d+)`)

// AFlist default AF list for check
var AFlist = []string{
	"GnomAD EAS AF",
	"GnomAD AF",
	"1000G AF",
	"ESP6500 AF",
	"ExAC EAS AF",
	"ExAC AF",
	"PVFD AF",
	"Panel AlleleFreq",
	"wgs_GnomAD_AF",
}
