package parking

type SpotSize string

const (
	MediumSpot SpotSize = "medium"
	SmallSpot  SpotSize = "small"
	LargeSpot  SpotSize = "large"
)

type ParkingSpot struct {
	id          int
	isAvailable bool
	size        SpotSize
	vehicle     *Vehicle
}

func (ps *ParkingSpot) Park(vehicle *Vehicle) {
	ps.isAvailable = false
	ps.vehicle = vehicle
}

func (ps *ParkingSpot) Leave() {
	ps.isAvailable = true
	ps.vehicle = nil
}
