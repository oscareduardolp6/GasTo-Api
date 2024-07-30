package gasrecord_test

import (
	app "gasto-api/src/GasRecord/App"
	gasrecord_infrastructure_inmemory "gasto-api/src/GasRecord/Infrastructure/InMemoryGasRecordRepository"
	share_infrastructure_inmemoryeventbus "gasto-api/src/Shared/Infrastructure/InMemoryEventBus"
	gasrecord_stub "gasto-api/test/GasRecord"
	"sync"
	"testing"
)

func TestMakeCreateGasRecord(t *testing.T) {
	var wg sync.WaitGroup
	inMemoryRepository := gasrecord_infrastructure_inmemory.NewGasRecordRepository()
	defer inMemoryRepository.Close()
	inMemoryEventBus := share_infrastructure_inmemoryeventbus.CreateInMemoryEventBus(&wg)
	create := app.MakeCreateGasRecord(inMemoryRepository, inMemoryEventBus)
	testPrimitives := gasrecord_stub.CreateBasicGasRecord()
	result := create(testPrimitives)
	if result != nil {
		t.Errorf("Gas Record Creation failed. Cause: %v", result)
	}

}
