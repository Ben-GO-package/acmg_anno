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

// log
var cycle1Count int

// var tier1Json *os.File
var tier1Data []map[string]string
