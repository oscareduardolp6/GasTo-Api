package gasrecord

import share "gasto-api/src/Share"

type GasRecordPrice float32

func CreateGasRecordPrice(price float32) (GasRecordPrice, error) {
	value, positiveNumberError := share.CreatePositiveNumber(price)
	if positiveNumberError != nil {
		return 0, positiveNumberError
	}
	return GasRecordPrice(value.Value()), nil
}
