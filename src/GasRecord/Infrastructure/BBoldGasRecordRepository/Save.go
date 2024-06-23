package gasrecord_infrastructure_bbold

import (
	"encoding/json"
	. "gasto-api/src/GasRecord"
	"log"

	"go.etcd.io/bbolt"
)

func (repo *bboldGasRepository) Save(gasRecord GasRecord) chan error {
	resultChan := make(chan error)

	go func() {
		updateError := repo.db.Update(func(transaction *bbolt.Tx) error {
			bucket := transaction.Bucket([]byte(bucket_name))
			bucketNotFound := bucket == nil
			if bucketNotFound {
				return bucketNotFoundError()
			}
			data := toBBoldGasRecord(gasRecord)
			encodedData, parsingError := json.Marshal(data)
			if parsingError != nil {
				return NewParsingRecordError([]byte{})
			}
			return bucket.Put([]byte(gasRecord.Id), encodedData)
		})

		var savingError error = nil

		if updateError != nil {
			log.Printf("Update error: %v", updateError.Error())
			savingError = NewRecordNotSaved(gasRecord)
		}

		resultChan <- savingError
		close(resultChan)
	}()

	return resultChan
}
