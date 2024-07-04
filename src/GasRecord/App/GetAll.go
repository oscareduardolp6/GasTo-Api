package app

import (
	domain "gasto-api/src/GasRecord"
)

func MakeGetAllGasRecords(repository domain.GasRecordRepository) func() []domain.GasRecord {
	return repository.GetAll
}
