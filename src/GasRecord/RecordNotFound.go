package gasrecord

import (
	"fmt"
)

type RecordNotFound struct {
	Id GasRecordId
}

func (err RecordNotFound) Error() string {
	return fmt.Sprintf("Record with ID: <%v> not found ", err.Id.Value())
}
