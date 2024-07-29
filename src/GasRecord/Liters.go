package gasrecord

import (
	"fmt"
)

type Liters float32

type InvalidLitersQuantity struct {
	quantity float32
}

func (err InvalidLitersQuantity) Error() string {
	return fmt.Sprintf("The Liters Quantity cannot be negative. The value <%v> is an invalid value", err.quantity)
}

func (liters Liters) Value() float32 {
	return float32(liters)
}

func CreateLiters(liters float32) (Liters, error) {
	if liters < 0 {
		return 0, InvalidLitersQuantity{liters}
	}
	return Liters(liters), nil
}

func (liters1 Liters) Substract(liters2 Liters) (Liters, error) {
	result := liters1.Value() - liters2.Value()
	return CreateLiters(result)
}
