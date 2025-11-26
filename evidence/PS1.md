# PS1 - 已知致病的氨基酸改变

判断条件:
1. 必须是错义突变|起始密码子变异|非编码RNA基因变异(missense|start_retained_variant|ncRNA)

2. 相同氨基酸位点的其他核苷酸改变已被确认为致病:
   - 检查 trans_phgvs (转录本:蛋白质改变)的致病记录数
   - 检查 trans_chgvs (转录本:核苷酸改变)的致病记录数
   - 如果 phgvs记录数 > chgvs记录数,则判定为PS1

注意事项:
- 数据来源于PS1PM5.database
- 仅考虑相同氨基酸位点的不同核苷酸改变
- 该证据项会输出对应证据
