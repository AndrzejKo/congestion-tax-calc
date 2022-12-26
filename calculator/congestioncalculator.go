package calculator

import (
	config "congestion-calculator/config"
	v "congestion-calculator/vehicle"
	"sort"
	"time"
)

func GetTax(cfg config.TaxConfig, vehicle v.Vehicle, dates []time.Time) int {
	if len(dates) == 0 {
		return 0
	}

	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})

	intervalStart := dates[0]
	intervalFee := 0
	totalFee := 0
	for _, date := range dates {
		nextFee := getTollFee(cfg, date, vehicle)
		diffInNanos := date.UnixNano() - intervalStart.UnixNano()
		seconds := diffInNanos / 1000000 / 1000

		if seconds <= int64(cfg.SingleChargeIntevalSeconds) {
			if nextFee > intervalFee {
				totalFee = totalFee - intervalFee + nextFee
				intervalFee = nextFee
			}
		} else {
			totalFee = totalFee + nextFee
			intervalFee = nextFee
			intervalStart = date
		}
	}

	if totalFee > cfg.MaxDaiyTax {
		totalFee = cfg.MaxDaiyTax
	}
	return totalFee
}

func getTollFee(config config.TaxConfig, t time.Time, v v.Vehicle) int {
	if config.FreeDayChecker.IsFreeDay(t) || config.FreeVehicleChecker.IsFreeVehicle(v) {
		return 0
	}

	return config.TimeRateChecker.GetTimeRate(t)
}
