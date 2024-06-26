package gasrecord_infrastructure_bbold

import (
	"encoding/json"
	"fmt"
	. "gasto-api/src/GasRecord"

	"go.etcd.io/bbolt"
)

func (repo *bboldGasRepository) GetById(id string) (GasRecord, error) {
	foundChan := make(chan GasRecord)
	errorChan := make(chan error)

	go func() {
		defer close(errorChan)
		transactionError := repo.db.View(func(transaction *bbolt.Tx) error {
			defer close(foundChan)
			bucket := transaction.Bucket([]byte(bucket_name))
			bucketNotFound := bucket == nil
			if bucketNotFound {
				return bucketNotFoundError()
			}
			byteRecord := bucket.Get([]byte(id))

			if byteRecord == nil {
				return RecordNotFound{id}
			}

			var rawRecord bboldGasRecord

			parsingError := json.Unmarshal(byteRecord, &rawRecord)

			if parsingError != nil {
				return newParsingRecordError(byteRecord)
			}

			record := fromBBoldGasRecord(rawRecord)

			foundChan <- record
			return nil
		})

		if transactionError != nil {
			errorChan <- transactionError
		}

	}()

	for {
		select {
		case foundRecord := <-foundChan:
			return foundRecord, nil
		case gettingError := <-errorChan:
			return GasRecord{}, gettingError
		}

	}
}

type RecordNotFound struct {
	Id string
}

func (err RecordNotFound) Error() string {
	return fmt.Sprintf("Record with ID: <%v> not found ", err.Id)
}
