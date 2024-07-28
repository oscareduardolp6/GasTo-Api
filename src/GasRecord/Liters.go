package gasrecord

import (
	"fmt"
)

type Liters struct {
	Value float32 `json:"value" validate:"required,gte=0"`
}

type GasTank struct {
	Initial Liters `json:"initial_liters"`
	Final   Liters `json:"final_liters"`
}

type InvalidLitersQuantity struct {
	quantity float32
}

func (err InvalidLitersQuantity) Error() string {
	return fmt.Sprintf("The Liters Quantity cannot be negative. The value <%v> is an invalid value", err.quantity)
}

func CreateLiters(liters float32) (Liters, error) {
	if liters < 0 {
		return Liters{}, InvalidLitersQuantity{liters}
	}
	return Liters{liters}, nil
}

func (liters1 Liters) Substract(liters2 Liters) (Liters, error) {
	result := liters1.Value - liters2.Value
	return CreateLiters(result)
}
