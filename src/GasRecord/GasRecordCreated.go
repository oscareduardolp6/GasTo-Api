package gasrecord

import shared "gasto-api/src/Shared"

const GAS_RECORD_CREATED_NAME = "GasRecordCreated"

func CreateGasRecordCreatedEvent(gasrecord GasRecordPrimitives) shared.Event {
	return shared.Event{
		Topic:   GAS_RECORD_CREATED_NAME,
		Payload: gasrecord,
	}
}
