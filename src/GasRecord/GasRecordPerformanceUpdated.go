package gasrecord

import share "gasto-api/src/Share"

const GAS_RECORD_PERFORMANCE_UPDATED = "GasRecordPerformanceUpdated"

type GasRecordPerformanceUpdatedPayload struct {
	RecordId    string
	Performance float32
}

func CreateGasRecordPerformanceUpdated(gasRecord GasRecord) share.Event {
	return share.Event{
		Topic: GAS_RECORD_CREATED_NAME,
		Payload: GasRecordPerformanceUpdatedPayload{
			RecordId:    gasRecord.Id.Value(),
			Performance: gasRecord.Performance.Value(),
		},
	}
}
