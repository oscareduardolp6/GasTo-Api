package app

import (
	. "gasto-api/src/GasRecord"
	share "gasto-api/src/Share"
)

func MakeUpdatePerformanceWhenGasRecordCreated(repo GasRecordRepository) func(share.Event) {
	return func(event share.Event) {
		if !canManageEvent(event) {
			return
		}
		updatePerformance := MakeUpdatePerformanceGasRecord(repo)
		gasRecord, _ := event.Payload.(GasRecord)
		updatePerformance(gasRecord.Id)
	}
}

func canManageEvent(event share.Event) bool {
	_, ok := event.Payload.(GasRecord)
	return ok && event.Topic == GAS_RECORD_CREATED_NAME
}
