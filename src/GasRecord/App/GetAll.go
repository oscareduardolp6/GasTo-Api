package app

import (
	. "gasto-api/src/GasRecord"
)

func MakeGetAllGasRecords(repository GasRecordRepository) func() chan []GasRecord {
	return repository.GetAll
}
