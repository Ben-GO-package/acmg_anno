# PVS1 - 导致蛋白质功能缺失的变异

判断条件：
1. 功能影响检查:
   - 必须是高度影响功能的变异类型:
     - splice-3
     - splice-5 
     - init-loss
     - alt-start
     - frameshift
     - nonsense
     - stop-gain
     - span

2. 基因敏感性检查:
   - 基因必须在LOF不耐受基因列表中

3. 额外条件(任一满足):
   - 突变位点后有重要的蛋白结构功能区域
   - 突变位点后有其他已知致病突变位点

4. 频率限制:
   - 所有人群频率数据库中频率均需 < 5%:
     - GnomAD EAS AF
     - GnomAD AF 
     - 1000G AF
     - ESP6500 AF
     - ExAC EAS AF
     - ExAC AF

注意事项:
- PVS1与PP3、PM4不能共同使用
- 如果使用AutoPVS1,其结果可能被降级为Strong/Moderate/Supporting 