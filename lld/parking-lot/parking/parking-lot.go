package parking

type ParkingLot struct {
	spots []*ParkingSpot
}

// add spot
func (pl *ParkingLot) AddSpot(spot *ParkingSpot) {
	pl.spots = append(pl.spots, spot)
}

// park
func (pl *ParkingLot) Park(vehicle *Vehicle, priority bool) bool {
	switch vehicle.Type {
	case Motorcycle:
		for _, spot := range pl.spots {
			if spot.size == SmallSpot && spot.isAvailable {
				spot.Park(vehicle)
				return true
			} else if priority && spot.size == MediumSpot && spot.isAvailable {
				spot.Park(vehicle)
				return true
			} else if priority && spot.size == LargeSpot && spot.isAvailable {
				spot.Park(vehicle)
				return true
			}
			return false
		}
	case Car:
		for _, spot := range pl.spots {
			if spot.size == MediumSpot && spot.isAvailable {
				spot.Park(vehicle)
				return true
			} else if priority && spot.size == LargeSpot && spot.isAvailable {
				spot.Park(vehicle)
				return true
			}
			return false
		}
	case Bus:
		for _, spot := range pl.spots {
			if spot.size == LargeSpot && spot.isAvailable {
				spot.Park(vehicle)
				return true
			}
			return false
		}
	default:
		return false
	}
	return false
}

// leave
func (pl *ParkingLot) Leave(spotId int) {
	for _, spot := range pl.spots {
		if spot.id == spotId {
			spot.Leave()
		}
	}
}
