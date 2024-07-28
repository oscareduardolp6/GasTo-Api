package gasrecord_infrastructure_bbold

import (
	"encoding/json"
	domain "gasto-api/src/GasRecord"

	"go.etcd.io/bbolt"
)

func (repo *bboldGasRepository) GetAll() []domain.GasRecord {
	records := make([]domain.GasRecord, 0)
	done := make(chan struct{})

	go func() {
		defer close(done)

		transactionError := repo.db.View(func(transaction *bbolt.Tx) error {
			bucket := transaction.Bucket([]byte(bucket_name))
			if bucket == nil {
				return bucketNotFoundError()
			}

			return bucket.ForEach(func(_, row []byte) error {
				var record domain.GasRecord
				parsingError := json.Unmarshal(row, &record)
				if parsingError != nil {
					return newParsingRecordError(row)
				}
				records = append(records, record)
				return nil
			})
		})

		handleError(transactionError)
	}()

	// Esperar a que la goroutine termine
	<-done

	return records
}
