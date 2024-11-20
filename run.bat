SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go env -w GOARCH=amd64
go env -w GOOS=windows
go build  -o acmg_anno.exe acmg.go   flag.go   global.go  io.go log.go  main.go 

SET CGO_ENABLED=0 
set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux
go build -o acmg_anno acmg.go   flag.go   global.go  io.go log.go  main.go 

.\acmg_anno.exe -snv test_data.tsv -acmg -autoPVS1 -tsv -json -outpred  -runPM1