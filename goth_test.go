package main

import (
	calc "congestion-calculator/calculator"
	config "congestion-calculator/config"
	v "congestion-calculator/vehicle"
	"testing"
	"time"
)

type testScenario struct {
	vehicle v.Vehicle
	dates   []time.Time
	tax     int
}

var cfg config.TaxConfig = config.GetGothConfig()

var scenarios = []testScenario{
	{ // Zero dates
		vehicle: v.Car{},
		dates:   []time.Time{},
		tax:     0,
	},
	{ // 2013-01-14 05:00:00.
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 01, 14, 5, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // 2013-01-14 08:00:00.
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 01, 14, 8, 0, 0, 0, time.Local)},
		tax:     13,
	},
	{ // Saturday
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 01, 19, 8, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // Sunday
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 01, 20, 8, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // Thursday 2013-05-09	Kristi himmelfärdsdag
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 5, 9, 8, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // Wednesday 2013-05-08	Day before Kristi himmelfärdsdag
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 5, 8, 8, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // 2013-01-14 08:00:00.
		vehicle: v.Emergency{},
		dates:   []time.Time{time.Date(2013, 01, 14, 8, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // 2013-01-14 08:00:00.
		vehicle: v.Motorcycle{},
		dates:   []time.Time{time.Date(2013, 01, 14, 8, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // 2013-01-14 21:00:00
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 01, 14, 21, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // Friday June 28
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 07, 1, 8, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // Monday July 1
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 07, 1, 8, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // Sunday July 14
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 07, 14, 9, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // Monday July 15
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 07, 15, 9, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // Wednesday July 31
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 07, 31, 8, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // Thursday August 1
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2013, 8, 1, 8, 0, 0, 0, time.Local)},
		tax:     13,
	},
	{ // Friday
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 6, 27, 00, 0, time.Local)},
		tax: 8,
	},
	{ // Friday. Edge case 5:59:59.999
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 5, 59, 59, 999999, time.Local)},
		tax: 0,
	},
	{ // Friday. Edge case 6:00:00
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 6, 0, 0, 0, time.Local)},
		tax: 8,
	},
	{ // Friday. Edge case 6:29:59
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 6, 29, 59, 999999, time.Local)},
		tax: 8,
	},
	{ // Friday. Edge case 6:30:00
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 6, 30, 0, 0, time.Local)},
		tax: 13,
	},
	{ // Friday. Edge case 6:59:59
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 6, 59, 59, 99, time.Local)},
		tax: 13,
	},
	{ // Friday. Edge case 7:00:00
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 7, 0, 0, 0, time.Local)},
		tax: 18,
	},
	{ // Friday. Edge case 7:59:59
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 7, 59, 59, 99, time.Local)},
		tax: 18,
	},
	{ // Friday. Edge case 8:00:00
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 8, 0, 0, 0, time.Local)},
		tax: 13,
	},
	{ // Friday. Edge case 8:29:59
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 8, 29, 59, 99, time.Local)},
		tax: 13,
	},
	{ // Friday. Edge case 8:30:00
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 8, 30, 0, 0, time.Local)},
		tax: 8,
	},
	{ // Friday. Edge case 14:59:59
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 14, 59, 59, 99, time.Local)},
		tax: 8,
	},
	{ // Friday. Edge case 15:00:00
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 15, 0, 0, 0, time.Local)},
		tax: 13,
	},
	{ // Friday. Edge case 15:29:59
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 15, 29, 59, 99, time.Local)},
		tax: 13,
	},
	{ // Friday. Edge case 15:30:00
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 15, 30, 0, 0, time.Local)},
		tax: 18,
	},
	{ // Friday. Edge case 16:59:59
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 16, 59, 59, 99, time.Local)},
		tax: 18,
	},
	{ // Friday. Edge case 17:00:00
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 17, 0, 0, 0, time.Local)},
		tax: 13,
	},
	{ // Friday. Edge case 17:59:59
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 17, 59, 59, 99, time.Local)},
		tax: 13,
	},
	{ // Friday. Edge case 18:00:00
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 18, 0, 0, 0, time.Local)},
		tax: 8,
	},
	{ // Friday. Edge case 18:29:59
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 18, 29, 59, 99, time.Local)},
		tax: 8,
	},
	{ // Friday. Edge case 18:30:00
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 18, 30, 0, 0, time.Local)},
		tax: 0,
	},
	{ // Friday. Three passes. Interva1: [5:30, 6:00]. Interval2: [7:00]
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 5, 30, 0, 0, time.Local), // 0
			time.Date(2013, 2, 8, 6, 0, 0, 0, time.Local),  // 8
			time.Date(2013, 2, 8, 7, 0, 0, 0, time.Local)}, // 18
		tax: 26,
	},
	{ // Friday. Two passes within 1h at 6:00 and 7:00
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 6, 0, 0, 0, time.Local),  // 8
			time.Date(2013, 2, 8, 7, 0, 0, 0, time.Local)}, // 18
		tax: 18,
	},
	{ // Friday. Two passes not within 1h at 6:00:00 and 7:00:01
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 6, 0, 0, 0, time.Local),  // 8
			time.Date(2013, 2, 8, 7, 0, 1, 0, time.Local)}, // 18
		tax: 26,
	},
	{ // Friday. Multiple passes not exceeting max fee of 60/day
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 6, 0, 0, 0, time.Local),   // 8
			time.Date(2013, 2, 8, 7, 0, 1, 0, time.Local),   // 18
			time.Date(2013, 2, 8, 15, 30, 0, 0, time.Local), // 18
			time.Date(2013, 2, 8, 17, 0, 0, 0, time.Local),  // 13
		},
		tax: 57,
	},
	{ // Friday. Multiple passes exceeting max fee of 60/day
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2013, 2, 8, 6, 0, 0, 0, time.Local),   // 8
			time.Date(2013, 2, 8, 7, 0, 1, 0, time.Local),   // 18
			time.Date(2013, 2, 8, 15, 30, 0, 0, time.Local), // 18
			time.Date(2013, 2, 8, 17, 0, 0, 0, time.Local),  // 13
			time.Date(2013, 2, 8, 18, 29, 0, 0, time.Local), // 8
		},
		tax: 60,
	},
}

func TestScenarios(t *testing.T) {
	for _, s := range scenarios {
		tax := calc.GetTax(cfg, s.vehicle, s.dates)
		if tax != s.tax {
			t.Errorf("GetTax for vehicle %s and dates: %v was %v; want %v", s.vehicle.GetVehicleType(), s.dates, tax, s.tax)
		}
	}
}
