package gasrecord

import shared "gasto-api/src/Shared"

type GasRecordRepository interface {
	Save(gasRecord GasRecord) error
	GetAll() []GasRecord
	Close() error
	Search(shared.Criteria[GasRecord]) []GasRecord
	GetById(id GasRecordId) (GasRecord, error)
}
