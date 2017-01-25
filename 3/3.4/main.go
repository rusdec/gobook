/**
Задание 3.4

Следуя подходу, использованному в примере с фигурами Лиссажу из раздела 1.7,
создайте веб-сервер, который вычисляетповерхности и возвращает клиенту
SVG-файл. Сервер должениспользовать в ответе заголовок ContentType наподобие
следующего:

w.Header().Set("ContentType", "image/svg+xml")

*/

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Polygon(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:2121", nil))
}
