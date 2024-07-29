package app

import (
	domain "gasto-api/src/GasRecord"
	shared "gasto-api/src/Shared"
	"log"
)

func MakeUpdatePerformanceWhenGasRecordCreated(repo domain.GasRecordRepository) func(shared.Event) {
	return func(event shared.Event) {
		if !canManageEvent(event) {
			return
		}
		updatePerformance := MakeUpdatePerformanceGasRecord(repo)
		gasRecord, _ := event.Payload.(domain.GasRecord)
		err := updatePerformance(gasRecord.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func canManageEvent(event shared.Event) bool {
	_, ok := event.Payload.(domain.GasRecord)
	return ok && event.Topic == domain.GAS_RECORD_CREATED_NAME
}
