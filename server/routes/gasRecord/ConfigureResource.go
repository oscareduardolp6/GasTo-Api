package routes

import (
	shareServer "gasto-api/server/share"
	domain "gasto-api/src/GasRecord"
	app "gasto-api/src/GasRecord/App"
	share "gasto-api/src/Share"
	"net/http"
)

func ConfigureResource(repo domain.GasRecordRepository, eventBus share.EventBus) {
	configureSuscriptions(eventBus, repo)
	handlers := make(shareServer.MethodRelation)
	getAllGasRecords := app.MakeGetAllGasRecords(repo)
	createGasRecords := app.MakeCreateGasRecord(repo, eventBus)
	handlers[http.MethodGet] = makeGetAllGetHandler(getAllGasRecords)
	handlers[http.MethodPut] = makeGasRecordPutHandler(createGasRecords)
	gasRecordResource := shareServer.DependingOnMethod(handlers)
	http.HandleFunc("/gasrecord", gasRecordResource)
}

func configureSuscriptions(eventBus share.EventBus, repo domain.GasRecordRepository) {
	updatePerformanceWhenGasRecordCreated := app.MakeUpdatePerformanceWhenGasRecordCreated(repo)
	eventBus.Suscribe(domain.GAS_RECORD_CREATED_NAME, updatePerformanceWhenGasRecordCreated)
}
