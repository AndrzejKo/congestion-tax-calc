package vehicle

import (
	"fmt"
)

type VehicleType int

const (
	CarType        VehicleType = iota
	MotorcycleType VehicleType = iota
	TractorType    VehicleType = iota
	EmergencyType  VehicleType = iota
	DiplomatType   VehicleType = iota
	ForeignType    VehicleType = iota
	MilitaryType   VehicleType = iota
)

func (vt VehicleType) String() string {
	switch int(vt) {
	case int(CarType):
		return "Car"
	case int(MotorcycleType):
		return "Motorcycle"
	case int(TractorType):
		return "Tractor"
	case int(EmergencyType):
		return "Emergency"
	case int(DiplomatType):
		return "Diplomat"
	case int(ForeignType):
		return "Foreign"
	case int(MilitaryType):
		return "Military"
	default:
		return fmt.Sprintf("%d", int(vt))
	}
}

type Vehicle interface {
	GetVehicleType() string
}
