package app

import (
	. "gasto-api/src/GasRecord"
	share "gasto-api/src/Share"
)

type gasRecordCreator func(primitives GasRecord) error

func MakeCreateGasRecord(repository GasRecordRepository, eventBus share.EventBus) gasRecordCreator {
	return func(primitives GasRecord) error {
		gasRecord, creationError := CreateGasRecord(primitives)
		if creationError != nil {
			return creationError
		}
		savingError := repository.Save(*gasRecord)

		if savingError != nil {
			return savingError
		}

		domainEvents := gasRecord.PullAllDomainEvents()

		for _, event := range domainEvents {
			eventBus.Publish(event)
		}
		return nil
	}
}
