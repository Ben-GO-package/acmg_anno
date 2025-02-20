# PP3 - 多个计算预测支持致病

判断条件:
1. 保守性预测(至少2个支持):
   - GERP++_RS_pred
   - PhyloP Vertebrates Pred
   - PhyloP Placental Mammals Pred

2. 致病性预测(满足以下任一):
   - 功能预测:
     - Ens Condel = deleterious
     - 至少2个预测致病:
       - SIFT
       - MutationTaster
       - Polyphen2 HVAR
   - 剪切位点预测:
     - 至少2个预测致病或SpliceAI预测致病:
       - dbscSNV_RF
       - dbscSNV_ADA
       - SpliceAI

3. 升级条件:
   - SpliceAI_Max_Score >= 0.2:
     - REVEL >= 0.932: 升级为Strong
     - 0.773 <= REVEL < 0.932: 升级为Moderate
     - REVEL < 0.773: 保持Supporting
     - 无REVEL时使用BayesDel:
       - >= 0.5: 升级为Strong
       - 0.27-0.5: 升级为Moderate
       - < 0.27: 保持Supporting

注意事项:
- 不能与PVS1共同使用
- 不能与PS3共同使用 