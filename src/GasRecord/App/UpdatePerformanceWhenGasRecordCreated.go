package app

import (
	domain "gasto-api/src/GasRecord"
	share "gasto-api/src/Share"
	"log"
)

func MakeUpdatePerformanceWhenGasRecordCreated(repo domain.GasRecordRepository) func(share.Event) {
	return func(event share.Event) {
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

func canManageEvent(event share.Event) bool {
	_, ok := event.Payload.(domain.GasRecord)
	return ok && event.Topic == domain.GAS_RECORD_CREATED_NAME
}
