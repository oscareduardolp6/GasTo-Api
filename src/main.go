package main

import (
	"fmt"
	gasrecord_routes "gasto-api/server/routes/gasRecord"
	infra "gasto-api/src/GasRecord/Infrastructure/BBoldGasRecordRepository"
	shared_infra "gasto-api/src/Shared/Infrastructure/InMemoryEventBus"
	"net/http"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	repo, creationRepoError := infra.NewGasRecordRepository()
	if creationRepoError != nil {
		panic("Error inicializando el repositorio")
	}
	defer repo.Close()
	eventBus := shared_infra.CreateInMemoryEventBus(&waitGroup)
	gasrecord_routes.ConfigureResource(repo, eventBus)
	port := ":8080"
	fmt.Printf("Server running in http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Print("Error running the server")
	}
}
