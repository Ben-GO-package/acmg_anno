# PM4 - 蛋白长度改变的非移码变异

判断条件:
1. 变异类型必须是以下之一:
   - cds-del
   - cds-ins
   - cds-indel
   - stop-loss
   - inframe_deletion
   - inframe_insertion
   - stop_lost
   - protein_altering_variant

2. 不能有重复序列标记(RepeatTag为空或.或-)

3. 排除条件:
   - 不能与PVS1共同使用
   - 不能与PP3共同使用

注意事项:
- 如果使用AutoPVS1且有结果,则不判定PM4 