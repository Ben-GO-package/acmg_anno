# PM2 - 在健康人群中罕见的变异

判断条件:
1. 常染色体隐性/X连锁隐性疾病:
   - 所有人群频率数据库中频率均需 < 0.005
   
2. 常染色体显性疾病:
   - 对于晚发病基因:频率需 < 0.00001
   - 对于其他基因:频率需 = 0

检查的频率数据库:
- ESP6500 AF
- 1000G AF
- GnomAD AF
- GnomAD EAS AF
~~- ExAC AF~~  #20250627 -央彩六期需求剔除该数据库
~~- ExAC EAS AF~~ # 20250627 -央彩六期需求剔除该数据库

注意事项:
- 满足条件时判定为Supporting级别
- 遗传方式为空时按隐性处理 