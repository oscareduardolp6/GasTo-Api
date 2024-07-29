package gasrecord

import share "gasto-api/src/Share"

type Kilometers float32

func (val Kilometers) Value() float32 {
	return float32(val)
}

func CreateKilometers(val float32) (Kilometers, error) {
	value, positiveNumError := share.CreatePositiveNumber(val)
	if positiveNumError != nil {
		return 0, positiveNumError
	}
	return Kilometers(value.Value()), nil
}
