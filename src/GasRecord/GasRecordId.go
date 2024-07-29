package gasrecord

import (
	"fmt"
	share "gasto-api/src/Share"

	"github.com/google/uuid"
)

type GasRecordId string

func (val GasRecordId) Value() string {
	return string(val)
}

type InvalidFormatGasRecordId struct {
	value string
}

func (err InvalidFormatGasRecordId) Error() string {
	return fmt.Sprintf("Invalid format for UUID <%v>", err.value)
}

func CreateGasRecordId(id string) (GasRecordId, error) {
	value, emptyError := share.CreateNonEmptyString(id)
	if emptyError != nil {
		return "", emptyError
	}
	validation := uuid.Validate(value.Value())
	if validation != nil {
		return "", InvalidFormatGasRecordId{id}
	}
	return GasRecordId(value.Value()), nil
}
