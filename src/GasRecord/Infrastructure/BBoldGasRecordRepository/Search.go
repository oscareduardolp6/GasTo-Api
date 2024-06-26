package gasrecord_infrastructure_bbold

import (
	. "gasto-api/src/GasRecord"
	share "gasto-api/src/Share"
	"sort"
)

func (repo *bboldGasRepository) Search(criteria share.Criteria[GasRecord]) []GasRecord {
	allRecords := repo.GetAll()

	filterRecords := []GasRecord{}

	for _, record := range allRecords {
		if criteria.Filter(record) {
			filterRecords = append(filterRecords, record)
		}
	}

	sort.Slice(filterRecords, func(i, j int) bool {
		val1 := filterRecords[i]
		val2 := filterRecords[j]
		return criteria.SortingLess(val1, val2)
	})
	return filterRecords
}
