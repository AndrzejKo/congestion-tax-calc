package config

import (
	v "congestion-calculator/vehicle"
	"time"
)

var tollFreeVehicles = []v.VehicleType{v.MotorcycleType, v.TractorType, v.EmergencyType, v.DiplomatType, v.ForeignType, v.MilitaryType}

func GetGothConfig() TaxConfig {
	return TaxConfig{
		ConfigName:                 "GOT",
		SingleChargeIntevalSeconds: 60 * 60,
		MaxDaiyTax:                 60,
		FreeVehicleChecker:         freeVehicleChecker{},
		FreeDayChecker:             freeDayChecker{},
		TimeRateChecker:            timeRateChecker{},
	}
}

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
	year := date.Year()
	month := date.Month()
	day := date.Day()

	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return true
	}

	if month == 7 { // July
		return true
	}

	if year == 2013 {
		/*
			2013-01-01	Nyårsdagen	1	Tisdag	1
			2013-01-06	Trettondedag jul	1	Söndag	6
			2013-03-29	Långfredagen	13	Fredag	88
			2013-03-31	Påskdagen	13	Söndag	90
			2013-04-01	Annandag påsk	14	Måndag	91
			2013-05-01	Första maj	18	Onsdag	121
			2013-05-09	Kristi himmelfärdsdag	19	Torsdag	129
			2013-05-19	Pingstdagen	20	Söndag	139
			2013-06-06	Sveriges nationaldag	23	Torsdag	157
			2013-06-22	Midsommardagen	25	Lördag	173
			2013-11-02	Alla helgons dag	44	Lördag	306
			2013-12-25	Juldagen	52	Onsdag	359
			2013-12-26	Annandag jul	52	Torsdag	360
		*/

		switch {
		case month == 1 && (day == 1 || day == 5 || day == 6):
			return true
		case month == 3 && (day == 28 || day == 29 || day == 30 || day == 31):
			return true
		case month == 4 && (day == 1 || day == 30):
			return true
		case month == 5 && (day == 1 || day == 8 || day == 9 || day == 18 || day == 19):
			return true
		case month == 6 && (day == 5 || day == 6 || day == 21 || day == 22):
			return true
		case month == 11 && (day == 1 || day == 2):
			return true
		case month == 12 && (day == 24 || day == 25 || day == 26 || day == 31):
			return true
		}
	}
	return false
}

type timeRateChecker struct{}

func (trc timeRateChecker) GetTimeRate(t time.Time) int {
	hour, minute := t.Hour(), t.Minute()

	switch {
	case hour == 6 && minute <= 29:
		return 8
	case hour == 6 && minute >= 30:
		return 13
	case hour == 7:
		return 18
	case hour == 8 && minute <= 29:
		return 13
	case hour >= 8 && hour < 15:
		return 8
	case hour == 15 && minute <= 29:
		return 13
	case hour >= 15 && hour < 17:
		return 18
	case hour == 17:
		return 13
	case hour == 18 && minute <= 29:
		return 8
	}

	return 0
}
