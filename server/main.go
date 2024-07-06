package main

import (
	"fmt"
	routes "gasto-api/server/routes/gasRecord"
	domain "gasto-api/src/GasRecord"
	app "gasto-api/src/GasRecord/App"
	infra "gasto-api/src/GasRecord/Infrastructure/BBoldGasRecordRepository"
	share "gasto-api/src/Share"
	share_infra "gasto-api/src/Share/Infrastructure/InMemoryEventBus"
	"net/http"
	"sync"
)

type methodRelation = map[string]http.HandlerFunc

func main() {
	var waitGroup sync.WaitGroup
	gasRecordsHandlers := make(methodRelation)
	repo, creationRepoError := infra.NewGasRecordRepository()
	if creationRepoError != nil {
		panic("Error inicializando el event bus o el repositorio")
	}
	defer repo.Close()
	eventBus := share_infra.CreateInMemoryEventBus(&waitGroup)
	configureSuscriptions(eventBus, repo)
	getAllGasRecords := app.MakeGetAllGasRecords(repo)
	createGasRecords := app.MakeCreateGasRecord(repo, eventBus)
	gasRecordsHandlers[http.MethodGet] = routes.MakeGetAllGetHandler(getAllGasRecords)
	gasRecordsHandlers[http.MethodPut] = routes.MakeGasRecordPutHandler(createGasRecords)
	gasRecordResource := dependingOnMethod(gasRecordsHandlers)
	http.HandleFunc("/gasrecord", gasRecordResource)
	port := ":8080"
	fmt.Printf("Server running in http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Print("Error running the server")
	}
}

func configureSuscriptions(eventBus share.EventBus, repo domain.GasRecordRepository) {
	updatePerformanceWhenGasRecordCreated := app.MakeUpdatePerformanceWhenGasRecordCreated(repo)
	eventBus.Suscribe(domain.GAS_RECORD_CREATED_NAME, updatePerformanceWhenGasRecordCreated)
}

func dependingOnMethod(handlers methodRelation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler, found := handlers[r.Method]
		if !found {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}
