package parking

type VehicleType string

const (
	Motorcycle VehicleType = "Motorcycle"
	Car        VehicleType = "Car"
	Bus        VehicleType = "Bus"
)

type Vehicle struct {
	LicenseNumber string
	Type          VehicleType
}
