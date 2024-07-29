package gasrecord

import shared "gasto-api/src/Shared"

type GasRecordPrice float32

func CreateGasRecordPrice(price float32) (GasRecordPrice, error) {
	value, positiveNumberError := shared.CreatePositiveNumber(price)
	if positiveNumberError != nil {
		return 0, positiveNumberError
	}
	return GasRecordPrice(value.Value()), nil
}
