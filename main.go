package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	calc "congestion-calculator/calculator"
	config "congestion-calculator/config"
	v "congestion-calculator/vehicle"
)

var configs = map[string]config.TaxConfig{}

func calculateGot(w http.ResponseWriter, req *http.Request) {
	// TODO: Read dates from POST or GET response
	tax := calc.GetTax(configs["GOT"], v.Car{}, []time.Time{
		time.Date(2013, 2, 7, 6, 23, 27, 0, time.Local),
		time.Date(2013, 2, 7, 15, 27, 0, 0, time.Local)})
	io.WriteString(w, fmt.Sprintf("Tax: %v", tax))
}

func main() {
	configs["GOT"] = config.GetGothConfig()
	//configs["STLM"] = config.GetStlmConfig() // TODO: Add configurations for other cities

	http.HandleFunc("/got", calculateGot)

	http.ListenAndServe(":8090", nil)
}
