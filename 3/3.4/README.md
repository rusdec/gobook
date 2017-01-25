## Задание 3.4

Следуя подходу, использованному в примере с фигурами Лиссажу из раздела 1.7,
создайте веб-сервер, который вычисляетповерхности и возвращает клиенту
SVG-файл. Сервер должениспользовать в ответе заголовок ContentType наподобие
следующего:

```
w.Header().Set("ContentType", "image/svg+xml")
```
Позвольте клиенту указывать разные параметры, такие как

* высота
* ширина
* цвет
* (частота)

в запросе HTTP

### Использование:

localhost:2121/?xyrange=N&colorMin=K&colorMax=L&size=S

где N	- частота колебаний					тип:ЦЕЛОЕ
где K,L	- цвет в шестнадцатиричном формате	тип:ЦЕЛОЕ
где S	- высота и ширина картинки			тип:ЦЕЛОЕ(диапазон:100<=S<=800)

### Примечанние

Ошибка в задании: заголовок нужен "Content-Type"