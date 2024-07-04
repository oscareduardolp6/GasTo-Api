package app

import (
	domain "gasto-api/src/GasRecord"
	"log"
	"time"
)

func MakeUpdatePerformanceGasRecord(repository domain.GasRecordRepository) func(createdRecordId string) error {
	return func(createdRecordId string) error {
		createdGasRecord, gettingError := repository.GetById(createdRecordId)
		if gettingError != nil {
			return gettingError
		}

		previousRecords := repository.Search(SortingByDateCriteriaAndBeforeThanADate{createdGasRecord.Date})
		if len(previousRecords) == 0 {
			log.Print("No performance update, this is the first record")
			return nil
		}
		recordBeforeCreatedRecord := previousRecords[len(previousRecords)-1]
		recordBeforeCreatedRecord.Performance = domain.CalculatePerformance(recordBeforeCreatedRecord, createdGasRecord)

		savingError := repository.Save(recordBeforeCreatedRecord)
		if savingError == nil {
			log.Printf("Performance Updated for Id: %v", recordBeforeCreatedRecord.Id)
		}
		return savingError
	}
}

type SortingByDateCriteriaAndBeforeThanADate struct {
	Date time.Time
}

func (criteria SortingByDateCriteriaAndBeforeThanADate) Filter(val domain.GasRecord) bool {
	return val.Date.Before(criteria.Date)
}

func (criteria SortingByDateCriteriaAndBeforeThanADate) SortingLess(val1, val2 domain.GasRecord) bool {
	return val1.Date.Before(val2.Date)
}