package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Программа привет + name
type Calculate struct {
	A         int    `json:"a"`
	B         int    `json:"b"`
	Operation string `json:"operation"`
	Result    int    `json:"result"`
}

func resultCalculate(w http.ResponseWriter, r *http.Request) {
	var calculate Calculate
	err := json.NewDecoder(r.Body).Decode(&calculate)
	if err != nil {
		fmt.Fprintln(w, "ошибка")
		return
	}
	switch calculate.Operation {
	case "plus":
		calculate.Result = calculate.A + calculate.B
	case "minys":
		calculate.Result = calculate.A - calculate.B
	case "umn":
		calculate.Result = calculate.A * calculate.B
	case "del":

		if calculate.B == 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "на нуль делить нельзя")
			return

		}
		calculate.Result = calculate.A / calculate.B

	}
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(calculate)
	w.Write(data)

}

func main() {
	log.Println("сервер запущен на порту 8080")
	http.HandleFunc("/calculate", resultCalculate)
	http.ListenAndServe(":8080", nil)

}
