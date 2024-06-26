package gasrecord

import share "gasto-api/src/Share"

const GAS_RECORD_CREATED_NAME = "GasRecordCreated"

func CreateGasRecordCreatedEvent(gasrecord GasRecord) share.Event {
	return share.Event{
		Topic:   GAS_RECORD_CREATED_NAME,
		Payload: gasrecord,
	}
}
