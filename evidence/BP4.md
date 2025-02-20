# BP4 - 多个计算预测支持良性

判断条件:
1. 保守性预测全部为"不保守":
   - GERP++_RS_pred
   - PhyloP Vertebrates Pred
   - PhyloP Placental Mammals Pred

2. 具体预测规则:
   - 剪切位点变异:
     - 如果预测致病则不判定BP4:
       - dbscSNV_RF_pred
       - dbscSNV_ADA_pred
       - SpliceAI Pred
     - 如果SpliceAI预测良性则判定BP4
   - 内含子变异:
     - 距离剪切位点1-10bp:
       - 同剪切位点变异规则
   - 其他变异:
     - 需同时满足:
       - SIFT预测良性
       - Polyphen2预测良性
       - MutationTaster预测良性
       - Ens Condel预测neutral

3. 升级条件:
   - BayesDel_noAF_score:
     - <= -0.36: 升级为Moderate
     - -0.36 到 -0.18: 保持Supporting
   - REVEL:
     - <= 0.016: 升级为Strong
     - 0.016-0.183: 升级为Moderate
     - 0.183-0.290: 保持Supporting

注意事项:
- 不能与BP3共同使用
- 不能与PVS1共同使用 