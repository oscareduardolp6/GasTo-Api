package gasrecord_infrastructure_inmemory

import (
	domain "gasto-api/src/GasRecord"
	shared "gasto-api/src/Shared"
	"sort"
)

type inMemoryGasRecordRepository struct {
	data []domain.GasRecord
}

func (repo *inMemoryGasRecordRepository) Save(gasRecord domain.GasRecord) error {
	repo.data = append(repo.data, gasRecord)
	return nil
}

func (repo *inMemoryGasRecordRepository) GetAll() []domain.GasRecord {
	records := make([]domain.GasRecord, 0)
	copy(records, repo.data)
	return records
}

func (repo *inMemoryGasRecordRepository) Close() error {
	repo.data = nil
	return nil
}

func (repo *inMemoryGasRecordRepository) Search(criteria shared.Criteria[domain.GasRecord]) []domain.GasRecord {
	allRecords := repo.GetAll()
	filteredValues := shared.Filter(criteria.Filter, allRecords...)

	sort.Slice(filteredValues, func(i, j int) bool {
		val1 := filteredValues[i]
		val2 := filteredValues[j]
		return criteria.SortingLess(val1, val2)
	})

	return filteredValues
}

func (repo *inMemoryGasRecordRepository) GetById(id domain.GasRecordId) (domain.GasRecord, error) {
	finded := shared.Find(func(record domain.GasRecord) bool {
		return record.Id.Equals(id)
	}, repo.data)
	if finded == nil {
		var nilValue domain.GasRecord
		return nilValue, domain.RecordNotFound{Id: id}
	}
	return *finded, nil
}

func NewGasRecordRepository() domain.GasRecordRepository {
	return &inMemoryGasRecordRepository{
		data: make([]domain.GasRecord, 0),
	}
}
