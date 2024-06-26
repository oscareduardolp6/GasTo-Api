package gasrecord

import share "gasto-api/src/Share"

type GasRecordRepository interface {
	Save(gasRecord GasRecord) error
	GetAll() []GasRecord
	Close() error
	Search(share.Criteria[GasRecord]) []GasRecord
	GetById(id string) (GasRecord, error)
}
