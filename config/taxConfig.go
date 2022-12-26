package config

import (
	v "congestion-calculator/vehicle"
	"time"
)

type FreeDayChecker interface {
	IsFreeDay(day time.Time) bool
}

type FreeVehicleChecker interface {
	IsFreeVehicle(vehicle v.Vehicle) bool
}

type TimeRateChecker interface {
	GetTimeRate(t time.Time) int
}

type TaxConfig struct {
	ConfigName                 string
	SingleChargeIntevalSeconds int
	MaxDaiyTax                 int
	FreeDayChecker             FreeDayChecker
	FreeVehicleChecker         FreeVehicleChecker
	TimeRateChecker            TimeRateChecker
}
