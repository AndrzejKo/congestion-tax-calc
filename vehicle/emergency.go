package vehicle

type Emergency struct {
}

func (e Emergency) GetVehicleType() string {
	return "Emergency"
}
