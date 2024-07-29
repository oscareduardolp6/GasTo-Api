package share

import "fmt"

type PositiveNumber float32

func (num PositiveNumber) Value() float32 {
	return float32(num)
}

type InvalidPositiveNumber float32

func (val InvalidPositiveNumber) Error() string {
	return fmt.Sprintf("The value <%f> is not a positive number", val)
}

func CreatePositiveNumber(num float32) (PositiveNumber, error) {
	if num < 0 {
		return 0, InvalidPositiveNumber(num)
	}
	return PositiveNumber(num), nil
}
