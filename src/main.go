package main

import (
	app "gasto-api/src/GasRecord/App"
	infra "gasto-api/src/GasRecord/Infrastructure/BBoldGasRecordRepository"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("file arg missing")
	}

	repo, createRepositoryError := infra.NewGasRecordRepository()
	if createRepositoryError != nil {
		log.Fatal(createRepositoryError)
	}
	defer repo.Close()
	csvFileName := args[1]

	executeBatchFromFile(csvFileName, csvMapper, repo)

	getAllData := app.MakeGetAllGasRecords(repo)

	log.Print("Inserted Data:")

	records := <-getAllData()

	for _, record := range records {
		log.Println(record)
	}
}
