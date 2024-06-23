package gasrecord

import "fmt"

type GasRecordNotSaved struct {
	GasRecord GasRecord
}

func NewRecordNotSaved(gasRecord GasRecord) error {
	return GasRecordNotSaved{gasRecord}
}

func (myError GasRecordNotSaved) Error() string {
	return fmt.Sprintf("Error saving the record (check the log for more details): %v", myError.GasRecord)
}
