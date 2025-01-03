package vehicles

import "fmt"

type Vehicle interface { //definimos una interfaz. y definimos un contrato que deben cumplir los tipos que la implementan.
	Distance() float64
}

const ( //Definimos constantes para los tipos de vehiculos
	CarVehicle        = "CAR"
	MotorCycleVehicle = "MOTORCYCLE"
	TruckVehicle      = "TRUCK"
)

// DEfinimos funcion que recibe un tipo de vehiculo y retorna un tipo de vehiculo. Como un tipo de Factory
func NewVehicle(vehicleType string, time int) (Vehicle, error) { //Retorna una interfaz y un error
	switch vehicleType { //Switch para determinar el tipo de vehiculo
	case CarVehicle:
		return &Car{Time: time}, nil //Retornamos un puntero a la estructura Car para que pueda tomarlo la interfaz Vehicle
	case MotorCycleVehicle:
		return &MotorCycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	default:
		return nil, fmt.Errorf("Vehicle of type %s not found", vehicleType)
	}
}

type Car struct {
	Time int
}

func (c Car) Distance() float64 {
	return (float64(c.Time) / 60) * 100
}

type MotorCycle struct {
	Time int
}

func (m MotorCycle) Distance() float64 {
	return (float64(m.Time) / 60) * 120
}

type Truck struct {
	Time int
}

func (t Truck) Distance() float64 {
	return (float64(t.Time) / 60) * 70
}
