package app

import (
	. "gasto-api/src/GasRecord"
)

func MakeGetAllGasRecords(repository GasRecordRepository) func() []GasRecord {
	return repository.GetAll
}
