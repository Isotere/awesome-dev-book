---
tags:
  - Architecture/Clean/Solid/ISP
---

## SOLID: Принцип разделения интерфейсов

```
ISP утверждает, что клиенты не должны зависеть от методов, которые они не используют.
```

_Предпочтение следует отдавать маленьким, связным интерфейсам._

Нарушения принципа ISP приводят к классам, зависящим от членов, в которых они не нуждаются: повышается связанность, 
уменьшается гибкость и ухудшается сопровождаемость.

