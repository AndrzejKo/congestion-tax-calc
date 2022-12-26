package calculator

import (
	config "congestion-calculator/config"
	v "congestion-calculator/vehicle"
	"testing"
	"time"
)

func getTaxConfig() config.TaxConfig {
	return config.TaxConfig{
		ConfigName:                 "STHLM",
		SingleChargeIntevalSeconds: 120 * 60,
		MaxDaiyTax:                 30,
		FreeVehicleChecker:         freeVehicleChecker{},
		FreeDayChecker:             freeDayChecker{},
		TimeRateChecker:            timeRateChecker{},
	}
}

var tollFreeVehicles = []v.VehicleType{v.MotorcycleType, v.TractorType, v.EmergencyType, v.DiplomatType, v.ForeignType, v.MilitaryType}

type freeVehicleChecker struct{}

func (fvc freeVehicleChecker) IsFreeVehicle(v v.Vehicle) bool {
	if v == nil {
		return false
	}
	vehicleType := v.GetVehicleType()

	for i := 0; i < len(tollFreeVehicles); i++ {
		tfv := tollFreeVehicles[i]
		if tfv.String() == vehicleType {
			return true
		}
	}

	return false
}

type freeDayChecker struct{}

func (fdc freeDayChecker) IsFreeDay(date time.Time) bool {
	month := date.Month()
	day := date.Day()

	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return true
	}

	if month == 8 && day == 19 {
		return true
	}

	return false
}

type timeRateChecker struct{}

func (trc timeRateChecker) GetTimeRate(t time.Time) int {
	hour, minute := t.Hour(), t.Minute()

	if hour == 6 && minute >= 0 && minute <= 29 {
		return 8
	}
	if hour == 6 && minute >= 30 && minute <= 59 {
		return 13
	}
	if hour == 7 && minute >= 0 && minute <= 59 {
		return 18
	}
	if hour == 8 && minute >= 0 && minute <= 29 {
		return 13
	}
	if (hour == 8 && minute >= 30) || (hour >= 9 && hour < 18) {
		return 8
	}

	return 0
}

type testScenario struct {
	vehicle v.Vehicle
	dates   []time.Time
	tax     int
}

var cfg config.TaxConfig = getTaxConfig()

var scenarios = []testScenario{
	{ // Zero dates
		vehicle: v.Car{},
		dates:   []time.Time{},
		tax:     0,
	},
	{ // Outside paid time
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2022, 12, 27, 5, 59, 59, 0, time.Local)},
		tax:     0,
	},
	{ // Car, during paid time, on a paid day
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2022, 12, 27, 6, 0, 0, 0, time.Local)},
		tax:     8,
	},
	{ // Car, during paid time, on a toll free day (****-08-19)
		vehicle: v.Car{},
		dates:   []time.Time{time.Date(2022, 8, 19, 7, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // Emergency, during paid time, on a paid day
		vehicle: v.Emergency{},
		dates:   []time.Time{time.Date(2022, 12, 27, 6, 0, 0, 0, time.Local)},
		tax:     0,
	},
	{ // Car, during paid time, on a paid day, two passes within SingleChargeIntevalSeconds=120*60
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2022, 12, 27, 6, 0, 0, 0, time.Local),      //8
			time.Date(2022, 12, 27, 7, 59, 59, 999, time.Local)}, //18
		tax: 18,
	},
	{ // Car, during paid time, on a paid day, two passes within SingleChargeIntevalSeconds=120*60. Passes on the edges of interval.
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2022, 12, 27, 6, 0, 0, 0, time.Local),  //8
			time.Date(2022, 12, 27, 8, 0, 0, 0, time.Local)}, //13
		tax: 13,
	},
	{ // First pass free (1st interval start), second pass paid (within 1st interval), third paid (second interval).
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2022, 12, 27, 5, 30, 0, 0, time.Local), //0
			time.Date(2022, 12, 27, 6, 0, 0, 0, time.Local),  //8
			time.Date(2022, 12, 27, 8, 0, 0, 0, time.Local)}, //13
		tax: 21,
	},
	{ // Car, during paid time, on a paid day, two passes spanning two intervals SingleChargeIntevalSeconds=120*60
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2022, 12, 27, 6, 0, 0, 0, time.Local),  //8
			time.Date(2022, 12, 27, 8, 0, 1, 0, time.Local)}, //13
		tax: 21,
	},
	{ // Car, during paid time, on a paid day, four passes spanning four intervals. Exceeding max daily fee.
		vehicle: v.Car{},
		dates: []time.Time{
			time.Date(2022, 12, 27, 6, 0, 0, 0, time.Local),   //8
			time.Date(2022, 12, 27, 8, 0, 1, 0, time.Local),   //13
			time.Date(2022, 12, 27, 13, 0, 0, 0, time.Local),  //8
			time.Date(2022, 12, 27, 17, 0, 0, 0, time.Local)}, //8
		tax: 30,
	},
}

func TestScenarios(t *testing.T) {
	for _, s := range scenarios {
		tax := GetTax(cfg, s.vehicle, s.dates)
		if tax != s.tax {
			t.Errorf("GetTax for vehicle %s and dates: %v was %v; want %v", s.vehicle.GetVehicleType(), s.dates, tax, s.tax)
		}
	}
}
