SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build  -o acmg_anno.exe acmg.go   flag.go   global.go  io.go log.go  main.go 

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o acmg_anno acmg.go   flag.go   global.go  io.go log.go  main.go 

.\acmg_anno.exe -snv vep_autopvs1.demo.tsv -acmg -autoPVS1 -tsv -json