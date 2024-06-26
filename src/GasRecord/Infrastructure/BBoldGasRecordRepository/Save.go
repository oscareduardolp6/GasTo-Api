package gasrecord_infrastructure_bbold

import (
	"encoding/json"
	. "gasto-api/src/GasRecord"
	"log"

	"go.etcd.io/bbolt"
)

func (repo *bboldGasRepository) Save(gasRecord GasRecord) error {
	updateErrorChan := make(chan error, 1) // Buffer de 1 para evitar bloqueo

	go func() {
		updateError := repo.db.Update(func(transaction *bbolt.Tx) error {
			bucket := transaction.Bucket([]byte(bucket_name))
			if bucket == nil {
				return bucketNotFoundError()
			}
			data := toBBoldGasRecord(gasRecord)
			encodedData, parsingError := json.Marshal(data)
			if parsingError != nil {
				return newParsingRecordError([]byte{})
			}
			return bucket.Put([]byte(gasRecord.Id), encodedData)
		})
		updateErrorChan <- updateError
		close(updateErrorChan)
	}()

	var savingError error = nil
	if updateError := <-updateErrorChan; updateError != nil {
		log.Printf("Update error: %v", updateError.Error())
		savingError = NewRecordNotSaved(gasRecord)
	}

	return savingError
}
