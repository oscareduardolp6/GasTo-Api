package main

import (
	domain "gasto-api/src/GasRecord"
	app "gasto-api/src/GasRecord/App"
	infra "gasto-api/src/GasRecord/Infrastructure/BBoldGasRecordRepository"
	share "gasto-api/src/Share"
	share_infrastructure_inmemoryeventbus "gasto-api/src/Share/Infrastructure/InMemoryEventBus"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

func main() {
	var waitGroup sync.WaitGroup
	eventBus := share_infrastructure_inmemoryeventbus.CreateInMemoryEventBus(&waitGroup)
	repo, createRepositoryError := infra.NewGasRecordRepository()
	if createRepositoryError != nil {
		log.Fatal(createRepositoryError)
	}
	defer repo.Close()
	configureSuscriptions(eventBus, repo)

	createGasRecord := app.MakeCreateGasRecord(repo, eventBus)

	err := createGasRecord(domain.GasRecord{
		Id:                 uuid.NewString(),
		Place:              "Test",
		Liters:             30,
		TotalPrice:         726,
		TraveledKilometers: 400,
		PriceByLiter:       24.2,
		Date:               time.Now(),
		RoadTrip:           false,
	})

	if err != nil {
		log.Fatal(err)
	}

	waitGroup.Wait()

}

func configureSuscriptions(eventBus share.EventBus, repo domain.GasRecordRepository) {
	updatePerformanceWhenGasRecordCreated := app.MakeUpdatePerformanceWhenGasRecordCreated(repo)
	eventBus.Suscribe(domain.GAS_RECORD_CREATED_NAME, updatePerformanceWhenGasRecordCreated)
}
