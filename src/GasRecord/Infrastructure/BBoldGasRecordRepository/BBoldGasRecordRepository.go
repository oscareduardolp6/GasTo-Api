package gasrecord_infrastructure_bbold

import (
	"fmt"
	. "gasto-api/src/GasRecord"
	"log"

	"go.etcd.io/bbolt"
)

const (
	read_and_write_permission = 0600
	bucket_name               = "GasRecords"
	db_path                   = "database.db"
)

type bboldGasRepository struct {
	db *bbolt.DB
}

func NewGasRecordRepository() (GasRecordRepository, error) {
	db, dbError := bbolt.Open(db_path, read_and_write_permission, nil)
	if dbError != nil {
		return nil, dbError
	}
	repo := &bboldGasRepository{db}
	initializationError := repo.initialize()
	if initializationError != nil {
		return nil, initializationError
	}
	return repo, nil
}

func bucketNotFoundError() error {
	return newBucketNotFound(bucket_name)
}

func handleError(err error) {
	if err == nil {
		return
	}

	switch e := err.(type) {
	case bucketNotFound, parsingError:
		message := fmt.Sprintf("Catched: %v", e)
		log.Fatal(message)
	default:
		message := fmt.Sprintf("Uncached: %v", err)
		log.Fatal(message)
	}
}

func (repo *bboldGasRepository) Close() error {
	return repo.db.Close()
}

func (repo *bboldGasRepository) initialize() error {
	err := repo.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket_name))
		if err != nil {
			log.Printf("Error creating bucket: %v", err)
		}
		return err
	})
	if err != nil {
		log.Printf("Error during initializacion %v", err)
	}
	log.Printf("inicialize correct")
	return err
}
