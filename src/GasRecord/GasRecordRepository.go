package gasrecord

type GasRecordRepository interface {
	Save(gasRecord GasRecord) chan error
	GetAll() chan []GasRecord
	Close() error
}
