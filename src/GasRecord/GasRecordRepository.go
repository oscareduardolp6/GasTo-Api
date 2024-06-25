package gasrecord

type GasRecordRepository interface {
	Save(gasRecord GasRecord, savedError chan<- error)
	GetAll() chan []GasRecord
	Close() error
}
