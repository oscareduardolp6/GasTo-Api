package gasrecord_infrastructure_bbold

import (
	"encoding/json"
	. "gasto-api/src/GasRecord"

	"go.etcd.io/bbolt"
)

func (repo *bboldGasRepository) GetAll() []GasRecord {
	var records []GasRecord

	go func() {

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

	}()

	return records
}
