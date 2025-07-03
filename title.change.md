v1.0.0.2 @ 2024.4.7
临床注释框架统一切换vep，会对输入文件title进行统一的命名更新，在这里记录以下title的更新规则，以及相关的acmg证据项
| 原title                    | 更新后title                           | acmg证据项目              |
| -------------------------- | ------------------------------------- | ------------------------- |
| ESP6500 AF                 | ESP6500_AF                            | BA1 ;BS1、BS2、 PM2、PVS1 |
| 1000G AF                   | AF                                    | BA1 ;BS1、BS2、 PM2、PVS1 |
| ExAC AF                    | ExAC_AF                               | BA1 ;BS1、BS2、 PM2、PVS1 |
| ExAC EAS AF                | ExAC_AF_EAS                           | BA1 ;BS1、BS2、 PM2、PVS1 |
| GnomAD AF                  | GnomAD_AF                             | BA1 ;BS1、BS2、 PM2、PVS1 |
| GnomAD EAS AF              | GnomAD_AF_eas                         | BA1 ;BS1、BS2、 PM2、PVS1 |
| Transcript                 | Feature                               |
| cHGVS                      | cHGVS                                 |
| Function                   | Consequence                           |
| Gene Symbol                | Symbol                                |
| RepeatTag                  | RepeatTag                             |
| VarType                    | VarType                               |
| GERP++_RS                  | GERP_RS                               |
| PhyloP Vertebrates         | PhyloP_Vertebrates                    |
| PhyloP Placental Mammals   | PhyloP_Placental_Mammals              |
| dbscSNV_RF_SCORE           | rf_score                              |
| dbscSNV_ADA_SCORE          | ada_score                             |
| SIFT Pred                  | SIFT_pred                             |
| Polyphen2_HVAR_Pred        | Polyphen2_HVAR_Pred                   | BP4 PP3                   |
| MutationTaster Pred        | MutationTaster_pred                   | BP4 PP3                   |
| Ens Condel Pred            | Condel                                | BP4  PP3                  |
| SpliceAI Pred              | SpliceAI_Pred                         | BP4 BP7 PP3               |
| ExAC HomoAlt Count         | ExAC_AC_Hom                           | BS2                       |
| GnomAD_nhomalt             | GnomAD_nhomalt                        | BS2                       |
| ModeInheritance            | Inheritance                           |
| Interpro_domain            | Interpro_domain                       |
| pfamId                     | 无法注释，PM1要弃用这个字段应该没用了 |
| #Chr                       | #Chr                                  |
| Start                      | Start                                 |
| Stop                       | Stop                                  |
| pHGVS                      | pHGVS                                 |
| pHGVS1                     | pHGVS1                                |
| AutoPVS1 Adjusted Strength | AutoPVS1_Adjusted_Strength            | PP3                       |
| MutationName               | MutationName                          |
| GWASdb_or                  | gwasCatalog_orOrBeta                  |
| Strand                     | STRAND                                |