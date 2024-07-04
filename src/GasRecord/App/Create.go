package app

import (
	domain "gasto-api/src/GasRecord"
	share "gasto-api/src/Share"
)

type gasRecordCreator func(primitives domain.GasRecord) error

func MakeCreateGasRecord(repository domain.GasRecordRepository, eventBus share.EventBus) gasRecordCreator {
	return func(primitives domain.GasRecord) error {
		gasRecord, creationError := domain.CreateGasRecord(primitives)
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
