package app

import (
	"fmt"
	domain "gasto-api/src/GasRecord"
	infra "gasto-api/src/GasRecord/Infrastructure"
)

func MakeBatchCreationFromFile(batchCreation func(reader <-chan domain.GasRecord) <-chan Result) func(filepath string) error {
	return func(filepath string) error {
		reader, errorReadingFile := infra.ReadRecordsFromTextFile(filepath)
		if errorReadingFile != nil {
			return errorReadingFile
		}

		resultsChan := batchCreation(reader)

		insertedIds := make([]string, 0)
		creationErrors := make([]error, 0)

		for result := range resultsChan {
			if result.CreatedSuccessfull {
				insertedIds = append(insertedIds, result.RecordId)
				continue
			}
			creationErrors = append(creationErrors, result.CreationError)
		}

		if len(creationErrors) == 0 {
			return nil
		}

		return incompleteBatchCreation{
			insertedIds,
			creationErrors,
		}
	}
}

type incompleteBatchCreation struct {
	insertedIds    []string
	creationErrors []error
}

func (err incompleteBatchCreation) Error() string {
	if len(err.insertedIds) == 0 {
		return "Error in the batch creation"
	}
	return fmt.Sprintf("Parcial error in batch creation, some inserts failed. This are the correct inserts ids (%v): %v. And The non created records failed by this reasons: %v", len(err.insertedIds), err.insertedIds, err.creationErrors)
}
