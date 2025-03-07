# acmg_anno <!-- omit in toc -->

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/Ben-GO-package/acmg_anno)


## PARAM

| arg         | type    | example                    | note                                                                                                                                                   |
| ----------- | ------- | -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| -acmg       | boolean |                            | 使用ACMG2015计算证据项PVS1, PS1,PS4, PM1,PM2,PM4,PM5 PP2,PP3, BA1, BS1,BS2, BP1,BP3,BP4,BP7                                                            |
| -autoPVS1   | boolean |                            | 使用autoPVS1结果处理证据项PVS1                                                                                                                         |
| -runPM1     | boolean |                            | ACMG判定过程中，是否启用证据项PM1 （默认不适用）                                                                                                       |
| -outpred    | boolean | false                      | 是否输出中间预测结果，"dbscSNV_ADA_pred", "dbscSNV_RF_pred", "GERP_RS_pred", "PhyloP_Vertebrates_Pred", "PhyloP_Placental_Mammals_Pred")（默认不输出） |
| -acmgDb     | string  | cfg/acmg.db.cfg            | acmg 各证据项分析时，需要提供的相关数据库文件                                                                                                          |
| -json       | boolean | prefix.acmg.json           | 输出json格式结果                                                                                                                                       |
| -tsv        | boolean | prefix.acmg.tsv            | 输出tsv格式结果                                                                                                                                        |
| -snv        | string  | snv1.txt,snv2.txt          | snv注释结果，如果多个文件，文件间使用逗号分割                                                                                                          |
| -prefix     | string  | outputPrefix               | 输出前缀，默认 -snv 第一个输入                                                                                                                         |
| -log        | string  | prefix.log                 | log输出文件                                                                                                                                            |
| -temp_title | string  | cfg/final_result_title.cfg | 用于指定 prefix.acmg.temp.tsv 文件需要输出保存的字段名称（可以指定证据项处理过程中的中间字段                                                           |


## demo
acmg_anno -snv vep_autopvs1.demo.tsv -acmg -autoPVS1 -tsv -json -runPM1  -outpred

## acmgDb数据库
cfg/acmg.db.cfg 调用的各个数据库构建说明参考底核心依赖库 [acmg2015说明文档](https://pkg.go.dev/github.com/Ben-GO-package/acmg2015#readme-heading)

## temp 文件
可以通过设置 final_result_title.cfg 文件，控制 acmg.temp.tsv 的输出字段，进而详细的查看每个证据项的评估结果。证据项值的含义说明

| 值         | 含义                                    |
| ---------- | --------------------------------------- |
| 0          | 证据项不支持                            |
| 1          | 证据项支持                              |
| Supporting | 证据项发生升降级，升降级后为 Supporting |
| Moderate   | 证据项发生升降级，升降级后为 Moderate   |
| Strong     | 证据项发生升降级，升降级后为 Strong     |
| -1         | 证据项和其他高优先级证据项不共用        |

- 不共用证据项清单
若强证据项支持，若证据项直接忽略不纳入考虑

| 强证据项 | 弱若证据项 | 更新时间 |
| -------- | ---------- | -------- |
| PVS1     | PM4        | 2023年前 |
| PVS1     | PP3        | 2023年前 |
| PM4      | PP3        | 2023年前 |
| PVS1     | BP4        | 2024.11  |
| PM1      | PP2        | 2023年前 |
| BP3      | BP4        | 2024.11  |
