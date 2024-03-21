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
	logFile             *os.File
	err                 error
	ts                  = []time.Time{time.Now()}
	step                = 0
	stats               = make(map[string]int)
	filterVariantsTitle []string
)

var transEN = map[string]string{
	"是":    "Yes",
	"否":    "No",
	"备注说明": "Note",
}

var InputTitle_check = []string{
	"AutoPVS1 Adjusted Strength",
	"cHGVS",
	"cHGVS_org",
	"#Chr",
	"dbscSNV_ADA_pred",
	"dbscSNV_RF_pred",
	"Ens Condel Pred",
	"Function",
	"Gene Symbol",
	"GERP++_RS_pred",
	"GWASdb_or",
	"Interpro_domain",
	"ModeInheritance",
	"MutationName",
	"MutationTaster Pred",
	"pfamId",
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

// log
var cycle1Count int

// var tier1Json *os.File
var tier1Data []map[string]string
