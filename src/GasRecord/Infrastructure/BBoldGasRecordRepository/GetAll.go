package gasrecord_infrastructure_bbold

import (
	"encoding/json"
	. "gasto-api/src/GasRecord"

	"go.etcd.io/bbolt"
)

func (repo *bboldGasRepository) GetAll() chan []GasRecord {
	resultChan := make(chan []GasRecord)

	go func() {
		var records []GasRecord

		transactionError := repo.db.View(func(transaction *bbolt.Tx) error {
			bucket := transaction.Bucket([]byte(bucket_name))
			bucketNotFound := bucket == nil
			if bucketNotFound {
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
				var noError error = nil
				return noError
			})
		})

		handleError(transactionError)

		resultChan <- records
		close(resultChan)
	}()

	return resultChan
}
