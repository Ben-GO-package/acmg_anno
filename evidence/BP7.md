# BP7 - 同义突变的剪切预测

判断条件:
1. 基本要求:
   - 必须是同义突变(coding-synon/synonymous)

2. 保守性预测全部为"不保守":
   - GERP++_RS_pred
   - PhyloP Vertebrates Pred
   - PhyloP Placental Mammals Pred

3. 剪切预测:
   - 如果预测致病则不判定BP7:
     - dbscSNV_RF_pred
     - dbscSNV_ADA_pred
     - SpliceAI Pred
   - 如果SpliceAI预测良性则判定BP7

注意事项:
- 仅适用于同义突变
- 需要满足保守性和剪切预测要求 