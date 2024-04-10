# acmg_anno <!-- omit in toc -->

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/Ben-GO-package/acmg_anno)


## PARAM

| arg         | type    | example                    | note                                                                                         |
| ----------- | ------- | -------------------------- | -------------------------------------------------------------------------------------------- |
| -acmg       | boolean |                            | 使用ACMG2015计算证据项PVS1, PS1,PS4, PM1,PM2,PM4,PM5 PP2,PP3, BA1, BS1,BS2, BP1,BP3,BP4,BP7  |
| -autoPVS1   | boolean |                            | 使用autoPVS1结果处理证据项PVS1                                                               |
| -acmgDb     | string  | cfg/acmg.db.cfg            | acmg 各证据项分析时，需要提供的相关数据库文件                                                |
| -json       | boolean | prefix.acmg.json           | 输出json格式结果                                                                             |
| -tsv        | boolean | prefix.acmg.tsv            | 输出tsv格式结果                                                                              |
| -snv        | string  | snv1.txt,snv2.txt          | snv注释结果，如果多个文件，文件间使用逗号分割                                                |
| -prefix     | string  | outputPrefix               | 输出前缀，默认 -snv 第一个输入                                                               |
| -log        | string  | prefix.log                 | log输出文件                                                                                  |
| -temp_title | string  | cfg/final_result_title.cfg | 用于指定 prefix.acmg.temp.tsv 文件需要输出保存的字段名称（可以指定证据项处理过程中的中间字段 |


## demo
acmg_anno -snv vep_autopvs1.demo.tsv -acmg -autoPVS1 -tsv -json