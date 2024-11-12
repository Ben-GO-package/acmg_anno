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

var transverTitle_relation_map = map[string]string{
	"Uploaded_variation":         "Uploaded_variation",
	"entrez_id":                  "Gene",
	"Gene Symbol":                "Symbol",
	"ESP6500 AF":                 "ESP6500_AF",                 // :"BA1 ;BS1、BS2、 PM2、PVS1"
	"1000G AF":                   "AF",                         //:"BA1 ;BS1、BS2、 PM2、PVS1",
	"ExAC AF":                    "ExAC_AF",                    //:"BA1 ;BS1、BS2、 PM2、PVS1",
	"ExAC EAS AF":                "ExAC_AF_EAS",                //:"BA1 ;BS1、BS2、 PM2、PVS1",
	"GnomAD AF":                  "GnomAD_AF",                  //:"BA1 ;BS1、BS2、 PM2、PVS1",
	"GnomAD EAS AF":              "GnomAD_AF_eas",              //:"BA1 ;BS1、BS2、 PM2、PVS1",
	"Transcript":                 "Feature",                    //PS1 ; PM5
	"Protein_position":           "Protein_position",           //PS1 ; PM5
	"cHGVS":                      "cHGVS",                      //PS1 ; PM5
	"Function":                   "Consequence",                //BP1 ; BP3 ; BP4 ; BP7 ; PM1 ; PM4 ; PM5 ;PP2 ; PP3 ; BP4 ; PS1
	"RepeatTag":                  "RepeatTag",                  // BP3 ; PM4
	"VarType":                    "VarType",                    // BP3
	"GERP++_RS":                  "GERP_RS",                    // BP4 ;BP7 ; PP3
	"PhyloP Vertebrates":         "PhyloP_Vertebrates",         // BP4 ;BP7 ; PP3
	"PhyloP Placental Mammals":   "PhyloP_Placental_Mammals",   // BP4 ;BP7 ; PP3
	"SIFT Pred":                  "SIFT_pred",                  //BP4 PP3
	"SpliceAI Pred":              "SpliceAI_Pred",              //:"BP4 BP7 PP3",
	"ExAC HomoAlt Count":         "ExAC_AC_Hom",                //:"BS2",
	"GnomAD HomoAlt Count":       "GnomAD_nhomalt",             //:"BS2",
	"ModeInheritance":            "Inheritance",                // PM2 ; BS2
	"Interpro_domain":            "Interpro_domain",            // PM1
	"#Chr":                       "#Chr",                       // PVS1
	"Start":                      "Start",                      // PVS1
	"Stop":                       "Stop",                       // PVS1
	"pHGVS":                      "pHGVS",                      // PM5 ; PS1 ;
	"pHGVS1":                     "pHGVS1",                     //PS1 ; PM5
	"AutoPVS1 Adjusted Strength": "AutoPVS1_Adjusted_Strength", //:"PP3", PM4
	"MutationName":               "MutationName",               // PS1 ; PVS1
	"GWASdb_or":                  "gwasCatalog_orOrBeta",       // PS4
	"Strand":                     "STRAND",                     // PVS1
	"entrezID_oe":                "entrezID_oe",                //PM1
	"SpliceAI_Max_Score":         "SpliceAI_Max_Score",         //BP4 PP3
	"BayesDel_noAF_score":        "BayesDel_noAF_score",        //BP4 PP3
	"REVEL":                      "REVEL",                      //BP4 PP3
}

var TempOutputTitle = []string{
	"#Chr", "Start", "Stop", "Ref", "Call",
	"Feature",
	"cHGVS",
	"pHGVS",
	"autoRuleName", "automated_judgment",
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
	"entrezID_oe",
	"SpliceAI_Max_Score",
	"BayesDel_noAF_score",
	"REVEL",
}

// log
var cycle1Count int

// var tier1Json *os.File
var WholeResultData []map[string]string
var ImporttempData []map[string]string
