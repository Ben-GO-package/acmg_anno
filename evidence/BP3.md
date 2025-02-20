# BP3 - 非框移的重复序列变异

判断条件:
1. 变异类型要求:
   - cds-del
   - cds-ins
   - cds-indel
   - inframe_deletion
   - inframe_insertion
   - protein_altering_variant

2. 重复序列要求:
   - 必须有RepeatTag标记
   - 如果是重复序列扩增:
     - 重复次数>=10判定为BP3
     - 重复次数<10不判定BP3

注意事项:
- 需要非SNV变异
- 需要有重复序列标记 