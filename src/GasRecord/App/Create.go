package app

import (
	domain "gasto-api/src/GasRecord"
	shared "gasto-api/src/Shared"
)

type gasRecordCreator func(primitives domain.GasRecordPrimitives) error

func MakeCreateGasRecord(repository domain.GasRecordRepository, eventBus shared.EventBus) gasRecordCreator {
	return func(primitives domain.GasRecordPrimitives) error {
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
