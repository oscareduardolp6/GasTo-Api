package gasrecord_infrastructure_bbold

import (
	"encoding/json"
	"fmt"
	. "gasto-api/src/GasRecord"

	"go.etcd.io/bbolt"
)

func (repo *bboldGasRepository) GetById(id string) (GasRecord, error) {
	var record GasRecord

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

		var rawRecord bboldGasRecord

		parsingError := json.Unmarshal(byteRecord, &rawRecord)

		if parsingError != nil {
			return newParsingRecordError(byteRecord)
		}

		record = fromBBoldGasRecord(rawRecord)

		return nil
	})

	if transactionError != nil {
		return GasRecord{}, transactionError
	}

	return record, nil

}

type RecordNotFound struct {
	Id string
}

func (err RecordNotFound) Error() string {
	return fmt.Sprintf("Record with ID: <%v> not found ", err.Id)
}
