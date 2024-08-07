package main

import (
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// os
var (
	ex, _   = os.Executable()
	exPath  = filepath.Dir(ex)
	cfgPath = filepath.Join(exPath, "cfg")
)

// regexp
var (
	isGz      = regexp.MustCompile(`\.gz$`)
	isComment = regexp.MustCompile(`^##`)
)

var snvs []string

var (
	logFile          *os.File
	err              error
	ts               = []time.Time{time.Now()}
	step             = 0
	stats            = make(map[string]int)
	finalOutputTitle []string
)

var transEN = map[string]string{
	"是":    "Yes",
	"否":    "No",
	"备注说明": "Note",
}

var InputTitle_check = []string{
	"Uploaded_variation",
	"AutoPVS1 Adjusted Strength",
	"cHGVS",
	"cHGVS_org",
	"#Chr",
	"dbscSNV_ADA_pred",
	"dbscSNV_RF_pred",
	"Ens Condel Pred",
	"Function",
	"Gene Symbol",
	"entrez_id",
	"GERP++_RS_pred",
	"GWASdb_or",
	"Interpro_domain",
	"ModeInheritance",
	"MutationName",
	"MutationTaster Pred",
	"pHGVS",
	"pHGVS1",
	"PhyloP Placental Mammals Pred",
	"PhyloP Vertebrates Pred",
	"PM5",
	"Polyphen2 HVAR Pred",
	"PS1",
	"PVS1",
	"RepeatTag",
	"SIFT Pred",
	"SpliceAI Pred",
	"Start",
	"Stop",
	"Transcript",
	"VarType",
}

var transverTitle_relation_map = map[string]string{
	"Uploaded_variation":         "Uploaded_variation",
	"ESP6500 AF":                 "ESP6500_AF",       // :"BA1 ;BS1、BS2、 PM2、PVS1"
	"1000G AF":                   "AF",               //:"BA1 ;BS1、BS2、 PM2、PVS1",
	"ExAC AF":                    "ExAC_AF",          //:"BA1 ;BS1、BS2、 PM2、PVS1",
	"ExAC EAS AF":                "ExAC_AF_EAS",      //:"BA1 ;BS1、BS2、 PM2、PVS1",
	"GnomAD AF":                  "GnomAD_AF",        //:"BA1 ;BS1、BS2、 PM2、PVS1",
	"GnomAD EAS AF":              "GnomAD_AF_eas",    //:"BA1 ;BS1、BS2、 PM2、PVS1",
	"Transcript":                 "Feature",          //PS1 ; PM5
	"Protein_position":           "Protein_position", //PS1 ; PM5
	"cHGVS":                      "cHGVS",            //PS1 ; PM5
	"Function":                   "Consequence",
	"Gene Symbol":                "Symbol",
	"RepeatTag":                  "RepeatTag",
	"VarType":                    "VarType",
	"GERP++_RS":                  "GERP_RS",
	"PhyloP Vertebrates":         "PhyloP_Vertebrates",
	"PhyloP Placental Mammals":   "PhyloP_Placental_Mammals",
	"dbscSNV_RF_SCORE":           "rf_score",
	"dbscSNV_ADA_SCORE":          "ada_score",
	"SIFT Pred":                  "SIFT_pred",
	"Polyphen2 HVAR Pred":        "Polyphen2_HVAR_Pred", //:"BP4 PP3",
	"MutationTaster Pred":        "MutationTaster_pred", //:"BP4 PP3",
	"Ens Condel Pred":            "Condel",              //:"BP4  PP3",
	"SpliceAI Pred":              "SpliceAI_Pred",       //:"BP4 BP7 PP3",
	"ExAC HomoAlt Count":         "ExAC_AC_Hom",         //:"BS2",
	"GnomAD HomoAlt Count":       "GnomAD_nhomalt",      //:"BS2",
	"ModeInheritance":            "Inheritance",
	"Interpro_domain":            "Interpro_domain",
	"#Chr":                       "#Chr",
	"Start":                      "Start",
	"Stop":                       "Stop",
	"pHGVS":                      "pHGVS",
	"pHGVS1":                     "pHGVS1",                     //PS1 ; PM5
	"AutoPVS1 Adjusted Strength": "AutoPVS1_Adjusted_Strength", //:"PP3",
	"MutationName":               "MutationName",
	"GWASdb_or":                  "gwasCatalog_orOrBeta",
	"Strand":                     "STRAND",
	"entrez_id":                  "Gene",
}

var TempOutputTitle = []string{
	"#Chr", "Start", "Stop", "Ref", "Call",
	"Feature",
	"cHGVS",
	"pHGVS",
	"autoRuleName", "自动化判断",
	"pHGVS1",
	"Symbol",
	"entrez_id",
	"Consequence",
	"RepeatTag",
	"VarType",
	"GERP_RS",
	"PhyloP_Vertebrates",
	"PhyloP_Placental_Mammals",
	"rf_score",
	"ada_score",
	"SIFT_pred",
	"Polyphen2_HVAR_Pred",
	"MutationTaster_pred",
	"Condel",
	"SpliceAI_Pred",
	"ExAC_AC_Hom",
	"GnomAD_nhomalt",
	"Inheritance",
	"Interpro_domain",
	"AutoPVS1_Adjusted_Strength",
	"MutationName",
	"gwasCatalog_orOrBeta",
	"STRAND",
	"ESP6500_AF",
	"AF",
	"ExAC_AF",
	"ExAC_AF_EAS",
	"GnomAD_AF",
	"GnomAD_AF_eas",
}

// log
var cycle1Count int

// var tier1Json *os.File
var WholeResultData []map[string]string
var ImporttempData []map[string]string
