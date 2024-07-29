package gasrecord_infrastructure_bbold

import (
	"encoding/json"
	"fmt"
	domain "gasto-api/src/GasRecord"

	"go.etcd.io/bbolt"
)

func (repo *bboldGasRepository) GetById(id domain.GasRecordId) (domain.GasRecord, error) {
	var record domain.GasRecord

	transactionError := repo.db.View(func(transaction *bbolt.Tx) error {
		bucket := transaction.Bucket([]byte(bucket_name))
		bucketNotFound := bucket == nil
		if bucketNotFound {
			return bucketNotFoundError()
		}

		byteRecord := bucket.Get([]byte(id))
		if byteRecord == nil {
			return RecordNotFound{id}
		}

		parsingError := json.Unmarshal(byteRecord, &record)

		if parsingError != nil {
			return newParsingRecordError(byteRecord)
		}

		return nil
	})

	if transactionError != nil {
		return domain.GasRecord{}, transactionError
	}

	return record, nil

}

type RecordNotFound struct {
	Id domain.GasRecordId
}

func (err RecordNotFound) Error() string {
	return fmt.Sprintf("Record with ID: <%v> not found ", err.Id)
}
