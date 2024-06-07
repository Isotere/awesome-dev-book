---
tags:
  - SystemDesign/General
aliases:
  - Общие положения
---
# Общие положения

Шаги, которые предпринимаются в ходе масштабирования системы для поддержки миллионов пользователей:

- веб-уровень не должен хранить состояния;
- резервирование должно быть предусмотрено на каждом уровне;
- кэширование данных следует проводить как можно более активно;
- система должна поддерживать больше одного центра обработки данных;
- статические ресурсы нужно хранить в CDN;
- для масштабирования данных следует применять шардинг;
- уровни должны быть разделены на отдельные сервисы;
- необходимо выполнять мониторинг системы и использовать средства автоматизации.

## Дополнительные материалы

1. Протокол передачи гипертекста: https://ru.wikipedia.org/wiki/HTTP
2. Should you go Beyond Relational Databases?: https://blog.teamtreehouse.com/should-you-go-beyond-relational-databases
3. Репликация: https://ru.wikipedia.org/wiki/Репликация_(вычислительная_техника)
4. Репликация с несколькими ведущими серверами: https://en.wikipedia.org/wiki/Multi-master_replication
5. NDB Cluster Replication: Multi-Master and Circular Replication: https://dev.mysql.com/doc/refman/5.7/en/mysql-cluster-replicationmulti-master.html
6. Caching Strategies and How to Choose the Right One: https://codeahoy.com/2017/08/11/caching-strategies-and-how-tochoose-the-right-one/
7. R. Nishtala, «Facebook, Scaling Memcache at», 10th USENIX Symposiumon Networked Systems Design and Implementation (NSDI ’13).
8. Единая точка отказа (англ.): https://en.wikipedia.org/wiki/Single_point_of_failure
9. Доставка динамического контента с Amazon CloudFront: https://aws.amazon.com/ru/cloudfront/dynamic-content/
10. Configure Sticky Sessions for Your Classic Load Balancer: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elbsticky-sessions.html
11. Active-Active for Multi-Regional Resiliency: https://netflixtechblog.com/active-active-for-multi-regional-resiliencyc47719f6685b
12. Инстансы Amazon EC2 High Memory: https://aws.amazon.com/ru/ec2/instance-types/high-memory/
13. What it takes to run Stack Overflow: http://nickcraver.com/blog/2013/11/22/what-it-takes-to-runstack-overflow
14. What The Heck Are You Actually Using NoSQL For: http://highscalability.com/blog/2010/12/6/what-the-heck-are-youactually-using-nosql-for.html