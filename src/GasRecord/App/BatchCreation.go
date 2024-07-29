package app

import (
	domain "gasto-api/src/GasRecord"
	"sync"
)

type Result struct {
	RecordId           string
	CreatedSuccessfull bool
	CreationError      error
}

func MakeBatchCreation(creationUseCase func(domain.GasRecordPrimitives) error) func(reader <-chan domain.GasRecordPrimitives) <-chan Result {
	return func(records <-chan domain.GasRecordPrimitives) <-chan Result {
		resultsChan := make(chan Result)
		var waitGroup sync.WaitGroup

		go func() {
			for record := range records {
				waitGroup.Add(1)
				go func(record domain.GasRecordPrimitives) {
					defer waitGroup.Done()
					creationError := creationUseCase(record)
					success := creationError == nil
					resultsChan <- Result{
						RecordId:           record.Id,
						CreatedSuccessfull: success,
						CreationError:      creationError,
					}
				}(record)
			}

			// Wait for all goroutines to finish and close the results channel
			waitGroup.Wait()
			close(resultsChan)
		}()

		return resultsChan
	}
}
