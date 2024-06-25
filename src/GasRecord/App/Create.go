package app

import (
	. "gasto-api/src/GasRecord"
)

type gasRecordCreator func(primitives GasRecord) error

func MakeCreateGasRecord(repository GasRecordRepository) gasRecordCreator {
	return func(primitives GasRecord) error {
		gasRecord, creationError := CreateGasRecord(primitives)
		if creationError != nil {
			return creationError
		}
		return <-repository.Save(gasRecord)
	}
}
