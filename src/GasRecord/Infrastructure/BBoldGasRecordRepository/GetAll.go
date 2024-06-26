package gasrecord_infrastructure_bbold

import (
	"encoding/json"
	. "gasto-api/src/GasRecord"

	"go.etcd.io/bbolt"
)

func (repo *bboldGasRepository) GetAll() []GasRecord {
	var records []GasRecord
	done := make(chan struct{})

	go func() {
		defer close(done)

		transactionError := repo.db.View(func(transaction *bbolt.Tx) error {
			bucket := transaction.Bucket([]byte(bucket_name))
			if bucket == nil {
				return bucketNotFoundError()
			}

			return bucket.ForEach(func(_, row []byte) error {
				var rawRecord bboldGasRecord
				parsingError := json.Unmarshal(row, &rawRecord)
				if parsingError != nil {
					return newParsingRecordError(row)
				}
				record := fromBBoldGasRecord(rawRecord)
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
