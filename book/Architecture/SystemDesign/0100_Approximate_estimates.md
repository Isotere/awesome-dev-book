---
tags:
  - SystemDesign/Estimates
aliases:
  - Системный дизайн - Приблизительные оценки
---
# Приблизительные оценки

> Вычисления на коленке — это оценки, основанные на мысленных экспериментах и типичных показателях производительности, которые дают хорошее представление о том, какие архитектуры соответствуют вашим требованиям

## Степень двойки

В распределенных системах данные могут достигать огромных размеров, но все вычисления сводятся к элементарным свойствам. Чтобы получить правильный результат, нужно обязательно знать объем данных, используя вторую степень.

## ПОКАЗАТЕЛИ ЛАТЕНТНОСТИ

Некоторые из приведенных ниже чисел устарели в связи с повышением производительности компьютеров. Но порядок +- остался тем же. 


| Название операции                                    | Время                   |
| ---------------------------------------------------- | ----------------------- |
| Обращение к кэшу L1                                  | 0,5 нс                  |
| Ошибочное предсказание перехода                      | 5 нс                    |
| Обращение к кэшу L2                                  | 7 нс                    |
| Блокирование/разблокирование мьютекса                | 100 нс                  |
| Обращение к основной памяти                          | 100 нс                  |
| Сжатие 1 Кб с помощью Zippy                          | 10 000 нс = 10 мкс      |
| Отправка 2 Кб по сети 1 Гбит/с                       | 20 000 нс = 20 мкс      |
| Последовательное чтение из памяти 1 Мб               | 250 000 нс = 250 мкс    |
| Перемещение пакета туда и обратно внутри одного ЦОД  | 500 000 нс = 500 мкс    |
| Время поиска по диску                                | 10 000 000 нс = 10 мс   |
| Последовательное чтение 1 Мб из сети                 | 10 000 000 нс = 10 мс   |
| Последовательное чтение 1 Мб с диска                 | 30 000 000 нс = 30 мс   |
| Передача пакета из Калифорнии в Нидерланды и обратно | 150 000 000 нс = 150 мс |
|                                                      |                         |

**Вывод:**
- память быстрая, а диск медленный;
- по возможности следует избегать поиска по диску;
- простые алгоритмы сжатия отличаются высокой скоростью;
- прежде чем отправлять данные по интернету, их по возможности нужно сжимать;
- центры обработки данных обычно находятся в разных регионах, и передача информации между ними занимает время.

## ПОКАЗАТЕЛИ ДОСТУПНОСТИ

Высокая доступность — это способность системы долго и непрерывно находиться в рабочем состоянии.

Поставщики сервисов часто используют такой термин, как соглашение об уровне услуг (service level agreement, SLA). Это соглашение между вами (поставщиком) и вашим клиентом, которое официально определяет уровень беспрерывной работы вашего сервиса.
