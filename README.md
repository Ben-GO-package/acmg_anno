# anno2xlsx <!-- omit in toc -->

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/liserjrqlxue/anno2xlsx)


## PARAM

| arg       | type    | example           | note                                                                                        |
| --------- | ------- | ----------------- | ------------------------------------------------------------------------------------------- |
| -json     | boolean |                   | 输出json格式结果                                                                            |
| -acmg     | boolean |                   | 使用ACMG2015计算证据项PVS1, PS1,PS4, PM1,PM2,PM4,PM5 PP2,PP3, BA1, BS1,BS2, BP1,BP3,BP4,BP7 |
| -autoPVS1 | boolean |                   | 使用autoPVS1结果处理证据项PVS1                                                              |
| -cfg      | string  | etc/config.toml   | toml配置文件                                                                                |
| -snv      | string  | snv1.txt,snv2.txt | snv注释结果，逗号分割                                                                       |
| -prefix   | string  | outputPrefix      | 输出前缀，默认 -snv 第一个输入                                                              |
| -log      | string  | prefix.log        | log输出文件                                                                                 |
| -tag      | string  | .tag              | tier1结果文件名加入额外标签，[prefix].Tier1[tag].xlsx                                       |


## demo
acmg_anno -snv vep_autopvs1.demo.tsv -acmg -autoPVS1 -tsv -json